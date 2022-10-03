job "docs" {
  datacenters = ["dc1"]

  group "example" {
    network {
      port "http" {
        static = "8090"
      }
    }
    task "server" {
      driver = "docker"

      config {
        image = "misnaged/microservice"
        ports = ["http"]
      }
    }
  }
}

