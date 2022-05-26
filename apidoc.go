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
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	dittoJson "github.com/99nil/ditto/json"
	"gopkg.in/yaml.v2"

	"github.com/zc2638/apidoc/resource"
	"github.com/zc2638/apidoc/swag"
)

var defaultTemplate = template.New("default").
	Funcs(template.FuncMap{
		"toLower":         strings.ToLower,
		"toHTML":          toHTML,
		"mdToHTML":        mdToHTML,
		"checkTag":        checkTag,
		"getEndpointSet":  getEndpointSet,
		"getParameters":   getParameters,
		"getBody":         getBody,
		"getResponse":     getResponse,
		"getBodyRows":     getBodyRows,
		"getResponseRows": getResponseRows,
	})

func Parse(content []byte) ([]byte, error) {
	var obj swag.API

	line, pos, ok := dittoJson.CheckBytes(content)
	if !ok {
		jsonErrStr := fmt.Sprintf("\njson check failed, line: %d, pos: %d", line, pos)
		if err := yaml.Unmarshal(content, &obj); err != nil {
			return nil, fmt.Errorf("%s\nyaml parse failed: %v", jsonErrStr, err)
		}
	} else {
		if err := json.Unmarshal(content, &obj); err != nil {
			return nil, fmt.Errorf("json parse failed: %v", err)
		}
	}
	obj.TransformSchemas()

	tpl, err := resource.ReadTemplate("default")
	if err != nil {
		return nil, err
	}
	t, err := defaultTemplate.Parse(string(tpl))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, &obj); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ParseFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, expect %d, actual %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	var body bytes.Buffer
	if _, err := io.Copy(&body, resp.Body); err != nil {
		return nil, err
	}
	return Parse(body.Bytes())
}
