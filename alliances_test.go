package gocrest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MAX_ALLIANCES_PAGE = 250

func TestGetAlliances(t *testing.T) {
	var expected_alliances int = MAX_ALLIANCES_PAGE

	alliances, err := GetAlliances()
	assert.Nil(t, err,
		fmt.Sprintf("gocrest: GetAlliances call failed. Could be networking related: %v", err))

	// If the total count is greater than what appears to be the default CREST limit && the length of
	// alliances is lower than the page, we're probably at the end...
	if (alliances.TotalCount > MAX_ALLIANCES_PAGE) && len(alliances.Alliances) <= MAX_ALLIANCES_PAGE {
		expected_alliances = len(alliances.Alliances)
	}
	assert.Equal(t, expected_alliances, len(alliances.Alliances), "total count did not match")
}
