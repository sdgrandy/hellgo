FROM golang:latest
ARG ENVIRONMENT

RUN echo "ENVIRONMENT: ${ENVIRONMENT}"

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN ls
RUN go build -o main .
CMD ["/app/main"]