package packet_test

import (
	"testing"

	packet "valksoftware.nl/koperwiek/Packet"
)

func TestSerializeWriteShort(t *testing.T) {
	s := packet.Serializer{0, []byte{}}

	s.WriteShort(uint16(0xFF))

	test := []byte{0x00, 0xFF}

	for i, v := range test {
		if s.Data[i] != v {
			t.FailNow()
		}
	}
}

func TestSerializeWriteInt(t *testing.T) {
	s := packet.Serializer{0, []byte{}}

	s.WriteInt(int32(0xFF))

	test := []byte{0x00, 0x00, 0x00, 0xFF}

	for i, v := range test {
		if s.Data[i] != v {
			t.FailNow()
		}
	}
}

func TestSerializeWriteVarint(t *testing.T) {
	s := packet.Serializer{0, []byte{}}

	s.WriteVarint(uint32(0xFF))

	test := []byte{0xFF, 0x01}

	for i, v := range test {
		if s.Data[i] != v {
			t.FailNow()
		}
	}
}

func TestSerializeWriteVarlong(t *testing.T) {
	s := packet.Serializer{0, []byte{}}

	s.WriteVarlong(uint64(0xFF))

	test := []byte{0xFF, 0x01}

	for i, v := range test {
		if s.Data[i] != v {
			t.FailNow()
		}
	}
}

func TestSerializeWriteVarlong2(t *testing.T) {
	s := packet.Serializer{0, []byte{}}

	s.WriteVarlong(uint64(2147483647))

	test := []byte{0xff, 0xff, 0xff, 0xff, 0x07}

	for i, v := range test {
		if s.Data[i] != v {
			t.FailNow()
		}
	}
}
