FROM golang:1.6-onbuild
EXPOSE 8888
# Image will do:
# COPY . /go/src/app,
# RUN go get -d -v, and
# RUN go install -v
# CMD ["app"]
