terraform {
  required_providers {
    docker = {
      source = "terraform-providers/docker"
    }
  }
}

provider "docker" {}

resource "docker_container" "authorize" {
  image = "authorize:1.0"
  name  = "auth"
  ports {
    internal = 8080
    external = 8080
  }
}
