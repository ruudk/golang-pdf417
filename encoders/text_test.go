package encoders

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTextEncoder_CanEncode(t *testing.T) {
	assert := assert.New(t)

	encoder := CreateTextEncoder()

	for ord := int(' '); ord < int('Z'); ord++ {
		chr := rune(ord)

		if chr == '"' {
			continue
		}

		assert.True(
			encoder.CanEncode(string(chr)),
			"Unable to encode: %d %c",
			ord,
			chr,
		)
	}
}

func TestTextEncoder_GetSwitchCode(t *testing.T) {
	assert := assert.New(t)

	encoder := CreateTextEncoder()

	assert.Equal(TEXT_SWITCH_CODE_WORD, encoder.GetSwitchCode("123"))
	assert.Equal(TEXT_SWITCH_CODE_WORD, encoder.GetSwitchCode("foo"))
}

func TestTextEncoder_Encode(t *testing.T) {
	assert := assert.New(t)

	encoder := CreateTextEncoder()

	assert.Equal([]int64{900, 567, 615, 137, 808, 760}, encoder.Encode("Super !", true))
	assert.Equal([]int64{567, 615, 137, 808, 760}, encoder.Encode("Super !", false))
}

func TestTextEncoder_Encode2(t *testing.T) {
	assert := assert.New(t)

	encoder := CreateTextEncoder()

	assert.Equal([]int64{900, 567, 615, 137, 809}, encoder.Encode("Super ", true))
	assert.Equal([]int64{567, 615, 137, 809}, encoder.Encode("Super ", false))
}
