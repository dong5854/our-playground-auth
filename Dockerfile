FROM golang:1.19.0-buster as build
  
# set go module mode without GOPATH
ENV GO111MODULE=on

WORKDIR /usr/src/app
ENV GONOSUMDB=github.com/useb-inc/*
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /usr/src/app/cmd/auth

RUN go build -o our-playground-auth

FROM golang:1.19.0-buster as stage
  
# remove apt install lists
RUN apt-get update && \
apt-get -y install uuid-runtime && \
rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app/
COPY --from=build /usr/src/app/private.key /usr/src/app/
COPY --from=build /usr/src/app/public.pem /usr/src/app/
COPY --from=build /usr/src/app/cmd/auth/our-playground-auth /usr/src/app/cmd/auth/our-playground-auth

RUN chmod +x /usr/src/app/cmd/auth/our-playground-auth

CMD [ "/usr/src/app/cmd/auth/our-playground-auth" ]
