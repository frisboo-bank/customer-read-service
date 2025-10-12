FROM gcr.io/distroless/static:nonroot

ENV GOTRACEBACK=single

COPY customer-read-service /services/customer-read-service

ENTRYPOINT [ "/services/customer-read-service" ]
