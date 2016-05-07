package pdf417

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateDataEncoder(t *testing.T) {
	assert := assert.New(t)

	encoder := CreateDataEncoder()

	assert.NotEmpty(encoder.Encoders)
}

func TestEncode(t *testing.T) {
	assert := assert.New(t)

	encoder := CreateDataEncoder()

	// When starting with text, the first code word does not need to be the switch
	result := encoder.Encode("ABC123")
	assert.NotEqual(TEXT_SWITCH_CODE_WORD, result[0])
	assert.Equal([]int{1, 89, 902, 1, 223}, result)

	// When starting with numbers, we do need to switchresult := encoder.Encode("ABC123")
	result = encoder.Encode("123ABC")
	assert.Equal(NUMBER_SWITCH_CODE_WORD, result[0])
	assert.Equal([]int{902, 1, 223, 900, 1, 89}, result)

	// Also with bytes
	//result = encoder.Encode("\x0B")
	//assert.Equal(BYTE_SWITCH_CODE_WORD, result[0])
	//assert.Equal([]int{901, 11}, result)

	// Alternate bytes switch code when number of bytes is divisble by 6
	//result = encoder.Encode("\x0B\x0B\x0B\x0B\x0B\x0B")
	//assert.Equal(BYTE_SWITCH_CODE_WORD_ALT, result[0])
	//assert.Equal([]int{924, 18, 455, 694, 754, 291}, result)
}
