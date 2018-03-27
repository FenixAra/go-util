package log

type Config struct {
	// The support log levels are as follows
	// DEBUG < INFO < WARN < ERROR < FATAL
	// If a log level is specified all logs with level below the specified level are ignored
	// Eg. If INFO is selected, All DEBUG logs are ignored
	// If ERROR is selected all logs except ERROR and FATAL are ignored
	Level int

	// Size of the file to be printed, there are two possible values FULL, SHORT
	// SHORT - Only the file name is displayed
	// FULL - File name along with full file path is specified
	// SHORT is used by default
	FileSize int

	// Log Reference (context) ID to be added to each log
	// This can be used to search relevent logs for the context
	Reference string
}

func NewConfig(ref string, level, fileSize int) *Config {
	return &Config{
		Reference: ref,
		Level:     level,
		FileSize:  fileSize,
	}
}

func (c *Config) SetLevel(level int) {
	c.Level = level
}

func (c *Config) SetFileSize(fileSize int) {
	c.FileSize = fileSize
}

func (c *Config) SetReference(ref string) {
	c.Reference = ref
}
