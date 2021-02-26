FROM golang
COPY ./calculator/calculatorServer /app/calculator/calculatorServer
COPY ./calculator/calculatorPB /app/calculator/calculatorPB
WORKDIR /app
RUN go build /app/calculator/calculatorServer && chmod +x /app/calculator/calculatorServer
ENTRYPOINT ["./calculatorServer"]
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./calculatorServer"]
EXPOSE 7777