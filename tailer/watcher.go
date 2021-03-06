// Copyright 2016-2018 The grok_exporter Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tailer

import (
	"github.com/sirupsen/logrus"
	"io"
)

// File system notification watcher.
type Watcher interface {
	io.Closer
	StartEventLoop() EventLoop
}

// Loop reading file system events from the operating system.
type EventLoop interface {
	io.Closer
	Errors() chan error
	Events() chan Events
}

// File system events.
// The operating system may return more than one event for each read, so it's plural.
type Events interface {
	Process(fileBefore *File, reader *lineReader, abspath string, logger logrus.FieldLogger) (file *File, lines []string, err error)
}
