FROM egaillardon/jmeter
STOPSIGNAL SIGINT

ADD PressureMeter /jmeter
ADD Config.yaml /jmeter

EXPOSE 8080
WORKDIR /jmeter
RUN mkdir Data
VOLUME [ "/jmeter/Data" ]
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["/jmeter/PressureMeter"]