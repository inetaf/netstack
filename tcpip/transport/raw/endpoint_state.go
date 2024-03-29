// Copyright 2018 The gVisor Authors.
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

package raw

import (
	"fmt"
	"time"

	"inet.af/netstack/tcpip"
	"inet.af/netstack/tcpip/buffer"
	"inet.af/netstack/tcpip/stack"
)

// saveReceivedAt is invoked by stateify.
func (p *rawPacket) saveReceivedAt() int64 {
	return p.receivedAt.UnixNano()
}

// loadReceivedAt is invoked by stateify.
func (p *rawPacket) loadReceivedAt(nsec int64) {
	p.receivedAt = time.Unix(0, nsec)
}

// saveData saves rawPacket.data field.
func (p *rawPacket) saveData() buffer.VectorisedView {
	// We cannot save p.data directly as p.data.views may alias to p.views,
	// which is not allowed by state framework (in-struct pointer).
	return p.data.Clone(nil)
}

// loadData loads rawPacket.data field.
func (p *rawPacket) loadData(data buffer.VectorisedView) {
	// NOTE: We cannot do the p.data = data.Clone(p.views[:]) optimization
	// here because data.views is not guaranteed to be loaded by now. Plus,
	// data.views will be allocated anyway so there really is little point
	// of utilizing p.views for data.views.
	p.data = data
}

// afterLoad is invoked by stateify.
func (e *endpoint) afterLoad() {
	stack.StackFromEnv.RegisterRestoredEndpoint(e)
}

// beforeSave is invoked by stateify.
func (e *endpoint) beforeSave() {
	e.freeze()
}

// Resume implements tcpip.ResumableEndpoint.Resume.
func (e *endpoint) Resume(s *stack.Stack) {
	e.net.Resume(s)

	e.thaw()
	e.stack = s
	e.ops.InitHandler(e, e.stack, tcpip.GetStackSendBufferLimits, tcpip.GetStackReceiveBufferLimits)

	if e.associated {
		netProto := e.net.NetProto()
		if err := e.stack.RegisterRawTransportEndpoint(netProto, e.transProto, e); err != nil {
			panic(fmt.Sprintf("e.stack.RegisterRawTransportEndpoint(%d, %d, _): %s", netProto, e.transProto, err))
		}
	}
}
