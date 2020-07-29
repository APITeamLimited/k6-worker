/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2016 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package js

import (
	"context"
	"os"
	"strconv"

	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
)

// console represents a JS console implemented as a logrus.Logger.
type console struct ***REMOVED***
	logger logrus.FieldLogger
***REMOVED***

// Creates a console with the standard logrus logger.
func newConsole() *console ***REMOVED***
	return &console***REMOVED***logrus.StandardLogger().WithField("source", "console")***REMOVED***
***REMOVED***

// Creates a console logger with its output set to the file at the provided `filepath`.
func newFileConsole(filepath string) (*console, error) ***REMOVED***
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644) //nolint:gosec
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	l := logrus.New()
	l.SetOutput(f)

	// TODO: refactor to not rely on global variables, albeit external ones
	l.SetFormatter(logrus.StandardLogger().Formatter)

	return &console***REMOVED***l***REMOVED***, nil
***REMOVED***

func (c console) log(ctx *context.Context, level logrus.Level, msgobj goja.Value, args ...goja.Value) ***REMOVED***
	if ctx != nil && *ctx != nil ***REMOVED***
		select ***REMOVED***
		case <-(*ctx).Done():
			return
		default:
		***REMOVED***
	***REMOVED***

	// TODO this is not how it works anywhere else
	// nodejs: https://nodejs.org/api/console.html#console_console_info_data_args
	// mdn: https://developer.mozilla.org/en-US/docs/Web/API/Console/log
	fields := make(logrus.Fields, len(args)+1)
	for i, arg := range args ***REMOVED***
		fields[strconv.Itoa(i)] = arg.String()
	***REMOVED***
	msg := msgobj.ToString()
	e := c.logger.WithFields(fields)
	switch level ***REMOVED*** //nolint:exhaustive
	case logrus.DebugLevel:
		e.Debug(msg)
	case logrus.InfoLevel:
		e.Info(msg)
	case logrus.WarnLevel:
		e.Warn(msg)
	case logrus.ErrorLevel:
		e.Error(msg)
	***REMOVED***
***REMOVED***

func (c console) Log(ctx *context.Context, msg goja.Value, args ...goja.Value) ***REMOVED***
	c.Info(ctx, msg, args...)
***REMOVED***

func (c console) Debug(ctx *context.Context, msg goja.Value, args ...goja.Value) ***REMOVED***
	c.log(ctx, logrus.DebugLevel, msg, args...)
***REMOVED***

func (c console) Info(ctx *context.Context, msg goja.Value, args ...goja.Value) ***REMOVED***
	c.log(ctx, logrus.InfoLevel, msg, args...)
***REMOVED***

func (c console) Warn(ctx *context.Context, msg goja.Value, args ...goja.Value) ***REMOVED***
	c.log(ctx, logrus.WarnLevel, msg, args...)
***REMOVED***

func (c console) Error(ctx *context.Context, msg goja.Value, args ...goja.Value) ***REMOVED***
	c.log(ctx, logrus.ErrorLevel, msg, args...)
***REMOVED***
