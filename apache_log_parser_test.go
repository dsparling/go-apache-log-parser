package apachelogparser

import (
	"log"
	"testing"
	"time"
)

type TestLine struct {
	RemoteHost string
	Time       time.Time
	Request    string
	Status     int
	Bytes      int
	Referer    string
	UserAgent  string
	Url        string
}

var value = "05/Oct/2014:04:34:35 -0500"
var layout = "02/Jan/2006:15:04:05 -0700"
var t, _ = time.Parse(layout, value)
var testLine = TestLine{"127.0.0.1", t, "GET /test.html HTTP/1.1", 200, 6776, "http://www.example.com/index.html", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:32.0) Gecko/20100101 Firefox/32.0", "/test.html"}

func TestParse(t *testing.T) {
	lines, err := Parse("test_log")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		if testLine.RemoteHost != line.RemoteHost {
			t.Errorf("remote host [%v]; want [%v]", testLine.RemoteHost, line.RemoteHost)
		}
		if testLine.Time != line.Time {
			t.Errorf("time [%v]; want [%v]", testLine.Time, line.Time)
		}
		if testLine.Request != line.Request {
			t.Errorf("request [%v]; want [%v]", testLine.Request, line.Request)
		}
		if testLine.Status != line.Status {
			t.Errorf("status [%v]; want [%v]", testLine.Status, line.Status)
		}
		if testLine.Bytes != line.Bytes {
			t.Errorf("bytes [%v]; want [%v]", testLine.Bytes, line.Bytes)
		}
		if testLine.Referer != line.Referer {
			t.Errorf("referer [%v]; want [%v]", testLine.Referer, line.Referer)
		}
		if testLine.UserAgent != line.UserAgent {
			t.Errorf("user agent [%v]; want [%v]", testLine.UserAgent, line.UserAgent)
		}
		if testLine.Url != line.Url {
			t.Errorf("url [%v]; want [%v]", testLine.Url, line.Url)
		}
	}
}
