// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as template from "@pulumi/terraform-template";

const config = new pulumi.Config("template");
const name = config.get("name") || "world";

const f = template.getFile({
    template: "hello, ${name}!",
    vars: {
        "name": name,
    },
});

export const rendered = pulumi.output(f).apply(f => f.rendered);
