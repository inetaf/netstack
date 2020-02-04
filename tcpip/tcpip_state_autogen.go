// automatically generated by stateify.

// +build go1.9
// +build !go1.15

package tcpip

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *PacketBuffer) save(m state.Map) {
	x.beforeSave()
	m.Save("Data", &x.Data)
	m.Save("DataOffset", &x.DataOffset)
	m.Save("DataSize", &x.DataSize)
	m.Save("Header", &x.Header)
	m.Save("LinkHeader", &x.LinkHeader)
	m.Save("NetworkHeader", &x.NetworkHeader)
	m.Save("TransportHeader", &x.TransportHeader)
}

func (x *PacketBuffer) afterLoad() {}
func (x *PacketBuffer) load(m state.Map) {
	m.Load("Data", &x.Data)
	m.Load("DataOffset", &x.DataOffset)
	m.Load("DataSize", &x.DataSize)
	m.Load("Header", &x.Header)
	m.Load("LinkHeader", &x.LinkHeader)
	m.Load("NetworkHeader", &x.NetworkHeader)
	m.Load("TransportHeader", &x.TransportHeader)
}

func (x *FullAddress) beforeSave() {}
func (x *FullAddress) save(m state.Map) {
	x.beforeSave()
	m.Save("NIC", &x.NIC)
	m.Save("Addr", &x.Addr)
	m.Save("Port", &x.Port)
}

func (x *FullAddress) afterLoad() {}
func (x *FullAddress) load(m state.Map) {
	m.Load("NIC", &x.NIC)
	m.Load("Addr", &x.Addr)
	m.Load("Port", &x.Port)
}

func (x *ControlMessages) beforeSave() {}
func (x *ControlMessages) save(m state.Map) {
	x.beforeSave()
	m.Save("HasTimestamp", &x.HasTimestamp)
	m.Save("Timestamp", &x.Timestamp)
	m.Save("HasInq", &x.HasInq)
	m.Save("Inq", &x.Inq)
	m.Save("HasTOS", &x.HasTOS)
	m.Save("TOS", &x.TOS)
	m.Save("HasTClass", &x.HasTClass)
	m.Save("TClass", &x.TClass)
}

func (x *ControlMessages) afterLoad() {}
func (x *ControlMessages) load(m state.Map) {
	m.Load("HasTimestamp", &x.HasTimestamp)
	m.Load("Timestamp", &x.Timestamp)
	m.Load("HasInq", &x.HasInq)
	m.Load("Inq", &x.Inq)
	m.Load("HasTOS", &x.HasTOS)
	m.Load("TOS", &x.TOS)
	m.Load("HasTClass", &x.HasTClass)
	m.Load("TClass", &x.TClass)
}

func init() {
	state.Register("tcpip.PacketBuffer", (*PacketBuffer)(nil), state.Fns{Save: (*PacketBuffer).save, Load: (*PacketBuffer).load})
	state.Register("tcpip.FullAddress", (*FullAddress)(nil), state.Fns{Save: (*FullAddress).save, Load: (*FullAddress).load})
	state.Register("tcpip.ControlMessages", (*ControlMessages)(nil), state.Fns{Save: (*ControlMessages).save, Load: (*ControlMessages).load})
}
