FROM vikings/alpine:v1.0.8-7-ge66a81d
COPY groot /groot
EXPOSE 8000
ENTRYPOINT /groot