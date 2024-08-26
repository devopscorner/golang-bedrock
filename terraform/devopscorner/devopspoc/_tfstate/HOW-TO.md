# Terraform Module TFState (Terraform Cloud TFState)

## How-to-Deploy

- Terraform Initialize

  ```
  tofu init
  ```

- List Existing Workspace

  ```
  tofu workspace list
  ```

- Create Workspace

  ```
  tofu workspace new [environment]
  ---
  eg:
  tofu workspace new lab
  tofu workspace new staging
  tofu workspace new prod
  ```

- Use Workspace

  ```
  tofu workspace select [environment]
  ---
  eg:
  tofu workspace select lab
  tofu workspace select staging
  tofu workspace select prod
  ```

- Terraform Planning

  ```
  tofu plan
  ```

- Terraform Provisioning

  ```
  tofu apply
  ```

## Migrate State

- Rename Backend

  ```
  mv backend.tf.example backend.tf
  ```

- Initiate Migrate

  ```
  tofu init --migrate-state
  ```

## Cleanup Environment

```
tofu destroy
```

## Copyright

- Author: **Dwi Fahni Denni (@zeroc0d3)**
- Vendor: **DevOps Corner Indonesia (devopscorner.id)**
- License: **Apache v2**
