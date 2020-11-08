## Counter

* Shows the time since the server start
* Shows numbers of visits (unique requests)

## Docker instructions

Instrucions for local build and run with Docker:

    docker build -t counter .
    docker run --publish 8080:8080 counter

Instructioins to use already compiled application from Docker Hub public repository:

    docker run --publish 8080:8080 fedir/counter:latest