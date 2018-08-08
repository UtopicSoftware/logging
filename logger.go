//
// Copyright © 2018 Anton Filatov
//
// This file is part of GoLogging project.
//
// GoLogging project is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// GoLogging project is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with GoLogging project. If not, see <https://www.gnu.org/licenses/lgpl-3.0>.
//

package logging

import (
	"fmt"
)

// Logger is a common logging API definition
type Logger interface {
	Trace(arg ...interface{})
	Tracef(pattern string, arg ...interface{})
	Logger(name string) Logger
}

// Factory is a basic logger constructor
type Factory interface {
	Logger(name string) Logger
	Close() error
}

// Level represents logs output threshold
type Level uint8

// Logging levels from least informative to most informative
const (
	OFF = Level(iota)
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
	ALL = ^Level(0)
)

func (l Level) String() string {
	switch l {
	case OFF:
		return "OFF"
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case ALL:
		return "ALL"
	default:
		return fmt.Sprintf("UNKNOWN_LEVEL_%d", l)
	}
}
