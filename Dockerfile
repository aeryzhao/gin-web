FROM scratch

WORKDIR $GOPATH/src/github.com/iszhaoxg/gin-web
COPY . $GOPATH/src/github.com/iszhaoxg/gin-web

EXPOSE 8000
CMD ["./go-gin-example"]