FROM golang:1.24 as base

# set /app as workdir
WORKDIR /app

# set timezone
ENV TZ=Asia/Tehran

# install git
# build-essential => -race option on golang
RUN apt-get update && apt-get install -y git build-essential

# set timezone
RUN ls /usr/share/zoneinfo && cp /usr/share/zoneinfo/Iran /etc/localtime && echo "Asia/Tehran" >  /etc/timezone

# install swag for api swagger doc 
RUN go install github.com/swaggo/swag/cmd/swag@latest

# copy go.mod and go.sum
COPY go.mod go.mod
COPY go.sum go.sum

# download required packages
RUN go mod download

# copy source code in container
COPY . .

# ============================== # 
#  dev # 
# ============================== # 

FROM base as dev

WORKDIR /app

RUN swag init --parseInternal --parseDepth 100 -d 'cmd/,srv/api,types' -o ./.swagger/docs

CMD ["go","run","./cmd"]

# ============================== # 
#  builder # 
# ============================== # 

FROM base AS builder

WORKDIR /app

COPY ./configs/prod.config.yaml ./configs/config.yaml

RUN swag init --parseInternal --parseDepth 100 -d 'cmd/,srv/api,types' -o ./.swagger/docs

RUN go build -o ./build/gobit ./cmd

# ============================== # 
#  production # 
# ============================== # 

FROM debian:bookworm-slim AS prod 

# set timezone
ENV TZ=Asia/Tehran
RUN ls /usr/share/zoneinfo && cp /usr/share/zoneinfo/Iran /etc/localtime && echo "Asia/Tehran" >  /etc/timezone

WORKDIR /app

# copy build and configs
COPY --from=builder /app/build /app/build
COPY --from=builder /app/configs /app/configs

# run gobit
CMD ["/app/build/gobit"]
