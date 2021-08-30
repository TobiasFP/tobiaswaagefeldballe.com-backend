FROM mongo:4.2
LABEL TobiasFP "tobias@osandweb.dk"
ADD init-mongo.js /docker-entrypoint-initdb.d/
CMD ["mongod", "--auth"]