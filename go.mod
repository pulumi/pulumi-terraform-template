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
	github.com/Nvveen/Gotty v0.0.0-20170406111628-a8b993ba6abd // indirect
	github.com/go-ini/ini v1.31.0 // indirect
	github.com/hashicorp/terraform v0.12.7
	github.com/pulumi/pulumi v0.17.28-0.20190731182900-6804d640fc7c
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190923175357-e7ab441da0dd
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-template v1.0.0
	gopkg.in/vmihailenco/msgpack.v2 v2.9.1 // indirect
)
