# Gobot
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE)

- - - -

# Setup

Gobot uses [govendor](https://github.com/kardianos/govendor/blob/master/doc/faq.md) to manage project dependencies.

    go run *.go

- - - -
# Hacking away at this project

1. Add this project to your GOPATH

    go get github.com/Swampy821/gobot

1. Install the vendored packages

    govendor sync

1. Add a new dependency that exists on your machine

    govendor add github.com/Swampy821/my-project

1. Fetch a remote dependency and add it to `vendor/vendor.json`

    govendor fetch github.com/Swampy821/some-project

1. Add the files to git and make a PR
