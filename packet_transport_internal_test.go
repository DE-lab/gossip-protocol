/*
 * Copyright 2018 De-labtory
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPacketTransport_listen(t *testing.T) {
	// given
	c := PacketTransportConfig{
		BindAddress: "127.0.0.1",
		BindPort:    1234,
	}

	// when
	p, err := NewPacketTransport(&c)
	assert.Nil(t, err)

	p.WriteTo([]byte{'a'}, "127.0.0.1:1234")

	// then
	packet := <-p.packetCh
	assert.NotNil(t, packet)
}
