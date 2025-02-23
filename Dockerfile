FROM golang:1.24

#COPYING HOST FILES INTO THE CONTAINER
COPY . .

#
RUN go mod download
CMD ["go", "run", "."]
 
