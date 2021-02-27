/*
Copyright 2015-2021 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plugin

import "github.com/gravitational/trace"

// Plugin describes interfaces of the teleport core plugin
type Plugin interface {
	// GetName returns plugin name
	GetName() string
	// RegisterProxyWebHandlers registers new methods on the proxy web handler
	RegisterProxyWebHandlers(handler interface{}) error
	// RegisterAuthWebHandlers registers new methods on the auth web handlers
	RegisterAuthWebHandlers(service interface{}) error
	// RegisterAuthServices registers new services on the auth GRPC server
	RegisterAuthServices(server interface{}) error
}

// Registry is the plugin registry
type Registry struct {
	plugins []Plugin
}

// Add adds plugin to the plugin registry
func (r *Registry) Add(p Plugin) error {
	if p == nil {
		return trace.BadParameter("missing plugin")
	}

	r.plugins = append(r.plugins, p)

	return nil
}

// RegisterProxyWebHandlers registers additional Proxy Web Handlers
func (r *Registry) RegisterProxyWebHandlers(hander interface{}) error {
	for _, p := range r.plugins {
		if err := p.RegisterProxyWebHandlers(hander); err != nil {
			return trace.Wrap(err, "plugin %v failed to register", p.GetName())
		}
	}

	return nil
}

// RegisterAuthWebHandlers registers additional Auth Web Handlers
func (r *Registry) RegisterAuthWebHandlers(handler interface{}) error {
	for _, p := range r.plugins {
		if err := p.RegisterAuthWebHandlers(handler); err != nil {
			return trace.Wrap(err, "plugin %v failed to register", p.GetName())
		}
	}

	return nil
}

// RegisterAuthServices registerse additional Auth Server Services
func (r *Registry) RegisterAuthServices(server interface{}) error {
	for _, p := range r.plugins {
		if err := p.RegisterAuthServices(server); err != nil {
			return trace.Wrap(err, "plugin %v failed to register", p.GetName())
		}
	}

	return nil
}
