# Yuno FAQman reciever
This applciation acts together with yuno-faqman, which acted as the frontend on a terminal.

This acts as a microservice and recieves the http requests.

It ends up writing or reading the data to a MongoDB database.

# Running

    go run .

    curl -X POST http://127.0.0.1:8221/thema/new \
     -H "Content-Type: application/json" \
     -d '{"title":"go"}'

    curl "http://127.0.0.1:8221/thema?id={uuid}"

    curl "http://127.0.0.1:8221/thema?title=go"
    
    curl http://127.0.0.1:8221/thema





# Images
Code runs in Docker / Containerd. Before running, you can start them with these commands

## Reciever (Go)
Version 1.25.5-trixie

nerdctl network create faqman

go get go.mongodb.org/mongo-driver/mongo

nerdctl run --rm \
    --init \
    --name faqman-reciever \
    --network faqman \
    -p 127.0.0.1:8221:8221 \
    --pull never \
    -v $(pwd):/app \
    -v go-mod-cache:/go/pkg/mod \
    -w /app \
    golang:1.25.5-trixie \
    go run .

## Mongo
Version 8.2.3

### Run container

nerdctl run -d \
    --restart=unless-stopped \
    --name faqman-db \
    --network faqman \
    -v /var/lib/mongo:/data/db \
    -p 127.0.0.1:8222:27017 \
    --pull never \
    -d mongo:8.2.3

### Use mongosh

nerdctl exec -it faqman-db mongosh
show dbs
use yuno-faqman
db.themas.find()