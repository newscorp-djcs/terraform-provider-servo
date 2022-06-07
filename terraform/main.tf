terraform {
  required_providers {
    servo = {
      version = "0.1"
      source  = "local/poc/servo"
    }
  }
}

provider "servo" {
  # token = var.SERVO_TOKEN
}

module "apps" {
  source = "./app"

  app_handle = "admin-djcss"
}

output "apps" {
  # value = module.apps.source
  value = module.apps.all_apps
}

variable "SERVO_TOKEN" {}
