// Generated by github.com/davyxu/cellnet/objprotogen
// DO NOT EDIT!
package proto

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/goobjfmt"
	"reflect"
)

func (self *TestMsgACK) String() string        { return goobjfmt.CompactTextString(self) }
func (self *BindClientREQ) String() string     { return goobjfmt.CompactTextString(self) }
func (self *BindClientACK) String() string     { return goobjfmt.CompactTextString(self) }
func (self *PID) String() string               { return goobjfmt.CompactTextString(self) }
func (self *Start) String() string             { return goobjfmt.CompactTextString(self) }
func (self *Stop) String() string              { return goobjfmt.CompactTextString(self) }
func (self *SystemExit) String() string        { return goobjfmt.CompactTextString(self) }
func (self *NexusOpen) String() string         { return goobjfmt.CompactTextString(self) }
func (self *NexusClose) String() string        { return goobjfmt.CompactTextString(self) }
func (self *RouteACK) String() string          { return goobjfmt.CompactTextString(self) }
func (self *DomainIdentifyACK) String() string { return goobjfmt.CompactTextString(self) }

func init() {

	cellnet.RegisterMessageMeta("binary", "proto.TestMsgACK", reflect.TypeOf((*TestMsgACK)(nil)).Elem(), 2238643133)
	cellnet.RegisterMessageMeta("binary", "proto.BindClientREQ", reflect.TypeOf((*BindClientREQ)(nil)).Elem(), 2423039609)
	cellnet.RegisterMessageMeta("binary", "proto.BindClientACK", reflect.TypeOf((*BindClientACK)(nil)).Elem(), 4214643227)
	cellnet.RegisterMessageMeta("binary", "proto.PID", reflect.TypeOf((*PID)(nil)).Elem(), 748122295)
	cellnet.RegisterMessageMeta("binary", "proto.Start", reflect.TypeOf((*Start)(nil)).Elem(), 3048584699)
	cellnet.RegisterMessageMeta("binary", "proto.Stop", reflect.TypeOf((*Stop)(nil)).Elem(), 2782698386)
	cellnet.RegisterMessageMeta("binary", "proto.SystemExit", reflect.TypeOf((*SystemExit)(nil)).Elem(), 3199756597)
	cellnet.RegisterMessageMeta("binary", "proto.NexusOpen", reflect.TypeOf((*NexusOpen)(nil)).Elem(), 2848479583)
	cellnet.RegisterMessageMeta("binary", "proto.NexusClose", reflect.TypeOf((*NexusClose)(nil)).Elem(), 801786219)
	cellnet.RegisterMessageMeta("binary", "proto.RouteACK", reflect.TypeOf((*RouteACK)(nil)).Elem(), 892591827)
	cellnet.RegisterMessageMeta("binary", "proto.DomainIdentifyACK", reflect.TypeOf((*DomainIdentifyACK)(nil)).Elem(), 3151462984)
}