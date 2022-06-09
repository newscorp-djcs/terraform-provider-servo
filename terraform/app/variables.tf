variable "region" {
    description = "Region where the app will be created, options: virginia, oregon"
  type = string
  default = "virginia"
}

variable "org" {
    description = "Org where the app will be created, options: dev, prod"
  type = string
  default = "dev"
}
