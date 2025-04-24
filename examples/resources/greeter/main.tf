provider "greeter" {
  greeting_suffix = "!!!"
}

resource "greeter_greeting" "example" {
  name = "Terraform"
}

output "greeting_message" {
  value = resource.greeter_greeting.example.greeting
}