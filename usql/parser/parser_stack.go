// Copyright (C) 2018 The go-sqlparser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

// ParserStack is a stack for any parser objects.
type ParserStack struct {
	objects []interface{}
}

// NewParserStack creates a new stack.
func NewParserStack() *ParserStack {
	s := &ParserStack{
		objects: make([]interface{}, 0),
	}
	return s
}

// PushObject adds a parser object to the stack.
func (s *ParserStack) PushObject(obj interface{}) {
	s.objects = append(s.objects, obj)
}

// PeekObject returns a top parser object
func (s *ParserStack) PeekObject() interface{} {
	objectCount := len(s.objects)
	if objectCount <= 0 {
		return nil
	}
	return s.objects[objectCount-1]
}

// PopObject removes and returns a top parser object
func (s *ParserStack) PopObject() interface{} {
	objectCount := len(s.objects)
	if objectCount <= 0 {
		return nil
	}
	obj := s.objects[objectCount-1]
	s.objects = s.objects[0:(objectCount - 1)]
	return obj
}

// Size return a count of the stack objects
func (s *ParserStack) Size() int {
	return len(s.objects)
}
