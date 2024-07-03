terraform {
  required_providers {
    coder = {
      source = "coder/coder"
    }
    docker = {
      source = "kreuzwerker/docker"
    }
  }
}

locals {
  username = data.coder_workspace.me.owner
}

data "coder_provisioner" "me" {
}

provider "coder" {
}

data "coder_external_auth" "github" {
  # Matches the ID of the external auth provider in Coder.
  id = "github"
}

provider "docker" {
}

data "coder_workspace" "me" {
}

resource "docker_image" "dev_container" {
  name = "coder-${data.coder_workspace.me.id}"
  build {
    context = "./build"
    build_args = {
      USER = local.username
    }
  }
  triggers = {
    dir_sha1 = sha1(join("", [for f in fileset(path.module, "build/*") : filesha1(f)]))
  }
}

resource "docker_network" "private_network" {
  name = "network-${data.coder_workspace.me.id}"
}


