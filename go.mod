module github.com/pulumi/pulumi-terraform-template

go 1.12

replace (
	git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
	github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v31.1.0+incompatible
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-aws => github.com/pulumi/terraform-provider-aws v0.0.0-20190912220144-b9ca5fbb6e9e
)

require (
	github.com/Azure/go-autorest/autorest v0.9.1 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/hashicorp/terraform v0.12.7
	github.com/pulumi/pulumi v1.1.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190923175357-e7ab441da0dd
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-template v1.0.0
)
