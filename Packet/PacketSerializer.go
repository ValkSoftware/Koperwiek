package packet

import (
	"log"

	"github.com/google/uuid"
)

type Serializer struct {
	Index uint64
	Data  []byte
}

func (s *Serializer) Clear() {
	s.Data = []byte{}
}

func (s *Serializer) Finish() []byte {
	b := []byte{}    // temp
	a := len(s.Data) // length of packet

	for {
		if (a & 0xFFFFFF80) == 0 {
			b = append(b, byte(a))
			break
		}
		b = append(b, byte(a&0x7F|0x80))
		a >>= 7
	}

	return append(b, s.Data...) // varint length + packet data
}

// gopls stop complaining
func (s *Serializer) WriteByte(a uint8) {
	s.Data = append(s.Data, a)
}

func (s *Serializer) WriteSignedByte(a int8) {
	s.Data = append(s.Data, byte(a))
}

func (s *Serializer) WriteShort(a uint16) {
	s.Data = append(s.Data, byte(a>>8), byte(a))
}

func (s *Serializer) WriteSignedShort(a int16) {
	s.Data = append(s.Data, byte(a>>8), byte(a))
}

func (s *Serializer) WriteInt(a int32) {
	s.Data = append(s.Data, byte(a>>24), byte(a>>16), byte(a>>8), byte(a))
}

func (s *Serializer) WriteLong(a int64) {
	s.Data = append(s.Data,
		byte(a>>56), byte(a>>48), byte(a>>40), byte(a>>32),
		byte(a>>24), byte(a>>16), byte(a>>8), byte(a),
	)
}

func (s *Serializer) WriteByteArray(a []byte) {
	s.Data = append(s.Data, a...)
}

func (s *Serializer) WriteUUID(a string) {
	u, err := uuid.Parse(a)
	if err != nil {
		log.Fatal("Failed to parse UUID:" + a)
	}

	r, err := u.MarshalBinary()
	if err != nil {
		log.Fatal("Failed to serialize UUID " + a + " to bytes")
	}

	s.Data = append(s.Data, r...)
}

func (s *Serializer) WriteDouble(a float64) {
	s.WriteLong(int64(a))
}

func (s *Serializer) WriteFloat(a float32) {
	s.WriteInt(int32(a))
}

func (s *Serializer) WriteVarint(a uint32) {
	for {
		if (a & 0xFFFFFF80) == 0 {
			s.Data = append(s.Data, byte(a))
			break
		}
		s.Data = append(s.Data, byte(a&0x7F|0x80))
		a >>= 7
	}
}

func (s *Serializer) WriteVarlong(a uint64) {
	for {
		if (a & 0xFFFFFFFFFFFFFF80) == 0 {
			s.Data = append(s.Data, byte(a))
			return
		}
		s.Data = append(s.Data, byte(a&0x7F|0x80))
		a >>= 7
	}
}

func (s *Serializer) WriteString(a string) {
	s.WriteVarint(uint32(len(a)))
	s.WriteByteArray([]byte(a))
}

func (s *Serializer) ReadByte() uint8 {
	s.Index++
	return s.Data[s.Index-1]
}

func (s *Serializer) ReadSignedByte() int8 {
	s.Index++
	return int8(s.Data[s.Index-1])
}

func (s *Serializer) ReadShort() uint16 {
	s.Index += 2
	return uint16(s.Data[s.Index-2])>>8 + uint16(s.Data[s.Index-1])
}

func (s *Serializer) ReadSignedShort() int16 {
	s.Index += 2
	return int16(s.Data[s.Index-2])>>8 + int16(s.Data[s.Index-1])
}

func (s *Serializer) ReadLong() int64 {
	s.Index += 8
	return int64(s.Data[s.Index-8])>>56 +
		int64(s.Data[s.Index-7])>>48 +
		int64(s.Data[s.Index-6])>>40 +
		int64(s.Data[s.Index-5])>>32 +
		int64(s.Data[s.Index-4])>>24 +
		int64(s.Data[s.Index-3])>>16 +
		int64(s.Data[s.Index-2])>>8 +
		int64(s.Data[s.Index-1])
}

func (s *Serializer) ReadString() string {
	return string(s.ReadByteArray(s.ReadVarint()))
}

func (s *Serializer) ReadByteArray(len int32) []byte {
	b := []byte{}
	for i := int32(0); i < len; i++ {
		b = append(b, s.Data[s.Index])
		s.Index++
	}
	return b
}

func (s *Serializer) ReadUUID() [16]byte {
	var res [16]byte = [16]byte{}
	copy(res[:], s.ReadByteArray(16))
	return res
}

func (s *Serializer) ReadBool() bool {
	s.Index++
	if s.Data[s.Index-1] == 0 {
		return false
	} else {
		return true
	}
}

func (s *Serializer) ReadVarint() int32 {
	pos := 0
	var val int32 = 0
	for {
		s.Index++
		val |= int32(s.Data[s.Index-1]&0x7F) << pos

		if (s.Data[s.Index-1] & 0x80) == 0 {
			return val
		}

		pos += 7

		if pos >= 32 {
			log.Fatal("Varint too big")
		}

	}
}
