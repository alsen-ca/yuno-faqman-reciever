# Yuno FAQman reciever
This applciation acts together with yuno-faqman, which acted as the frontend on a terminal.

This acts as a microservice and recieves the http requests.

It ends up writing or reading the data to a MongoDB database.

# Running

    go run main.go

    go build main.go
    ./main

    curl -X POST http://127.0.0.1:3200/thema \
     -H "Content-Type: application/json" \
     -d '{"title":"My First Thema"}'
