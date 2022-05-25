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

package app

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/zc2638/apidoc"
)

type Option struct {
	Template string
	Src      string
	Dest     string
	Format   string // format, default is pdf.
}

func NewServerCommand() *cobra.Command {
	opt := &Option{}
	cmd := &cobra.Command{
		Use:          "apidoc",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opt.Template != "default" {
				return fmt.Errorf("custom templates will be supported soon")
			}
			if opt.Format != "pdf" {
				return fmt.Errorf("other formats will be supported soon")
			}

			var data []byte
			uri, err := url.Parse(opt.Src)
			if err == nil && uri.Host != "" {
				data, err = apidoc.ParseFromURL(opt.Src)
				if err != nil {
					return err
				}
			} else {
				fileData, err := os.ReadFile(opt.Src)
				if err != nil {
					return fmt.Errorf("read src file failed: %v", err)
				}
				data, err = apidoc.Parse(fileData)
				if err != nil {
					return err
				}
			}

			if err := os.MkdirAll(opt.Dest, os.ModePerm); err != nil {
				return fmt.Errorf("create dest dir failed: %v", err)
			}
			base := filepath.Base(opt.Src)
			ext := filepath.Ext(base)
			target := strings.TrimSuffix(base, ext) + ".pdf"
			if err := apidoc.SaveToPDF(data, filepath.Join(opt.Dest, target)); err != nil {
				return fmt.Errorf("save failed: %v", err)
			}
			return nil
		},
	}
	completionFlags(cmd, opt)
	return cmd
}

func completionFlags(cmd *cobra.Command, opt *Option) {
	cmd.Flags().StringVar(&opt.Template, "template", "default", "Specify the template file, the built-in `default` is used by default")
	cmd.Flags().StringVar(&opt.Format, "format", "pdf", "Specify the output file format, the default is pdf")
	cmd.Flags().StringVar(&opt.Src, "src", "", "Specify the swagger configuration file path")
	cmd.Flags().StringVar(&opt.Dest, "dest", "dist", "Specify output path.")
}
