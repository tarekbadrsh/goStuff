package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCompressHandler(t *testing.T) {
	tt := []struct {
		name  string
		value string
		res   string
		err   string
	}{
		{name: "two", value: "2", res: "C"},
		{name: "from 1 to 9", value: "123456879", res: "v18WH"},
		{name: "missing value", value: "", err: "missing value"},
		{name: "not a number", value: "x", err: "not a number: x"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:5050/compress?v="+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()
			compressHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status Bad Request; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}

			if string(bytes.TrimSpace(b)) != tc.res {
				t.Fatalf("expected %v; got %v", tc.res, tc.res)
			}
		})
	}
}

func TestUncompressHandler(t *testing.T) {
	tt := []struct {
		name  string
		value string
		res   int64
		err   string
	}{
		{name: "two", value: "C", res: 2},
		{name: "test v18WH", value: "v18WH", res: 123456879},
		{name: "missing value", value: "", err: "missing value"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:5050/uncompress?v="+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()
			uncompressHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status Bad Request; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}

			d, err := strconv.ParseInt(string(bytes.TrimSpace(b)), 10, 64)
			if err != nil {
				t.Fatalf("expected an integer; got %s", b)
			}
			if d != tc.res {
				t.Fatalf("expected %v; got %v", tc.res, tc.res)
			}
		})
	}
}

func TestRoutingCompress(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/compress?v=2", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	if string(bytes.TrimSpace(b)) != "C" {
		t.Fatalf("expected double to be 4; got %s", b)
	}
}

func TestRoutingUncompress(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/uncompress?v=C", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
	if err != nil {
		t.Fatalf("expected an integer; got %s", b)
	}
	if d != 2 {
		t.Fatalf("expected double to be 2; got %v", d)
	}
}
