terraform {
  required_providers {
    servo = {
      version = "0.1"
      source  = "local/poc/servo"
    }
  }
}

provider "servo" {
  # token = ""
}

# module "apps" {
#   source = "./app"

#   app_handle = "admin-djcss"
# }

# output "apps" {
#   value = module.apps.all_apps
# }

resource "servo_app" "test" {
  app {
    handle = "provider-test"
    # source = "https://github.com/test/myapp.git"
  }

}

# output "edu_order" {
#   value = servo_app.test
# }

