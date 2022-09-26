package plugin

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"
)

func TestRecords(t *testing.T) {
	// instantiate a test http server
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// return a json as response
		w.Write([]byte(`[{
			"dnsName" : "test.example.com"
		}]`))
	}))
	defer svr.Close()

	provider := NewPluginProvider(svr.URL)
	endpoints, err := provider.Records(context.TODO())
	require.Nil(t, err)
	require.NotNil(t, endpoints)
	require.Equal(t, []*endpoint.Endpoint{&endpoint.Endpoint{
		DNSName: "test.example.com",
	}}, endpoints)
}
