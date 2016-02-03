// Copyright (c) 2014 The gomqtt Authors. All rights reserved.
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

package client

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestFutureStore(t *testing.T) {
	f := &ConnectFuture{}

	s := newFutureStore()
	assert.Equal(t, 0, len(s.all()))

	s.put(1, f)
	assert.Equal(t, f, s.get(1))
	assert.Equal(t, 1, len(s.all()))

	s.del(1)
	assert.Nil(t, s.get(1))
	assert.Equal(t, 0, len(s.all()))
}

func TestFutureStoreAwait(t *testing.T) {
	f := &ConnectFuture{}
	f.initialize()

	s := newFutureStore()
	assert.Equal(t, 0, len(s.all()))

	s.put(1, f)
	assert.Equal(t, f, s.get(1))
	assert.Equal(t, 1, len(s.all()))

	go func(){
		time.Sleep(1 * time.Millisecond)
		f.complete()
		s.del(1)
	}()

	err := s.await(10 * time.Millisecond)
	assert.NoError(t, err)
}

func TestFutureStoreAwaitTimeout(t *testing.T) {
	f := &ConnectFuture{}
	f.initialize()

	s := newFutureStore()
	assert.Equal(t, 0, len(s.all()))

	s.put(1, f)
	assert.Equal(t, f, s.get(1))
	assert.Equal(t, 1, len(s.all()))

	err := s.await(10 * time.Millisecond)
	assert.Equal(t, ErrTimeoutExceeded, err)
}

func TestCounter(t *testing.T) {
	c := newCounter()
	assert.Equal(t, uint16(0), c.next())

	c.id = math.MaxUint16
	assert.Equal(t, uint16(math.MaxUint16), c.next())
	assert.Equal(t, uint16(0), c.next())
}
