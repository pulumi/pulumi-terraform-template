// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package terraform-template

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Renders a template from a file.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-template/blob/master/website/docs/d/file.html.markdown.
func LookupFile(ctx *pulumi.Context, args *GetFileArgs) (*GetFileResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["filename"] = args.Filename
		inputs["template"] = args.Template
		inputs["vars"] = args.Vars
	}
	outputs, err := ctx.Invoke("terraform-template:index/getFile:getFile", inputs)
	if err != nil {
		return nil, err
	}
	return &GetFileResult{
		Filename: outputs["filename"],
		Rendered: outputs["rendered"],
		Template: outputs["template"],
		Vars: outputs["vars"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getFile.
type GetFileArgs struct {
	// _Deprecated, please use `template` instead_. The filename for
	// the template. Use [path variables](https://www.terraform.io/docs/configuration/interpolation.html#path-variables) to make
	// this path relative to different path roots.
	Filename interface{}
	// The contents of the template. These can be loaded
	// from a file on disk using the [`file()` interpolation
	// function](https://www.terraform.io/docs/configuration/interpolation.html#file_path_).
	Template interface{}
	// Variables for interpolation within the template. Note
	// that variables must all be primitives. Direct references to lists or maps
	// will cause a validation error.
	Vars interface{}
}

// A collection of values returned by getFile.
type GetFileResult struct {
	Filename interface{}
	// The final rendered template.
	Rendered interface{}
	// See Argument Reference above.
	Template interface{}
	// See Argument Reference above.
	Vars interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
