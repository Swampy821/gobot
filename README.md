# Gobot
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE)

- - - -

# Setup

Gobot uses [govendor](https://github.com/kardianos/govendor/blob/master/doc/faq.md) to manage project dependencies.

    go run *.go

- - - -
# Hacking away at this project

    go get github.com/Swampy821/gobot
    govendor add github.com/thoj/go-ircevent
    govendor add github.com/BurntSushi/toml
    govendor add github.com/Jeffail/gabs
    go run *.go
