package ownersmodels

import (
	"encoding/json"
	"testing"
)

func TestOwnerRetainsInactiveUserIdentity(t *testing.T) {
	var owner Owner
	if err := json.Unmarshal([]byte(`{"id":"17","userIdIncludingInactive":123,"archived":true}`), &owner); err != nil {
		t.Fatal(err)
	}
	if owner.UserIDIncludingInactive != 123 {
		t.Fatalf("inactive user = %d", owner.UserIDIncludingInactive)
	}
}
