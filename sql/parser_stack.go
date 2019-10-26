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
