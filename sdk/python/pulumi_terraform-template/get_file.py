# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetFileResult:
    """
    A collection of values returned by getFile.
    """
    def __init__(__self__, filename=None, rendered=None, template=None, vars=None, id=None):
        if filename and not isinstance(filename, str):
            raise TypeError("Expected argument 'filename' to be a str")
        __self__.filename = filename
        if rendered and not isinstance(rendered, str):
            raise TypeError("Expected argument 'rendered' to be a str")
        __self__.rendered = rendered
        """
        The final rendered template.
        """
        if template and not isinstance(template, str):
            raise TypeError("Expected argument 'template' to be a str")
        __self__.template = template
        """
        See Argument Reference above.
        """
        if vars and not isinstance(vars, dict):
            raise TypeError("Expected argument 'vars' to be a dict")
        __self__.vars = vars
        """
        See Argument Reference above.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetFileResult(GetFileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileResult(
            filename=self.filename,
            rendered=self.rendered,
            template=self.template,
            vars=self.vars,
            id=self.id)

def get_file(filename=None,template=None,vars=None,opts=None):
    """
    Renders a template from a file.
    
    :param str filename: _Deprecated, please use `template` instead_. The filename for
           the template. Use [path variables](https://www.terraform.io/docs/configuration/interpolation.html#path-variables) to make
           this path relative to different path roots.
    :param str template: The contents of the template. These can be loaded
           from a file on disk using the [`file()` interpolation
           function](https://www.terraform.io/docs/configuration/interpolation.html#file_path_).
    :param dict vars: Variables for interpolation within the template. Note
           that variables must all be primitives. Direct references to lists or maps
           will cause a validation error.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-template/blob/master/website/docs/d/file.html.markdown.
    """
    __args__ = dict()

    __args__['filename'] = filename
    __args__['template'] = template
    __args__['vars'] = vars
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('terraform-template:index/getFile:getFile', __args__, opts=opts).value

    return AwaitableGetFileResult(
        filename=__ret__.get('filename'),
        rendered=__ret__.get('rendered'),
        template=__ret__.get('template'),
        vars=__ret__.get('vars'),
        id=__ret__.get('id'))
