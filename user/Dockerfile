FROM alpine
WORKDIR /Initial
COPY ./bin/user .
COPY ./etc/test.yaml .
RUN  mkdir etc && mv test.yaml etc/user.yaml
EXPOSE 8080
ENTRYPOINT ["./user"]