// Copyright © 2022 zc2638 <zc2638@qq.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swag

import (
	"bytes"
	"encoding/json"

	"gopkg.in/yaml.v2"
)

// API is the root document object for the specification.
// It combines what previously was the Resource Listing
// and API Declaration (version 1.2 and earlier) together into one document.
type API struct {
	set    map[string]interface{}
	rowSet map[string][]Row

	// Required. Specifies the Swagger Specification version being used.
	// It can be used by the Swagger UI and other clients to interpret the API listing.
	// The value MUST be "2.0".
	Swagger string `json:"swagger,omitempty"`

	// Required. Provides metadata about the API.
	// The metadata can be used by the clients if needed.
	Info Info `json:"info"`

	// The host (name or ip) serving the API.
	// This MUST be the host only and does not include the scheme nor sub-paths.
	// It MAY include a port.
	// If the host is not included, the host serving the documentation is to be used (including the port).
	// The host does not support path templating.
	Host string `json:"host,omitempty"`

	// The base path on which the API is served, which is relative to the host.
	// If it is not included, the API is served directly under the host.
	// The value MUST start with a leading slash (/).
	// The basePath does not support path templating.
	BasePath string `json:"basePath,omitempty"`

	// The transfer protocol of the API.
	// Values MUST be from the list: "http", "https", "ws", "wss".
	// If the schemes is not included, the default scheme to be used is the one used to access the Swagger definition itself.
	Schemes []string `json:"schemes,omitempty"`

	// Required. The available paths and operations for the API.
	Paths map[string]*Endpoints `json:"paths,omitempty"`

	// An object to hold data types produced and consumed by operations.
	Definitions map[string]*Schema `json:"definitions,omitempty"`

	// A list of tags used by the specification with additional metadata.
	// The order of the tags can be used to reflect on their order by the parsing tools.
	// Not all tags that are used by the Operation Object must be declared.
	// The tags that are not declared may be organized randomly or based on the tools' logic.
	// Each tag name in the list MUST be unique.
	Tags []Tag `json:"tags,omitempty"`

	// Security scheme definitions that can be used across the specification.
	SecurityDefinitions map[string]*SecurityScheme `json:"securityDefinitions,omitempty"`

	// A declaration of which security schemes are applied for the API as a whole.
	// The list of values describes alternative security schemes that can be used
	// (that is, there is a logical OR between the security requirements).
	// Individual operations can override this definition.
	Security *SecurityRequirement `json:"security,omitempty"`
}

func (a *API) TransformSchemas() {
	a.set = ConvertSchemaToMap(a.Definitions)
	a.rowSet = ConvertSchemaToRowSet(a.Definitions)
}

func (a *API) GetObject(schema *Schema) interface{} {
	if schema == nil {
		return nil
	}
	if obj, ok := ConvertSchemaToValue(a.set, schema); ok {
		return obj
	}
	return nil
}

func (a *API) GetRows(schema *Schema) []Row {
	if schema == nil {
		return nil
	}
	if out, ok := ConvertSchemaToRow(a.rowSet, schema, nil, false); ok {
		current := make([]Row, 0, len(out))
		for _, v := range out {
			if v.Name == "" {
				continue
			}
			current = append(current, v)
		}
		return current
	}
	return nil
}

// Info provides metadata about the API.
// The metadata can be used by the clients if needed, and can be presented in the Swagger-UI for convenience.
type Info struct {
	// Required. The title of the application.
	Title string `json:"title,omitempty"`

	// Required Provides the version of the application API (not to be confused with the specification version).
	Version string `json:"version,omitempty"`

	// A short description of the application. GFM syntax can be used for rich text representation.
	Description string `json:"description,omitempty"`

	// The Terms of Service for the API.
	TermsOfService string `json:"termsOfService,omitempty"`

	// The contact information for the exposed API.
	Contact *Contact `json:"contact,omitempty"`

	// The license information for the exposed API.
	License License `json:"license"`
}

// Contact information for the exposed API.
type Contact struct {
	// The identifying name of the contact person/organization.
	Name string `json:"name,omitempty"`

	// The URL pointing to the contact information. MUST be in the format of a URL.
	URL string `json:"url,omitempty"`

	// The email address of the contact person/organization. MUST be in the format of an email address.
	Email string `json:"email,omitempty"`
}

// License represents the license entity from the swagger definition; used by Info
type License struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Endpoints represents all the swagger endpoints associated with a particular path
type Endpoints struct {
	Delete  *Endpoint `json:"delete,omitempty"`
	Head    *Endpoint `json:"head,omitempty"`
	Get     *Endpoint `json:"get,omitempty"`
	Options *Endpoint `json:"options,omitempty"`
	Post    *Endpoint `json:"post,omitempty"`
	Put     *Endpoint `json:"put,omitempty"`
	Patch   *Endpoint `json:"patch,omitempty"`
	Trace   *Endpoint `json:"trace,omitempty"`
	Connect *Endpoint `json:"connect,omitempty"`
}

// Endpoint represents an endpoint from the swagger doc
type Endpoint struct {
	Tags        []string             `json:"tags,omitempty"`
	Path        string               `json:"-"`
	Method      string               `json:"-"`
	Summary     string               `json:"summary,omitempty"`
	Description string               `json:"description,omitempty"`
	OperationID string               `json:"operationId,omitempty"`
	Produces    []string             `json:"produces,omitempty"`
	Consumes    []string             `json:"consumes,omitempty"`
	Handler     interface{}          `json:"-"`
	Parameters  []Parameter          `json:"parameters,omitempty"`
	Responses   map[string]*Response `json:"responses,omitempty"`

	// swagger spec requires security to be an array of objects
	Security   *SecurityRequirement `json:"security,omitempty"`
	Deprecated bool                 `json:"deprecated,omitempty"`
}

// Parameter represents a parameter from the swagger doc
type Parameter struct {
	Name        string `json:"name,omitempty"`
	In          string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required"`

	Type    ParameterType `json:"type,omitempty"`
	Format  string        `json:"format,omitempty"`
	Default string        `json:"default,omitempty"`

	Schema *Schema `json:"schema,omitempty"`
}

// Schema represents a schema from the swagger doc
type Schema struct {
	Type       ParameterType      `json:"type,omitempty"`
	Required   []string           `json:"required,omitempty"`
	Ref        string             `json:"$ref,omitempty"`
	Properties map[string]*Schema `json:"properties,omitempty"`
	Items      *Schema            `json:"items,omitempty"`

	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Format      string   `json:"format,omitempty"`
	Example     string   `json:"example,omitempty"`
}

// Response represents a response from the swagger doc
type Response struct {
	Description string            `json:"description"`
	Schema      *Schema           `json:"schema,omitempty"`
	Headers     map[string]Header `json:"headers,omitempty"`
}

// Header represents a response header
type Header struct {
	Type        string `json:"type"`
	Format      string `json:"format"`
	Description string `json:"description"`
}

// Tag represents a swagger tag
type Tag struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Docs        *TagDocs `json:"externalDocs,omitempty"`
}

// TagDocs represents tag docs from the swagger definition
type TagDocs struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

// SecurityScheme represents a security scheme from the swagger definition.
type SecurityScheme struct {
	Type             string            `json:"type"`
	Description      string            `json:"description,omitempty"`
	Name             string            `json:"name,omitempty"`
	In               string            `json:"in,omitempty"`
	Flow             string            `json:"flow,omitempty"`
	AuthorizationURL string            `json:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty"`
}

type SecurityRequirement struct {
	Requirements    []map[string][]string
	DisableSecurity bool
}

func (s *SecurityRequirement) MarshalJSON() ([]byte, error) {
	if s.DisableSecurity {
		return []byte("[]"), nil
	}
	return json.Marshal(s.Requirements)
}

func (s *SecurityRequirement) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	if bytes.Equal([]byte("[]"), b) {
		s.DisableSecurity = true
		return nil
	}
	return json.Unmarshal(b, &s.Requirements)
}

func (s *SecurityRequirement) MarshalYAML() (interface{}, error) {
	if s.DisableSecurity {
		return []byte("[]"), nil
	}
	return yaml.Marshal(s.Requirements)
}

func (s *SecurityRequirement) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := unmarshal(&s.Requirements); err != nil {
		return err
	}
	if s.Requirements == nil {
		s.DisableSecurity = true
	}
	return nil
}
