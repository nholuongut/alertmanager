// Code generated by go-swagger; DO NOT EDIT.

// Copyright Nho luong DevOps
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/nholuongut/alertmanager/api/v2/client/alert"
	"github.com/nholuongut/alertmanager/api/v2/client/alertgroup"
	"github.com/nholuongut/alertmanager/api/v2/client/general"
	"github.com/nholuongut/alertmanager/api/v2/client/receiver"
	"github.com/nholuongut/alertmanager/api/v2/client/silence"
)

// Default alertmanager API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v2/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new alertmanager API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *AlertmanagerAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new alertmanager API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *AlertmanagerAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new alertmanager API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *AlertmanagerAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(AlertmanagerAPI)
	cli.Transport = transport
	cli.Alert = alert.New(transport, formats)
	cli.Alertgroup = alertgroup.New(transport, formats)
	cli.General = general.New(transport, formats)
	cli.Receiver = receiver.New(transport, formats)
	cli.Silence = silence.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// AlertmanagerAPI is a client for alertmanager API
type AlertmanagerAPI struct {
	Alert alert.ClientService

	Alertgroup alertgroup.ClientService

	General general.ClientService

	Receiver receiver.ClientService

	Silence silence.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *AlertmanagerAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Alert.SetTransport(transport)
	c.Alertgroup.SetTransport(transport)
	c.General.SetTransport(transport)
	c.Receiver.SetTransport(transport)
	c.Silence.SetTransport(transport)
}
