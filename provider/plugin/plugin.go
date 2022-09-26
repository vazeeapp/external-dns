package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type PluginProvider struct {
	provider.BaseProvider
	client       *http.Client
	remoteServer string
}

func NewPluginProvider(url string) *PluginProvider {
	return &PluginProvider{
		client:       &http.Client{},
		remoteServer: url,
	}
}

// Records will make a GET call to /records
func (p PluginProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	req, err := http.NewRequest("GET", p.remoteServer, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))

	endpoints := []*endpoint.Endpoint{}
	err = json.Unmarshal(b, &endpoints)
	if err != nil {
		return nil, err
	}
	return endpoints, nil
}

// ApplyChanges will make a POST to /records
func (p PluginProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return nil
}
