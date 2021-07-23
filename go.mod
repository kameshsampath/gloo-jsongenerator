module github.com/kameshsampath/gloo-jsongenerator

go 1.16

require (
	github.com/iancoleman/orderedmap v0.2.0
	github.com/iancoleman/strcase v0.1.3
	github.com/solo-io/gloo-mesh v1.0.11
	github.com/solo-io/solo-apis v1.6.30
	k8s.io/api v0.21.0 // indirect
	k8s.io/apiextensions-apiserver v0.21.3
	k8s.io/apimachinery v0.21.0
	sigs.k8s.io/controller-runtime v0.9.0-alpha.1.0.20210412152200-442d3cad1e99 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.20.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.4
	k8s.io/client-go => k8s.io/client-go v0.20.4
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.7.0
)
