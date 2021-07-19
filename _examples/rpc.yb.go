// Package rpc is generated by yobuffer@0.1.0
package  rpc

import (
	"reflect"

	yb "github.com/cupen/yobuffer"
	"github.com/cupen/yobuffer/encodings/tiny"
)

var (
	structs map[int]*reflect.Value
	tinyEncoder = tiny.NewEncoder()
	tinyDecoder = tiny.NewDecoder()
)

func init() {
	structs = map[int]*reflect.Value{
		0: reflect.TypeOf(&PlayerInfo{}),
	}
}

// PlayerInfo 
type PlayerInfo struct {
	UserId string "yobuffer:\"1\""
	Name string "yobuffer:\"2\""
	Avatar string "yobuffer:\"3\""
	ShortId int64 "yobuffer:\"4\""
	Level int16 "yobuffer:\"5\""
	WinRate float64 "yobuffer:\"6\""
	IsAdmin bool "yobuffer:\"7\""
	
}

func (t *PlayerInfo) Size() (size int) {
	size += 4 + len(t.UserId) // UserId<string>

	size += 4 + len(t.Name) // Name<string>

	size += 4 + len(t.Avatar) // Avatar<string>

	size += 8 // ShortId<int64>

	size += 2 // Level<int16>

	size += 8 // WinRate<float64>

	size += 1 // IsAdmin<bool>

	return
}

func (t *PlayerInfo) MarshalYB() ([]byte, error) {
	dAtA := make([]byte, t.Size())
	i := 0
	// UserId<string>
	tinyEncoder.String(dAtA[i:], t.UserId)
	i += 4+len(UserId)

	// Name<string>
	tinyEncoder.String(dAtA[i:], t.Name)
	i += 4+len(Name)

	// Avatar<string>
	tinyEncoder.String(dAtA[i:], t.Avatar)
	i += 4+len(Avatar)

	// ShortId<int64>
	tinyEncoder.Int64(dAtA[i:], t.ShortId)
	i += 8

	// Level<int16>
	tinyEncoder.Int16(dAtA[i:], t.Level)
	i += 2

	// WinRate<float64>
	tinyEncoder.Float64(dAtA[i:], t.WinRate)
	i += 8

	// IsAdmin<bool>
	tinyEncoder.Bool(dAtA[i:], t.IsAdmin)
	i += 1

	if i != len(dAtA) {
		return nil, fmt.Errorf("invalid data size. expected:%d  actual:%d", len(dAtA), i)
	}
	return dAtA, nil
}

func (t *PlayerInfo) UnmarshalYB(data []byte) error {
	i := 0
	// UserId<string>
	t.UserId = tinyDecoder.String(dAtA[i:])
	i += 4+len(UserId)

	// Name<string>
	t.Name = tinyDecoder.String(dAtA[i:])
	i += 4+len(Name)

	// Avatar<string>
	t.Avatar = tinyDecoder.String(dAtA[i:])
	i += 4+len(Avatar)

	// ShortId<int64>
	t.ShortId = tinyDecoder.Int64(dAtA[i:])
	i += 8

	// Level<int16>
	t.Level = tinyDecoder.Int16(dAtA[i:])
	i += 2

	// WinRate<float64>
	t.WinRate = tinyDecoder.Float64(dAtA[i:])
	i += 8

	// IsAdmin<bool>
	t.IsAdmin = tinyDecoder.Bool(dAtA[i:])
	i += 1

	return dAtA, nil

}
