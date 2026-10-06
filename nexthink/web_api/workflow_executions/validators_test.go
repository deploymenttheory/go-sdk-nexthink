package workflow_executions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateTimelineOptions(t *testing.T)   { require.Error(t, validateTimelineOptions(nil)) }
func TestValidateListOptions(t *testing.T)       { require.Error(t, validateListOptions(nil)) }
func TestValidateExecuteRequest(t *testing.T)    { require.Error(t, validateExecuteRequest(nil)) }
func TestValidateExecuteNQLRequest(t *testing.T) { require.Error(t, validateExecuteNQLRequest(nil)) }
