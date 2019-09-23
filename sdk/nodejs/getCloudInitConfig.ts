// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-template/blob/master/website/docs/d/cloudinit_config.html.markdown.
 */
export function getCloudInitConfig(args: GetCloudInitConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudInitConfigResult> & GetCloudInitConfigResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetCloudInitConfigResult> = pulumi.runtime.invoke("terraform-template:index/getCloudInitConfig:getCloudInitConfig", {
        "base64Encode": args.base64Encode,
        "gzip": args.gzip,
        "parts": args.parts,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getCloudInitConfig.
 */
export interface GetCloudInitConfigArgs {
    /**
     * Base64 encoding of the rendered output. Default to `true`
     */
    readonly base64Encode?: boolean;
    /**
     * Specify whether or not to gzip the rendered output. Default to `true`
     */
    readonly gzip?: boolean;
    /**
     * One may specify this many times, this creates a fragment of the rendered cloud-init config file. The order of the parts is maintained in the configuration is maintained in the rendered template.
     */
    readonly parts: inputs.GetCloudInitConfigPart[];
}

/**
 * A collection of values returned by getCloudInitConfig.
 */
export interface GetCloudInitConfigResult {
    readonly base64Encode?: boolean;
    readonly gzip?: boolean;
    readonly parts: outputs.GetCloudInitConfigPart[];
    /**
     * The final rendered multi-part cloudinit config.
     */
    readonly rendered: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
