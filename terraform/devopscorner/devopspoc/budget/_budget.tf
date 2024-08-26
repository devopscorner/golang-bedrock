# ==========================================================================
#  Budget: _budget.tf
# --------------------------------------------------------------------------
#  Description:
#    Budget Setup
# --------------------------------------------------------------------------
#    - Name Budget
#    - Limit Budget (Number)
#    - Limit Unit (USD)
#    - Time Budget
#    - Alert Notification (Threshold)
# ==========================================================================

# --------------------------------------------------------------------------
#  Resources Tags
# --------------------------------------------------------------------------
locals {
  resources_tags = {
    Name          = "budget-${var.workspace_env[local.env]}",
    ResourceGroup = "${var.environment[local.env]}-BUD-INFRA"
  }
}

resource "aws_budgets_budget" "monthly_forcasted" {
  name              = var.monthly_forcasted_name
  budget_type       = var.monthly_forcasted_type
  limit_amount      = var.monthly_forcasted_limit_amount
  limit_unit        = var.monthly_forcasted_limit_unit
  time_unit         = var.monthly_forcasted_time_unit
  time_period_start = var.monthly_forcasted_time_period_start

  notification {
    comparison_operator        = var.monthly_forcasted_notification_comparison
    threshold                  = var.monthly_forcasted_notification_threshold
    threshold_type             = var.monthly_forcasted_notification_threshold_type
    notification_type          = var.monthly_forcasted_notification_type
    subscriber_email_addresses = ["${var.monthly_forcasted_notification_subscriber}"]
  }

}

resource "aws_budgets_budget" "monthly_budget" {
  name              = var.monthly_budget_name
  budget_type       = var.monthly_budget_type
  limit_amount      = var.monthly_budget_limit_amount
  limit_unit        = var.monthly_budget_limit_unit
  time_unit         = var.monthly_budget_time_unit
  time_period_start = var.monthly_budget_time_period_start

  notification {
    comparison_operator        = var.monthly_budget_notification_comparison
    threshold                  = var.monthly_budget_notification_threshold
    threshold_type             = var.monthly_budget_notification_threshold_type
    notification_type          = var.monthly_budget_notification_type
    subscriber_email_addresses = ["${var.monthly_budget_notification_subscriber}"]
  }

}
