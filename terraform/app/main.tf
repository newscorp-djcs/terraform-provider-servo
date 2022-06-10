terraform {
  required_providers {
    servo = {
      version = "0.1"
      source  = "local/poc/servo"
    }
  }
}

resource "servo_app" "test" {
  region = var.region
  org    = var.org
  handle = "terraform-provider-test"
  source = "https://github.dowjones.net/servo3/example"
}

output "app_handle" {
  value = servo_app.test.app.handle
}
