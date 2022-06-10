resource "servo_app" "example" {
  org    = var.org
  region = var.region
  handle = "example_app"
  source = "https://github.dowjones.net/servo3/example"
}

data "servo_app" "example1" {
  context = "servo.dev.virginia"
  handle  = "example_app"
}

data "servo_app" "example2" {
  org    = "dev"
  region = "virginia"
  handle = "example_app"
}
