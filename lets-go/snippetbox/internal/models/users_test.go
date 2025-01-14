package models

import (
	"krikchaip/snippetbox/internal/assert"
	"krikchaip/snippetbox/internal/testutils"
	"testing"
)

func TestUserModelExists(t *testing.T) {
	// Skip the test if the "-short" flag is provided when running the test.
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	cases := []struct {
		name   string
		userID int
		want   bool
	}{
		{name: "Valid ID", userID: 1, want: true},
		{name: "Zero ID", userID: 0, want: false},
		{name: "Non-existent ID", userID: 2, want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			db := testutils.NewTestDB(t)
			users := NewUserModel(db)

			got, err := users.Exists(c.userID)

			assert.Equal(t, got, c.want)
			assert.NilError(t, err)
		})
	}
}
