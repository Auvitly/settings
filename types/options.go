package types

type Options string

// General option types
const (
	TimeFormat      Options = "time_format"
	ProcessingMode  Options = "processing_mode"
	LoggerHook      Options = "logger_hook"
	LoggerInstance  Options = "logger_instance"
	ValidatorEnable Options = "validator_enable"
)

// ProcessingMode
const (
	OverwritingMode string = "overwriting"
	ComplementMode  string = "complement"
)
