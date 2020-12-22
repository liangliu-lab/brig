// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	capnp2 "github.com/sahib/brig/catfs/nodes/capnp"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Change describes a single change
type Change struct{ capnp.Struct }

// Change_TypeID is the unique identifier for the type Change.
const Change_TypeID = 0x9592300df48789af

func NewChange(s *capnp.Segment) (Change, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return Change{st}, err
}

func NewRootChange(s *capnp.Segment) (Change, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return Change{st}, err
}

func ReadRootChange(msg *capnp.Message) (Change, error) {
	root, err := msg.RootPtr()
	return Change{root.Struct()}, err
}

func (s Change) String() string {
	str, _ := text.Marshal(0x9592300df48789af, s.Struct)
	return str
}

func (s Change) Mask() uint64 {
	return s.Struct.Uint64(0)
}

func (s Change) SetMask(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Change) Head() (capnp2.Node, error) {
	p, err := s.Struct.Ptr(0)
	return capnp2.Node{Struct: p.Struct()}, err
}

func (s Change) HasHead() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Change) SetHead(v capnp2.Node) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewHead sets the head field to a newly
// allocated capnp2.Node struct, preferring placement in s's segment.
func (s Change) NewHead() (capnp2.Node, error) {
	ss, err := capnp2.NewNode(s.Struct.Segment())
	if err != nil {
		return capnp2.Node{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Change) Next() (capnp2.Node, error) {
	p, err := s.Struct.Ptr(1)
	return capnp2.Node{Struct: p.Struct()}, err
}

func (s Change) HasNext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Change) SetNext(v capnp2.Node) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewNext sets the next field to a newly
// allocated capnp2.Node struct, preferring placement in s's segment.
func (s Change) NewNext() (capnp2.Node, error) {
	ss, err := capnp2.NewNode(s.Struct.Segment())
	if err != nil {
		return capnp2.Node{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Change) Curr() (capnp2.Node, error) {
	p, err := s.Struct.Ptr(2)
	return capnp2.Node{Struct: p.Struct()}, err
}

func (s Change) HasCurr() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Change) SetCurr(v capnp2.Node) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewCurr sets the curr field to a newly
// allocated capnp2.Node struct, preferring placement in s's segment.
func (s Change) NewCurr() (capnp2.Node, error) {
	ss, err := capnp2.NewNode(s.Struct.Segment())
	if err != nil {
		return capnp2.Node{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Change) MovedTo() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s Change) HasMovedTo() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Change) MovedToBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s Change) SetMovedTo(v string) error {
	return s.Struct.SetText(3, v)
}

func (s Change) WasPreviouslyAt() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s Change) HasWasPreviouslyAt() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Change) WasPreviouslyAtBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s Change) SetWasPreviouslyAt(v string) error {
	return s.Struct.SetText(4, v)
}

// Change_List is a list of Change.
type Change_List struct{ capnp.List }

// NewChange creates a new list of Change.
func NewChange_List(s *capnp.Segment, sz int32) (Change_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return Change_List{l}, err
}

func (s Change_List) At(i int) Change { return Change{s.List.Struct(i)} }

func (s Change_List) Set(i int, v Change) error { return s.List.SetStruct(i, v.Struct) }

func (s Change_List) String() string {
	str, _ := text.MarshalList(0x9592300df48789af, s.List)
	return str
}

// Change_Promise is a wrapper for a Change promised by a client call.
type Change_Promise struct{ *capnp.Pipeline }

func (p Change_Promise) Struct() (Change, error) {
	s, err := p.Pipeline.Struct()
	return Change{s}, err
}

func (p Change_Promise) Head() capnp2.Node_Promise {
	return capnp2.Node_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Change_Promise) Next() capnp2.Node_Promise {
	return capnp2.Node_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Change_Promise) Curr() capnp2.Node_Promise {
	return capnp2.Node_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

// Patch contains a single change
type Patch struct{ capnp.Struct }

// Patch_TypeID is the unique identifier for the type Patch.
const Patch_TypeID = 0x927c7336e3054805

func NewPatch(s *capnp.Segment) (Patch, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Patch{st}, err
}

func NewRootPatch(s *capnp.Segment) (Patch, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Patch{st}, err
}

func ReadRootPatch(msg *capnp.Message) (Patch, error) {
	root, err := msg.RootPtr()
	return Patch{root.Struct()}, err
}

func (s Patch) String() string {
	str, _ := text.Marshal(0x927c7336e3054805, s.Struct)
	return str
}

func (s Patch) FromIndex() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Patch) SetFromIndex(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s Patch) CurrIndex() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Patch) SetCurrIndex(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

func (s Patch) Changes() (Change_List, error) {
	p, err := s.Struct.Ptr(0)
	return Change_List{List: p.List()}, err
}

func (s Patch) HasChanges() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Patch) SetChanges(v Change_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewChanges sets the changes field to a newly
// allocated Change_List, preferring placement in s's segment.
func (s Patch) NewChanges(n int32) (Change_List, error) {
	l, err := NewChange_List(s.Struct.Segment(), n)
	if err != nil {
		return Change_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Patch_List is a list of Patch.
type Patch_List struct{ capnp.List }

// NewPatch creates a new list of Patch.
func NewPatch_List(s *capnp.Segment, sz int32) (Patch_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return Patch_List{l}, err
}

func (s Patch_List) At(i int) Patch { return Patch{s.List.Struct(i)} }

func (s Patch_List) Set(i int, v Patch) error { return s.List.SetStruct(i, v.Struct) }

func (s Patch_List) String() string {
	str, _ := text.MarshalList(0x927c7336e3054805, s.List)
	return str
}

// Patch_Promise is a wrapper for a Patch promised by a client call.
type Patch_Promise struct{ *capnp.Pipeline }

func (p Patch_Promise) Struct() (Patch, error) {
	s, err := p.Pipeline.Struct()
	return Patch{s}, err
}

// Patches contains several patches
type Patches struct{ capnp.Struct }

// Patches_TypeID is the unique identifier for the type Patches.
const Patches_TypeID = 0xc2984f083ea5351d

func NewPatches(s *capnp.Segment) (Patches, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Patches{st}, err
}

func NewRootPatches(s *capnp.Segment) (Patches, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Patches{st}, err
}

func ReadRootPatches(msg *capnp.Message) (Patches, error) {
	root, err := msg.RootPtr()
	return Patches{root.Struct()}, err
}

func (s Patches) String() string {
	str, _ := text.Marshal(0xc2984f083ea5351d, s.Struct)
	return str
}

func (s Patches) Patches() (Patch_List, error) {
	p, err := s.Struct.Ptr(0)
	return Patch_List{List: p.List()}, err
}

func (s Patches) HasPatches() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Patches) SetPatches(v Patch_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewPatches sets the patches field to a newly
// allocated Patch_List, preferring placement in s's segment.
func (s Patches) NewPatches(n int32) (Patch_List, error) {
	l, err := NewPatch_List(s.Struct.Segment(), n)
	if err != nil {
		return Patch_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Patches_List is a list of Patches.
type Patches_List struct{ capnp.List }

// NewPatches creates a new list of Patches.
func NewPatches_List(s *capnp.Segment, sz int32) (Patches_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Patches_List{l}, err
}

func (s Patches_List) At(i int) Patches { return Patches{s.List.Struct(i)} }

func (s Patches_List) Set(i int, v Patches) error { return s.List.SetStruct(i, v.Struct) }

func (s Patches_List) String() string {
	str, _ := text.MarshalList(0xc2984f083ea5351d, s.List)
	return str
}

// Patches_Promise is a wrapper for a Patches promised by a client call.
type Patches_Promise struct{ *capnp.Pipeline }

func (p Patches_Promise) Struct() (Patches, error) {
	s, err := p.Pipeline.Struct()
	return Patches{s}, err
}

const schema_b943b54bf1683782 = "x\xda\x84\x93\xc1k\x13O\x1c\xc5\xbf\xef;\xbb\xbf\xb4" +
	"?R\xd319\x88T\xba\x88\x97z\xb0-E\x05\x0f" +
	"\xd6Z\x05E\xc5\x8c(\xeaEX7\xd3n0\xd9\x84" +
	"\xcc6V\xb0\x04z\x11=V\x85\x16\x14+\xb4b\xa1" +
	"R/\x82\x1e\xbc\x08\xfe\x0b\xfe\x03=\x89'\xc1K{" +
	"Y\x99$\xa6\xa5\x16{[>\xf3\xd8\xd9\xcf\xe3\xed\x10" +
	"p\x86\x87\xdd\x0c\x88T\xda\xfd/q/\xb8\xeb'\xcc" +
	"\xc39R}\xe0d\xf6d\xf8\xf3\xd2\x87\xf1O\xe4\"" +
	"E4\\\xde\x0f9\x93\x923\xfdre\x8d\x90\xac=" +
	"~\xf4\xabgh\xee\xb9\xcdb[\xd6M\x11\x8d\xdc\xc6" +
	"Ad\x8bHe\x8b\xe8\x1fY\xc0M\x10\x92C\xc7\x97" +
	"Ow]\x9d\xffB\xb2\x0f;\xde=\xf2\x9d\x0f#\xbb" +
	"\xc1\xa9\xec\x06\xf7g\x07\xc4(!\x09\xfcx\xc2\x0c\xd6" +
	"\x03a\x06\x03\xbf\x1aU\x07\xab~\x1c\x84\xc7\x9a\xcf\xa7" +
	"\xf2~\x1c \xcc\x03\xca\x01'w\x9e\xbeR\x9f\xbf=" +
	"\xf9J\xcaa\x8c\xf5\x01i\"\x89\xcd\xc4\xa6B/\xa8" +
	"p\x14\xfb\xc5\xc8x\xbeg\x8a\xd1dI{\xa3A\xe8" +
	"G\x93\xdaZ\x0b\x87\xc8\x01\x91<\x7f\x8dH\x9d\x13P" +
	"y\x86\x04r\xb0\xf0\x8a\x85\x97\x05\xd4-\x068\x07&" +
	"\x927\xce\x12\xa9\xbc\x80*1\x92\x89Z\xa5|1*" +
	"h\xc24\\b\xb8\xf6\xcb\xa7j\xb5\x1d\xac\xd1\xba\xd0" +
	"`\x1f!/\x80\xde\xad\xfa\x08\x16\xfe[w<\xf4#" +
	"1\xa9w\xf7\xf5\x9a\xbe\xc3\xf8\x1f\xc9x\xf3\x16\xaf " +
	"\xb4\x09j\xc5\xbbz\x9br\xdb\x18\xea@\xc7x\xe1(" +
	"\x91z&\xa0\x16\x19\x7f\x84_Z6/\xa0\x96\x18\x92" +
	"\xd12~m\xe1\x0b\x01\xf5\x96!\x05\xe7 \x88\xe4\xb2" +
	"\x85\x8b\x02j\x95!\x1d\x91\x83C$Wl7K\x02" +
	"\xea=C\xbaN\x0e.\x91|7K\xa4V\x05\xd4G" +
	"F\xa6\xec\x9b{\xe8&F7!\x13j\xbf\x80\xded" +
	"}s\xa2\xda\xf81\xf0\xc6v\xd1K\xc8Dz:\xde" +
	"\x05\xdb^\xff\xc6\x8dr\xa5\xae\x0b\xd7+H\x13#M" +
	"H\xee\xfb&_\xd3\xf5\"*S\xa6\xf4`,\xa6\xce" +
	"\xc9\x9e\x93J\x85\xda\xecYrsU\xdax\"\xa8\xb4" +
	"wet]\xd7\xfc\x92Wm\x9d\x10\x94\xd3)\xb9\xc7" +
	"6\xd2%\xa0\x8e0\x1a\xed\xc0\xd6\x0a:?\\k\x05" +
	"\xbf\x03\x00\x00\xff\xff\xe5\xc7\xe9\x1d"

func init() {
	schemas.Register(schema_b943b54bf1683782,
		0x927c7336e3054805,
		0x9592300df48789af,
		0xc2984f083ea5351d)
}
