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
	"runtime"
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
			cfg.Writer.Write(msg.format(&cfg))
			f.wg.Done()
		}
	}()
	return f, nil
}

func (f *loggerFactory) Logger(name string) logging.Logger {
	if err := f.errorIfClosed("LoggerFactory is closed"); err != nil {
		panic(err)
	}
	return &logger{
		name:    name,
		factory: f,
	}
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

func (f *loggerFactory) accept(msg *loggerMessage) {
	if f.cfg.Level < msg.level {
		return
	}
	if err := f.errorIfClosed("LoggerFactory is closed"); err != nil {
		panic(err)
	}
	if f.cfg.Flags&SourceLine != 0 {
		var _, file, line, ok = runtime.Caller(2)
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
