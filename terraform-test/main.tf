terraform {
  required_providers {
    greeter = {
      source = "jbrcoleman/greeter"
      version = "0.1.0"
    }
  }
}

provider "greeter" {
  greeting_suffix = "!!!"
}

resource "greeter_greeting" "example" {
  name = "World"
}

data "greeter_greeting" "example" {
  name = "Data"
}

output "resource_greeting" {
  value = greeter_greeting.example.greeting
}

output "data_greeting" {
  value = data.greeter_greeting.example.greeting
}