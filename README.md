Chartmuseum Authserver
======================

A JWT authentication server for [Chartmuseum](https://chartmuseum.com/).

Implementation follows the [official example](https://github.com/chartmuseum/auth-server-example).

## Key handling

The Authserver needs a private key generated and the Chartmuseum instance needs the corresponding public key.

Generating keys:

```
openssl genrsa -out key.pem 2048
openssl rsa -in key.pem -outform PEM -pubout -out public.pem
```

## Building

The project is set up to be used with the [GB](https://getgb.io) build tool.

## Runtime configuration

Use the following command-line arguments:

| Argument              | Description                                            |
| --------------------- |:-------------------------------------------------------|
| `token-expiry`        | The duration that the generated token is valid for     |
| `required-grant-type` | The grant type name to request a token from the server |
| `master-access-key`   | The key used to request a token from the server        |
| `private-key-path`    | The file path to the private key file                  |
| `service-port`        | The HTTP port to bind to                               |
