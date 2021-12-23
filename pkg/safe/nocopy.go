/*
 * Copyright 2021 The KunStack Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package safe

import "sync"

var _ sync.Locker = &NoCopy{}

// NoCopy prohibits the copying of struct objects
// sample:
// type Foo struct {
// noCopy
//}
// a go vet grammar check will prompt a warning
type NoCopy struct{}

// Lock noop lock
func (*NoCopy) Lock() {}

// Unlock noop Unlock
func (*NoCopy) Unlock() {}