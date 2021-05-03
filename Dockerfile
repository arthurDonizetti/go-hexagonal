FROM golang:1.15

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go mod init appproduct && \
  go get github.com/spf13/cobra/cobra@v1.1.3 && \
  go get github.com/golang/mock/mockgen@v1.5.0

RUN apt-get update && apt-get install sqlite3 -y

RUN touch db.sqlite

CMD ["tail", "-f", "/dev/null"]
