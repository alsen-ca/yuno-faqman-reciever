# Yuno FAQman reciever
This applciation acts together with yuno-faqman, which acted as the frontend on a terminal.

This acts as a microservice and recieves the http requests.

It ends up writing or reading the data to a MongoDB database.

# Running

    go run .

    go build main.go
    ./main

    curl -X POST http://127.0.0.1:8221/thema \
     -H "Content-Type: application/json" \
     -d '{"title":"My First Thema"}'

# Images
## Reciever (Go)
Version 1.25.5-trixie

nerdctl network create faqman

go get go.mongodb.org/mongo-driver/mongo

nerdctl run \
    --name faqman-reciever \
    --network faqman \
    -p 127.0.0.1:8221:8221 \
    --pull never \
    -d golang:1.25.5-trixie \
    # Copy source
    # Download mod
    # Build binary
    # Run binary

## Mongo
Version 8.2.3

nerdctl run \
    --name faqman-db \
    --network faqman \
    -v /var/lib/mongo:/data/db \
    -p 127.0.0.1:8222:8222 \
    --pull never \
    -d mongo:8.2.3
