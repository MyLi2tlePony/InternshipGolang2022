# Собираем в го
FROM golang:1.18.3 as build

ENV BIN_FILE /bin/app
ENV CODE_DIR /go/src/

# Кэшируем слои с модулями
WORKDIR ${CODE_DIR}
COPY go.mod .
COPY go.sum .
RUN go mod download

# Копипуем код
COPY . ${CODE_DIR}

# Собираем бинарник Go
RUN go build -o ${BIN_FILE} ./cmd/migration

ENV CONFIG_FILE configs/test/config.toml

#CMD sleep 10 && ${BIN_FILE} -config ${CONFIG_FILE}
CMD ${BIN_FILE} -config ${CONFIG_FILE} && go test -tags=integration ${CODE_DIR}/...