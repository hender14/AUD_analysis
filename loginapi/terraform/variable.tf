variable "project_name" {
  default = {
    tf-login = "gold-cycling-307817"
  }
}

variable "credential" {
  default = {
    data = "../app/key.json"
  }
}

# variable "webhook" {
#   default = {
#     url = "<your-webhook-url>"
#   }
# }