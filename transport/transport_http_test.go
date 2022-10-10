package transport

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/log"
	"github.com/ong-gtp/play-service/service"
)

type serviceRequest struct {
	method, url, body string
	want              int
}

func TestHTTP(t *testing.T) {
	t.Setenv("RANDOM_CHOICE_URL", "https://codechallenge.boohma.com/random")
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "env", "testing", "caller", log.DefaultCaller)

	s := service.NewService()
	r := NewHttpServer(s, logger)
	srv := httptest.NewServer(r)

	srs := []serviceRequest{
		{method: "GET", url: srv.URL + "/playsv/health", body: "", want: http.StatusOK},
		{method: "POST", url: srv.URL + "/play", body: `{"player": 1}`, want: http.StatusOK},
		{method: "POST", url: srv.URL + "/play", body: `{"player": "1"}`, want: http.StatusBadRequest},
		{method: "POST", url: srv.URL + "/play", body: `{"player": 1, "opponent":5}`, want: http.StatusOK},
		{method: "GET", url: srv.URL + "/play", body: `{"player": 1}`, want: http.StatusMethodNotAllowed},
		{method: "GET", url: srv.URL + "/play", body: `{"player": 1, "opponent":5}`, want: http.StatusMethodNotAllowed},
	}

	for _, testcase := range srs {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Errorf("error : %s resp: %s", err, resp.Body)
		}
		if testcase.want != resp.StatusCode {
			t.Errorf("%s %s: want %d have %d", testcase.method, testcase.url, testcase.want, resp.StatusCode)
			t.Errorf("%s %s: want %d have %d", testcase.method, testcase.url, testcase.want, resp.StatusCode)
		}
	}
}
