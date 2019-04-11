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

package logasync

import (
	"strings"
	"time"

	"github.com/UtopicSoftware/logging"
)

type logger struct {
	name    string
	factory *loggerFactory
}

func (l *logger) NewLogger(name ...string) (logging.Logger, error) {
	return l.factory.NewLogger(strings.Join([]string{l.name, strings.Join(name, l.factory.cfg.NamingDelimiter)}, l.factory.cfg.NamingDelimiter))
}

func (l *logger) Log(level logging.Level, arg ...interface{}) {
	l.log(level, nil, arg...)
}

func (l *logger) Logf(level logging.Level, pattern string, arg ...interface{}) {
	l.log(level, &pattern, arg...)
}

func (l *logger) log(level logging.Level, pattern *string, arg ...interface{}) {
	t := time.Now()
	l.factory.accept(&loggerMessage{
		level:   level,
		ts:      t,
		name:    l.name,
		pattern: pattern,
		args:    &arg,
	}, 4)
}
