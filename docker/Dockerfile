FROM ubuntu:24.04
WORKDIR /app
COPY backend /app/
COPY assets /app/
RUN chmod +x /app/backend
ENTRYPOINT [ "/app/backend" ]