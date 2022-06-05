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

data "servo_apps" "all" {}

# Returns all apps
output "all_apps" {
  value = data.servo_apps.all
}

# Only returns context app admin-djcs
# output "app" {
#   value = { 
#     for app in data.servo_apps.all :
#     app.context => app
#     # if app.handle == var.app_handle
#   }
# }
