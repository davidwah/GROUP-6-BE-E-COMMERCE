FROM golang:1.17

# create directory app
RUN mkdir /app

# set or make /app our working directory
WORKDIR /app

# copy all files to /app
COPY . .

RUN go build -o g6-ecommerce

CMD [ "./g6-ecommerce" ]