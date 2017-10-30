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
	"testing"
	"time"

	"github.com/256dpi/gomqtt/tools"
	"github.com/stretchr/testify/assert"
)

func TestGenericFutureBind(t *testing.T) {
	f1 := newGenericFuture()
	f1.Cancel()

	f2 := newGenericFuture()
	go f2.Bind(f1)

	err := f2.Wait(10 * time.Millisecond)
	assert.Equal(t, tools.ErrFutureCanceled, err)
}

func TestSubscribeFutureBind(t *testing.T) {
	f1 := newSubscribeFuture()
	f1.Cancel()

	f2 := newSubscribeFuture()
	go f2.Bind(f1)

	err := f2.Wait(10 * time.Millisecond)
	assert.Equal(t, tools.ErrFutureCanceled, err)
}
