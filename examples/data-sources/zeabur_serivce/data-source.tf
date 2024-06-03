data "zeabur_user" "me" {}

data "zeabur_project" "example" {
  name  = "<service-name>"
  owner = data.zeabur_user.me.username
}