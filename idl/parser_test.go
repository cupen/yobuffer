package idl

import (
	"errors"
	"testing"

	"github.com/cupen/yobuffer/idl/errorhandler"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {

	// slog.SetDefault(slog.NewLogLogger(slog.Default().Handler(), slog.LevelDebug))
	assert := require.New(t)
	p := New()
	c, err := p.Parse([]byte(testdata_full))
	assert.NoError(err)
	assert.NotNil(c)
	assert.NotNil(c.Package)
	assert.Equal("pkgname", c.Package.Name)
	assert.Equal(nil, c)
}

func TestParserPackage(t *testing.T) {
	assert := require.New(t)
	p := New()
	// assert.PanicsWithError("package statement should end with ';'", func() {
	c, err := p.Parse([]byte("package abc"))
	assert.Error(err)
	assert.True(errors.Is(err, errorhandler.ErrSytax))
	// assert.ErrorIs(errorhandler.ErrSytax, err)
	assert.Nil(c)
	// })

	c, err = p.Parse([]byte("package abc;"))
	assert.NoError(err)
	assert.NotNil(c)
	assert.NotNil(c.Package)
	assert.Equal(c.Package.Name, "abc")
}
