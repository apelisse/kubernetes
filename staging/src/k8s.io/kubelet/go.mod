// This is a generated file. Do not edit directly.

module k8s.io/kubelet

go 1.12

require (
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
)

replace (
	github.com/gogo/protobuf => github.com/apelisse/protobuf v0.0.0-20190410021324-0ad0d52e9ce345a781f6c002748fc399d4efb611
	golang.org/x/sync => golang.org/x/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190209173611-3b5209105503
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190313210603-aa82965741a9
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/kubelet => ../kubelet
)
