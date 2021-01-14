FROM golang:latest as go
ENV CGO_ENABLED 0
WORKDIR /package
COPY go.* FirestoreKey.json main.go /package/
RUN go get github.com/gin-gonic/gin
RUN go build

FROM gcr.io/distroless/static-debian10
ARG PROJECT_ID
COPY --from=go /package/authorize /package/*.json /
ENV GOOGLE_APPLICATION_CREDENTIALS=/FirestoreKey.json
ENV GOOGLE_PROJECT_ID=$PROJECT_ID
CMD ["/authorize"]
