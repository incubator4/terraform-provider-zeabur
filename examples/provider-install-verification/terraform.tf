terraform {
  required_providers {
    zeabur = {
      source = "incubator4/zeabur"

    }
  }
}

provider "zeabur" {
  # Configuration options
  #   api_token = "sk-xstgoddeu4576dqi43owqa2w74ilh"
}

# resource "zeabur_project" "test" {
#   name   = "test_project"
#   region = "hkg1"
# }

data "zeabur_project" "example" {
  id = "66549e889348c9bd9d2821a1"
}

output "project" {
  value = data.zeabur_project.example
}
