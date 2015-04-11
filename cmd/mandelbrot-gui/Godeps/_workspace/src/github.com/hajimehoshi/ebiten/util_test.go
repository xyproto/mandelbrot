// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ebiten_test

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/gopherjs/gopherjs/js"
	"io"
	"io/ioutil"
	"strings"
)

// NOTE: Please update if any of testdata/* is updated.
var data = map[string]string{
	"testdata/ebiten.png": `
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
SUVORK5CYII=`,
}

func readFile(path string) (io.Reader, error) {
	if js.Global == nil {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(b), nil
	}
	// According to https://github.com/gopherjs/gopherjs/blob/master/doc/syscalls.md,
	// syscall can be available on the unstable version of npm.
	// Let's not use os package here.
	str, ok := data[path]
	if !ok {
		return nil, errors.New("not found")
	}
	return base64.NewDecoder(base64.StdEncoding, strings.NewReader(str)), nil
}
