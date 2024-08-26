// utility/loki.go
package utility

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/devopscorner/golang-restfulapi-bedrock/src/config"
	"github.com/gin-gonic/gin"
	"github.com/grafana/loki-client-go/loki"
	"github.com/prometheus/common/model"
)

var lokiClient *loki.Client
var lokiURL *url.URL

func InitLokiLogger(cfg *config.Config) error {
	lokiURL, err := url.Parse(cfg.LokiEndpoint)
	if err != nil {
		return fmt.Errorf("invalid Loki URL: %v", err)
	}

	lokiConfig := loki.Config{
		ClientConfig: loki.ClientConfig{
			URL:            lokiURL,
			BatchWait:      1 * time.Second,
			BatchSize:      1024 * 1024,
			Client:         nil,
			BackoffConfig:  nil,
			ExternalLabels: nil,
			Timeout:        0,
		},
	}

	client, err := loki.NewWithConfig(lokiConfig)
	if err != nil {
		return fmt.Errorf("failed to create Loki client: %v", err)
	}

	lokiClient = client
	return nil
}

func LogWithLoki(ctx *gin.Context, level, message string, err error) {
	if lokiClient == nil {
		fmt.Println("Loki client not initialized")
		return
	}

	labels := model.LabelSet{
		"level":    model.LabelValue(level),
		"service":  "file-upload-service",
		"endpoint": model.LabelValue(ctx.FullPath()),
		"method":   model.LabelValue(ctx.Request.Method),
	}

	if err != nil {
		labels["error"] = model.LabelValue(err.Error())
	}

	_ = lokiClient.Handle(labels, time.Now(), message)
}

func LogInfo(ctx *gin.Context, message string) {
	LogWithLoki(ctx, "info", message, nil)
}

func LogWarn(ctx *gin.Context, message string) {
	LogWithLoki(ctx, "warn", message, nil)
}

func LogError(ctx *gin.Context, message string, err error) {
	LogWithLoki(ctx, "error", message, err)
}

func LogFatal(ctx *gin.Context, message string, err error) {
	LogWithLoki(ctx, "fatal", message, err)
	os.Exit(1)
}
