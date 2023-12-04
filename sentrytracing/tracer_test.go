package sentrytracing_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler/testserver"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/benevolent0505/gqlgen-tracing-sentry/sentrytracing"
)

func Test_NoError(t *testing.T) {
	mux := http.NewServeMux()

	srv := testserver.New()
	srv.AddTransport(transport.POST{})
	srv.Use(sentrytracing.Tracer{})

	mux.Handle("/query", srv)

	resp := doRequest(mux, http.MethodPost, "/query", `{ "query": "{ name find(id: 1) }"}`)
	assert.Equal(t, http.StatusOK, resp.Code, resp.Body.String())

	var respData struct {
		Errors gqlerror.List
	}
	require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &respData))
}

func doRequest(handler http.Handler, method, target, body string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)
	return w
}
