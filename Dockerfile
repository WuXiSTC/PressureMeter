FROM yindaheng98/go-git
ADD . .

RUN go get -d -v ./... && \
    cd main && \
    go build -v -o /PressureMeter

FROM egaillardon/jmeter
STOPSIGNAL SIGINT

RUN mkdir /jmeter/Data && chmod a+rw /jmeter/Data
COPY --from=0 /PressureMeter /jmeter
COPY entrypoint.sh /usr/local/bin/
WORKDIR /jmeter

EXPOSE 8080
VOLUME [ "/jmeter/Data" ]
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["/jmeter/PressureMeter"]