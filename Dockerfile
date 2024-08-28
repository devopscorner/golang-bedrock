### Builder ###
FROM golang:1.21.4-alpine3.18 as builder

WORKDIR /go/src/app
ENV GIN_MODE=release
ENV GOPATH=/go

RUN apk add --no-cache \
        build-base \
        git \
        curl \
        make \
        bash \
        ca-certificates

COPY src /go/src/app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    cd /go/src/app && \
        go build -mod=readonly -ldflags="-s -w" -o goapp

### Binary ###
FROM golang:1.21.4-alpine3.18

ARG BUILD_DATE
ARG BUILD_VERSION
ARG GIT_COMMIT
ARG GIT_URL

ENV VENDOR="DevOpsCornerId"
ENV AUTHOR="DevOpsCorner.id <support@devopscorner.id>"
ENV IMG_NAME="alpine"
ENV IMG_VERSION="3.18"
ENV IMG_DESC="Docker GO App Alpine 3.18"
ENV IMG_ARCH="linux/amd64"

ENV ALPINE_VERSION="3.18"
ENV GIN_MODE=release
ENV GOLANG_VERSION=1.21.4
ENV APP_URL=${APP_URL:-http://0.0.0.0}
ENV APP_PORT=${APP_PORT:-8080}
ENV DB_CONNECTION=${DB_CONNECTION:-sqlite}
ENV DB_HOST=${DB_HOST:-0.0.0.0}
ENV DB_PORT=${DB_PORT:-5000}
ENV DB_DATABASE=${DB_DATABASE:-golang-bedrock.db}
ENV DB_USERNAME=${DB_USERNAME:-root}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV JWT_AUTH_USERNAME=${JWT_AUTH_USERNAME:-devopscorner}
ENV JWT_AUTH_PASSWORD=${JWT_AUTH_PASSWORD:-DevOpsCorner2024}
ENV JWT_SECRET=${JWT_SECRET:-s3cr3t}

# Logger
ENV LOG_LEVEL=${LOG_LEVEL:-INFO}

# AWS Credentials
ENV AWS_REGION=${AWS_REGION:-us-west-2}
ENV AWS_ACCESS_KEY=${AWS_ACCESS_KEY:-YOUR_AWS_ACCESS_KEY}
ENV AWS_SECRET_KEY_ID=${AWS_SECRET_KEY_ID:-YOUR_AWS_SECRET_KEY_ID}
ENV AWS_BUCKET_NAME=${AWS_BUCKET_NAME:-devopscorner-bedrock}

# OpenSearch
ENV OPENSEARCH_ENDPOINT=${OPENSEARCH_ENDPOINT:-https://opensearch.us-west-2.es.amazonaws.com}
ENV OPENSEARCH_USERNAME=${OPENSEARCH_USERNAME:-devopscorner}
ENV OPENSEARCH_PASSWORD=${OPENSEARCH_PASSWORD:-DevOpsCorner2024}

# Prometheus
ENV PROMETHEUS_ENDPOINT=${PROMETHEUS_ENDPOINT:-http://localhost:9090}
ENV PROMETHEUS_PORT=${PROMETHEUS_PORT:-9090}

ENV LOKI_ENDPOINT=${LOKI_ENDPOINT:-http://0.0.0.0:3100}
ENV LOKI_PORT=${LOKI_PORT:-3100}

# Grafana
ENV GRAFANA_ENDPOINT=${GRAFANA_ENDPOINT:-http://localhost:3000}
ENV GRAFANA_PORT=${GRAFANA_PORT:-3000}
ENV GRAFANA_API_KEY=${GRAFANA_API_KEY:-YOUR_GRAFANA_API_KEY}

# OpenTelemetry
ENV OTEL_INSTRUMENTATION_METRIC_ENABLED=${OTEL_INSTRUMENTATION_METRIC_ENABLED:-true}
ENV OTEL_INSTRUMENTATION_TRACE_ENABLED=${OTEL_INSTRUMENTATION_TRACE_ENABLED:-true}
ENV OTEL_INSTRUMENTATION_LOG_ENABLED=${OTEL_INSTRUMENTATION_LOG_ENABLED:-true}

# Trace Type: xray / jaeger
ENV OTEL_INSTRUMENTATION_TRACE_NAME=${OTEL_INSTRUMENTATION_TRACE_NAME:-jaeger}

ENV OTEL_ENVIRONMENT=${OTEL_ENVIRONMENT:-nonprod}
ENV OTEL_SERVICE_NAME=${OTEL_SERVICE_NAME:-golang-bedrock}
ENV OTEL_EXPORTER_OTLP_ENDPOINT=${OTEL_EXPORTER_OTLP_ENDPOINT:-http://localhost:4317}
ENV OTEL_EXPORTER_OTLP_INSECURE=${OTEL_EXPORTER_OTLP_INSECURE:-true}
ENV OTEL_EXPORTER_OTLP_HEADERS=${OTEL_EXPORTER_OTLP_HEADERS}
ENV OTEL_RESOURCE_ATTRIBUTES=${OTEL_RESOURCE_ATTRIBUTES}
ENV OTEL_TIME_INTERVAL=${OTEL_TIME_INTERVAL:-1}
ENV OTEL_RANDOM_TIME_ALIVE_INCREMENTER=${OTEL_RANDOM_TIME_ALIVE_INCREMENTER:-1}
ENV OTEL_RANDOM_TOTAL_HEAP_SIZE_UPPER_BOUND=${OTEL_RANDOM_TOTAL_HEAP_SIZE_UPPER_BOUND:-100}
ENV OTEL_RANDOM_THREAD_ACTIVE_UPPOR_BOUND=${OTEL_RANDOM_THREAD_ACTIVE_UPPOR_BOUND:-10}
ENV OTEL_RANDOM_CPU_USAGE_UPPER_BOUND=${OTEL_RANDOM_CPU_USAGE_UPPER_BOUND:-100}

# X-Ray
ENV XRAY_VERSION=${XRAY_VERSION:-latest}
ENV XRAY_DAEMON_ENDPOINT=${XRAY_DAEMON_ENDPOINT:-https://xray.us-west-2.amazonaws.com}
ENV XRAY_DAEMON_PORT=${XRAY_DAEMON_PORT:-2000}

# Jaeger
ENV JAEGER_AGENT_PORT=${JAEGER_AGENT_PORT:-6831}
ENV JAEGER_SAMPLER_TYPE=${JAEGER_SAMPLER_TYPE:-const}
ENV JAEGER_SAMPLER_PARAM=${JAEGER_SAMPLER_PARAM:-1}
ENV JAEGER_SAMPLER_MANAGER_HOST_PORT=${JAEGER_SAMPLER_MANAGER_HOST_PORT}
ENV JAEGER_REPORTER_LOG_SPANS=${JAEGER_REPORTER_LOG_SPANS:-true}
ENV JAEGER_REPORTER_BUFFER_FLUSH_INTERVAL=${JAEGER_REPORTER_BUFFER_FLUSH_INTERVAL:-5}
ENV JAEGER_REPORTER_MAX_QUEUE_SIZE=${JAEGER_REPORTER_MAX_QUEUE_SIZE:-100}
ENV JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT=${JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT}
ENV JAEGER_REPORTER_COLLECTOR_ENDPOINT=${JAEGER_REPORTER_COLLECTOR_ENDPOINT:-http://localhost:14268/api/traces}
ENV JAEGER_REPORTER_COLLECTOR_USER=${JAEGER_REPORTER_COLLECTOR_USER:-devopscorner}
ENV JAEGER_REPORTER_COLLECTOR_PASSWORD=${JAEGER_REPORTER_COLLECTOR_PASSWORD:-DevOpsCorner2024}
ENV JAEGER_TAGS=${JAEGER_TAGS:-golang,otel,restful,api,bookstore}

LABEL maintainer="$AUTHOR" \
        architecture="$IMG_ARCH" \
        alpine-version="$ALPINE_VERSION" \
        org.label-schema.build-date="$BUILD_DATE" \
        org.label-schema.name="$IMG_NAME" \
        org.label-schema.description="$IMG_DESC" \
        org.label-schema.vcs-ref="$GIT_COMMIT" \
        org.label-schema.vcs-url="$GIT_URL" \
        org.label-schema.vendor="$VENDOR" \
        org.label-schema.version="$BUILD_VERSION" \
        org.label-schema.schema-version="$IMG_VERSION" \
        org.opencontainers.image.authors="$AUTHOR" \
        org.opencontainers.image.description="$IMG_DESC" \
        org.opencontainers.image.vendor="$VENDOR" \
        org.opencontainers.image.version="$IMG_VERSION" \
        org.opencontainers.image.revision="$GIT_COMMIT" \
        org.opencontainers.image.created="$BUILD_DATE" \
        fr.hbis.docker.base.build-date="$BUILD_DATE" \
        fr.hbis.docker.base.name="$IMG_NAME" \
        fr.hbis.docker.base.vendor="$VENDOR" \
        fr.hbis.docker.base.version="$BUILD_VERSION"

RUN apk add --no-cache \
        build-base \
        git \
        curl \
        make \
        bash; sync

## Install Golang with GVM ##
RUN curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash
RUN gvm install go${GOLANG_VERSION}; sync
RUN gvm use go${GOLANG_VERSION} --default; sync

WORKDIR /go

COPY --from=builder /go/src/app/goapp /go/goapp
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=devopscorner/k8s-context:latest /usr/local/bin/k8s-context /usr/local/bin/k8s-context
COPY --from=devopscorner/k8s-context:latest /usr/local/bin/kc /usr/local/bin/kc

COPY src /go/src
COPY _infra /go/src
COPY .aws /go/src
COPY src/.env.example /go/.env
COPY entrypoint.sh /go/entrypoint

ENTRYPOINT ["/go/goapp"]
EXPOSE 80 443 8080
