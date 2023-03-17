package timeonly

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTimeOnly(t *testing.T) {
	_, err := NewTimeOnly(26, 12, 2)
	require.True(t, errors.Is(err, Error_TimeOnlyParse))

	timeOnly, err := NewTimeOnly(12, 12, 0)
	require.Nil(t, err)
	require.NotNil(t, timeOnly)

	jsonV, _ := json.Marshal(timeOnly)
	require.Equal(t, string(jsonV), `"12:12:00"`)
	require.Equal(t, "12:12:00", timeOnly.ToString())

	require.Equal(t, timeOnly.Hour(), 12)
	require.Equal(t, timeOnly.Minute(), 12)
	require.Equal(t, timeOnly.Second(), 0)
}
