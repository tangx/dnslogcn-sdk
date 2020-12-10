package dnslog

import (
	"fmt"
	"net/http"
)

func dnslogGet(url string, session string) (resp *http.Response, err error) {

	//
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", session)

	return http.DefaultClient.Do(req)

}

func trans(body []byte) []byte {
	header := `{"records":`
	footer := `}`

	r := fmt.Sprintf("%s%s%s", header, string(body), footer)

	return []byte(r)
}
