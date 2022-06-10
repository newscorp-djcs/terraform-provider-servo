resource "servo_stack" "dev" {
  app_context = data.servo_app.example.context
  handle      = "dev"
  ticket      = true
}

data "servo_app" "example" {
  org    = "dev"
  region = "virginia"
  handle = "example_app"
}

data "servo_stack" "dev1" {
  context = "servo:dev:virginia.example_app"
  handle  = "dev"
}

data "servo_stack" "dev2" {
  org        = "dev"
  region     = "virginia"
  app_handle = "example_app"
  handle     = "dev"
}
