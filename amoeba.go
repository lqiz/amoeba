package amoeba

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

var (
	PanicHandler   func(interface{})
	UpdateInterval int
)

const (
	amoeba   = "amoeba"
	pathReq  = "%s_req"
	pathResp = "%s_resp"
)

func ReverseProxy(c *gin.Context) {
	remote, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Hostname()
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
		req.RequestURI = req.URL.Path
	}

	proxy.Transport = &transport{http.DefaultTransport}
	proxy.ServeHTTP(c.Writer, c.Request)
}

type transport struct {
	http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	pathKey := req.Header.Get(amoeba)
	if len(pathKey) <= 0 {
		return nil, errors.New("amoeba not set")
	}

	errReq := requestHandler(pathKey, req)
	if errReq != nil {
		// TODO ERROR
	}
	resp, err = t.RoundTripper.RoundTrip(req)
	errResp := responseHandler(pathKey, resp)
	if errResp != nil {
		// TODO ERROR
	}
	return
}

func requestHandler(pathKey string, req *http.Request) error {
	body := req.Body
	if body == nil {
		return nil
	}
	bodyByteArray, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	schema, ok := schemaMap[fmt.Sprintf(pathReq, pathKey)]
	if !ok {
		return errors.New("no request schema found")
	}
	result, err := Marshal(schema, string(bodyByteArray))
	if err != nil {
		// TODO ERROR
	}
	req.ContentLength = int64(len(result))
	req.Body = ioutil.NopCloser(bytes.NewReader(result))
	return nil
}

func responseHandler(pathKey string, resp *http.Response) error {
	bRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	schema, ok := schemaMap[fmt.Sprintf(pathResp, pathKey)]
	if !ok {
		return errors.New("no response schema found")
	}

	bNew, err := Marshal(schema, string(bRaw))
	if err != nil {
		// TODO ERROR
	}

	body := ioutil.NopCloser(bytes.NewReader(bNew))
	resp.Body = body
	resp.ContentLength = int64(len(bRaw))
	resp.Header.Set("Content-Length", strconv.Itoa(len(bNew)))
	return nil
}
