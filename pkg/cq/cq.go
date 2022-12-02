package cq

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type cq struct {
	client  *http.Client
	url     string
	debug   bool
	cqData  map[string]string
	errDeal func(rep *http.Response) error
}

var (
	Cq *cq
)

func Init(url string, group_id int, debug bool) *cq {
	Cq = &cq{
		client: &http.Client{},
		url:    url,
		debug:  debug,
		cqData: map[string]string{
			"group_id":    strconv.Itoa(group_id),
			"message":     "defalt msg",
			"auto_escape": "false",
		},
	}
	Cq.errDeal = errDeal
	if debug {
		Cq.errDeal = errDealWithBug
	}
	return Cq

}

func (c *cq) Send(msg string) (err error) {
	c.cqData["message"] = msg
	req, err := http.NewRequest("POST", c.url, urlEncode(c.cqData))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", `application/x-www-form-urlencoded`)
	rep, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer rep.Body.Close()
	return c.errDeal(rep)
}

func urlEncode(m map[string]string) io.Reader {
	params := url.Values{}
	for k, v := range m {
		params.Add(k, v)
	}
	return bytes.NewBufferString(params.Encode())
}

func readAllString(i io.Reader) string {
	bs, err := io.ReadAll(i)
	if err != nil {
		return ""
	}
	return string(bs)
}

func errDeal(rep *http.Response) error {
	if rep.StatusCode != http.StatusOK {
		return fmt.Errorf(strconv.Itoa(rep.StatusCode))
	}
	return nil
}
func errDealWithBug(rep *http.Response) error {
	if rep.StatusCode != http.StatusOK {
		fmt.Println(readAllString(rep.Body))
		return fmt.Errorf(strconv.Itoa(rep.StatusCode))
	}
	return nil
}
