import re

ICAL_DATETIME_START = "DTSTART"
ICAL_DATETIME_END = "DTEND"
ICAL_DATETIME_STAMP = "DTSTAMP"
ICAL_SUMMARY = "SUMMARY"
ICAL_DESCRIPTION = "DESCRIPTION"
ICAL_ATTENDEE = "ATTENDEE"
ICAL_UID = "UID"
ICAL_SEQUENCE = "SEQUENCE"
ICAL_RECURRENCE_ID = "RECURRENCE-ID"
ICAL_RRULE = "RRULE"
ICAL_UNTIL = "UNTIL"
ICAL_LAST_MODIFIED = "LAST-MODIFIED"
ICAL_LOCATION = "LOCATION"
ICAL_PRIORITY = "PRIORITY"
ICAL_STATUS = "STATUS"
ICAL_STATUS_CANCELLED = "CANCELLED"
ICAL_COMPONENT_VEVENT = "VEVENT"
RE_PRIORITY = re.compile(r"^\[L(\d+)\]")
RE_EVENT_UID_EXPORT = re.compile(r"([\w\d]+)-(\d+)-([\w\d]+)")
RE_EVENT_UID_V1 = re.compile(r"amixr-([\w\d-]+)-U(\d+)-E(\d+)-S(\d+)")
RE_EVENT_UID_V2 = re.compile(r"oncall-([\w\d-]+)-PK([\w\d]+)-U(\d+)-E(\d+)-S(\d+)")

CALENDAR_TYPE_FINAL = "final"

EXPORT_WINDOW_DAYS_AFTER = 180
EXPORT_WINDOW_DAYS_BEFORE = 15

SCHEDULE_ONCALL_CACHE_KEY_PREFIX = "schedule_oncall_users_"
SCHEDULE_ONCALL_CACHE_TTL = 15 * 60  # 15 minutes in seconds

PREFETCHED_SHIFT_SWAPS = "prefetched_shift_swaps"