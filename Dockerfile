FROM gcr.io/distroless/static:nonroot

ENV GOTRACEBACK=single

COPY customers-service /services/customers-service

ENTRYPOINT [ "/services/customers-service" ]
