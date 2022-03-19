terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.13.0"
    }
  }
}

provider "docker" {}


resource "docker_image" "traefik" {
  name         = "traefik:latest"
  keep_locally = false
}

resource "docker_container" "traefik" {
  image = docker_image.traefik.latest
  name  = "traefik"
  ports {
    internal = var.ports
    external = var.ports
  }
}

resource "docker_image" "consul" {
  name         = "consul:latest"
  keep_locally = false
}

resource "docker_container" "consul" {
  image = docker_image.consul.latest
  name  = "consul"
  ports {
    internal = var.cports
    external = var.cports
  }
}

resource "docker_image" "influxdb"{
    name = "influxdb:latest"
    keep_locally = false
}

resource "docker_container" "influxdb"{
    image =  docker_image.influxdb.latest
    name  = "influxdb"
    ports{
        internal = var.influx
        external = var.influx
    }
}


