// Package traefik_extraheaders - forward the TCP Source Port of the Client to a Service, as well as the HTTP Protocol Version
package traefik_extraheaders

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"text/template"
)

type Config struct {
	ClientPortHeader string `json:"clientPortHeader,omitempty"`
	HTTPVerHeader    string `json:"httpVerHeader,omitempty"`
}

var (
	Logger = log.New(os.Stdout, "extraheaders: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func CreateConfig() *Config {
	return &Config{
		ClientPortHeader: "X-Forwarded-Clientport",
		HTTPVerHeader:    "X-Forwarded-HTTP-Ver",
	}
}

type extraheaders struct {
	next             http.Handler
	clientPortHeader string
	httpVerHeader    string
	name             string
	template         *template.Template
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	Logger.Println("extraheaders initialized - MRAM")

	return &extraheaders{
		clientPortHeader: config.ClientPortHeader,
		httpVerHeader:    config.HTTPVerHeader,
		next:             next,
		name:             name,
		template:         template.New("demo").Delims("[[", "]]"),
	}, nil
}

func (a *extraheaders) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if _, clientPort, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		req.Header.Set(a.clientPortHeader, clientPort)
	}
	req.Header.Set(a.httpVerHeader, req.Proto)
	a.next.ServeHTTP(rw, req)
}
