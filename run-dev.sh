#!/bin/bash

sudo nerdctl --address /run/k3s/containerd/containerd.sock run --rm \
    -it \
    --name faqman-reciever \
    --network faqman \
    -p 127.0.0.1:8221:8221 \
    --pull never \
    -v $(pwd):/app \
    -v go-mod-cache:/go/pkg/mod \
    -w /app \
    golang:1.25.5-trixie \
    bash
