package encoders

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNumberEncoder_CanEncode(t *testing.T) {
	assert := assert.New(t)

	encoder := CreateNumberEncoder()

	assert.True(encoder.CanEncode("0"))
	assert.True(encoder.CanEncode("1"))
	assert.True(encoder.CanEncode("2"))
	assert.True(encoder.CanEncode("3"))
	assert.True(encoder.CanEncode("4"))
	assert.True(encoder.CanEncode("5"))
	assert.True(encoder.CanEncode("6"))
	assert.True(encoder.CanEncode("7"))
	assert.True(encoder.CanEncode("8"))
	assert.True(encoder.CanEncode("9"))

	assert.False(encoder.CanEncode(""))
	assert.False(encoder.CanEncode("a"))
	assert.False(encoder.CanEncode("abc"))
	assert.False(encoder.CanEncode("123"))
}

func TestNumberEncoder_GetSwitchCode(t *testing.T) {
	assert := assert.New(t)

	encoder := CreateNumberEncoder()

	assert.Equal(NUMBER_SWITCH_CODE_WORD, encoder.GetSwitchCode("123"))
	assert.Equal(NUMBER_SWITCH_CODE_WORD, encoder.GetSwitchCode("foo"))
}

func TestNumberEncoder_Encode(t *testing.T) {
	assert := assert.New(t)

	encoder := CreateNumberEncoder()

	assert.Equal([]int64{902, 112, 434}, encoder.Encode("01234", true))
	assert.Equal([]int64{112, 434}, encoder.Encode("01234", false))

}
