FROM golang:1.14-alpine as build
WORKDIR /src/sliceConverter
COPY . .
RUN go build

FROM alpine:latest
COPY --from=build /src/sliceConverter .
CMD ["./sliceConverter", "-port", "3000"]
EXPOSE 3000
