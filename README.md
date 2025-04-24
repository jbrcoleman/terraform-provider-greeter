# Terraform Provider Greeter

This is a sample Terraform provider for educational purposes. It demonstrates how to build a custom Terraform provider from scratch.

## Features

- Manages "greeting" resources
- Provides a "greeting" data source
- Configurable greeting suffix

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

## Building The Provider

1. Clone the repository
```sh
git clone https://github.com/jbrcoleman/terraform-provider-greeter.git
cd terraform-provider-greeter
```

2. Build the provider
```sh
go build -o terraform-provider-greeter
```

## Local Development Setup

To use your locally built provider, configure Terraform to use it instead of downloading from the registry.

Create a `.terraformrc` file in your home directory (or `terraform.rc` in `%APPDATA%` on Windows):

```hcl
provider_installation {
  dev_overrides {
    "yourusername/greeter" = "/path/to/your/terraform-provider-greeter"
  }
  direct {}
}
```

Replace `/path/to/your/terraform-provider-greeter` with the absolute path to the directory containing your built provider binary.

## Using the Provider

```hcl
terraform {
  required_providers {
    greeter = {
      source = "yourusername/greeter"
      version = "0.1.0"
    }
  }
}

provider "greeter" {
  greeting_suffix = "!"
}

resource "greeter_greeting" "example" {
  name = "World"
}

output "greeting" {
  value = greeter_greeting.example.greeting
}
```

## Testing

To run the provider tests:

```sh
go test -v ./internal/provider
```

To run the acceptance tests:

```sh
TF_ACC=1 go test -v ./internal/provider
```

## Examples

See the `examples/` directory for example configurations.

## Project Structure

```
terraform-provider-greeter/
├── examples/                  # Example usage of our provider
│   ├── resources/             # Examples for resources
│   │   └── greeter/           # Examples for greeter resource
│   │       └── resource.tf    # Example resource configuration
│   └── data-sources/          # Examples for data sources
│       └── greeter/           # Examples for greeter data source
│           └── data-source.tf # Example data source configuration
├── go.mod                     # Go module definition
├── go.sum                     # Go module dependencies
├── internal/                  # Internal provider code
│   └── provider/              # Provider implementation
│       ├── greeter_resource.go        # Implementation of the greeter resource
│       ├── greeter_data_source.go     # Implementation of the greeter data source
│       ├── greeter_resource_test.go   # Tests for the greeter resource
│       └── provider.go                # Main provider implementation
├── test/                      # Test configurations
│   └── main.tf                # Main test configuration
└── main.go                    # Provider entrypoint
```

## Next Steps for Learning

After exploring this simple provider, you might want to:

1. Add more complex resources with nested schema attributes
2. Implement a real API client to manage external resources
3. Add more comprehensive validation logic
4. Implement import functionality for existing resources
5. Add more detailed documentation and examples
6. Learn about publishing to the Terraform Registry