/*
Copyright 2013 The Camlistore Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package misc contains stuff which should probably move elsewhere.
//
// This is a gross place to put code.
package misc

import (
	"crypto/sha1"
	"fmt"
)

// SHA1Prefix computes the SHA-1 digest of data and returns
// the first ten digits of its lowercase hex string.
func SHA1Prefix(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))[:10]
}
