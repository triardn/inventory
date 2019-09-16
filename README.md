# Inventory

A simple API for inventory

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

## How to
You need to duplicate `inventory.toml.sample` and rename it to `.inventory.toml` first, and then running the app.

Running it then should be as simple as:

```console
$ make build
```

or if you too lazy to duplicate the `inventory.toml.sample` file, you can just use this command:

```console
$ make start-dev
```

You can see the Postman Collection from the API at `_postman`