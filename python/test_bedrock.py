import boto3
import json
import os
from botocore.exceptions import ClientError

# Set up AWS credentials and region
aws_access_key = os.getenv('AWS_ACCESS_KEY_ID')
aws_secret_key = os.getenv('AWS_SECRET_ACCESS_KEY')
aws_region = os.getenv('AWS_REGION', 'us-west-2')  # Replace with your region if different
model_id = os.getenv('AMAZON_BEDROCK_MODEL_ID', 'anthropic.claude-3-sonnet-20240229-v1:0')

# Initialize Bedrock clients
bedrock_runtime = boto3.client(
    service_name='bedrock-runtime',
    region_name=aws_region,
    aws_access_key_id=aws_access_key,
    aws_secret_access_key=aws_secret_key
)

bedrock = boto3.client(
    service_name='bedrock',
    region_name=aws_region,
    aws_access_key_id=aws_access_key,
    aws_secret_access_key=aws_secret_key
)

def list_models():
    try:
        response = bedrock.list_foundation_models()
        print("\nAvailable Bedrock models:")
        for model in response['modelSummaries']:
            print(f"- {model['modelId']}")
    except ClientError as e:
        print(f"Error listing models: {e}")

def invoke_model(model_id):
    try:
        body = json.dumps({
            "anthropic_version": "bedrock-2023-05-31",
            "max_tokens": 100,
            "messages": [
                {
                    "role": "user",
                    "content": "Hello, Claude. Please respond with 'Hello, Human!'"
                }
            ]
        })

        response = bedrock_runtime.invoke_model_with_response_stream(
            modelId=model_id,
            contentType='application/json',
            accept='application/json',
            body=body
        )

        for event in response['body']:
            chunk = json.loads(event['chunk']['bytes'].decode())
            if chunk['type'] == 'content_block_delta':
                print(chunk['delta']['text'], end='')
            elif chunk['type'] == 'message_stop':
                print()  # New line at the end of the response

    except ClientError as e:
        print(f"Error invoking model: {e}")

if __name__ == "__main__":
    print(f"Testing Bedrock in region: {aws_region}")
    print(f"Using model ID: {model_id}")

    list_models()
    print("\nAttempting to invoke the model:")
    invoke_model(model_id)