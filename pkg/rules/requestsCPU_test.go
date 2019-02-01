package rules

import (
	"github.com/ghodss/yaml"
	"testing"
)

func Test_RequestsCPU_Pod(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: c1
    resources:
      limits:
       cpu: 300m
      requests:
       cpu: 300m
  - name: c2
    resources:
      requests:
  - name: c3
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	containers := RequestsCPU(json)
	if containers != 1 {
		t.Errorf("Got %v containers wanted %v", containers, 1)
	}
}
