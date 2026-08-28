package datautils

import (
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeRoundTrip(t *testing.T) {
	id, error := uuid.NewV7()
	if error != nil {
		t.Fatal("Unable to generate uuid")
	}
	original := Data{
		id,
		"charles.bedard@gmail.com",
		"charles bedard",
		1787507854,
		31,
	}

	encoded, err := EncodeKey(original)
	if err != nil {
		t.Fatal("Error during encoding/decoding")
	}
	updated, err := Decode(encoded)
	if err != nil {
		t.Fatal("Error during encoding/decoding")
	}

	assert.Equal(t, updated, original)
}
