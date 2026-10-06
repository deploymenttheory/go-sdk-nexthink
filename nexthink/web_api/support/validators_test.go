package support

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateSearchRequest(t *testing.T) { require.Error(t, validateSearchRequest(nil)) }
