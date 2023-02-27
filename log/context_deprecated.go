/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package log

import (
	"github.com/containerd/log"
)

// Fields type to pass to "WithFields".
type Fields = log.Fields

// Entry is a logging entry.
type Entry = log.Entry

// RFC3339NanoFixed is [time.RFC3339Nano] with nanoseconds padded using
// zeros to ensure the formatted time is always the same number of
// characters.

// Level is a logging level.
type Level = log.Level

// Supported log levels.
const (
	// TraceLevel level.
	TraceLevel Level = log.TraceLevel

	// DebugLevel level.
	DebugLevel Level = log.DebugLevel

	// InfoLevel level.
	InfoLevel Level = log.InfoLevel

	// WarnLevel level.
	WarnLevel Level = log.WarnLevel

	// ErrorLevel level
	ErrorLevel Level = log.ErrorLevel

	// FatalLevel level.
	FatalLevel Level = log.FatalLevel

	// PanicLevel level.
	PanicLevel Level = log.PanicLevel
)

// SetLevel sets log level globally. It returns an error if the given
// level is not supported.
func SetLevel(level string) error {
	return log.SetLevel(level)
}

// GetLevel returns the current log level.
func GetLevel() log.Level {
	return log.GetLevel()
}

// OutputFormat specifies a log output format.
type OutputFormat = log.OutputFormat

// SetFormat sets the log output format.
func SetFormat(format OutputFormat) error {
	return log.SetFormat(format)
}
