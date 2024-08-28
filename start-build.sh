#!/usr/bin/env sh
# -----------------------------------------------------------------------------
#  Docker Build Container (Linux/AMD64)
# -----------------------------------------------------------------------------
#  Author     : Dwi Fahni Denni
#  License    : Apache v2
# -----------------------------------------------------------------------------
set -e

export AWS_ACCOUNT_ID=$1
export AWS_DEFAULT_REGION="us-west-2"
export CI_REGISTRY="$AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com"

export CI_PROJECT_PATH="devopscorner"
export CI_PROJECT_NAME="golang-bedrock"

export IMAGE="$CI_PROJECT_PATH/$CI_PROJECT_NAME"
export ECR_IMAGE="$CI_REGISTRY/$CI_PROJECT_PATH/$CI_PROJECT_NAME"

export TARGETPLATFORM="linux/amd64"
export ARCH="amd64"

export VERSION="1.1.4"

build_golang_bedrock() {
    TAGS="3.20 \
        alpine-3.20 \
        alpine-latest \
        alpine \
        latest \
        $VERSION "

    ### DockerHub Build Images ###
    for TAG in $TAGS; do
        echo " Build Image => $IMAGE:$TAG"
        docker build \
            -f Dockerfile-$ARCH \
            -t $IMAGE:$TAG .
        echo ''
    done

    ### ECR Build Images ###
    for TAG in $TAGS; do
        echo " Build Image => $ECR_IMAGE:$TAG"
        docker build \
            -f Dockerfile-$ARCH \
            -t $ECR_IMAGE:$TAG .
        echo ''
    done
}

docker_build() {
    build_golang_bedrock
}

docker_clean() {
    echo "Cleanup Unknown Tags"
    echo "docker images -a | grep none | awk '{ print $3; }' | xargs docker rmi"
    docker images -a | grep none | awk '{ print $3; }' | xargs docker rmi
    echo ''
}

main() {
    docker_build
    docker_clean
    echo ''
    echo '-- ALL DONE --'
}

### START HERE ###
main