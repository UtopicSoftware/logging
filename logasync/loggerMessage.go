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
	"fmt"
	"os"
	"time"

	"github.com/UtopicSoftware/logging"
)

type loggerMessage struct {
	level   logging.Level
	ts      time.Time
	name    string
	pattern *string
	file    *string
	line    int
	args    *[]interface{}
}

func (m *loggerMessage) format(cfg *Config) []byte {
	var log string
	if p := m.pattern; p != nil {
		log = fmt.Sprintf(*p, (*m.args)...)
	} else {
		log = fmt.Sprintln((*m.args)...)
	}
	if cfg.Flags&SourceLine != 0 {
		var file string
		if nil == m.file {
			file = "???"
		} else {
			file = *m.file
			if cfg.Flags&SourceShort != 0 {
				short := file
				for i := len(file) - 1; i > 0; i-- {
					if file[i] == '/' || file[i] == os.PathSeparator {
						short = file[i+1:]
						break
					}
				}
				file = short
			}
		}
		return []byte(fmt.Sprintf(cfg.OutPattern, m.ts.Format(cfg.TimePattern), m.name, m.level, log, file, m.line))
	}
	return []byte(fmt.Sprintf(cfg.OutPattern, m.ts.Format(cfg.TimePattern), m.name, m.level, log))
}
