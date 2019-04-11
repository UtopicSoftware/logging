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
	"runtime"
	"strings"
	"sync"

	"github.com/UtopicSoftware/logging"
)

type messageChannel chan *loggerMessage

type loggerFactory struct {
	cfg      Config
	ch       messageChannel
	wg       sync.WaitGroup
	isClosed bool
}

// NewFactoryDefault creates asyncronous logging factory with defaultParameters
func NewFactoryDefault() (logging.Factory, error) {
	return NewFactory(DefaultConfig)
}

// NewFactory creates asyncronous logging factory
func NewFactory(config ConfigProvider) (logging.Factory, error) {
	cfg := Config{}
	config(&cfg)
	f := &loggerFactory{
		cfg: cfg,
		ch:  make(messageChannel, cfg.QueueSize),
	}
	go func() {
		for msg := range f.ch {
			if _, err := cfg.Writer.Write(msg.format(&cfg)); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
			f.wg.Done()
		}
	}()
	return f, nil
}

func (f *loggerFactory) NewLogger(name ...string) (logging.Logger, error) {
	if err := f.errorIfClosed("LoggerFactory is closed"); err != nil {
		return nil, err
	}
	return logging.NewLogger(&logger{
		name:    strings.Join(name, f.cfg.NamingDelimiter),
		factory: f,
	}, nil)
}

func (f *loggerFactory) Close() error {
	if err := f.errorIfClosed("LoggerFactory is already closed"); err != nil {
		return err
	}
	f.isClosed = true
	f.wg.Wait()
	close(f.ch)
	return nil
}

func (f *loggerFactory) accept(msg *loggerMessage, skip int) {
	if f.cfg.Level < msg.level {
		return
	}
	if err := f.errorIfClosed("LoggerFactory is closed"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	if f.cfg.Flags&SourceLine != 0 {
		var _, file, line, ok = runtime.Caller(skip)
		if ok {
			msg.file = &file
			msg.line = line
		}
	}
	f.wg.Add(1)
	f.ch <- msg
}

func (f *loggerFactory) errorIfClosed(message string) *loggerError {
	if f.isClosed {
		return &loggerError{message}
	}
	return nil
}
