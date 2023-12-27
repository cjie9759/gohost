package client

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetHostInfo(t *testing.T) {
	info := getHostInfo()
	d, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s", d)

}
