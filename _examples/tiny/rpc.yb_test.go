package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerInfo_marshal(t *testing.T) {
	assert := assert.New(t)

	p := PlayerInfo{
		UserId:  "userId",
		Name:    "name",
		Avatar:  "avatar",
		ShortId: 12123,
		Level:   12,
		WinRate: 95.27,
		IsAdmin: true,
	}
	data, err := p.Marshal()
	assert.NoError(err)
	assert.NotNil(data)
	assert.Equal(p.Size(), len(data))

	p2 := PlayerInfo{}
	err = p2.Unmarshal(data)
	assert.NoError(err)
	assert.Equal(p, p2)
}
