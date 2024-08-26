# Terraform Module TFState (Terraform Cloud TFState)

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 4.50.0, < 5.7 |
| <a name="requirement_random"></a> [random](#requirement\_random) | >= 2.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 4.50.0, < 5.7 |
| <a name="provider_random"></a> [random](#provider\_random) | >= 2.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_dynamodb"></a> [dynamodb](#module\_dynamodb) | github.com/terraform-aws-modules/terraform-aws-dynamodb-table.git | v3.3.0 |
| <a name="module_s3_bucket"></a> [s3\_bucket](#module\_s3\_bucket) | github.com/terraform-aws-modules/terraform-aws-s3-bucket.git | v3.14.0 |

## Resources

| Name | Type |
|------|------|
| [aws_iam_role.this](https://registry.terraform.io/providers/opentofu/aws/latest/docs/resources/iam_role) | resource |
| [random_pet.this](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/pet) | resource |
| [aws_canonical_user_id.current](https://registry.terraform.io/providers/opentofu/aws/latest/docs/data-sources/canonical_user_id) | data source |
| [aws_cloudfront_log_delivery_canonical_user_id.cloudfront](https://registry.terraform.io/providers/opentofu/aws/latest/docs/data-sources/cloudfront_log_delivery_canonical_user_id) | data source |
| [aws_iam_policy_document.bucket_policy](https://registry.terraform.io/providers/opentofu/aws/latest/docs/data-sources/iam_policy_document) | data source |
| [aws_kms_key.cmk_key](https://registry.terraform.io/providers/opentofu/aws/latest/docs/data-sources/kms_key) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_aws_account_id"></a> [aws\_account\_id](#input\_aws\_account\_id) | The AWS Account ID to deploy the ACM certificate in | `string` | n/a | yes |
| <a name="input_aws_account_profile"></a> [aws\_account\_profile](#input\_aws\_account\_profile) | The AWS Profile to deploy the ACM certificate in | `string` | n/a | yes |
| <a name="input_aws_region"></a> [aws\_region](#input\_aws\_region) | The AWS region to deploy the ACM certificate in | `string` | n/a | yes |
| <a name="input_department"></a> [department](#input\_department) | Department Owner | `string` | `"DEVOPS"` | no |
| <a name="input_environment"></a> [environment](#input\_environment) | Target Environment (tags) | `map(string)` | <pre>{<br>  "default": "DEF",<br>  "lab": "RND",<br>  "prod": "PROD",<br>  "staging": "STG"<br>}</pre> | no |
| <a name="input_kms_env"></a> [kms\_env](#input\_kms\_env) | KMS Key Environment | `map(string)` | <pre>{<br>  "lab": "RnD",<br>  "prod": "Production",<br>  "staging": "Staging"<br>}</pre> | no |
| <a name="input_kms_key"></a> [kms\_key](#input\_kms\_key) | KMS Key References | `map(string)` | <pre>{<br>  "default": "arn:aws:kms:aws_region:YOUR_AWS_ACCOUNT:key/CMK_KEY_HASH",<br>  "lab": "arn:aws:kms:aws_region:YOUR_AWS_ACCOUNT:key/CMK_KEY_HASH",<br>  "prod": "arn:aws:kms:aws_region:YOUR_AWS_ACCOUNT:key/CMK_KEY_HASH",<br>  "staging": "arn:aws:kms:aws_region:YOUR_AWS_ACCOUNT:key/CMK_KEY_HASH"<br>}</pre> | no |
| <a name="input_tfstate_bucket"></a> [tfstate\_bucket](#input\_tfstate\_bucket) | Name of bucket to store tfstate | `string` | `"devopscorner-tf-remote-state"` | no |
| <a name="input_tfstate_dynamodb_table"></a> [tfstate\_dynamodb\_table](#input\_tfstate\_dynamodb\_table) | Name of dynamodb table to store tfstate | `string` | `"devopscorner-tf-state-lock"` | no |
| <a name="input_tfstate_encrypt"></a> [tfstate\_encrypt](#input\_tfstate\_encrypt) | Name of bucket to store tfstate | `bool` | `true` | no |
| <a name="input_tfstate_path"></a> [tfstate\_path](#input\_tfstate\_path) | Path .tfstate in Bucket | `string` | `"tfstate/terraform.tfstate"` | no |
| <a name="input_workspace_env"></a> [workspace\_env](#input\_workspace\_env) | Workspace Environment Selection | `map(string)` | <pre>{<br>  "default": "default",<br>  "lab": "rnd",<br>  "prod": "prod",<br>  "staging": "staging"<br>}</pre> | no |
| <a name="input_workspace_name"></a> [workspace\_name](#input\_workspace\_name) | Workspace Environment Name | `string` | `"default"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_dynamodb_table_arn"></a> [dynamodb\_table\_arn](#output\_dynamodb\_table\_arn) | ARN of the DynamoDB table |
| <a name="output_dynamodb_table_id"></a> [dynamodb\_table\_id](#output\_dynamodb\_table\_id) | ID of the DynamoDB table |
| <a name="output_dynamodb_table_stream_arn"></a> [dynamodb\_table\_stream\_arn](#output\_dynamodb\_table\_stream\_arn) | The ARN of the Table Stream. Only available when var.stream\_enabled is true |
| <a name="output_dynamodb_table_stream_label"></a> [dynamodb\_table\_stream\_label](#output\_dynamodb\_table\_stream\_label) | A timestamp, in ISO 8601 format of the Table Stream. Only available when var.stream\_enabled is true |
| <a name="output_s3_bucket_arn"></a> [s3\_bucket\_arn](#output\_s3\_bucket\_arn) | The ARN of the bucket. Will be of format arn:aws:s3:::bucketname. |
| <a name="output_s3_bucket_id"></a> [s3\_bucket\_id](#output\_s3\_bucket\_id) | The name of the bucket. |
| <a name="output_s3_bucket_region"></a> [s3\_bucket\_region](#output\_s3\_bucket\_region) | The AWS region this bucket resides in. |
<!-- END_TF_DOCS -->