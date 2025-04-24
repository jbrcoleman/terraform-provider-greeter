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

resource "greeter_greeting" "test" {
  name = "World"
}

data "greeter_greeting" "test" {
  name = "Data"
}

output "resource_greeting" {
  value = greeter_greeting.test.greeting
}

output "data_greeting" {
  value = data.greeter_greeting.test.greeting
}