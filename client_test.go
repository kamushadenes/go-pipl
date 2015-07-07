package pipl

import (
	"os"
	"testing"
	"time"
)

var client *V4Client

func init() {
	key := os.Getenv("PIPL_KEY")
	client = NewV4Client(key)
}

func TestKey(t *testing.T) {
	if client.ApiKey == "" {
		t.Error("api key not set, you must set the environment variable PIPL_KEY before running")
		t.FailNow()
	}
}

func TestSearchPerson(t *testing.T) {
	timeout := time.Duration(20 * time.Second)
	p := Person{}
	p.Phones = append(p.Phones, Phone{Raw: "2069735100"})
	_, err := client.SearchPerson(p, timeout, make(map[string]string))
	if err != nil {
		t.Error(err)
	}
}
