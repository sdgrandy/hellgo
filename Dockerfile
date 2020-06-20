FROM golang:latest
ARG PORT
ENV PORT="${PORT}"
RUN echo "PORT: ${PORT}"
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN ls
RUN pwd
RUN /app/env.sh
RUN go build -o main .
CMD ["/app/main"]