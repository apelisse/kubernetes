// This is a generated file. Do not edit directly.

module k8s.io/api

go 1.12

require (
	github.com/gogo/protobuf v0.0.0-20171007142547-342cbe0a0415
	github.com/stretchr/testify v1.2.2
	k8s.io/apimachinery v0.0.0
)

replace (
	github.com/gogo/protobuf => github.com/apelisse/protobuf v0.0.0-20190410021324-0ad0d52e9ce345a781f6c002748fc399d4efb611
	golang.org/x/sync => golang.org/x/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190209173611-3b5209105503
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190313210603-aa82965741a9
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
)
