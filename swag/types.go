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
	"path"
	"strconv"
)

type ParameterType string

func (pt ParameterType) String() string {
	return string(pt)
}

const (
	Integer ParameterType = "integer"
	Number  ParameterType = "number"
	Boolean ParameterType = "boolean"
	String  ParameterType = "string"

	Array  ParameterType = "array"
	Object ParameterType = "object"

	File ParameterType = "file"
)

func ConvertSchemaToMap(schemas map[string]*Schema) map[string]interface{} {
	// 解析为 ref => obj
	set := make(map[string]interface{})
	extra := make(map[string]*Schema)
	for name, s := range schemas {
		extra[name] = s
	}
	for len(extra) > 0 {
		empty := make(map[string]*Schema)
		for name, s := range extra {
			ref := path.Join("#/definitions", name)
			if s.Ref != "" {
				empty[name] = s
				continue
			}
			obj, ok := ConvertSchemaToValue(set, s)
			if !ok {
				empty[name] = s
				continue
			}
			set[ref] = obj
		}
		extra = empty
	}
	return set
}

func ConvertSchemaToValue(set map[string]interface{}, schema *Schema) (interface{}, bool) {
	if schema.Ref != "" {
		if obj, ok := set[schema.Ref]; ok {
			return obj, true
		}
		return nil, false
	}
	var value interface{}
	switch schema.Type {
	case String:
		value = schema.Example
	case Integer:
		value, _ = strconv.Atoi(schema.Example)
	case Number:
		value, _ = strconv.ParseFloat(schema.Example, 64)
	case Boolean:
		value, _ = strconv.ParseBool(schema.Example)
	case Array:
		obj, ok := ConvertSchemaToValue(set, schema.Items)
		if !ok {
			return nil, false
		}
		value = []interface{}{obj}
	case Object:
		objs := make(map[string]interface{})
		for name, s := range schema.Properties {
			obj, ok := ConvertSchemaToValue(set, s)
			if !ok {
				return nil, false
			}
			objs[name] = obj
		}
		value = objs
	}
	return value, true
}
