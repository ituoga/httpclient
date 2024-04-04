package httpclient_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ituoga/httpclient"
)

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"hello"}`))
	}))
	defer srv.Close()

	type Response struct {
		Message string `json:"message"`
	}

	r, err := httpclient.Get[Response](srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	if r.Message != "hello" {
		t.Fatalf("unexpected response: %v", r)
	}
}

func TestPost(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.Copy(w, r.Body)
	}))
	defer srv.Close()

	type Response struct {
		Message string `json:"message"`
	}

	r, err := httpclient.Post[Response](srv.URL, "application/json", io.NopCloser(strings.NewReader(`{"message":"hello"}`)))
	if err != nil {
		t.Fatal(err)
	}
	if r.Message != "hello" {
		t.Fatalf("unexpected response: %v", r)
	}
}

func TestError(t *testing.T) {
	type Response struct {
		Message string `json:"message"`
	}

	_, err := httpclient.Get[Response]("http://asdasd")
	if err == nil {
		t.Fatal("expected an error")
	}

	_, err = httpclient.Post[Response]("http://asdasd", "application/json", io.NopCloser(strings.NewReader(`{"message":"hello"}`)))
	if err == nil {
		t.Fatal("expected an error")
	}

	_, err = httpclient.Post[Response]("http://asdasd", "application/json", io.NopCloser(strings.NewReader(`{"message:"hello"}`)))
	if err == nil {
		t.Fatal("expected an error")
	}
}
