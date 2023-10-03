# civo-versions

A basic api to return a list of available Kubernetes versions for Civo.

## Features

Exposes the following endpoints:

* `/` - Lists all active releases
* `/k3s/` - Lists all active k3s releases
* `/k3s/stable` - Lists all stable k3s releases
* `/k3s/development` - Lists all development k3s releases
* `/talos/` - Lists all Talos releases
* `/talos/stable` - Lists all stable Talos releases
* `/talos/development` - Lists all development Talos releases

## Building from source

With Docker:

```sh
make docker-build
```

Standalone:

```sh
make build
```

## Resources

* [civogo](https://github.com/civo/civogo)

## Contributing

If you find a bug or have an idea for a new feature please [raise an issue](issues/new) to discuss it.

Pull requests are welcomed but please try and follow similar code style as the rest of the project and ensure all tests and code checkers are passing.

Thank you 💛

## License

See [LICENSE](LICENSE)
