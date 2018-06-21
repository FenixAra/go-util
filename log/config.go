package log

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

	RemoteLogger bool

	RemoteLoggerURL string

	RemoteToken string

	RemoteUserName string
}

func NewConfig(ref, levelStr, filePathSizeStr, appName, remoteLoggerURL, token string) *Config {
	var level Level
	var filePathSize int
	switch levelStr {
	case Debug:
		level = DEBUG
	case Info:
		level = INFO
	case Warn:
		level = WARN
	case Error:
		level = ERROR
	case Fatal:
		level = FATAL
	default:
		level = INFO
	}

	var remoteLogger bool
	if remoteLoggerURL != "" {
		remoteLogger = true
	}

	switch filePathSizeStr {
	case FilePathShort:
		filePathSize = SHORT
	case FilePathFull:
		filePathSize = FULL
	default:
		filePathSize = SHORT
	}

	return &Config{
		Reference:       ref,
		Level:           level,
		FilePathSize:    filePathSize,
		AppName:         appName,
		RemoteLogger:    remoteLogger,
		RemoteLoggerURL: remoteLoggerURL,
		RemoteToken:     token,
	}
}

func (c *Config) SetLevel(level Level) {
	c.Level = level
}

func (c *Config) SetFilePathSize(filePathSize int) {
	c.FilePathSize = filePathSize
}

func (c *Config) SetReference(ref string) {
	c.Reference = ref
}
