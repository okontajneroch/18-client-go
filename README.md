# REST Client a REST Mapper

Zdrojové kódy z článku **Custom Resource Definition** na [okontajneroch.sk](https://okontajneroch.sk). Na spustenie je potrebné mať nainštalované:

- Go 1.23
- deepcopy-gen (`go install k8s.io/code-generator/cmd/deepcopy-gen@latest`)
- controller-gen (`go install sigs.k8s.io/controller-tools/cmd/controller-gen`)
- k8s cluster (napr kind)

V adresári `./k8s` sa nachádza vygenerovaný CRD. Na pregenerovanie je možné 
použiť `make generate` alebo `go generate`

Samozrejme, najprv je potrebne mať teda cluster a aplikovat CRD definície:

```bash
$ kind create cluster
$ kubectl apply -f ./k8s/starwars.okontajneroch.sk_starfighter.yaml
```

Potom môžeme spustiť príklady. Napríklad:

```
go run ./examples/rest-client
```
