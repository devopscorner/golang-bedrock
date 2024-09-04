// utility/loki.go
package utility

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/devopscorner/golang-bedrock/src/config"
	"github.com/gin-gonic/gin"
	"github.com/grafana/loki-client-go/loki"
	"github.com/grafana/loki-client-go/pkg/urlutil"
	"github.com/prometheus/common/model"
)

var lokiClient *loki.Client
var lokiURL *url.URL

func InitLokiLogger(cfg *config.Config) error {
	if cfg.LokiEndpoint == "" {
		return fmt.Errorf("❗ Loki endpoint is not configured")
	}

	var err error
	lokiURL, err = url.Parse(cfg.LokiEndpoint)
	if err != nil {
		return fmt.Errorf("❗ Invalid Loki URL: %v", err)
	}

	if lokiURL.Scheme == "" {
		return fmt.Errorf("❗ Loki URL scheme is missing (e.g., http:// or https://)")
	}

	lokiConfig := loki.Config{
		URL: urlutil.URLValue{
			URL: lokiURL,
		},
		BatchWait: 1 * time.Second,
		BatchSize: 1024 * 1024,
		Timeout:   10 * time.Second,
	}

	client, err := loki.New(lokiConfig)
	if err != nil {
		return fmt.Errorf("✗ Failed to create Loki client: %v", err)
	}

	lokiClient = client
	return nil
}

func LogWithLoki(ctx *gin.Context, level, message string, err error) {
	if lokiClient == nil {
		fmt.Println("❗ Loki client not initialized")
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
