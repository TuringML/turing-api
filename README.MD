# TuringML API

### How to run the APIs

I assume you have docker installed and you are using a Mac. 
In you terminal, go to the root of the project and then type:

```bash
$: cd api/
$: make build
$: make docker-image
```

The image will be created under this name `turingml/api`.
To run the container use the following:

```bash
$: docker run --rm -it -p 8080:8080 turingml/api
```

In the Docker Image, **Gin** is set up in `release` mode. If you don't want
such a thing, you can change it with `debug` or comment it out.

## Vault initialization

Exec into vault with `docker exec -it vault sh`. Run the following:

```bash
$: vault status
$: vault operator init  // store the Root key and the unsealing keys
$: vault operator unseal // repeat this command 3 times
$: vault login // insert the root key
```