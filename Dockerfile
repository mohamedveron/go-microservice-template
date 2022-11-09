FROM golang:1.17-alpine as builder

RUN apk add --no-cache git
RUN apk add --update make
RUN apk add --no-cache openssh

# Move to working directory /app
WORKDIR /app


# create ssh directory
RUN mkdir ~/.ssh

ARG SSH_PRIVATE_KEY
RUN echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_ed25519
RUN chmod 0600 ~/.ssh/id_ed25519
RUN touch ~/.ssh/known_hosts
RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

RUN git config --global url.git@github.com:.insteadOf https://github.com/

ENV GO111MODULE "on"
ENV GONOPROXY "github.com/mohamedveron/*"
ENV GONOSUMDB "github.com/mohamedveron/*"
ENV GOPRIVATE "github.com/mohamedveron/*"
ENV GOPROXY "https://proxy.golang.org,direct"

# Copy the code into the container
COPY . .

RUN go mod download

RUN make build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/app /bin/app

# Export necessary port
EXPOSE 9090
# Command to run when starting the container
CMD ["/bin/app"]