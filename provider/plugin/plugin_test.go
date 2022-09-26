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
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

func TestApplyChanges(t *testing.T) {
	successfulApplyChanges := true
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if successfulApplyChanges {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer svr.Close()

	provider := NewPluginProvider(svr.URL)
	err := provider.ApplyChanges(context.TODO(), nil)
	require.Nil(t, err)

	successfulApplyChanges = false

	err = provider.ApplyChanges(context.TODO(), nil)
	require.NotNil(t, err)
}
