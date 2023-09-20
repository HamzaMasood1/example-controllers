// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/HamzaMasood1/example-controllers/basic-controller/github.com/myid/myresource-crd/pkg/apis/mygroup.example.com/v1alpha1"
	"github.com/HamzaMasood1/example-controllers/basic-controller/github.com/myid/myresource-crd/pkg/clientset/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type MygroupV1alpha1Interface interface {
	RESTClient() rest.Interface
	MyResourcesGetter
}

// MygroupV1alpha1Client is used to interact with features provided by the mygroup.example.com group.
type MygroupV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MygroupV1alpha1Client) MyResources(namespace string) MyResourceInterface {
	return newMyResources(c, namespace)
}

// NewForConfig creates a new MygroupV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*MygroupV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new MygroupV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*MygroupV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &MygroupV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MygroupV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MygroupV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MygroupV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MygroupV1alpha1Client {
	return &MygroupV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MygroupV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
