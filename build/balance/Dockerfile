# Собираем в го
FROM golang:1.18.3 as build

ENV MIGRATION_BIN_FILE /bin/migration
ENV BALANCE_BIN_FILE /bin/balance
ENV CODE_DIR /go/src/

# Кэшируем слои с модулями
WORKDIR ${CODE_DIR}
COPY go.mod .
COPY go.sum .
RUN go mod download

# Копипуем код
COPY . ${CODE_DIR}

# Собираем бинарник Go
RUN go build -o ${MIGRATION_BIN_FILE} ./cmd/migration
RUN go build -o ${BALANCE_BIN_FILE} ./cmd/balance

ENV MIGRATION_CONFIG_FILE configs/migration/config.toml
ENV BALANCE_CONFIG_FILE configs/balance/config.toml

CMD ${MIGRATION_BIN_FILE} -config ${MIGRATION_CONFIG_FILE} && ${BALANCE_BIN_FILE} -config ${BALANCE_CONFIG_FILE}