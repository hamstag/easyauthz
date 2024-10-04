package easyauthz

import "testing"

func TestAuthorize(t *testing.T) {
	result := Authorize("edit-setting")
	if result != true {
		t.Error("Not return true")
	}
}
