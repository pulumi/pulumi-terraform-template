// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package terraform-template

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-template/blob/master/website/docs/r/dir.html.markdown.
type Dir struct {
	s *pulumi.ResourceState
}

// NewDir registers a new resource with the given unique name, arguments, and options.
func NewDir(ctx *pulumi.Context,
	name string, args *DirArgs, opts ...pulumi.ResourceOpt) (*Dir, error) {
	if args == nil || args.DestinationDir == nil {
		return nil, errors.New("missing required argument 'DestinationDir'")
	}
	if args == nil || args.SourceDir == nil {
		return nil, errors.New("missing required argument 'SourceDir'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["destinationDir"] = nil
		inputs["sourceDir"] = nil
		inputs["vars"] = nil
	} else {
		inputs["destinationDir"] = args.DestinationDir
		inputs["sourceDir"] = args.SourceDir
		inputs["vars"] = args.Vars
	}
	s, err := ctx.RegisterResource("terraform-template:index/dir:Dir", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Dir{s: s}, nil
}

// GetDir gets an existing Dir resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDir(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DirState, opts ...pulumi.ResourceOpt) (*Dir, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["destinationDir"] = state.DestinationDir
		inputs["sourceDir"] = state.SourceDir
		inputs["vars"] = state.Vars
	}
	s, err := ctx.ReadResource("terraform-template:index/dir:Dir", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Dir{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Dir) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Dir) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Path to the directory where the templated files will be written.
func (r *Dir) DestinationDir() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["destinationDir"])
}

// Path to the directory where the files to template reside.
func (r *Dir) SourceDir() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceDir"])
}

// Variables for interpolation within the template. Note
// that variables must all be primitives. Direct references to lists or maps
// will cause a validation error.
func (r *Dir) Vars() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["vars"])
}

// Input properties used for looking up and filtering Dir resources.
type DirState struct {
	// Path to the directory where the templated files will be written.
	DestinationDir interface{}
	// Path to the directory where the files to template reside.
	SourceDir interface{}
	// Variables for interpolation within the template. Note
	// that variables must all be primitives. Direct references to lists or maps
	// will cause a validation error.
	Vars interface{}
}

// The set of arguments for constructing a Dir resource.
type DirArgs struct {
	// Path to the directory where the templated files will be written.
	DestinationDir interface{}
	// Path to the directory where the files to template reside.
	SourceDir interface{}
	// Variables for interpolation within the template. Note
	// that variables must all be primitives. Direct references to lists or maps
	// will cause a validation error.
	Vars interface{}
}
