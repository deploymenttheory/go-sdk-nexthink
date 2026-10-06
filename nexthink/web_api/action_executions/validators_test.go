package action_executions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateListOptions(t *testing.T)    { require.Error(t, validateListOptions(nil)) }
func TestValidateDetailsRequest(t *testing.T) { require.Error(t, validateDetailsRequest(nil)) }
func TestValidateNQLRequest(t *testing.T)     { require.Error(t, validateNQLRequest(nil)) }
func TestValidateExecuteRequest(t *testing.T) { require.Error(t, validateExecuteRequest(nil)) }
func TestValidateHistoryRequest(t *testing.T) { require.Error(t, validateHistoryRequest(nil)) }
