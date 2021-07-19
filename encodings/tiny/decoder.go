package tiny

type Decoder struct{}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (dec *Decoder) Int8(b []byte) int8 {
	return int8(b[0])
}
func (dec *Decoder) Int16(b []byte) int16 {
	_ = b[1]
	return int16(b[0]) | int16(b[1])<<8
}

func (dec *Decoder) Int32(b []byte) int32 {
	_ = b[3]
	return int32(b[0]) | int32(b[1])<<8 |
		int32(b[2])<<16 | int32(b[3])<<24
}

func (dec *Decoder) Int64(b []byte) int64 {
	_ = b[7]
	return int64(b[0]) | int64(b[1])<<8 |
		int64(b[2])<<16 | int64(b[3])<<24 |
		int64(b[4])<<32 | int64(b[5])<<40 |
		int64(b[6])<<48 | int64(b[7])<<56
}

func (dec *Decoder) UInt8(b []byte) uint8 {
	return uint8(b[0])
}
func (dec *Decoder) UInt16(b []byte) uint16 {
	_ = b[1]
	return uint16(b[0]) | uint16(b[1])<<8
}
func (dec *Decoder) UInt32(b []byte) uint32 {
	_ = b[3]
	return uint32(b[0]) | uint32(b[1])<<8 |
		uint32(b[2])<<16 | uint32(b[3])<<24

}
func (dec *Decoder) UInt64(b []byte) uint64 {
	_ = b[7]
	return uint64(b[0]) | uint64(b[1])<<8 |
		uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 |
		uint64(b[6])<<48 | uint64(b[7])<<56
}

func (dec *Decoder) Bool(b []byte) bool {
	return b[0] == 1
}

func (dec *Decoder) String(b []byte) string {
	l := dec.UInt32(b)
	return string(b[3:l])
}
