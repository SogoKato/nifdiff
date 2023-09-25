FROM gcr.io/distroless/static-debian11
COPY nifdiff /
ENTRYPOINT [ "/nifdiff" ]
