package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, 1)
	p := New()
	p.Parse([]byte("1+2"))
}

func TestParserYobuffer(t *testing.T) {
	assert := require.New(t)
	p := New()
	c, err := p.Parse([]byte("package abc"))
	assert.NoError(err)
	assert.NotNil(c)
	assert.NotNil(c.Package)
	assert.Equal(c.Package.Name, "abc")
}
