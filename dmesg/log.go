package dmesg

import (
	"io"
)

/*
[LOG_EMERG]   = { "emerg", N_("system is unusable") },
  [LOG_ALERT]   = { "alert", N_("action must be taken immediately") },
  [LOG_CRIT]    = { "crit",  N_("critical conditions") },
  [LOG_ERR]     = { "err",   N_("error conditions") },
  [LOG_WARNING] = { "warn",  N_("warning conditions") },
  [LOG_NOTICE]  = { "notice",N_("normal but significant condition") },
  [LOG_INFO]    = { "info",  N_("informational") },
  [LOG_DEBUG]   = { "debug", N_("debug-level messages") }
*/
type Level int

const (
	LOG_EMERG   Level = iota
	LOG_ALERT   Level = iota
	LOG_CRIT    Level = iota
	LOG_ERR     Level = iota
	LOG_WARNING Level = iota
	LOG_NOTICE  Level = iota
	LOG_INFO    Level = iota
	LOG_DEBUG   Level = iota
)

func (l Level) String() string {
	switch l {
	case LOG_EMERG:
		return "emerg"
	case LOG_ALERT:
		return "alert"
	case LOG_CRIT:
		return "crit"
	case LOG_ERR:
		return "err"
	case LOG_WARNING:
		return "warn"
	case LOG_NOTICE:
		return "notice"
	case LOG_INFO:
		return "info"
	case LOG_DEBUG:
		return "debug"
	default:
		return "unknown"
	}
}

/*
[FAC_BASE(LOG_KERN)]     = { "kern",     N_("kernel messages") },
  [FAC_BASE(LOG_USER)]     = { "user",     N_("random user-level messages") },
  [FAC_BASE(LOG_MAIL)]     = { "mail",     N_("mail system") },
  [FAC_BASE(LOG_DAEMON)]   = { "daemon",   N_("system daemons") },
  [FAC_BASE(LOG_AUTH)]     = { "auth",     N_("security/authorization messages") },
  [FAC_BASE(LOG_SYSLOG)]   = { "syslog",   N_("messages generated internally by syslogd") },
  [FAC_BASE(LOG_LPR)]      = { "lpr",      N_("line printer subsystem") },
  [FAC_BASE(LOG_NEWS)]     = { "news",     N_("network news subsystem") },
  [FAC_BASE(LOG_UUCP)]     = { "uucp",     N_("UUCP subsystem") },
  [FAC_BASE(LOG_CRON)]     = { "cron",     N_("clock daemon") },
  [FAC_BASE(LOG_AUTHPRIV)] = { "authpriv", N_("security/authorization messages (private)") },
  [FAC_BASE(LOG_FTP)]      = { "ftp",      N_("FTP daemon") },
*/
type Facility int

const (
	LOG_KERN Facility = iota
	LOG_USER
	LOG_MAIL
	LOG_DAEMON
	LOG_AUTH
	LOG_SYSLOG
	LOG_LPR
	LOG_NEWS
	LOG_UUCP
	LOG_CRON
	LOG_AUTHPRIV
	LOG_FTP
)

func (f Facility) String() string {
	switch f {
	case LOG_KERN:
		return "kern"
	case LOG_USER:
		return "user"
	case LOG_MAIL:
		return "mail"
	case LOG_DAEMON:
		return "daemon"
	case LOG_AUTH:
		return "auth"
	case LOG_SYSLOG:
		return "syslog"
	case LOG_LPR:
		return "lpr"
	case LOG_NEWS:
		return "news"
	case LOG_UUCP:
		return "uucp"
	case LOG_CRON:
		return "cron"
	case LOG_AUTHPRIV:
		return "authpriv"
	case LOG_FTP:
		return "ftp"
	default:
		return "unknown"
	}
}

type Timestamp int64

type LogEntry struct {
	Level     Level
	Facility  Facility
	SeqNum    int
	Timestamp Timestamp
	Message   string
	Tags      map[string]string
}

type LogReader interface {
	Read() (*LogEntry, error)
}

type LogReadCloser interface {
	LogReader
	io.Closer
}
