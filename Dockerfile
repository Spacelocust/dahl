FROM alpine
COPY dahl /usr/bin/dahl
ENTRYPOINT ["dahl"]