# Golang RESTful API with Amazon Bedrock

ObservabilityAI: Memanfaatkan Amazon Bedrock Untuk Memantau Kinerja RESTful API Golang

[![goreport](https://goreportcard.com/badge/github.com/devopscorner/golang-bedrock/src)](https://goreportcard.com/badge/github.com/devopscorner/golang-bedrock/src)
[![all contributors](https://img.shields.io/github/contributors/devopscorner/golang-bedrock)](https://github.com/devopscorner/golang-bedrock/graphs/contributors)
[![tags](https://img.shields.io/github/v/tag/devopscorner/golang-bedrock?sort=semver)](https://github.com/devopscorner/golang-bedrock/releases)
[![docker pulls](https://img.shields.io/docker/pulls/devopscorner/golang-bedrock.svg)](https://hub.docker.com/r/devopscorner/golang-bedrock/)
![download all](https://img.shields.io/github/downloads/devopscorner/golang-bedrock/total.svg)
![download latest](https://img.shields.io/github/downloads/devopscorner/golang-bedrock/1.1.4/total)
![view](https://views.whatilearened.today/views/github/devopscorner/golang-bedrock.svg)
![clone](https://img.shields.io/badge/dynamic/json?color=success&label=clone&query=count&url=https://github.com/devopscorner/golang-bedrock/blob/master/clone.json?raw=True&logo=github)
[![issues](https://img.shields.io/github/issues/devopscorner/golang-bedrock)](https://github.com/devopscorner/golang-bedrock/issues)
[![pull requests](https://img.shields.io/github/issues-pr/devopscorner/golang-bedrock)](https://github.com/devopscorner/golang-bedrock/pulls)
[![forks](https://img.shields.io/github/forks/devopscorner/golang-bedrock)](https://github.com/devopscorner/golang-bedrock/fork)
[![stars](https://img.shields.io/github/stars/devopscorner/golang-bedrock)](https://github.com/devopscorner/golang-bedrock/stargazers)
[![license](https://img.shields.io/github/license/devopscorner/golang-bedrock)](https://img.shields.io/github/license/devopscorner/golang-bedrock)

---

## Available Tags

### Alpine GOLANG

| Image name | Size |
|------------|------|
| `devopscorner/golang-bedrock:latest` | [![docker image size](https://img.shields.io/docker/image-size/devopscorner/golang-bedrock/latest.svg?label=Image%20size&logo=docker)](https://hub.docker.com/repository/docker/devopscorner/golang-bedrock/tags?page=1&ordering=last_updated&name=latest) ![default-latest](https://img.shields.io/static/v1?label=latest&message=default&color=brightgreen) ![latest](https://img.shields.io/static/v1?label=latest&message=alpine&color=orange) |
| `devopscorner/golang-bedrock:alpine` | [![docker image size](https://img.shields.io/docker/image-size/devopscorner/golang-bedrock/alpine.svg?label=Image%20size&logo=docker)](https://hub.docker.com/repository/docker/devopscorner/golang-bedrock/tags?page=1&ordering=last_updated&name=alpine) ![latest](https://img.shields.io/static/v1?label=latest&message=alpine&color=orange) |
| `devopscorner/golang-bedrock:alpine-latest` | [![docker image size](https://img.shields.io/docker/image-size/devopscorner/golang-bedrock/alpine-latest.svg?label=Image%20size&logo=docker)](https://hub.docker.com/repository/docker/devopscorner/golang-bedrock/tags?page=1&ordering=last_updated&name=alpine-latest) |
| `devopscorner/golang-bedrock:alpine-3.18` | [![docker image size](https://img.shields.io/docker/image-size/devopscorner/golang-bedrock/alpine-3.18.svg?label=Image%20size&logo=docker)](https://hub.docker.com/repository/docker/devopscorner/golang-bedrock/tags?page=1&ordering=last_updated&name=alpine-3.18) |

---

## GO Repository Pattern

- Folder Structure

```
.
├── cmd
│   └── migrate_file_upload.go
├── config
│   ├── config.go
│   ├── const.go
│   ├── logger.go
│   └── value.go
├── controller
│   ├── file_controller.go
│   └── login_controller.go
├── driver
│   ├── db.go
│   ├── mysql.go
│   ├── psql.go
│   └── sqlite.go
├── go.mod
├── go.sum
├── golang-bedrock.db
├── .env.example
├── main.go
├── middleware
│   └── auth_middleware.go
├── model
│   └── file.go
├── repository
│   └── file_repository.go
├── routes
│   ├── file_routes.go
│   └── main_routes.go
├── utility
│   ├── bedrock.go
│   ├── genid.go
│   ├── loki.go
│   ├── otel.go
│   ├── prometheus.go
│   └── s3.go
└── view
    ├── error_view.go
    ├── file_view.go
    └── login_view.go

11 directories, 30 files
```

## Coverages:
### AWS Services
- Amazon Elastic Container Registry (ECR)
- Amazon EKS

### Others:
- Docker
- Docker-Compose
- Prometheus
- Loki
- Fluentbit
- OpenTelemetry
- Golang
- HelmChart
- Postman
- Terraform

## Environment Variables

- Default Value
  ```
  GIN_MODE=release
  APP_URL=http://0.0.0.0
  APP_PORT=8080

  AWS_REGION=us-west-2
  AWS_ACCESS_KEY=YOUR_AWS_KEY
  AWS_SECRET_KEY_ID=YOUR_SECRET_KEY
  AWS_BUCKET_NAME=devopscorner-bedrock

  JWT_AUTH_USERNAME=devopscorner
  JWT_AUTH_PASSWORD=DevOpsCorner2024
  JWT_SECRET=s3cr3t

  LOG_LEVEL=INFO
  ```

- Multi Driver Connection
  ```
  # default
  DB_CONNECTION=sqlite
  ---
  Available for:
  - sqlite
  - mysql
  - postgres
  ```

- AWS X-Ray (AWS Distribution Tracing System)
  - `XRAY_VERSION`: Daemon Version X-Ray
    - `latest`: Use the latest version of the AWS X-Ray daemon available.
    - `3.x`: Use version 3.x of the AWS X-Ray daemon.
    - `2.x`: Use version 2.x of the AWS X-Ray daemon.
    - `1.x`: Use version 1.x of the AWS X-Ray daemon.

  - `XRAY_DAEMON_ENDPOINT: Daemon Endpoint of X-Ray
    ```
    XRAY_DAEMON_ENDPOINT=http://localhost:2000
    ```
  - `XRAY_DAEMON_PORT: Daemon Endpoint Port of X-Ray
    ```
    XRAY_DAEMON_PORT=2000
    ```

- Prometheus
  ```
  PROMETHEUS_ENDPOINT=http://localhost:9090
  PROMETHEUS_PORT=9090
  ```

- Loki
  ```
  LOKI_ENDPOINT=http://0.0.0.0:3100
  LOKI_PORT=3100
  ```

- Grafana
  ```
  GRAFANA_ENDPOINT=http://localhost:3000
  GRAFANA_PORT=3000
  GRAFANA_API_KEY=YOUR_GRAFANA_API_KEY
  ```

- OpenTelemetry
  ```
  OTEL_INSTRUMENTATION_METRIC_ENABLED=true
  OTEL_INSTRUMENTATION_TRACE_ENABLED=true
  OTEL_INSTRUMENTATION_LOG_ENABLED=true

  # Trace Type: xray / jaeger
  OTEL_INSTRUMENTATION_TRACE_NAME=jaeger

  OTEL_ENVIRONMENT=nonprod
  OTEL_SERVICE_NAME=golang-bedrock
  OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
  OTEL_EXPORTER_OTLP_PORT=4317
  OTEL_EXPORTER_OTLP_INSECURE=true
  OTEL_EXPORTER_OTLP_HEADERS=
  OTEL_RESOURCE_ATTRIBUTES=
  ```

- Jaeger Environment
  - `JAEGER_SERVICE_NAME`: The name of the service being instrumented (`JAEGER_SERVICE_NAME` = `OTEL_SERVICE_NAME`).
  - `JAEGER_AGENT_HOST`: The host name or IP address of the Jaeger agent to use for sending trace data.
  - `JAEGER_AGENT_PORT`: The port number of the Jaeger agent to use for sending trace data.
  - `JAEGER_SAMPLER_TYPE`: The type of sampling to use for tracing (e.g. const, probabilistic, rateLimiting, remote).
  - `JAEGER_SAMPLER_PARAM`: The parameter value to use for the selected sampler type.
  - `JAEGER_SAMPLER_MANAGER_HOST_PORT`: The host name and port of the remote sampling manager to use (if using remote sampling).
  - `JAEGER_REPORTER_LOG_SPANS`: Whether to log spans instead of sending them to a Jaeger agent (e.g. true or false).
  - `JAEGER_REPORTER_MAX_QUEUE_SIZE`: The maximum number of spans that can be queued before they are sent to a Jaeger agent.
  - `JAEGER_REPORTER_FLUSH_INTERVAL`: The interval at which to flush the span queue and send spans to a Jaeger agent (e.g. 1 * time.Second).
  - `JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT`: The host and port of the local Jaeger agent, if using UDP instead of HTTP.
  - `JAEGER_REPORTER_COLLECTOR_ENDPOINT`: The endpoint URL of the Jaeger collector (e.g. `http://localhost:14268/api/traces`).
  - `JAEGER_REPORTER_COLLECTOR_USER`: The username for authenticating with the Jaeger collector, if required.
  - `JAEGER_REPORTER_COLLECTOR_PASSWORD`: The password for authenticating with the Jaeger collector, if required.
  - `JAEGER_TAGS`: Additional tags to attach to the trace (e.g. key1=value1,key2=value2).


## Tested Environment

### Versioning

- Docker version

  ```
  docker version

  Client:
    Version:           27.1.1-rd
    API version:       1.46
    Go version:        go1.21.12
    Git commit:        cdc3063
    Built:             Wed Jul 24 17:06:24 2024
    OS/Arch:           darwin/arm64
    Context:           default
  ```

- Docker-Compose version

  ```
  docker-compose -v
  ---
  Docker Compose version v2.29.1
  ```

- AWS Cli

  ```
  aws --version
  ---
  aws-cli/2.9.16 Python/3.9.11 Darwin/23.6.0 exe/x86_64 prompt/off
  ```

- Terraform Cli

  ```
  terraform version
  ---
  Terraform v1.8.4
  on darwin_arm64
  - provider registry.terraform.io/hashicorp/aws v3.74.3
  - provider registry.terraform.io/hashicorp/local v2.1.0
  - provider registry.terraform.io/hashicorp/null v3.1.0
  - provider registry.terraform.io/hashicorp/random v3.1.0
  - provider registry.terraform.io/hashicorp/time v0.7.2
  ```

- OpenTofu Cli

  ```
  tofu --version                                                                                                                                                                                                             ─╯
  ---
  OpenTofu v1.7.3
  on darwin_arm64
  ```

- Terraform / OpenTofu / Terragrunt Environment Cli

  ```
  tenv --version                                                                                                                                                                                                             ─╯
  ---
  tenv version 3.1.0
  ```

- Golang Version Manager (GVM)

  ```
  gvm version
  ---
  Go Version Manager v1.0.22 installed at /Users/devopscorner/.gvm
  ```

- Golang Version
  ```
  go version
  ---
  go version go1.21.4 darwin/arm64
  ```

## Security Check

Make sure that you didn't push sensitive information in this repository

- [ ] AWS Credentials (AWS_ACCESS_KEY, AWS_SECRET_KEY)
- [ ] AWS Account ID
- [ ] AWS Resources ARN
- [ ] Username & Password
- [ ] Private (id_rsa) & Public Key (id_rsa.pub)
- [ ] DNS Zone ID
- [ ] APP & API Key

## Copyright

- Author: **Dwi Fahni Denni (@zeroc0d3)**
- Vendor: **DevOps Corner Indonesia (devopscorner.id)**
- License: **Apache v2**
