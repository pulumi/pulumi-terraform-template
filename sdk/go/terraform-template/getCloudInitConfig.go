// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package terraform-template

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Renders a multi-part cloud-init config from source files.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-template/blob/master/website/docs/d/cloudinit_config.html.markdown.
func LookupCloudInitConfig(ctx *pulumi.Context, args *GetCloudInitConfigArgs) (*GetCloudInitConfigResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["base64Encode"] = args.Base64Encode
		inputs["gzip"] = args.Gzip
		inputs["parts"] = args.Parts
	}
	outputs, err := ctx.Invoke("terraform-template:index/getCloudInitConfig:getCloudInitConfig", inputs)
	if err != nil {
		return nil, err
	}
	return &GetCloudInitConfigResult{
		Base64Encode: outputs["base64Encode"],
		Gzip: outputs["gzip"],
		Parts: outputs["parts"],
		Rendered: outputs["rendered"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getCloudInitConfig.
type GetCloudInitConfigArgs struct {
	// Base64 encoding of the rendered output. Default to `true`
	Base64Encode interface{}
	// Specify whether or not to gzip the rendered output. Default to `true`
	Gzip interface{}
	// One may specify this many times, this creates a fragment of the rendered cloud-init config file. The order of the parts is maintained in the configuration is maintained in the rendered template.
	Parts interface{}
}

// A collection of values returned by getCloudInitConfig.
type GetCloudInitConfigResult struct {
	Base64Encode interface{}
	Gzip interface{}
	Parts interface{}
	// The final rendered multi-part cloudinit config.
	Rendered interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
