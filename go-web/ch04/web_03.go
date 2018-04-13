package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

var _ string = `
type Cookie struct {
    Name  string
    Value string

    Path       string    // optional
    Domain     string    // optional
    Expires    time.Time // optional
    RawExpires string    // for reading cookies only

    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int
    Secure   bool
    HttpOnly bool
    Raw      string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}
    A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an
    HTTP response or the Cookie header of an HTTP request.

    See http://tools.ietf.org/html/rfc6265 for details.


func (c *Cookie) String() string
`

func cookieHandler(rw http.ResponseWriter, req *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "first_cookie",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "second_cookie",
		HttpOnly: true,
	}

	rw.Header().Set("Set-Cookie", c1.String())
	rw.Header().Add("Set-Cookie", c2.String())

	http.SetCookie(rw, &c1)
	http.SetCookie(rw, &c2)
}

func getCookieHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, req.Header["cookie"])

	c1, err := req.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(rw, "can not get first cookie")
	}
	cs := req.Cookies()
	fmt.Fprintln(rw, c1)
	fmt.Fprintln(rw, cs)
}

func setMessage(rw http.ResponseWriter, req *http.Request) {
	msg := []byte("hello night")
	c := http.Cookie{
		Name:  "msg",
		Value: base64.URLEncoding.EncodeToString(msg),
	}

	http.SetCookie(rw, &c)
}

func showMessage(rw http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("msg")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(rw, "no cookie found")
		}
	} else {
		// expire immediately
		rc := http.Cookie{
			Name:    "msg",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}

		http.SetCookie(rw, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(rw, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/set_cookie", cookieHandler)
	http.HandleFunc("/get_cookie", getCookieHandler)
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
