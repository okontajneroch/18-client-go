package main

import (
	"context"
	"fmt"
	starwarsv1 "github.com/okontajneroch/starwars/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// potrebujem schemu a serializer, ktory bude vediet konvertovat 'starwars.io'
	// objekty na externu reprezentaciu (JSON,Protobuf,YAML) a naspat
	scheme := runtime.NewScheme()
	starwarsv1.AddToScheme(scheme)

	// ziskam konfiguraciu standardne z $HOME/.kube/config
	config, _ := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		nil,
	).ClientConfig()
	config.APIPath = "/apis"

	// vytvorim konfiguraciu, ktora je platna len pre 'starwars.io/v1'
	starwarsConfig := *config // kopia hlavnej konfiguracie
	starwarsConfig.GroupVersion = &schema.GroupVersion{
		Group:   "starwars.okontajneroch.sk",
		Version: "v1",
	}
	starwarsConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme)

	// vytvorim REST clienta, ktory je platny pre 'starwars.io/v1'
	starwarsClient, _ := rest.RESTClientFor(&starwarsConfig)

	// ziskam povodny bjekt 'x-wing-1` priamo z k8s clustra
	result := &starwarsv1.Starfighter{}
	err := starwarsClient.
		Get().
		Namespace("default").
		Resource("starfighters").
		Name("x-wing-1").
		Do(context.TODO()).
		Into(result)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// zmenim
	result.Spec.Pilot = "Wedge Antilles"

	// vykonam update objektu 'x-wing-1' v k8s clustri pomocou PUT requestu
	err = starwarsClient.
		Put().
		Namespace("default").
		Resource("starfighters").
		Name("x-wing-1").
		Body(result).
		Do(context.TODO()).
		Into(result)

	fmt.Printf("Updated Starfighter: %s %s \n", result.Name, result.UID)
}
