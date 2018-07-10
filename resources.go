// Copyright 2016-2018, Pulumi Corporation.
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

package template

import (
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/tokens"
	tftemplate "github.com/terraform-providers/terraform-provider-template/template"
)

// all of the template token components used below.
const (
	// packages:
	templatePkg = "terraform-template"

	// modules:
	templateMod = "index"
)

// templateMember manufactures a member token for the template package and the given member name.
func templateMember(mod, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(templatePkg + ":" + mod + ":" + mem)
}

// templateType manufactures a type token for the template package and the given module and type.
func templateType(mod, typ string) tokens.Type {
	return tokens.Type(templateMember(mod, typ))
}

// templateDataSource manufactures a standard resource token given a resource name. It automatically uses the template
// package and names the file by simply lower casing the data source's first character.
func templateDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return templateMember(mod+"/"+fn, res)
}

// templateResource manufactures a standard resource token given a module and resource name. It automatically uses the template
// package and names the file by simply lower casing the resource's first character.
func templateResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return templateType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// Provider returns additional overlaid schema and metadata associated with the template package.
func Provider() tfbridge.ProviderInfo {
	p := tftemplate.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "template",
		Description: "A Pulumi package for creating and managing Terraform template resources.",
		Keywords:    []string{"pulumi", "terraform", "template"},
		Homepage:    "https://pulumi.io/terraform-template",
		Repository:  "https://github.com/pulumi/pulumi-terraform-template",
		Resources: map[string]*tfbridge.ResourceInfo{
			"template_file": {Tok: templateResource(templateMod, "File")},
			"template_cloudinit_config": {Tok: templateResource(templateMod, "CloudInitConfig")},
			"template_dir": {Tok: templateResource(templateMod, "Dir")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"template_file": {Tok: templateDataSource(templateMod, "getFile")},
			"template_cloudinit_config": {Tok: templateDataSource(templateMod, "getCloudInitConfig")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^0.14.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=0.14.2,<0.15.0",
			},
		},
	}

	return prov
}
