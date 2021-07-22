package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iancoleman/orderedmap"
	"github.com/kameshsampath/gloo-jsongenerator/jsonschema"
	_ "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
	meshEntNetworkv1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1"
	meshNetworkv1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	glooRbacv1 "github.com/solo-io/gloo-mesh/pkg/api/rbac.enterprise.mesh.gloo.solo.io/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	"reflect"
)

func arrayOrStringMapper(i reflect.Type) *jsonschema.Type {
	//if (i == reflect.TypeOf(meshNetworkv1.AccessPolicySpec{})) {
	//	properties := orderedmap.New()
	//	properties.Set("source_selector", &jsonschema.Type{
	//		Type: "object",
	//	})
	//	properties.Set("destination_selector", &jsonschema.Type{
	//		Type: "object",
	//	})
	//	properties.Set("allowed_paths", &jsonschema.Type{
	//		Type: "array",
	//		Items: &jsonschema.Type{
	//			Type: "string",
	//		},
	//	})
	//	properties.Set("allowed_methods", &jsonschema.Type{
	//		Type: "array",
	//		Items: &jsonschema.Type{
	//			Type: "string",
	//		},
	//	})
	//	properties.Set("allowed_ports", &jsonschema.Type{
	//		Type: "array",
	//		Items: &jsonschema.Type{
	//			Type: "integer",
	//		},
	//	})
	//	return &jsonschema.Type{
	//		Type:                 "object",
	//		AdditionalProperties: []byte("false"),
	//		Required:             []string{"source_selector", "destination_selector"},
	//		Properties:           properties}
	//}
	if (i == reflect.TypeOf(v1.Duration{})) {
		return &jsonschema.Type{
			Type:    "string",
			Pattern: "^[-+]?([0-9]*(\\.[0-9]*)?(ns|us|µs|μs|ms|s|m|h))+$",
		}
	}
	if (i == reflect.TypeOf(v1.Time{})) {
		return &jsonschema.Type{
			Type:   "string",
			Format: "data-time",
		}
	}
	//TODO clean up
	//if (i == reflect.TypeOf(common.IdentitySelector{}) ||
	//	i == reflect.TypeOf(common.WorkloadSelector{}) ||
	//	i == reflect.TypeOf(common.DestinationSelector{})) {
	//	return &jsonschema.Type{
	//		Type:                 "object",
	//		AdditionalProperties: []byte("true"),
	//		Properties:           orderedmap.New()}
	//}
	//if (i == reflect.TypeOf(knative.VolatileTime{})) {
	//	return &jsonschema.Type{
	//		Type: "string",
	//		Format: "data-time",
	//	}
	//}
	//if (i == reflect.TypeOf(v1beta1.WhenExpression{})) {
	//	properties := orderedmap.New()
	//	properties.Set("input", &jsonschema.Type{
	//		Type: "string",
	//	})
	//	properties.Set("operator", &jsonschema.Type{
	//		Type: "string",
	//	})
	//	properties.Set("values", &jsonschema.Type{
	//		Type: "array",
	//		Items: &jsonschema.Type{
	//			Type: "string",
	//		},
	//	})
	//	return &jsonschema.Type{
	//		Type: "object",
	//		AdditionalProperties: []byte("false"),
	//		Required: []string{"input", "operator", "values"},
	//		Properties: properties}
	//}
	if (i == reflect.TypeOf(runtime.RawExtension{})) {
		return &jsonschema.Type{
			Type:                 "object",
			AdditionalProperties: []byte("true"),
			Properties:           orderedmap.New()}
	}
	if (i == reflect.TypeOf(apiextensionsv1.JSON{})) {
		return &jsonschema.Type{
			OneOf: []*jsonschema.Type{
				{
					Type: "boolean",
				},
				{
					Type: "integer",
				},
				{
					Type: "number",
				},
				{
					Type: "string",
				},
				{
					Type: "array",
					Items: &jsonschema.Type{
						OneOf: []*jsonschema.Type{
							{
								Type: "string",
							},
							{
								Type:                 "object",
								AdditionalProperties: []byte("true"),
								Properties:           orderedmap.New(),
							},
						},
					},
				},
				{
					Type:                 "object",
					AdditionalProperties: []byte("true"),
					Properties:           orderedmap.New(),
				},
				{
					Type: "null",
				},
			},
		}
	}
	return nil
}

func dump(v interface{}, apiVersion string, kind string) {
	fmt.Printf("Starting generation of %s %s\n", apiVersion, kind)
	filename := fmt.Sprintf("%s_%s.json", apiVersion, kind)
	reflector := jsonschema.Reflector{
		TypeMapper: arrayOrStringMapper,
	}
	reflect := reflector.Reflect(v)
	JSON, _ := reflect.MarshalJSON()
	file, _ := os.Create(filename)
	defer file.Close()
	var out bytes.Buffer
	json.Indent(&out, JSON, "", "  ")
	out.WriteTo(file)
	index, _ := os.OpenFile("index.properties", os.O_WRONLY|os.O_APPEND, 0)
	index.WriteString(filename)
	index.WriteString("\n")
}

func main() {
	os.Create("index.properties")
	os.Mkdir("networking.mesh.gloo.solo.io", os.ModePerm)
	os.Mkdir("rbac.enterprise.mesh.gloo.solo.io", os.ModePerm)
	os.Mkdir("networking.enterprise.mesh.gloo.solo.io", os.ModePerm)
	dump(&meshNetworkv1.AccessPolicy{}, "networking.mesh.gloo.solo.io/v1", "AccessPolicy")
	dump(&meshNetworkv1.TrafficPolicy{}, "networking.mesh.gloo.solo.io/v1", "TrafficPolicy")
	dump(&glooRbacv1.Role{}, "rbac.enterprise.mesh.gloo.solo.io/v1", "Role")
	dump(&glooRbacv1.RoleBinding{}, "rbac.enterprise.mesh.gloo.solo.io/v1", "RoleBinding")
	dump(&meshEntNetworkv1beta1.VirtualDestination{}, "networking.enterprise.mesh.gloo.solo.io/v1beta1", "VirtualDestination")

}
