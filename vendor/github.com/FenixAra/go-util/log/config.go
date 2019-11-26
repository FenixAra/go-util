package log

import uuid "github.com/satori/go.uuid"

type Config struct {
	// The supported log levels are as follows
	// DEBUG < INFO < WARN < ERROR < FATAL
	// If a log level is specified all logs with level below the specified level are ignored
	// Eg. If INFO is selected, All DEBUG logs are ignored
	// If ERROR is selected all logs except ERROR and FATAL are ignored
	Level Level

	// Log levels in string format.
	// The supported log level strings are Debug, Info, Warn, Error, Fatal
	// You can specify log level using Level Enum or string
	// The Enum value is given first preference
	LevelStr string

	// Size of the file to be printed, there are two possible values FULL, SHORT
	// SHORT - Only the file name is displayed
	// FULL - File name along with full file path is specified
	// SHORT is used by default
	FilePathSize int

	// Log Reference (context) ID to be added to each log
	// This can be used to search relevent logs for the context
	Reference string

	AppName string

	remoteLogger bool

	RemoteLoggerURL string

	RemoteToken string

	RemoteUserName string
}

func NewConfig(name string) *Config {
	uuid := uuid.NewV4()
	return &Config{
		Reference:       uuid.String(),
		Level:           DEBUG,
		FilePathSize:    SHORT,
		AppName:         name,
		remoteLogger:    false,
		RemoteLoggerURL: "",
		RemoteToken:     "",
		RemoteUserName:  "",
	}
}

func (c *Config) SetRemoteConfig(url, token, uname string) {
	if url != "" {
		c.remoteLogger = true
	}

	c.RemoteLoggerURL = url
	c.RemoteToken = token
	c.RemoteUserName = uname
}

func (c *Config) SetLevel(level Level) {
	c.Level = level
}

func (c *Config) SetLevelStr(lvl string) {
	switch lvl {
	case Debug:
		c.Level = DEBUG
	case Info:
		c.Level = INFO
	case Warn:
		c.Level = WARN
	case Error:
		c.Level = ERROR
	case Fatal:
		c.Level = FATAL
	default:
		c.Level = INFO
	}
}

func (c *Config) SetFilePathSize(filePathSize int) {
	c.FilePathSize = filePathSize
}

func (c *Config) SetFilePathSizeStr(fps string) {
	switch fps {
	case FilePathShort:
		c.FilePathSize = SHORT
	case FilePathFull:
		c.FilePathSize = FULL
	default:
		c.FilePathSize = SHORT
	}
}

func (c *Config) SetReference(ref string) {
	if ref == "" {
		return
	}

	c.Reference = ref
}
