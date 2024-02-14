# Reviewing Terraform Plan and Apply Logs in Terraform Cloud UI

This guide provides step-by-step instructions for reviewing Terraform plan and apply logs within the Terraform Cloud UI.

## Step 1: Log in to Terraform Cloud

1. Navigate to the Terraform Cloud login page at [app.terraform.io](https://app.terraform.io/).
2. Enter your credentials to log in, or sign up if you do not have an account.

## Step 2: Access Your Workspace

1. Upon logging in, you'll be directed to the "Workspaces" dashboard.
2. Select the workspace you wish to review logs for by clicking on its name.

## Step 3: Review Runs

1. Inside the workspace, the "Runs" tab displays all historical runs.
2. Each run is listed with a timestamp and a brief description.

## Step 4: Select a Run

1. Locate the run you're interested in by scrolling or using the search function.
2. Click on the run to view its details.

## Step 5: Review the Plan or Apply Log

- **Plan Stage**: Click on the "Plan" section to view the `terraform plan` command output, showing intended infrastructure changes.
- **Apply Stage**: After a plan is applied, click on the "Apply" section to view the `terraform apply` command output, showing actual changes made to the infrastructure.

## Step 6: Analyze the Logs

Scroll through the detailed log output in both the Plan and Apply sections.
When reviewing the Terraform plan output, you'll encounter various symbols preceding the resource names. These symbols indicate the action Terraform intends to take:

- `+` : **Create** - The resource will be created.

- `-` : **Destroy** - The resource will be deleted.

- `~` : **Update in-place** - The resource will be updated without being destroyed and re-created.

- `-/+` : **Destroy and then create** - The resource will be destroyed first and then created anew.

