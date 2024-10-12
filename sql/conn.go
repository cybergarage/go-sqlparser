// Copyright (C) 2019 The go-sqlparser Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

import (
	"net"
	"time"

	"github.com/cybergarage/go-tracing/tracer"
	"github.com/google/uuid"
)

// ConnID represents a connection ID.
type ConnID uint64

// Conn represents a connection.
type Conn interface {
	tracer.Context
	// RemoteAddr returns the remote address.
	RemoteAddr() net.Addr
	// Close closes the connection.
	Close() error
	// SetDatabase sets a database name.
	SetDatabase(db string)
	// Database returns the database name.
	Database() string
	// Timestamp returns the timestamp.
	Timestamp() time.Time
	// UUID returns the UUID.
	UUID() uuid.UUID
	// ID returns the ID.
	ID() ConnID
	// SetSpanContext sets a span context.
	SetSpanContext(ctx tracer.Context)
	// SpanContext returns a span context.
	SpanContext() tracer.Context
}
