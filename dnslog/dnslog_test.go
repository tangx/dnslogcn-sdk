package dnslog

import (
	"fmt"
	"testing"
)

func Test_GetDomain(t *testing.T) {
	d, c := GetDomain()
	fmt.Printf("domain=%s\n", d)
	fmt.Printf("cookie=%s\n", c)

}

func Test_GetRecord(t *testing.T) {
	session := `PHPSESSID=2c531vlond0nc0ogq07r37u8l6`

	rs := GetRecords(session)
	fmt.Println(rs)
}
