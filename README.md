## tfdev

A tool to make Terraform Provider development easier on Terraform 0.14+ - by templating the `providers.tf` to allow easier toggling between Development and Production binaries.

### Usage

```
$ tfdev [PROVIDER] [MODE]
```

* Provider being the name of the Provider (e.g. `azurerm`, `aws`)
* Mode being either `dev` (to use the dev override) or `prod` (to use the prod value)

### Setup

Run:

```
go install github.com/tombuildsstuff/tfdev
```

Assuming the following `~/.terraformrc`:

```
provider_installation {
   dev_overrides {
     "tombuildsstuff/azurerm" = "/Users/tharvey/code/bin"
   }

   # For all other providers, install them directly from their origin provider
   # registries as normal. If you omit this, Terraform will _only_ use
   # the dev_overrides block, and so no other providers will be available.
   direct {}
}
```

Add a `.tfdev.hcl` file to your home directory:

```
$ cat ~/.tfdev.hcl
provider "azurerm" {
  dev  = "tombuildsstuff/azurerm"
  prod = "hashicorp/azurerm"
}
```

At which point running `tfdev azurerm dev` or `tfdev azurerm prod` will output a `providers.tf` containing the required Provider configuration:

```
terraform {
  required_providers {
    azurerm = {
      source = "tombuildsstuff/azurerm"
    }
  }
}
```

### Future Enhancements

- Looking up the latest version of a given release and pinning to that, e.g. `1.x`

### Licence

Apache 2
