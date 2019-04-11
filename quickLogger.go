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

// QuickLogger is an utility class for logging shortcuts
type QuickLogger interface {
	Fatal(arg ...interface{})
	Fatalf(pattern string, arg ...interface{})
	Error(arg ...interface{})
	Errorf(pattern string, arg ...interface{})
	Warn(arg ...interface{})
	Warnf(fapattern string, rg ...interface{})
	Info(arg ...interface{})
	Infof(fapattern string, rg ...interface{})
	Debug(arg ...interface{})
	Debugf(pattern string, arg ...interface{})
	Trace(arg ...interface{})
	Tracef(pattern string, arg ...interface{})
	Tracee(arg ...interface{})
	Traceef(pattern string, arg ...interface{})
	Traceee(arg ...interface{})
	Traceeef(pattern string, arg ...interface{})
	QuickLoggerProvider
	Logger
}

// QuickLoggerProvider can create QuickLogger objects
type QuickLoggerProvider interface {
	NewQuickLogger(name ...string) (QuickLogger, error)
}

type quickLogger struct {
	logger Logger
}

// NewQuickLogger creates new QuickLogger wrapper for Logger
func NewQuickLogger(loggerBase Logger, err error) (QuickLogger, error) {
	if err != nil {
		return nil, err
	}
	// Check if logger is wrapped by logging.logger
	if loggerWrap, ok := loggerBase.(*logger); ok {
		return &quickLogger{
			logger: loggerWrap.logger,
		}, err
	} else
	// Check if logger is wrapped by us
	if loggerWrap, ok := loggerBase.(*quickLogger); ok {
		return &quickLogger{
			logger: loggerWrap.logger,
		}, err
	}
	// Assuming loggerBase is not wrapped
	return &quickLogger{
		logger: loggerBase,
	}, err
}

func (l *quickLogger) NewQuickLogger(name ...string) (QuickLogger, error) {
	return NewQuickLogger(l.logger.NewLogger(name...))
}

func (l *quickLogger) NewLogger(name ...string) (Logger, error) {
	return NewLogger(l.logger.NewLogger(name...))
}

func (l *quickLogger) Log(level Level, arg ...interface{}) {
	l.logger.Log(level, arg...)
}

func (l *quickLogger) Fatal(arg ...interface{}) {
	l.logger.Log(FATAL, arg...)
}

func (l *quickLogger) Error(arg ...interface{}) {
	l.logger.Log(ERROR, arg...)
}

func (l *quickLogger) Warn(arg ...interface{}) {
	l.logger.Log(WARN, arg...)
}

func (l *quickLogger) Info(arg ...interface{}) {
	l.logger.Log(INFO, arg...)
}

func (l *quickLogger) Debug(arg ...interface{}) {
	l.logger.Log(DEBUG, arg...)
}

func (l *quickLogger) Trace(arg ...interface{}) {
	l.logger.Log(TRACE, arg...)
}

func (l *quickLogger) Tracee(arg ...interface{}) {
	l.logger.Log(TRACEE, arg...)
}

func (l *quickLogger) Traceee(arg ...interface{}) {
	l.logger.Log(TRACEEE, arg...)
}

func (l *quickLogger) Logf(level Level, pattern string, arg ...interface{}) {
	l.logger.Logf(level, pattern, arg...)
}

func (l *quickLogger) Fatalf(pattern string, arg ...interface{}) {
	l.logger.Logf(FATAL, pattern, arg...)
}

func (l *quickLogger) Errorf(pattern string, arg ...interface{}) {
	l.logger.Logf(ERROR, pattern, arg...)
}

func (l *quickLogger) Warnf(pattern string, arg ...interface{}) {
	l.logger.Logf(WARN, pattern, arg...)
}

func (l *quickLogger) Infof(pattern string, arg ...interface{}) {
	l.logger.Logf(INFO, pattern, arg...)
}

func (l *quickLogger) Debugf(pattern string, arg ...interface{}) {
	l.logger.Logf(DEBUG, pattern, arg...)
}

func (l *quickLogger) Tracef(pattern string, arg ...interface{}) {
	l.logger.Logf(TRACE, pattern, arg...)
}

func (l *quickLogger) Traceef(pattern string, arg ...interface{}) {
	l.logger.Logf(TRACEE, pattern, arg...)
}

func (l *quickLogger) Traceeef(pattern string, arg ...interface{}) {
	l.logger.Logf(TRACEEE, pattern, arg...)
}
