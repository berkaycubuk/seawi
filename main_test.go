package main

import "testing"

func TestGetConfiguration(t *testing.T) {
	ans, err := getConfiguration("./test_data.yml")
	if err != nil {
		t.Error(err)
	}

	if ans.FirstName != "Berkay" {
		t.Error("FirstName is not correct")
	}

	if ans.LastName != "Çubuk" {
		t.Error("LastName is not correct")
	}
}
