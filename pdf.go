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

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func SaveToPDF(data []byte, isGray bool) ([]byte, error) {
	gen, err := pdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	gen.Dpi.Set(300)
	gen.Grayscale.Set(isGray)

	gen.AddPage(pdf.NewPageReader(bytes.NewReader(data)))
	if err := gen.Create(); err != nil {
		return nil, err
	}
	return gen.Bytes(), nil
}

func SaveToPDFFile(data []byte, isGray bool, to string) error {
	pdfData, err := SaveToPDF(data, isGray)
	if err != nil {
		return err
	}

	to = strings.TrimSuffix(to, ".pdf") + ".pdf"
	outFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = outFile.Write(pdfData)
	return err
}
