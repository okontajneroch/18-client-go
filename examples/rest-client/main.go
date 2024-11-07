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
	starwarsConfig := *config // kopia hlavnej konfigurace
	starwarsConfig.GroupVersion = &schema.GroupVersion{
		Group:   "starwars.okontajneroch.sk",
		Version: "v1",
	}
	starwarsConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme)

	// vytvorim REST clienta pre 'starwars.io/v1'
	starwarsClient, _ := rest.RESTClientFor(&starwarsConfig)

	// vytvorim si novy Starfighter objekt
	xwing := &starwarsv1.Starfighter{}
	xwing.Name = "x-wing-1"
	xwing.Spec.Pilot = "Luke Skywalker"
	xwing.Spec.Type = "X-Wing"
	xwing.Spec.Faction = starwarsv1.Rebellion

	// poslem objekt do k8s pomocou POST requestu
	result := &starwarsv1.Starfighter{}
	err := starwarsClient.
		Post().
		Namespace("default").
		Resource("starfighters").
		Name("x-wing-1").
		Body(xwing).
		Do(context.TODO()).
		Into(result)

	// spracujem error alebo vysledok
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Updated Starfighter: %s %s \n", result.Name, result.UID)
}
