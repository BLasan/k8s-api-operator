// +build tools

package tools

import (
       _ "k8s.io/code-generator/cmd/client-gen"
       _ "k8s.io/code-generator/cmd/deepcopy-gen"
       _ "k8s.io/code-generator/cmd/defaulter-gen"
       _ "k8s.io/code-generator/cmd/informer-gen"
       _ "k8s.io/code-generator/cmd/lister-gen"
       _ "k8s.io/kube-openapi/cmd/openapi-gen"
)
