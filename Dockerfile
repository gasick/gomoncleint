FROM golang:1.16 as build
RUN apt update && apt install -y git
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download 
COPY . .
RUN go build -o /out/app /app/cmd/client


FROM fedora
COPY --from=build /out/app /app
CMD ["/app"]