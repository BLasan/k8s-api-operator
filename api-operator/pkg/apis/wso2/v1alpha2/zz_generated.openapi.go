// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/wso2/v1alpha2.API":                  schema_pkg_apis_wso2_v1alpha2_API(ref),
		"./pkg/apis/wso2/v1alpha2.APISpec":              schema_pkg_apis_wso2_v1alpha2_APISpec(ref),
		"./pkg/apis/wso2/v1alpha2.APIStatus":            schema_pkg_apis_wso2_v1alpha2_APIStatus(ref),
		"./pkg/apis/wso2/v1alpha2.RateLimiting":         schema_pkg_apis_wso2_v1alpha2_RateLimiting(ref),
		"./pkg/apis/wso2/v1alpha2.RateLimitingSpec":     schema_pkg_apis_wso2_v1alpha2_RateLimitingSpec(ref),
		"./pkg/apis/wso2/v1alpha2.RateLimitingStatus":   schema_pkg_apis_wso2_v1alpha2_RateLimitingStatus(ref),
		"./pkg/apis/wso2/v1alpha2.Security":             schema_pkg_apis_wso2_v1alpha2_Security(ref),
		"./pkg/apis/wso2/v1alpha2.SecuritySpec":         schema_pkg_apis_wso2_v1alpha2_SecuritySpec(ref),
		"./pkg/apis/wso2/v1alpha2.SecurityStatus":       schema_pkg_apis_wso2_v1alpha2_SecurityStatus(ref),
		"./pkg/apis/wso2/v1alpha2.TargetEndpoint":       schema_pkg_apis_wso2_v1alpha2_TargetEndpoint(ref),
		"./pkg/apis/wso2/v1alpha2.TargetEndpointSpec":   schema_pkg_apis_wso2_v1alpha2_TargetEndpointSpec(ref),
		"./pkg/apis/wso2/v1alpha2.TargetEndpointStatus": schema_pkg_apis_wso2_v1alpha2_TargetEndpointStatus(ref),
	}
}

func schema_pkg_apis_wso2_v1alpha2_API(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "API is the Schema for the apis API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wso2/v1alpha2.APISpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wso2/v1alpha2.APIStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wso2/v1alpha2.APISpec", "./pkg/apis/wso2/v1alpha2.APIStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wso2_v1alpha2_APISpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APISpec defines the desired state of API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"mode": {
						SchemaProps: spec.SchemaProps{
							Description: "Mode of the API. The mode from the swagger definition will be overridden by this value. Supports \"privateJet\", \"sidecar\", \"<empty>\". Default value \"<empty>\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"updateTimeStamp": {
						SchemaProps: spec.SchemaProps{
							Description: "Update API definition creating a new docker image. Make a rolling update to the existing API. with prefixing the timestamp value. Default value \"<empty>\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replica count of the API.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"definition": {
						SchemaProps: spec.SchemaProps{
							Description: "Definition of the API.",
							Ref:         ref("./pkg/apis/wso2/v1alpha2.Definition"),
						},
					},
					"override": {
						SchemaProps: spec.SchemaProps{
							Description: "Override the exiting API docker image. Default value \"false\".",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version of the API. The version from the swagger definition will be overridden by this value. Default value \"<empty>\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"environmentVariables": {
						SchemaProps: spec.SchemaProps{
							Description: "Environment variables to be added to the API deployment. Default value \"<empty>\".",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Docker image of the API to be deployed. If specified, ignores the values of `UpdateTimeStamp`, `Override`. Uses the given image for the deployment. Default value \"<empty>\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiEndPoint": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ingressHostname": {
						SchemaProps: spec.SchemaProps{
							Description: "Ingress Hostname that the API is being exposed. Default value \"<empty>\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"swaggerConfigMapName": {
						SchemaProps: spec.SchemaProps{
							Description: "Config map name of which the project zip or swagger file is included",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"swaggerConfigMapName"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wso2/v1alpha2.Definition"},
	}
}

func schema_pkg_apis_wso2_v1alpha2_APIStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIStatus defines the observed state of API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "replicas field in the status sub-resource will define the initial replica count allocated to the API.This will be the minimum replica count for a single API",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"replicas"},
			},
		},
	}
}

func schema_pkg_apis_wso2_v1alpha2_RateLimiting(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RateLimiting is the Schema for the ratelimitings API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wso2/v1alpha2.RateLimitingSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wso2/v1alpha2.RateLimitingSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wso2_v1alpha2_RateLimitingSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RateLimitingSpec defines the desired state of RateLimiting",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"timeUnit": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"unitTime": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"requestCount": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wso2/v1alpha2.RequestCount"),
						},
					},
					"stopOnQuotaReach": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"bandwidth": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wso2/v1alpha2.Bandwidth"),
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wso2/v1alpha2.Conditions"),
						},
					},
				},
				Required: []string{"type", "timeUnit", "unitTime", "requestCount", "stopOnQuotaReach", "description", "bandwidth", "conditions"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wso2/v1alpha2.Bandwidth", "./pkg/apis/wso2/v1alpha2.Conditions", "./pkg/apis/wso2/v1alpha2.RequestCount"},
	}
}

func schema_pkg_apis_wso2_v1alpha2_RateLimitingStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RateLimitingStatus defines the observed state of RateLimiting",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_wso2_v1alpha2_Security(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Security is the Schema for the securities API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wso2/v1alpha2.SecuritySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wso2/v1alpha2.SecurityStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wso2/v1alpha2.SecuritySpec", "./pkg/apis/wso2/v1alpha2.SecurityStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wso2_v1alpha2_SecuritySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecuritySpec defines the desired state of Security",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"securityConfig": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/wso2/v1alpha2.SecurityConfig"),
									},
								},
							},
						},
					},
				},
				Required: []string{"type", "securityConfig"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wso2/v1alpha2.SecurityConfig"},
	}
}

func schema_pkg_apis_wso2_v1alpha2_SecurityStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecurityStatus defines the observed state of Security",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_wso2_v1alpha2_TargetEndpoint(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TargetEndpoint is the Schema for the targetendpoints API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wso2/v1alpha2.TargetEndpointSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wso2/v1alpha2.TargetEndpointStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wso2/v1alpha2.TargetEndpointSpec", "./pkg/apis/wso2/v1alpha2.TargetEndpointStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wso2_v1alpha2_TargetEndpointSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TargetEndpointSpec defines the desired state of TargetEndpoint",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"applicationProtocol": {
						SchemaProps: spec.SchemaProps{
							Description: "Protocol of the application. Supports \"http\" and \"https\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ports": {
						SchemaProps: spec.SchemaProps{
							Description: "List of optional ports of the target endpoint. First port should be the port of the target endpoint which is referred in swagger definition.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/wso2/v1alpha2.Port"),
									},
								},
							},
						},
					},
					"deploy": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployment details.",
							Ref:         ref("./pkg/apis/wso2/v1alpha2.Deploy"),
						},
					},
					"mode": {
						SchemaProps: spec.SchemaProps{
							Description: "Mode of the Target Endpoint. Supports \"privateJet\", \"sidecar\", \"serverless\". Default value \"privateJet\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"applicationProtocol", "ports", "deploy"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wso2/v1alpha2.Deploy", "./pkg/apis/wso2/v1alpha2.Port"},
	}
}

func schema_pkg_apis_wso2_v1alpha2_TargetEndpointStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TargetEndpointStatus defines the observed state of TargetEndpoint",
				Type:        []string{"object"},
			},
		},
	}
}
