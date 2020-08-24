## Mock API

A fake API that enables me to test and utilise a multitude of other services, such as GitLab and Docker.

Various usages will be pushed into particular branches.

## Docker

The `Dockerfile` has been annotated as much as possible to provide a level of detail into what the concept of each command is and does.

The overall concept is that there is a multi-stage build. The first stage utilises the latest Golang image, which contains all of relevant files to build a Go application into a static binary file. Afterwards, a minimal Alpine Linux image is used to run the binary file, since Go places all of the relevant dependencies within the binary itself.

The decision of using Alpine Linux vs a `FROM scratch` (completely blank) image with this Golang application is that access to the `jsonplaceholder` site requires HTTPS, meaning that TLS/SSL certificates are required - these are not available on a scratch image and must be installed as an aside, therefore it is simpler to use Alpine Linux, which is a minimal image in itself. A Go application which did not require HTTPS, such as only HTTP between other internal services or nothing at all, could use a blank image and see the benefit of a smaller overall image.

## TODO:
* GitLab CI 
* Docker/containerise: initial use @ `feature/containerise`
* Implement monitoring (Kibana)