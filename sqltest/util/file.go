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

package util

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// File represents a file or director.
type File struct {
	path string
}

const (
	errorFileIsNotDirectory = "%s is not a directory"
)

// NewFileWithPath return a file with the specified path.
func NewFileWithPath(path string) *File {
	file := &File{
		path: path,
	}
	return file
}

// Path returns the path.
func (file *File) Path() string {
	return file.path
}

// Name returns only the name.
func (file *File) Name() string {
	return filepath.Base(file.path)
}

// Ext returns only the extension.
func (file *File) Ext() string {
	return filepath.Ext(file.path)
}

// IsDir returns true when the file represents a directory, otherwise false.
func (file *File) IsDir() bool {
	fi, err := os.Stat(file.path)
	if err != nil {
		return false
	}
	if !fi.IsDir() {
		return false
	}
	return true
}

// ListFilesWithExtention returns files which has the specified extensions in the directory.
func (file *File) ListFilesWithExtention(targetExt string) ([]*File, error) {
	if !file.IsDir() {
		return nil, fmt.Errorf(errorFileIsNotDirectory, file.path)
	}

	rootPath := file.path
	files := []*File{}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if 0 < len(targetExt) {
			fileExt := filepath.Ext(path)
			if !strings.HasSuffix(fileExt, targetExt) {
				return nil
			}
		}
		files = append(files, NewFileWithPath(path))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// ListFilesWithExtention returns files which has the specified extensions in the directory.
func (file *File) ListFilesWithRegexp(re *regexp.Regexp) ([]*File, error) {
	if !file.IsDir() {
		return nil, fmt.Errorf(errorFileIsNotDirectory, file.path)
	}

	rootPath := file.path
	files := []*File{}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !re.MatchString(path) {
			return nil
		}
		files = append(files, NewFileWithPath(path))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// ListFiles returns a file list in the directory.
func (file *File) ListFiles() ([]*File, error) {
	return file.ListFilesWithExtention("")
}
