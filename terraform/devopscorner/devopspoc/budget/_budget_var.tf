# ==========================================================================
#  Budget: _budget_var.tf
# --------------------------------------------------------------------------
#  Description:
#    Budget Variable
# --------------------------------------------------------------------------
#    - Monthly Forcasted Name
#    - Monthly Forcasted Type
#    - Monthly Forcasted Limit Amount
#    - Monthly Forcasted Limit Unit
#    - Monthly Forcasted Time Unit
#    - Monthly Forcasted Time Period Start
#    - Monthly Budget Name
#    - Monthly Budget Type
#    - Monthly Budget Limit Amount
#    - Monthly Budget Limit Unit
#    - Monthly Budget Time Unit
#    - Monthly Budget Time Period Start
# ==========================================================================

# --------------------------------------------------------------------------
#  Budget
# --------------------------------------------------------------------------
###  Monthly Forcasted ###
variable "monthly_forcasted_name" {
  description = "Monthly Forcasted Name"
  type        = string
  default     = "monthly_forcasted_100"
}

variable "monthly_forcasted_type" {
  description = "Monthly Forcasted Type"
  type        = string
  default     = "COST"
}

variable "monthly_forcasted_limit_amount" {
  description = "Monthly Forcasted Limit Amount"
  type        = string
  default     = "100"
}

variable "monthly_forcasted_limit_unit" {
  description = "Monthly Forcasted Limit Unit"
  type        = string
  default     = "USD"
}

variable "monthly_forcasted_time_unit" {
  description = "Monthly Forcasted Time Unit"
  type        = string
  default     = "MONTHLY"
}

variable "monthly_forcasted_time_period_start" {
  description = "Monthly Forcasted Time Period Start"
  type        = string
  default     = "2023-01-01_00:00"
}

variable "monthly_forcasted_notification_comparison" {
  description = "Monthly Forcasted Notification Comparison"
  type        = string
  default     = "GREATER_THAN"
}

variable "monthly_forcasted_notification_threshold" {
  description = "Monthly Forcasted Notification Threshold"
  type        = number
  default     = 80
}

variable "monthly_forcasted_notification_threshold_type" {
  description = "Monthly Forcasted Notification Threshold Type"
  type        = string
  default     = "PERCENTAGE"
}

variable "monthly_forcasted_notification_type" {
  description = "Monthly Forcasted Notification Type"
  type        = string
  default     = "FORECASTED"
}

variable "monthly_forcasted_notification_subscriber" {
  description = "Monthly Forcasted Notification Subscriber Email Address"
  type        = string
  default     = "support@devopscorner.id"
}

###  Monthly Budget Billing ###
variable "monthly_budget_name" {
  description = "Monthly Budget Name"
  type        = string
  default     = "monthly_budget_150"
}

variable "monthly_budget_type" {
  description = "Monthly Budget Type"
  type        = string
  default     = "COST"
}

variable "monthly_budget_limit_amount" {
  description = "Monthly Budget Limit Amount"
  type        = string
  default     = "150"
}

variable "monthly_budget_limit_unit" {
  description = "Monthly Budget Limit Unit"
  type        = string
  default     = "USD"
}

variable "monthly_budget_time_unit" {
  description = "Monthly Budget Time Unit"
  type        = string
  default     = "MONTHLY"
}

variable "monthly_budget_time_period_start" {
  description = "Monthly Budget Time Period Start"
  type        = string
  default     = "2023-01-01_00:00"
}

variable "monthly_budget_notification_comparison" {
  description = "Monthly Budget Notification Comparison"
  type        = string
  default     = "GREATER_THAN"
}

variable "monthly_budget_notification_threshold" {
  description = "Monthly Budget Notification Threshold"
  type        = number
  default     = 80
}

variable "monthly_budget_notification_threshold_type" {
  description = "Monthly Budget Notification Threshold Type"
  type        = string
  default     = "PERCENTAGE"
}

variable "monthly_budget_notification_type" {
  description = "Monthly Budget Notification Type"
  type        = string
  default     = "ACTUAL"
}

variable "monthly_budget_notification_subscriber" {
  description = "Monthly Budget Notification Subscriber Email Address"
  type        = string
  default     = "support@devopscorner.id"
}