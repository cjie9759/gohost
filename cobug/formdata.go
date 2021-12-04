package cobug

import (
	"bytes"
	"fmt"
)

func NewFormData() *FormData {
	return &FormData{
		body:     &bytes.Buffer{},
		boundary: `-----------------------------9051914041544843365972754266`,
	}
}

type FormData struct {
	body     *bytes.Buffer
	boundary string
}

func (t *FormData) Add(m map[string]string) {
	for k, v := range m {
		t.body.WriteString(fmt.Sprint(t.boundary, "\n", `Content-Disposition: form-data; name="`, k,
			`"`, "\n\n", v, "\n"))
	}
}
func (t *FormData) AddFile(name, filename string, body []byte) {
	t.body.WriteString(fmt.Sprint(t.boundary, "\n", `Content-Disposition: form-data; name="`, name, `"; filename="`,
		filename, `"`, "\nContent-Type: application/octet-stream", "\n\n"))
	n, err := t.body.Write(body)
	if err != nil {
		fmt.Printf("err in whrite file  %v %s", n, err)
	}
	t.body.WriteString("\n")
}

func (f *FormData) GetBody() []byte {
	f.body.WriteString(f.boundary)
	f.body.WriteString("--")
	return f.body.Bytes()
}

func (f *FormData) GetRaw() *bytes.Buffer {
	f.body.WriteString(f.boundary)
	f.body.WriteString("--")
	return f.body
}
