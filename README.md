# Yuno FAQman reciever
This application acts together with [yuno-faqman](https://github.com/alsen-ca/yuno-faqman), which is writtten to act as the frontend as a REPL.

This acts as a microservice and recieves the http requests.

It ends up writing or reading the data to/from a MongoDB database.

# Images
Code runs in Docker / Containerd. You can start them with the commands below.

You can pull images

    nerdctl pull golang:1.25.5-trixie
    nerdctl pull mongo:8.2.3

Or simply delete the '--pull never' line from the run commands.

Note that this setup is, for now, intended for development use only.

## Reciever (Go)
Version 1.25.5-trixie

First time starting the containers:

    nerdctl network create faqman

Inside container the first time (Dependencies are saved on go-mod-cache volume):

    go get go.mongodb.org/mongo-driver/mongo

nerdctl run --rm \
    --init \
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

Before starting app, be sure you have [Mongo running](#run-container)

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

## Testing
Tests need refinement.
For now, one can:

    go test ./...
    go test ./internal/http/thema

Or when debugging:

    go test -v ./...



# Classes
Replace id={uuid} with the actual uuid, for example:

    /thema?id=25bbe563-2a67-4cf3-86b4-e945c41814d7

## Themas
### Create
    curl -X POST http://127.0.0.1:8221/thema \
     -H "Content-Type: application/json" \
     -d '{"title":"go"}'

    curl -X POST http://127.0.0.1:8221/thema \
     -H "Content-Type: application/json" \
     -d '{"title":"rust"}'

### Get
#### ID / UUID
    curl http://127.0.0.1:8221/thema?id={uuid}

#### Title
    curl http://127.0.0.1:8221/thema?title="go"
    
#### All
    curl http://127.0.0.1:8221/thema

### Update
    curl -X PUT http://127.0.0.1:8221/thema?id={uuid} \
     -H "Content-Type: application/json" \
     -d '{"title": "golang"}'

    curl -X PUT http://127.0.0.1:8221/thema?title="rust" \
     -H "Content-Type: application/json" \
     -d '{"title": "Rust"}'

### Delete
    curl -X DELETE http://127.0.0.1:8221/thema?id={uuid}

    curl -X DELETE http://127.0.0.1:8221/thema?title=Rust


## Tags
### Create
    curl -X POST http://127.0.0.1:8221/tag \
     -H "Content-Type: application/json" \
     -d '{"en_og": "Example", "de_trans": "Beispiel", "es_trans": "Ejemplo"}'

### Get
#### All
    curl http://127.0.0.1:8221/tag

#### ID / UUID
    curl http://127.0.0.1:8221/tag?id={uuid}

#### By Language

    curl http://127.0.0.1:8221/tag?en_og=Example
    curl http://127.0.0.1:8221/tag?de_trans=Beispiel
    curl http://127.0.0.1:8221/tag?es_trans=Ejemplo

### Delete
    curl -X DELETE http://127.0.0.1:8221/tag?id={uuid}

### Update
    curl -X PUT http://127.0.0.1:8221/tag?id={uuid} \
     -H "Content-Type: application/json" \
     -d '{"en_og": "New Tag Name", "de_trans": "neuer Tag", "es_trans": "nuevo"}'
    

## QA
Note that, for simplict's sake, the 'lang' is saved to the database as in, instead of as an enum (int).
If you have larger datasets and require more efficiency, it might be useful to implement enums here.

### Create
    curl -X POST http://127.0.0.1:8221/qa \
     -H "Content-Type: application/json" \
     -d '{"question":"How to perform POST?",
     "question_weights": [{"word": "how", "weight": 1.0},{"word": "to", "weight": 1.0},{"word": "perform", "weight": 1.0},{"word": "post", "weight": 1.0}],
     "answer": "curl something something",
     "lang": "en",
     "thema_id": "f242c924-aaf4-4c94-b4b9-368a7b1b919c",
     "tag_ids": [{9440a01f-929c-4f66-9edc-70a29606568a}]}'

    curl -X POST http://127.0.0.1:8221/qa \
     -H "Content-Type: application/json" \
     -d '{"question":"What is needed to change?",
     "question_weights": [{"word": "what", "weight": 1.0},{"word": "is", "weight": 1.0},{"word": "neeeded", "weight": 1.0},{"word": "to", "weight": 1.0},
        {"word": "change", "weight": 0.1}],
     "answer": "no idea",
     "lang": "en",
     "thema_id": "f242c924-aaf4-4c94-b4b9-368a7b1b919c",
     "tag_ids": ["9440a01f-929c-4f66-9edc-70a29606568a","6af90e78-ce11-4c5e-8c91-19a4fe3dcb3c"]}'
    
    curl -X POST http://127.0.0.1:8221/qa \
     -H "Content-Type: application/json" \
     -d '{"question":"How to add every file to git?",
     "question_weights": [{"word": "how", "weight": 1.0},{"word": "to", "weight": 1.0},
        {"word": "add", "weight": 1.0},{"word": "every", "weight": 1.0},{"word": "file", "weight": 1.0},
        {"word": "to", "weight": 0.1},{"word": "git", "weight": 4.0}],
     "answer": "git add file",
     "lang": "en"}'


### Get
#### All
    curl http://127.0.0.1:8221/qa

#### ID / UUID
    curl http://127.0.0.1:8221/qa?id={uuid}

#### By Question
(Question must match exactly)

    curl http://127.0.0.1:8221/qa?question=How%20to%20%init%Git%3D


#### Update
    curl "http://127.0.0.1:8221/qa?question=How%20to%20add%20every%20file%20to%20git%3F"

    curl -X PUT http://127.0.0.1:8221/qa?id={uuid} \
     -H "Content-Type: application/json" \
     -d '{"question":"How to add every file to git?",
     "question_weights": [{"word": "how", "weight": 1.0},{"word": "to", "weight": 1.0},
        {"word": "add", "weight": 1.0},{"word": "every", "weight": 1.0},{"word": "file", "weight": 1.0},
        {"word": "to", "weight": 0.1},{"word": "git", "weight": 4.0}],
     "answer": "git add .",
     "lang": "en"}'
