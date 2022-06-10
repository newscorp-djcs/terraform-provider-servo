terraform {
  required_providers {
    servo = {
      version = "0.1"
      source  = "local/poc/servo"
    }
  }
}

provider "servo" {
  token = var.SERVO_TOKEN
}

module "app" {
  source = "./app"
  region = "virginia"
  org    = "dev"
}

output "app_handle" {
  value = module.app.app_handle
}
