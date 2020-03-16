FROM golang:alpine
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./... && go build -v -o /PressureMeter

FROM egaillardon/jmeter
STOPSIGNAL SIGINT

COPY --from=0 /PressureMeter /jmeter
ADD Config.yaml /jmeter

EXPOSE 8080
WORKDIR /jmeter
RUN mkdir Data
VOLUME [ "/jmeter/Data" ]
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["/jmeter/PressureMeter"]