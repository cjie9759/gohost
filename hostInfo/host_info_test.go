package hostinfo_test

import (
	"encoding/json"
	"fmt"
	hostinfo "gohost/hostInfo"
	"testing"
)

func Test(t *testing.T) {
	h := hostinfo.GetHostInfo()
	d, err := json.MarshalIndent(h, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s", d)
}
