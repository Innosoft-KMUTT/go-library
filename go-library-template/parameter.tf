locals {
  workspace_path     = "go-library"
  workspace_git_repo = "https://github.com/Innosoft-KMUTT/go-library.git"
}

data "coder_parameter" "workspace_git_branch" {
  name        = "Workspace Git Branch"
  default     = "main"
  description = "Git branch for Workspace"
  type        = "string"
  mutable     = true
}

data "coder_parameter" "workspace_git_clone" {
  name        = "Workspace Git Clone Command"
  default     = <<-EOT
    if [ ! -d ~/${local.workspace_path} ]; then
      git clone -b ${data.coder_parameter.workspace_git_branch.value} --recurse-submodules ${local.workspace_git_repo} ~/${local.workspace_path}
      cd ~/${local.workspace_path}
      ./git_all.sh git switch dev && ./git_all.sh git pull --rebase 
    fi
  EOT
  description = "Command to clone git"
  type        = "string"
  mutable     = true
}

data "coder_parameter" "frontend_run_command" {
  name        = "Frontend Run Command"
  default     = ""
  description = "Command to run when start container"
  type        = "string"
  mutable     = true
}

data "coder_parameter" "backend_run_command" {
  name        = "Backend Run Command"
  default     = ""
  description = "Command to run when start container"
  type        = "string"
  mutable     = true
}

data "coder_parameter" "nginx_external_port" {
  name        = "Nginx External Port"
  default     = "8081"
  description = "Port that exponse on host"
  type        = "number"
  mutable     = true
}
