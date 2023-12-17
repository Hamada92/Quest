package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testSecretKey = []byte("md91034fj01934n1")

func TestJWT(t *testing.T) {
	_, err := NewJWT(testSecretKey)

	require.NoError(t, err)
}
