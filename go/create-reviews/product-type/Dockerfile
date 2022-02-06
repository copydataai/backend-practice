FROM public.ecr.aws/bitnami/golang:1.16 as build-image
# ENVS PG
ENV USER_PG="user_read"
ENV PASSWORD_PG="user_read"
ENV DB_PG="reviews-camera"
ENV HOST_PG="localhost"
ENV PORT_PG="5436"


WORKDIR /go/src
COPY go.mod go.sum main.go ./

RUN go build -o ../bin

FROM public.ecr.aws/lambda/go:1

COPY --from=build-image /go/bin/ /var/task/

# Command can be overwritten by providing a different command in the template directly.
CMD ["product-type"]
