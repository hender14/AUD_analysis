variable "project_name" {
  default = {
    tf-keyword = "gold-cycling-307817"
  }
}

variable "credential" {
  default = {
    data = "../app/src/credentials/key.json"
  }
}

# variable "webhook" {
#   default = {
#     url = "<your-webhook-url>"
#   }
# }