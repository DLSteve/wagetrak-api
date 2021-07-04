package currency

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClientNeedsRefresh_goodUpdated(t *testing.T) {
	updated, _ := time.Parse("2006-01-02", "2021-06-28")
	refreshTime := updated.Add(time.Hour * 4)
	now := updated.Add(time.Hour * 10)

	needsRefresh := needsRefresh(updated, refreshTime, now)
	assert.Equalf(t, false, needsRefresh, "Should be false")
}

func TestClientNeedsRefresh_staleUpdated(t *testing.T) {
	updated, _ := time.Parse("2006-01-02", "2021-06-28")
	refreshTime := updated.Add(time.Hour * 1)
	now := updated.Add(time.Hour * 24)

	needsRefresh := needsRefresh(updated, refreshTime, now)
	assert.Equalf(t, true, needsRefresh, "Should be true")
}

func TestClientNeedsRefresh_goodRefreshed(t *testing.T) {
	updated, _ := time.Parse("2006-01-02", "2021-06-28")
	refreshTime := updated.Add(time.Hour * 27)
	now := updated.Add(time.Hour * 28)

	needsRefresh := needsRefresh(updated, refreshTime, now)
	assert.Equalf(t, false, needsRefresh, "Should be true")
}

func TestClientNeedsRefresh_staleRefreshed(t *testing.T) {
	updated, _ := time.Parse("2006-01-02", "2021-06-28")
	refreshTime := updated.Add(time.Hour * 3)
	now := updated.Add(time.Hour * 24)

	needsRefresh := needsRefresh(updated, refreshTime, now)
	assert.Equalf(t, true, needsRefresh, "Should be true")
}
