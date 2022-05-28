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

package main

import (
	"log"
	"os"

	"github.com/zc2638/apidoc"
)

func main() {
	fileData, err := os.ReadFile("testdata/swagger.json")
	if err != nil {
		log.Fatal(err)
	}
	data, err := apidoc.Parse(fileData)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll("dist", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("dist/swagger.html", data, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := apidoc.SaveToPDFFile(data, false, "dist/swagger.pdf"); err != nil {
		log.Fatal(err)
	}
}
