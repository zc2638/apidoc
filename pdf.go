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
	"os"
	"strings"

	pdf "github.com/adrg/go-wkhtmltopdf"
)

func SaveToPDF(data []byte, to string) error {
	if err := pdf.Init(); err != nil {
		return err
	}
	defer pdf.Destroy()

	object, err := pdf.NewObjectFromReader(bytes.NewReader(data))
	if err != nil {
		return err
	}

	converter, err := pdf.NewConverter()
	if err != nil {
		return err
	}
	defer converter.Destroy()

	converter.Add(object)

	converter.PaperSize = pdf.A4
	converter.Orientation = pdf.Landscape
	converter.MarginTop = "1cm"
	converter.MarginBottom = "1cm"
	converter.MarginLeft = "10mm"
	converter.MarginRight = "10mm"

	to = strings.TrimSuffix(to, ".pdf") + ".pdf"
	outFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return converter.Run(outFile)
}
