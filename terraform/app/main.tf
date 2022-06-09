terraform {
  required_providers {
    servo = {
      version = "0.1"
      source  = "local/poc/servo"
    }
  }
}

variable "app_handle" {
  type    = string
  default = "admin-djcss"
}

# data "servo_app" "all" {}

# Returns all apps
# output "all_apps" {
#   value = data.servo_app.all
# }

