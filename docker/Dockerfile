FROM golang:alpine3.18 as build
RUN adduser -D nonroot
WORKDIR /app/
COPY . .
RUN apk add --update make
RUN make build


FROM busybox:stable as run
WORKDIR /app/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /app/bin/main /app/.env ./
USER nonroot
EXPOSE 9010
CMD [ "./main" ]
