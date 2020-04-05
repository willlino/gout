FROM golang:alpine as build

ADD . /src
# RUN apk --no-cache add build-base git bzr mercurial gcc upx && \
#     cd /src && \
#     GOOS=linux go build -ldflags="-s -w" -o /src/main.go && \
#     upx --brute /src/main.go

WORKDIR /src/src
RUN go test 

# FROM scratch
# WORKDIR /usr/bin
# COPY --from=build /src/main.go /usr/bin

ENTRYPOINT ["go", "main.go"]