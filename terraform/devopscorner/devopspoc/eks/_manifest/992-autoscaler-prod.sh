#!/bin/sh

export AWS_REGION="us-west-2"
export ACCOUNT_ID="YOUR_AWS_ACCOUNT"
export EKS_CLUSTER="devopspoc-nonprod"
export EKS_VPC_ID="vpc-0987612345"
export SSL_CERT_ARN="arn:aws:acm:us-west-2:${ACCOUNT_ID}:certificate/HASH_NUMBER"

kubectl config use-context arn:aws:eks:us-west-2:${ACCOUNT_ID}:cluster/${EKS_CLUSTER}

kubectl apply -f  https://raw.githubusercontent.com/kubernetes/autoscaler/master/cluster-autoscaler/cloudprovider/aws/examples/cluster-autoscaler-autodiscover.yaml

kubectl -n kube-system patch deployment cluster-autoscaler --patch \
'{
    "spec": {
        "template": {
            "metadata": {
                "annotations": {
                    "cluster-autoscaler.kubernetes.io/safe-to-evict": "false"
                }
            },
            "spec": {
                "containers": [
                    {
                        "image": "k8s.gcr.io/autoscaling/cluster-autoscaler:v1.29.0",
                        "name": "cluster-autoscaler",
                        "resources": {
                            "limits": {
                                "cpu": "300m",
                                "memory": "600Mi"
                            },
                            "requests": {
                                "cpu": "100m",
                                "memory": "300Mi"
                            }
                        },
                        "env": [
                            {
                                "name": "AWS_STS_REGIONAL_ENDPOINTS",
                                "value": "regional"
                            },
                            {
                                "name": "AWS_DEFAULT_REGION",
                                "value": "us-west-2"
                            },
                            {
                                "name": "AWS_REGION",
                                "value": "us-west-2"
                            },
                            {
                                "name": "AWS_ROLE_ARN",
                                "value": "arn:aws:iam::YOUR_AWS_ACCOUNT:role/cluster-autoscaler-devopscorner-nonprod-role"
                            }
                        ],
                        "command": [
                            "./cluster-autoscaler",
                            "--v=4",
                            "--stderrthreshold=info",
                            "--cloud-provider=aws",
                            "--skip-nodes-with-local-storage=false",
                            "--expander=least-waste",
                            "--node-group-auto-discovery=asg:tag=k8s.io/cluster-autoscaler/enabled,k8s.io/cluster-autoscaler/devopscorner-nonprod",
                            "--balance-similar-node-groups",
                            "--skip-nodes-with-system-pods=false"
                        ]
                    }
                ]
            }
        }
    }
}'

kubectl patch ServiceAccount cluster-autoscaler -n kube-system --patch '{"metadata":{"annotations":{"eks.amazonaws.com/role-arn": "arn:aws:iam::YOUR_AWS_ACCOUNT:role/cluster-autoscaler-devopscorner-nonprod-role"}}}'

kubectl apply -f 09-metrics-server.yaml