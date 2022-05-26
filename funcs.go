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

package apidoc

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/russross/blackfriday/v2"

	"github.com/zc2638/apidoc/swag"
)

func toHTML(s string) template.HTML {
	return template.HTML(s)
}

func mdToHTML(s string) template.HTML {
	desc := blackfriday.Run([]byte(s))
	return template.HTML(desc)
}

func checkTag(e *swag.Endpoint, tag string) bool {
	for _, v := range e.Tags {
		if tag == v {
			return true
		}
	}
	return false
}

func getEndpointSet(es *swag.Endpoints) []*swag.Endpoint {
	set := make([]*swag.Endpoint, 0, 9)
	if es.Get != nil {
		es.Get.Method = http.MethodGet
		set = append(set, es.Get)
	}
	if es.Post != nil {
		es.Post.Method = http.MethodPost
		set = append(set, es.Post)
	}
	if es.Delete != nil {
		es.Delete.Method = http.MethodDelete
		set = append(set, es.Delete)
	}
	if es.Put != nil {
		es.Put.Method = http.MethodPut
		set = append(set, es.Put)
	}
	if es.Patch != nil {
		es.Patch.Method = http.MethodPatch
		set = append(set, es.Patch)
	}
	if es.Options != nil {
		es.Options.Method = http.MethodOptions
		set = append(set, es.Options)
	}
	if es.Head != nil {
		es.Head.Method = http.MethodHead
		set = append(set, es.Head)
	}
	if es.Connect != nil {
		es.Connect.Method = http.MethodConnect
		set = append(set, es.Connect)
	}
	if es.Trace != nil {
		es.Trace.Method = http.MethodTrace
		set = append(set, es.Trace)
	}
	return set
}

func getParameters(e *swag.Endpoint) []swag.Parameter {
	parameters := make([]swag.Parameter, 0, len(e.Parameters))
	for _, p := range e.Parameters {
		if p.In == "body" {
			continue
		}
		parameters = append(parameters, p)
	}
	return parameters
}

func getBody(api *swag.API, e *swag.Endpoint) string {
	for _, p := range e.Parameters {
		if p.In != "body" {
			continue
		}
		if p.Schema == nil {
			return ""
		}

		body := api.GetObject(p.Schema)
		data, err := json.MarshalIndent(body, "", "    ")
		if err != nil {
			return ""
		}
		return string(data)
	}
	return ""
}

func getResponse(api *swag.API, res *swag.Response) string {
	if res.Schema == nil {
		return ""
	}
	body := api.GetObject(res.Schema)
	data, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		return ""
	}
	return string(data)
}

func getBodyRows(api *swag.API, e *swag.Endpoint) []swag.Row {
	for _, p := range e.Parameters {
		if p.In != "body" {
			continue
		}
		return api.GetRows(p.Schema)
	}
	return nil
}

func getResponseRows(api *swag.API, res *swag.Response) []swag.Row {
	return api.GetRows(res.Schema)
}
