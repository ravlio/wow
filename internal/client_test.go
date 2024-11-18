package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateHashcash(t *testing.T) {
	require.Equal(t, generateHashcash("example_question", 2), 23)

}
