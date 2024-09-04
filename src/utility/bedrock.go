// utility/bedrock.go
package utility

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/smithy-go"
	"github.com/devopscorner/golang-bedrock/src/config"
	"github.com/sirupsen/logrus"
)

var (
	bedrockClient *bedrockruntime.Client
	logger        *logrus.Logger
)

// Metrics struct to hold various metrics
type Metrics struct {
	TotalLatency    time.Duration
	UploadLatency   time.Duration
	AnalysisLatency time.Duration
	InputTokens     int
	OutputTokens    int
}

// InitLogger should be called from your main application to set up the logger
func InitLogger(l *logrus.Logger) {
	logger = l
}

func InitBedrock(cfg *config.Config) error {
	bedrockCfg, err := awsCfg.LoadDefaultConfig(context.TODO(),
		awsCfg.WithRegion(cfg.AWSRegion),
		awsCfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AWSAccessKey,
			cfg.AWSSecretKey,
			"",
		)),
	)
	if err != nil {
		return err
	}

	bedrockClient = bedrockruntime.NewFromConfig(bedrockCfg)
	return nil
}

func AnalyzeWithBedrock(ctx context.Context, content string) (string, Metrics, error) {
	startTime := time.Now()
	metrics := Metrics{}

	modelId := config.AmazonBedrockModelId()
	if modelId == "" {
		return "", metrics, fmt.Errorf("Bedrock model ID is not configured")
	}

	if logger != nil {
		logger.WithFields(logrus.Fields{
			"modelId": modelId,
		}).Info("Analyzing with Bedrock")
	}

	var inputBytes []byte
	var err error

	uploadStart := time.Now()
	if isClaudeV3Model(modelId) {
		inputBytes, err = constructClaudeV3Input(content)
	} else {
		inputBytes, err = constructStandardInput(content)
	}
	metrics.UploadLatency = time.Since(uploadStart)

	if err != nil {
		return "", metrics, fmt.Errorf("✗ Failed to construct input: %w", err)
	}

	metrics.InputTokens = estimateTokenCount(content)

	input := &bedrockruntime.InvokeModelInput{
		Body:        inputBytes,
		ModelId:     aws.String(modelId),
		ContentType: aws.String("application/json"),
	}

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	analysisStart := time.Now()
	output, err := bedrockClient.InvokeModel(ctx, input)
	metrics.AnalysisLatency = time.Since(analysisStart)

	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) {
			if logger != nil {
				logger.WithFields(logrus.Fields{
					"errorCode":    ae.ErrorCode(),
					"errorMessage": ae.ErrorMessage(),
				}).Error("✗ Bedrock API error")
			}

			switch ae.ErrorCode() {
			case "ResourceNotFoundException":
				return "", metrics, fmt.Errorf("✗ Bedrock model not found. Please check your model ID and permissions")
			case "AccessDeniedException":
				return "", metrics, fmt.Errorf("✗ Access denied to Bedrock model. Please check your IAM permissions")
			case "ValidationException":
				return "", metrics, fmt.Errorf("✗ Invalid input for Bedrock model: %s", ae.ErrorMessage())
			default:
				return "", metrics, fmt.Errorf("❗ Bedrock error: %s", ae.ErrorMessage())
			}
		}

		if ctx.Err() == context.DeadlineExceeded {
			return "", metrics, fmt.Errorf("❗ Bedrock analysis timed out")
		}

		return "", metrics, fmt.Errorf("✗ Failed to invoke Bedrock model: %w", err)
	}

	response, err := parseBedrockResponse(output.Body, modelId)
	if err != nil {
		return "", metrics, err
	}

	metrics.OutputTokens = estimateTokenCount(response)
	metrics.TotalLatency = time.Since(startTime)

	return response, metrics, nil
}

// estimateTokenCount is a simple function to estimate the number of tokens in a string
// This is a very rough estimate and should be replaced with a proper tokenizer for production use
func estimateTokenCount(s string) int {
	return len(s) / 4 // Assuming an average of 4 characters per token
}

func isClaudeV3Model(modelId string) bool {
	return modelId == "anthropic.claude-3-sonnet-20240229-v1:0" ||
		modelId == "anthropic.claude-3-haiku-20240307-v1:0" ||
		modelId == "anthropic.claude-v2"
}

func constructClaudeV3Input(content string) ([]byte, error) {
	input := struct {
		BedrockVersion string    `json:"anthropic_version"`
		MaxTokens      int       `json:"max_tokens"`
		Temperature    float64   `json:"temperature"`
		TopP           float64   `json:"top_p"`
		Messages       []Message `json:"messages"`
	}{
		BedrockVersion: config.AmazonBedrockVersion(),
		MaxTokens:      500,
		Temperature:    0.7,
		TopP:           1,
		Messages: []Message{
			{
				Role:    "user",
				Content: content,
			},
		},
	}

	return json.Marshal(input)
}

func constructStandardInput(content string) ([]byte, error) {
	input := struct {
		Prompt            string  `json:"prompt"`
		BedrockVersion    string  `json:"anthropic_version"`
		MaxTokensToSample int     `json:"max_tokens_to_sample"`
		Temperature       float64 `json:"temperature"`
		TopP              float64 `json:"top_p"`
	}{
		Prompt:            fmt.Sprintf("Human: %s\n\nAssistant: Here's my analysis:", content),
		BedrockVersion:    config.AmazonBedrockVersion(),
		MaxTokensToSample: 500,
		Temperature:       0.7,
		TopP:              1,
	}

	return json.Marshal(input)
}

func parseBedrockResponse(responseBody []byte, modelId string) (string, error) {
	if isClaudeV3Model(modelId) {
		var response struct {
			Content []struct {
				Text string `json:"text"`
			} `json:"content"`
		}
		if err := json.Unmarshal(responseBody, &response); err != nil {
			return "", fmt.Errorf("✗ Failed to unmarshal Claude V3 response: %w", err)
		}
		if len(response.Content) == 0 || response.Content[0].Text == "" {
			return "", fmt.Errorf("Empty response from Claude V3")
		}
		return response.Content[0].Text, nil
	} else {
		var response struct {
			Completion string `json:"completion"`
		}
		if err := json.Unmarshal(responseBody, &response); err != nil {
			return "", fmt.Errorf("✗ Failed to unmarshal standard response: %w", err)
		}
		return response.Completion, nil
	}
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// LogMetrics logs the collected metrics
func LogMetrics(metrics Metrics) {
	if logger != nil {
		logger.WithFields(logrus.Fields{
			"totalLatency":    metrics.TotalLatency,
			"uploadLatency":   metrics.UploadLatency,
			"analysisLatency": metrics.AnalysisLatency,
			"inputTokens":     metrics.InputTokens,
			"outputTokens":    metrics.OutputTokens,
		}).Info("Bedrock analysis metrics")
	}
}
