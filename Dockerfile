FROM golang:1.6-onbuild
EXPOSE 8080
# Image will do:
# COPY . /go/src/app,
# RUN go get -d -v, and
# RUN go install -v
# CMD ["app"]
