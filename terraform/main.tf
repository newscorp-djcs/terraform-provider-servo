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

variable "REGION" {
  type = string
  default = "virginia"
}

variable "ORG" {
  type = string
  default = "dev"
}

variable "SERVO_TOKEN" {}

resource "servo_app" "test" {
  region = var.REGION
  org = var.ORG
  app = {
    handle = "terraform-provider-test-fred"
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

 
