// utility/bedrock.go
package utility

import (
	"context"
	"encoding/json"
	"fmt"

	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/devopscorner/golang-bedrock/src/config"
)

var bedrockClient *bedrockruntime.Client

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

func AnalyzeWithBedrock(ctx context.Context, content string) (string, error) {
	input := &bedrockruntime.InvokeModelInput{
		Body:        []byte(constructPrompt(content)),
		ModelId:     ptrString("anthropic.claude-v2"),
		ContentType: ptrString("application/json"),
	}

	output, err := bedrockClient.InvokeModel(ctx, input)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(output.Body, &response); err != nil {
		return "", err
	}

	return response["completion"].(string), nil
}

func constructPrompt(content string) string {
	return fmt.Sprintf(`{
		"prompt": "Human: %s\n\nAssistant: Here's my analysis:",
		"max_tokens_to_sample": 500,
		"temperature": 0.7,
		"top_p": 1
	}`, content)
}

func ptrString(s string) *string {
	return &s
}
