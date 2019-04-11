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
	"io"
	"os"

	"github.com/UtopicSoftware/logging"
)

//Flags is a logger working mode settings
type Flags uint8

//Basic flags
const (
	SourceLine = Flags(1 << iota)
	SourceShort
)

// Config contains settings for logger factory
type Config struct {
	QueueSize       int
	Writer          io.Writer
	NamingDelimiter string
	OutPattern      string
	TimePattern     string
	Level           logging.Level
	Flags           Flags
}

// ConfigProvider provides configuration for logger factory
type ConfigProvider func(cfg *Config) *Config

// DefaultConfig is a default configuration provider for logger factory
func DefaultConfig(cfg *Config) *Config {
	cfg.QueueSize = 64
	cfg.NamingDelimiter = "."
	cfg.OutPattern = "[%[1]v] %[3]v %[2]v: %[4]v"
	cfg.TimePattern = "2006-01-02 15:04:05.000000-07:00"
	cfg.Writer = os.Stdout
	cfg.Level = logging.INFO
	return cfg
}
