package ping

import (
	"bytes"
	"html/template"
	"regexp"
	"strconv"
	"time"
)

// Protocol ...
type Protocol int

const (
	// TCP is tcp protocol
	TCP Protocol = iota
	// HTTP is http protocol
	HTTP
	// HTTPS is https protocol
	HTTPS
)

func (protocol Protocol) String() string {
	switch protocol {
	case TCP:
		return "tcp"
	case HTTP:
		return "http"
	case HTTPS:
		return "https"
	}
	return "unkown"
}

// Ping is a ping interface
type Ping interface {
	Start() <-chan struct{}

	Host() string
	Port() int
	Protocol() Protocol
	Counter() int

	Stop()

	Result() Result
}

// Result ...
type Result struct {
	Counter        int
	SuccessCounter int
	Pinger         Ping

	MinDuration   time.Duration
	MaxDuration   time.Duration
	TotalDuration time.Duration
}

// Avg return the average time of ping
func (result Result) Avg() time.Duration {
	if result.SuccessCounter == 0 {
		return 0
	}
	return result.TotalDuration / time.Duration(result.SuccessCounter)
}

// Failed return failed counter
func (result Result) Failed() int {
	return result.Counter - result.SuccessCounter
}

func (result Result) String() string {
	const resultTpl = `
Ping statistics {{.Pinger.Protocol}}://{{.Pinger.Host}}:{{.Pinger.Port}}
	{{.Counter}} probes sent.
	{{.SuccessCounter}} successful, {{.Failed}} failed.
Approximate trip times:
	Minimum = {{.MinDuration}}, Maximum = {{.MaxDuration}}, Average = {{.Avg}}`
	t := template.Must(template.New("result").Parse(resultTpl))
	res := bytes.NewBufferString("")
	t.Execute(res, result)
	return res.String()
}

// CheckURI check uri
func CheckURI(uri string) (schema, host string, port int, matched bool) {
	const reExp = `^((?P<schema>((ht|f)tp(s?))|tcp)\://)?((([a-zA-Z0-9_\-]{2,}\.)+[a-zA-Z]{2,})|((?:(?:25[0-5]|2[0-4]\d|[01]\d\d|\d?\d)((\.?\d)\.)){4})|(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9]))(:([0-9]+))?(/[a-zA-Z0-9\-\._\?\,\'/\\\+&amp;%\$#\=~]*)?$`
	pattern := regexp.MustCompile(reExp)
	res := pattern.FindStringSubmatch(uri)
	if len(res) == 0 {
		return
	}
	matched = true
	schema = res[2]
	if schema == "" {
		schema = "tcp"
	}
	host = res[6]
	if res[17] == "" {
		if schema == HTTPS.String() {
			port = 443
		} else {
			port = 80
		}
	} else {
		port, _ = strconv.Atoi(res[17])
	}

	return
}