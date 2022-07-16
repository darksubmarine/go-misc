package ds_

import (
	"github.com/rs/zerolog"
	"io"
	"os"
)

// Log returns a zerolog.Logger instance the will write to stdout
func Log() zerolog.Logger {
	return LogToStdOut()
}

// LogTo returns a zerolog.Logger instance
func LogTo(w io.Writer) zerolog.Logger {
	return zerolog.New(w).With().Timestamp().Logger()
}

// LogToStdErr returns a zerolog.Logger instance the will write to stderr
func LogToStdErr() zerolog.Logger {
	return LogTo(os.Stderr)
}

// LogToStdOut returns a zerolog.Logger instance the will write to stdout
func LogToStdOut() zerolog.Logger {
	return LogTo(os.Stdout)
}

