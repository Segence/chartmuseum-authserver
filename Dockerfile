FROM alpine
ADD main /
ENTRYPOINT ["/main"]
