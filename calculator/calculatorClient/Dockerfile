FROM golang
COPY ./calculator/calculatorClient /app/calculator/calculatorClient
COPY ./calculator/calculatorPB /app/calculator/calculatorPB
WORKDIR /app
RUN go build /app/calculator/calculatorClient && chmod +x /app/calculator/calculatorClient
ENTRYPOINT ["./calculatorClient"]
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./calculatorClient"]
EXPOSE 7777