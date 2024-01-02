# scammmateo
All scams welcome!

All rights reserved, BS workshops. Don't call us, we'll call you!

## Setting up a development environment

 - [Get direnv here] (https://direnv.net/#basic-installation "link to direnv installation instructions")
 - [Get devbox here] (https://www.jetpack.io/devbox "link to devbox installation instructions")
 - Download all the go package dependencies: `go get -v ./...`

## Running in development mode

We use air to perform file watches and recompile/reload on changes. Run it via:

`air`

This starts a local server [which you can click here to open in a web browser] (127.0.0.1:1443 "link to a running dev server")

A> The dev server runs with a self-signed key/certificate located at scammmateo/etc/tls/key.pem and scammmateo/etc/tls/cert.pem. This is not secure!
