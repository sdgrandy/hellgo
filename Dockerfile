FROM golang:latest
ENV PORT=9876
RUN echo "PORT: ${PORT}"
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN ls
RUN pwd
RUN go build -o main .
CMD ["/app/main"]