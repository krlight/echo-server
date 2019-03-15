###################################################################################################
# docker build -t got -f build.Dockerfile . --target runtime                                      #
# docker run -p 8102:8102 got                                                                     #
###################################################################################################
FROM golang:1.12-alpine3.9 AS build

WORKDIR /build/
COPY main.go .
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o main.exe main.go

FROM alpine:3.9 AS runtime
RUN apk add --no-cache tzdata

WORKDIR /runtime/
COPY --from=build /build/main.exe .
ENV PORT 8102

CMD [ "sh", "-c", "./main.exe"]
