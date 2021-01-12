FROM golang:latest as go
ARG PROJECT_ID
ENV CGO_ENABLED 0
WORKDIR /package
COPY go.* FirestoreKey.json /package/
RUN go get github.com/gin-gonic/gin

COPY main.go /package/
RUN go build

FROM gcr.io/distroless/static-debian10
COPY --from=go /package/authorize /package/*.json /
ENV GOOGLE_APPLICATION_CREDENTIALS=/FirestoreKey.json
ENV GOOGLE_PROJECT_ID=$PROJECT_ID
EXPOSE 8080
CMD ["/authorize"]
