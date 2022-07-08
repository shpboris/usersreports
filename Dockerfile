##
## Build
##
FROM golang:1.18-bullseye AS build


WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o usersreports .
WORKDIR /dist
RUN cp /build/usersreports .

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /dist/usersreports /
EXPOSE 8001
##USER nonroot:nonroot
ENTRYPOINT ["/usersreports"]