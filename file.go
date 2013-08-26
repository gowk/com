// Copyright 2013 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package com

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

// get filepath base name
func Basename(file string) string {
	return path.Base(file)
}

// get filepath dir name
func Dirname(file string) string {
	return path.Dir(file)
}

// get absolute filepath, based on built excutable file
func RealPath(file string) (string, error) {
	if path.IsAbs(file) {
		return file, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, file), err
}

// get file modified time
func FileMTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// copy file to new path
func Copy(file string, to string) (bool, error) {
	sf, e := os.Open(file)
	if e != nil {
		return false, e
	}
	defer sf.Close()
	df, e2 := os.Create(to)
	if e2 != nil {
		return false, e2
	}
	defer df.Close()
	// buffer reader, do chunk transfer
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := sf.Read(buf)
		if err != nil && err != io.EOF {
			return false, err
		}
		if n == 0 {
			break
		}
		// write a chunk
		if _, err := df.Write(buf[:n]); err != nil {
			return false, err
		}
	}
	return true, nil
}

// move file to new path
func Move(file string, to string) (bool, error) {
	// copy
	r, e := Copy(file, to)
	if e != nil {
		return r, e
	}
	// then remove
	e = os.Remove(file)
	if e != nil {
		r = false
	}
	return r, e
}

// delete file
func Unlink(file string) error {
	return os.Remove(file)
}

// rename file name
func Rename(file string, to string) error {
	return os.Rename(file, to)
}

// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// SaveFile saves content type '[]byte' to file by given path.
// It returns error when fail to finish operation.
func SaveFile(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

// SaveFileS saves content type 'string' to file by given path.
// It returns error when fail to finish operation.
func SaveFileS(filePath string, s string) (int, error) {
	return SaveFile(filePath, []byte(s))
}

// ReadFile reads data type '[]byte' from file by given path.
// It returns error when fail to finish operation.
func ReadFile(filePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}

// ReadFileS reads data type 'string' from file by given path.
// It returns error when fail to finish operation.
func ReadFileS(filePath string) (string, error) {
	b, err := ReadFile(filePath)
	return string(b), err
}
