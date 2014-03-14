package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"strings"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	match := flag.String("match", ".clouddrive.com:443", "Only allow requests destined to this suffix match")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose
	proxy.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		if strings.HasSuffix(host, *match) {
			log.Printf("Allowing - %s %s %s", ctx.Req.RemoteAddr, ctx.Req.Method, ctx.Req.RequestURI)
			return goproxy.OkConnect, host
		} else {
			log.Printf("Rejecting - %s %s %s", ctx.Req.RemoteAddr, ctx.Req.Method, ctx.Req.RequestURI)
			return goproxy.RejectConnect, host
		}

	})
	proxy.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if strings.HasSuffix(r.URL.Host, *match) {
			log.Printf("Allowing - %s %s %s", ctx.Req.RemoteAddr, ctx.Req.Method, ctx.Req.RequestURI)
			return r, nil
		} else {
			// Note that unless the request container :443 we'll reject it!
			log.Printf("Rejecting - %s %s %s", ctx.Req.RemoteAddr, ctx.Req.Method, ctx.Req.RequestURI)
			return r, goproxy.NewResponse(r, goproxy.ContentTypeText,
				http.StatusForbidden, "Endpoint not allowed.")
		}
	})
	log.Printf("container-sync proxy starting up on '%v'\n", *addr)
	log.Printf("endpoints restricted to: '%v'\n", *match)
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
