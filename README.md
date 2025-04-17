# Terraform Provider Version Lifecycle

This Terraform provider manages application version lifecycle.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider

```bash
go build -o terraform-provider-versionlifecycle
```

## Using the Provider

```hcl
terraform {
  required_providers {
    versionlifecycle = {
      source = "jamieedwards/versionlifecycle"
    }
  }
}

provider "versionlifecycle" {}

resource "versionlifecycle_version" "example" {
  app_version        = "1.0.0"
  default_app_version = "1.0.0"
}
```

## Argument Reference

The following arguments are supported:

- `app_version` - (Required) The current application version
- `default_app_version` - (Required) The default application version to fall back to

## Attributes Reference

The following attributes are exported:

- `current_app_version` - The current active application version 