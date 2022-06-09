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

resource "servo_app" "test" {
  app = {
    handle = "terraform-provider-test"
    source = "https://github.dowjones.net/servo3/example"
  }
}

# module "apps" {
#   source = "./app"

#   app_handle = "admin-djcss"
# }

# output "apps" {
#   # value = module.apps.source
#   value = module.apps.all_apps
# }

# variable "SERVO_TOKEN" {}

 
