
// traefik ports 
variable "ports" {
  type    = string
  default = "8080"
}

// consul ports
variable "cports" {
  type    = string
  default = "8400"
}


// explorer ports
variable "eports" {
  type    = string
  default = "3000"
}


// influx ports
variable "influx"{
    type = string
    default = "8089"
}



