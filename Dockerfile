FROM golang:1.17-alpine

WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o /d-crud101-jcc

EXPOSE 8080
CMD [ "/d-crud101-jcc" ]