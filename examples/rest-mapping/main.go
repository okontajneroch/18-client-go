package main

import (
	"fmt"
	"k8s.io/client-go/discovery/cached/memory"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// ziskam konfiguraciu standardne z $HOME/.kube/config
	config, _ := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		nil,
	).ClientConfig()

	// vytvorim discovery rest mapper, ktory mi pomoze ziskat mapping pre
	// resources priamo z mojho k8s clustra, pomocou discovery API
	discoveryClient, _ := discovery.NewDiscoveryClientForConfig(config)
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(
		memory.NewMemCacheClient(discoveryClient),
	)

	// ziskam mapping pre moj vlastny resource 'Starfighter'
	mapping, _ := mapper.RESTMapping(schema.GroupKind{
		Group: "starwars.okontajneroch.sk",
		Kind:  "Starfighter",
	})

	// vypisem informacie o resource
	fmt.Printf("Version: %s\n", mapping.GroupVersionKind.Version)
	fmt.Printf("Resource: %s\n", mapping.Resource.Resource)
}
