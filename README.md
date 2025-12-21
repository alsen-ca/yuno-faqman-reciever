# Yuno FAQman reciever
This applciation acts together with yuno-faqman, which acted as the frontend on a terminal.

This acts as a microservice and recieves the http requests.

It ends up writing or reading the data to a MongoDB database.

# Running

    go run main.go

    go build main.go
    ./main

    curl -X POST http://127.0.0.1:8221/thema \
     -H "Content-Type: application/json" \
     -d '{"title":"My First Thema"}'

# Structure


root
| README
| service.go
| db.go
| main.go
|   models
    |   - thema
            | struct.go
            | controller.go
            - api (/thema)
                - newThema (/new)
                - getThema (/get/:title) (returns uuid)
                - updateThema (/update/:uuid)
                - deleteThema (/delete/:uuid)
                - allThemas (/all/1) (Only returns 10)
            - validations
                - correctLengths (new/update) (2..100)
                - uniqueTitle (new/update)
                - meaningfulGet (get) (:title)
                - meaningfulUuidLength (update/delete)

    |   - tag
            | struct.go
            | controller.go
            - api (/tag)
                - newTag (/new)
                - getTag (/get/:{en_og,de_trans,es_trans}) (returns uuid)
                - updateTag (/update/:uuid)
                - deleteTag (/delete/:uuid)
                - allTags (/all/1) (Only returns 10)
            - validations
                - correctLenghts (new/update) (2..50)
                - uniqueTag (new/update) (Checks uniqueness across all languages)
                - minOneLangProvided (new/update)
                - meaningfulGet (get) (:en_og, :de_trans, :es_trans)
                - meaningfulUuidLength (update/delete)

    |   - qa
            | struct.go
            | controller.go
            - api (/qa)
                - newQa (/new)
                - getQa (/get/:question) (returns uuid)
                - updateQa (/update/:uuid)
                - deleteQa (/delete/:uuid)
            - validations
                - correctWeights (new/update) (Defaults missing weights to 1)
                - correctLengths (new/update) (2..4000)
                - correctLanguages (new/update) (enum en/de/es. Defaults to en)
                - meaningfulGet (get) (:question)
                - meaningfulUuidLength (update/delete)

|   algo    (Algorithm that expects some keywords and finds a combination of QA.
            More concretely, it tries to find based on words given, which Questions is searched.
            Returns a list of most probable answers.)

        | run.go
