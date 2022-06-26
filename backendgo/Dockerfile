FROM golang:1.18-alpine as build
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o /app/backendgo
FROM golang:1.18-alpine
WORKDIR /app
COPY --from=build /app/backendgo /app
EXPOSE 8081
ENTRYPOINT ["sh", "-c"]
CMD [ "/app/backendgo" ]