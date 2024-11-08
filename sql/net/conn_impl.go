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

package net

import (
	"net"
	"time"

	"github.com/cybergarage/go-tracing/tracer"
	"github.com/google/uuid"
)

// ConnOption represents a connection option.
type ConnOption = func(*conn)

// conn represents a connection.
type conn struct {
	net.Conn
	isClosed bool
	db       string
	ts       time.Time
	uuid     uuid.UUID
	id       ConnID
	tracer.Context
}

// NewConnWith returns a connection with a raw connection.
func NewConnWith(netConn net.Conn, opts ...ConnOption) Conn {
	conn := &conn{
		Conn:     netConn,
		isClosed: false,
		db:       "",
		ts:       time.Now(),
		uuid:     uuid.New(),
		id:       0,
		Context:  nil,
	}
	for _, opt := range opts {
		opt(conn)
	}
	return conn
}

// WithConnDatabase sets a database name.
func WithConnDatabase(name string) func(*conn) {
	return func(conn *conn) {
		conn.db = name
	}
}

// WithConnTracer sets a tracer context.
func WithConnTracer(t tracer.Context) func(*conn) {
	return func(conn *conn) {
		conn.Context = t
	}
}

// WithConnID sets a connection ID.
func WithConnID(v ConnID) func(*conn) {
	return func(conn *conn) {
		conn.id = v
	}
}

// Close closes the connection.
func (conn *conn) Close() error {
	if conn.isClosed {
		return nil
	}
	if err := conn.Conn.Close(); err != nil {
		return err
	}
	conn.isClosed = true
	return nil
}

// SetDatabase sets the database name.
func (conn *conn) SetDatabase(db string) {
	conn.db = db
}

// Database returns the database name.
func (conn *conn) Database() string {
	return conn.db
}

// Timestamp returns the creation time of the connection.
func (conn *conn) Timestamp() time.Time {
	return conn.ts
}

// UUID returns the UUID of the connection.
func (conn *conn) UUID() uuid.UUID {
	return conn.uuid
}

// ID returns the ID of the connection.
func (conn *conn) ID() ConnID {
	return ConnID(conn.id)
}

// SetSpanContext sets the tracer span context of the connection.
func (conn *conn) SetSpanContext(ctx tracer.Context) {
	conn.Context = ctx
}

// SpanContext returns the tracer span context of the connection.
func (conn *conn) SpanContext() tracer.Context {
	return conn.Context
}
