// +k8s:deepcopy-gen=package
// +groupName=starwars.okontajneroch.sk
//
//go:generate deepcopy-gen . --output-file zz_generated.deepcopy.go
//go:generate controller-gen crd paths=./... output:crd:dir=../../k8s
package v1
