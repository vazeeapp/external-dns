package plugin

import (
	"context"
	"net/http"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type PluginProvider struct {
	client       *http.Client
	remoteServer string
}

func NewPluginProvider() *PluginProvider {
	return &PluginProvider{
		client: &http.Client{},
	}
}

// Records will make a GET call to /records
func (p PluginProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	p.client.Get(p.)
}

// ApplyChanges will make a POST to /records
func (p PluginProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {

}

// PropertyValuesEqual makes a GET to /propertyvaluesequal 
func (p PluginProvider) PropertyValuesEqual(name string, previous string, current string) bool {
}

// AdjustEndpoints makes a PUT to /endpoints
func (p PluginProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
}

// TODO how to map that?
func (p PluginProvider) GetDomainFilter() endpoint.DomainFilterInterface {
}
