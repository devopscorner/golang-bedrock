// config/config.go
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	GinMode                           string `mapstructure:"GIN_MODE"`
	AppUrl                            string `mapstructure:"APP_URL"`
	AppPort                           int    `mapstructure:"APP_PORT"`
	DbConnection                      string `mapstructure:"DB_CONNECTION"`
	DbHost                            string `mapstructure:"DB_HOST"`
	DbPort                            int    `mapstructure:"DB_PORT"`
	DbDatabase                        string `mapstructure:"DB_DATABASE"`
	DbUsername                        string `mapstructure:"DB_USERNAME"`
	DbPassword                        string `mapstructure:"DB_PASSWORD"`
	JwtAuthUsername                   string `mapstructure:"JWT_AUTH_USERNAME"`
	JwtAuthPassword                   string `mapstructure:"JWT_AUTH_PASSWORD"`
	JwtIssuer                         string `mapstructure:"JWT_AUTH_USERNAME"`
	JwtSecret                         string `mapstructure:"JWT_SECRET"`
	LogLevel                          string `mapstructure:"LOG_LEVEL"`
	AWSRegion                         string `mapstructure:"AWS_REGION"`
	AWSAccessKey                      string `mapstructure:"AWS_ACCESS_KEY"`
	AWSSecretKey                      string `mapstructure:"AWS_SECRET_KEY_ID"`
	AWSBucketName                     string `mapstructure:"AWS_BUCKET_NAME"`
	AmazonBedrockAgentId              string `mapstructure:"AMAZON_BEDROCK_AGENT_ID"`
	AmazonBedrockModelId              string `mapstructure:"AMAZON_BEDROCK_MODEL_ID"`
	AmazonBedrockVersion              string `mapstructure:"AMAZON_BEDROCK_VERSION"`
	OpenSearchEndpoint                string `mapstructure:"OPENSEARCH_ENDPOINT"`
	OpenSearchUsername                string `mapstructure:"OPENSEARCH_USERNAME"`
	OpenSearchPassword                string `mapstructure:"OPENSEARCH_PASSWORD"`
	PrometheusEndpoint                string `mapstructure:"PROMETHEUS_ENDPOINT"`
	PrometheusPort                    int    `mapstructure:"PROMETHEUS_PORT"`
	LokiEndpoint                      string `mapstructure:"LOKI_ENDPOINT"`
	LokiPort                          int    `mapstructure:"LOKI_PORT"`
	GrafanaEndpoint                   string `mapstructure:"GRAFANA_ENDPOINT"`
	GrafanaPort                       int    `mapstructure:"GRAFANA_PORT"`
	GrafanaApiKey                     string `mapstructure:"GRAFANA_API_KEY"`
	OtelEnvironment                   string `mapstructure:"OTEL_ENVIRONMENT"`
	OtelServiceName                   string `mapstructure:"OTEL_SERVICE_NAME"`
	OtelMetricEnable                  string `mapstructure:"OTEL_INSTRUMENTATION_METRIC_ENABLED"`
	OtelTraceEnable                   string `mapstructure:"OTEL_INSTRUMENTATION_TRACE_ENABLED"`
	OtelTraceName                     string `mapstructure:"OTEL_INSTRUMENTATION_TRACE_NAME"`
	OtelLogEnable                     string `mapstructure:"OTEL_INSTRUMENTATION_LOG_ENABLED"`
	OtelOtlpEndpoint                  string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	OtelOtlpPort                      int    `mapstructure:"OTEL_EXPORTER_OTLP_PORT"`
	OtelOtlpInsecure                  string `mapstructure:"OTEL_EXPORTER_OTLP_INSECURE"`
	OtelOtlpHeader                    string `mapstructure:"OTEL_EXPORTER_OTLP_HEADERS"`
	OtelAttributes                    string `mapstructure:"OTEL_RESOURCE_ATTRIBUTES"`
	OtelTimeInterval                  int64  `mapstructure:"OTEL_TIME_INTERVAL"`
	OtelTimeAliveIncrementer          int64  `mapstructure:"OTEL_RANDOM_TIME_ALIVE_INCREMENTER"`
	OtelTotalHeapSizeUpperBound       int64  `mapstructure:"OTEL_RANDOM_TOTAL_HEAP_SIZE_UPPER_BOUND"`
	OtelThreadsActiveUpperBound       int64  `mapstructure:"OTEL_RANDOM_THREAD_ACTIVE_UPPER_BOUND"`
	OtelCpuUsageUpperBound            int64  `mapstructure:"OTEL_RANDOM_CPU_USAGE_UPPER_BOUND"`
	JaegerAgentPort                   int    `mapstructure:"JAEGER_AGENT_PORT"`
	JaegerSamplerType                 string `mapstructure:"JAEGER_SAMPLER_TYPE"`
	JaegerSamplerParam                int    `mapstructure:"JAEGER_SAMPLER_PARAM"`
	JaegerSamplerManagerHostPort      string `mapstructure:"JAEGER_SAMPLER_MANAGER_HOST_PORT"`
	JaegerReporterLogSpan             string `mapstructure:"JAEGER_REPORTER_LOG_SPANS"`
	JaegerReporterBufferFlushInterval int    `mapstructure:"JAEGER_REPORTER_BUFFER_FLUSH_INTERVAL"`
	JaegerReporterMaxQueueSize        int    `mapstructure:"JAEGER_REPORTER_MAX_QUEUE_SIZE"`
	JaegerReporterLocalAgentHostPort  string `mapstructure:"JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT"`
	JaegerReporterCollectorEndpoint   string `mapstructure:"JAEGER_REPORTER_COLLECTOR_ENDPOINT"`
	JaegerReporterCollectorUser       string `mapstructure:"JAEGER_REPORTER_COLLECTOR_USER"`
	JaegerReporterCollectorPassword   string `mapstructure:"JAEGER_REPORTER_COLLECTOR_PASSWORD"`
	JaegerTags                        string `mapstructure:"JAEGER_TAGS"`
	XRayDaemonEndpoint                string `mapstructure:"XRAY_DAEMON_ENDPOINT"`
	XRayDaemonPort                    int    `mapstructure:"XRAY_DAEMON_PORT"`
	XRayVersion                       string `mapstructure:"XRAY_VERSION"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.SetDefault("GIN_MODE", "release")

	viper.SetDefault("APP_URL", "http://0.0.0.0")
	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("DB_CONNECTION", "sqlite")
	viper.SetDefault("DB_HOST", "0.0.0.0")
	viper.SetDefault("DB_PORT", 5000)
	viper.SetDefault("DB_DATABASE", "golang-bedrock.db")
	viper.SetDefault("DB_USERNAME", "root")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("JWT_AUTH_USERNAME", "devopscorner")
	viper.SetDefault("JWT_AUTH_PASSWORD", "DevOpsCorner2024")
	viper.SetDefault("JWT_SECRET", "s3cr3t")

	// LOG_LEVEL: DEBUG | INFO | WARN | ERROR
	viper.SetDefault("LOG_LEVEL", "INFO")

	// AWS MANAGED SERVICES
	viper.SetDefault("AWS_REGION", "us-west-2")
	viper.SetDefault("AWS_ACCESS_KEY", "YOUR_AWS_ACCESS_KEY")
	viper.SetDefault("AWS_SECRET_KEY_ID", "YOUR_AWS_SECRET_KEY_ID")
	viper.SetDefault("AWS_BUCKET_NAME", "devopscorner-bedrock")
	viper.SetDefault("AMAZON_BEDROCK_AGENT_ID", "YOUR_AMAZON_BEDROCK_AGENT_ID")
	viper.SetDefault("AMAZON_BEDROCK_MODEL_ID", "anthropic.claude-3-haiku-20240307-v1:0")
	viper.SetDefault("AMAZON_BEDROCK_VERSION", "bedrock-2023-05-31")

	viper.SetDefault("OPENSEARCH_ENDPOINT", "https://opensearch.us-west-2.es.amazonaws.com")
	viper.SetDefault("OPENSEARCH_USERNAME", "OPENSEARCH_USERNAME")
	viper.SetDefault("OPENSEARCH_PASSWORD", "OPENSEARCH_PASSWORD")

	viper.SetDefault("PROMETHEUS_ENDPOINT", "http://0.0.0.0:9090")
	viper.SetDefault("PROMETHEUS_PORT", 9090)

	viper.SetDefault("LOKI_ENDPOINT", "http://0.0.0.0:3100")
	viper.SetDefault("LOKI_PORT", 3100)

	viper.SetDefault("GRAFANA_ENDPOINT", "http://0.0.0.0:3000")
	viper.SetDefault("GRAFANA_PORT", 3000)
	viper.SetDefault("GRAFANA_API_KEY", "GRAFANA_API_KEY")

	// OPEN TELEMETRY (OTEL)
	viper.SetDefault("OTEL_ENVIRONMENT", "nonprod")                 // Environment OTEL?
	viper.SetDefault("OTEL_INSTRUMENTATION_METRIC_ENABLED", "true") // Prometheus Enable?
	viper.SetDefault("OTEL_INSTRUMENTATION_TRACE_ENABLED", "true")  // Tracing Enable?
	viper.SetDefault("OTEL_INSTRUMENTATION_LOG_ENABLED", "true")    // Logging Enable?
	viper.SetDefault("OTEL_SERVICE_NAME", "golang-bedrock")         // Service Name OTEL
	viper.SetDefault("OTEL_EXPORTER_OTLP_ENDPOINT", "http://0.0.0.0:4317")
	viper.SetDefault("OTEL_EXPORTER_OTLP_PORT", 4317)
	viper.SetDefault("OTEL_EXPORTER_OTLP_INSECURE", "true")
	viper.SetDefault("OTEL_EXPORTER_OTLP_HEADERS", "")
	viper.SetDefault("OTEL_RESOURCE_ATTRIBUTES", "")

	// TRACING with OTEL
	viper.SetDefault("OTEL_TIME_INTERVAL", 1)
	viper.SetDefault("OTEL_RANDOM_TIME_ALIVE_INCREMENTER", 1)
	viper.SetDefault("OTEL_RANDOM_TOTAL_HEAP_SIZE_UPPER_BOUND", 100)
	viper.SetDefault("OTEL_RANDOM_THREAD_ACTIVE_UPPER_BOUND", 10)
	viper.SetDefault("OTEL_RANDOM_CPU_USAGE_UPPER_BOUND", 100)

	// TRACING with XRAY
	viper.SetDefault("XRAY_VERSION", "latest")
	viper.SetDefault("XRAY_DAEMON_ENDPOINT", "https://xray.us-west-2.amazonaws.com")
	viper.SetDefault("XRAY_DAEMON_PORT", 2000)

	// TRACING with JAEGER
	viper.SetDefault("JAEGER_AGENT_PORT", 6831)
	viper.SetDefault("JAEGER_SAMPLER_TYPE", "const")
	viper.SetDefault("JAEGER_SAMPLER_PARAM", 1)
	viper.SetDefault("JAEGER_SAMPLER_MANAGER_HOST_PORT", "http://0.0.0.0:5778")
	viper.SetDefault("JAEGER_REPORTER_LOG_SPANS", "true")
	viper.SetDefault("JAEGER_REPORTER_BUFFER_FLUSH_INTERVAL", 5)
	viper.SetDefault("JAEGER_REPORTER_MAX_QUEUE_SIZE", 100)
	viper.SetDefault("JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT", "http://0.0.0.0:6831")
	viper.SetDefault("JAEGER_REPORTER_COLLECTOR_ENDPOINT", "http://0.0.0.0:14268/api/traces")
	viper.SetDefault("JAEGER_REPORTER_COLLECTOR_USER", "")
	viper.SetDefault("JAEGER_REPORTER_COLLECTOR_PASSWORD", "")
	viper.SetDefault("JAEGER_TAGS", "golang,otel,restful,api,filestore")

	config := &Config{
		GinMode:                           viper.GetString("GIN_MODE"),
		AppUrl:                            viper.GetString("APP_URL"),
		AppPort:                           viper.GetInt("APP_PORT"),
		DbConnection:                      viper.GetString("DB_CONNECTION"),
		DbHost:                            viper.GetString("DB_HOST"),
		DbPort:                            viper.GetInt("DB_PORT"),
		DbDatabase:                        viper.GetString("DB_DATABASE"),
		DbUsername:                        viper.GetString("DB_USERNAME"),
		DbPassword:                        viper.GetString("DB_PASSWORD"),
		JwtAuthUsername:                   viper.GetString("JWT_AUTH_USERNAME"),
		JwtAuthPassword:                   viper.GetString("JWT_AUTH_PASSWORD"),
		JwtSecret:                         viper.GetString("JWT_SECRET"),
		AWSRegion:                         viper.GetString("AWS_REGION"),
		AWSAccessKey:                      viper.GetString("AWS_ACCESS_KEY"),
		AWSSecretKey:                      viper.GetString("AWS_SECRET_KEY_ID"),
		AWSBucketName:                     viper.GetString("AWS_BUCKET_NAME"),
		AmazonBedrockAgentId:              viper.GetString("AMAZON_BEDROCK_AGENT_ID"),
		AmazonBedrockModelId:              viper.GetString("AMAZON_BEDROCK_MODEL_ID"),
		OpenSearchEndpoint:                viper.GetString("OPENSEARCH_ENDPOINT"),
		OpenSearchUsername:                viper.GetString("OPENSEARCH_USERNAME"),
		OpenSearchPassword:                viper.GetString("OPENSEARCH_PASSWORD"),
		PrometheusEndpoint:                viper.GetString("PROMETHEUS_ENDPOINT"),
		PrometheusPort:                    viper.GetInt("PROMETHEUS_PORT"),
		GrafanaEndpoint:                   viper.GetString("GRAFANA_ENDPOINT"),
		GrafanaApiKey:                     viper.GetString("GRAFANA_API_KEY"),
		LokiEndpoint:                      viper.GetString("LOKI_ENDPOINT"),
		LokiPort:                          viper.GetInt("LOKI_PORT"),
		OtelEnvironment:                   viper.GetString("OTEL_ENVIRONMENT"),
		OtelMetricEnable:                  viper.GetString("OTEL_INSTRUMENTATION_METRIC_ENABLED"),
		OtelTraceEnable:                   viper.GetString("OTEL_INSTRUMENTATION_TRACE_ENABLED"),
		OtelTraceName:                     viper.GetString("OTEL_INSTRUMENTATION_TRACE_NAME"),
		OtelLogEnable:                     viper.GetString("OTEL_INSTRUMENTATION_LOG_ENABLED"),
		OtelServiceName:                   viper.GetString("OTEL_SERVICE_NAME"),
		OtelOtlpEndpoint:                  viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT"),
		OtelOtlpPort:                      viper.GetInt("OTEL_EXPORTER_OTLP_PORT"),
		OtelOtlpInsecure:                  viper.GetString("OTEL_EXPORTER_OTLP_INSECURE"),
		OtelOtlpHeader:                    viper.GetString("OTEL_EXPORTER_OTLP_HEADERS"),
		OtelAttributes:                    viper.GetString("OTEL_RESOURCE_ATTRIBUTES"),
		OtelTimeInterval:                  viper.GetInt64("OTEL_TIME_INTERVAL"),
		OtelTimeAliveIncrementer:          viper.GetInt64("OTEL_RANDOM_TIME_ALIVE_INCREMENTER"),
		OtelTotalHeapSizeUpperBound:       viper.GetInt64("OTEL_RANDOM_TOTAL_HEAP_SIZE_UPPER_BOUND"),
		OtelThreadsActiveUpperBound:       viper.GetInt64("OTEL_RANDOM_THREAD_ACTIVE_UPPER_BOUND"),
		OtelCpuUsageUpperBound:            viper.GetInt64("OTEL_RANDOM_CPU_USAGE_UPPER_BOUND"),
		JaegerAgentPort:                   viper.GetInt("JAEGER_AGENT_PORT"),
		JaegerSamplerType:                 viper.GetString("JAEGER_SAMPLER_TYPE"),
		JaegerSamplerParam:                viper.GetInt("JAEGER_SAMPLER_PARAM"),
		JaegerSamplerManagerHostPort:      viper.GetString("JAEGER_SAMPLER_MANAGER_HOST_PORT"),
		JaegerReporterLogSpan:             viper.GetString("JAEGER_REPORTER_LOG_SPANS"),
		JaegerReporterBufferFlushInterval: viper.GetInt("JAEGER_REPORTER_BUFFER_FLUSH_INTERVAL"),
		JaegerReporterMaxQueueSize:        viper.GetInt("JAEGER_REPORTER_MAX_QUEUE_SIZE"),
		JaegerReporterLocalAgentHostPort:  viper.GetString("JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT"),
		JaegerReporterCollectorEndpoint:   viper.GetString("JAEGER_REPORTER_COLLECTOR_ENDPOINT"),
		JaegerReporterCollectorUser:       viper.GetString("JAEGER_REPORTER_COLLECTOR_USER"),
		JaegerReporterCollectorPassword:   viper.GetString("JAEGER_REPORTER_COLLECTOR_PASSWORD"),
		JaegerTags:                        viper.GetString("JAEGER_TAGS"),
		XRayVersion:                       viper.GetString("XRAY_VERSION"),
		XRayDaemonEndpoint:                viper.GetString("XRAY_DAEMON_ENDPOINT"),
		XRayDaemonPort:                    viper.GetInt("XRAY_DAEMON_PORT"),
	}

	return config, nil
}
