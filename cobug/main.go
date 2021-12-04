package cobug

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

func NewBug() *Bug {
	return &Bug{
		Client: client,
		Header: map[string]string{
			`User-Agent`: `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:94.0) Gecko/20100101 Firefox/94.0`,
		},
	}
}

type Bug struct {
	Client   *http.Client
	Request  *http.Request
	Response *http.Response
	Header   map[string]string
	Err      error
}

func (b *Bug) Get(url string, body io.Reader) *Bug {
	b.Request, b.Err = http.NewRequest("GET", url, body)
	return b.do()
}

func (b *Bug) Post(url string, body io.Reader) *Bug {
	if _, ok := b.Header["Content-Type"]; !ok {
		b.Header["Content-Type"] = `application/x-www-form-urlencoded`
	}
	b.Request, b.Err = http.NewRequest("POST", url, body)
	return b.do()
}

func (b *Bug) SetHeader(m map[string]string) *Bug {
	for k, v := range m {
		b.Header[k] = v
	}
	return b
}

func (b *Bug) DisAutoLoacationClien() *Bug {
	b.Client = disAutoLoacationClien
	return b
}

func (b *Bug) Body() []byte {
	if b.Err != nil {
		return nil
	}

	var r []byte

	r, b.Err = ioutil.ReadAll(b.Response.Body)
	return r
}

func (b *Bug) GetCookie() string {
	if b.Err != nil {
		return ""
	}
	cs := b.Response.Header.Values("Set-Cookie")
	c := ""
	re := regexp.MustCompile(`^[^\ ]*?;`)
	for _, v := range cs {
		c += re.FindString(v)
	}
	return c
}

func (b *Bug) do() *Bug {
	if b.Err != nil {
		fmt.Println("err in Bug do  ", b.Err)
		return b
	}

	for k, v := range b.Header {
		b.Request.Header.Set(k, v)
	}
	for i := 0; i < 10; i++ {
		b.Response, b.Err = b.Client.Do(b.Request)
		if b.Err == nil {
			break
		}
	}
	return b
}

func UrlEncode(m map[string]string) io.Reader {
	params := url.Values{}
	for k, v := range m {
		params.Add(k, v)
	}
	return bytes.NewBufferString(params.Encode())
}
