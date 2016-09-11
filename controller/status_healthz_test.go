package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestStatusHealthzGet(t *testing.T) {
	type (
		in struct {
			method, path string
		}
		out struct {
			body       string
			statusCode int
		}
	)

	tests := map[string]struct {
		in
		out
	}{
		"case-01": {
			in{"GET", "/_status/healthz"},
			out{"", 204},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, out := test.in, test.out

			ctrl := &StatusHealthzCtrl{}

			ps := httprouter.Params{}
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(in.method, in.path, nil)
			ctrl.Get(w, r, ps)

			if body := w.Body.String(); body != out.body {
				t.Errorf("actual body %s, expected body %s", body, out.body)
			}
			if statusCode := w.Code; statusCode != out.statusCode {
				t.Errorf("actual status code %d, expected status code %d", statusCode, out.statusCode)
			}
		})
	}
}
