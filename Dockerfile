FROM sierratecnologia/golang

MAINTAINER Ricardo Sierra <ricardo@sierratecnologia.com.br>

RUN go get -v github.com/neelance/graphql-go && \
    go get -v golang.org/x/crypto/bcrypt && \
    go get -v github.com/jmoiron/sqlx && \
    go get -v github.com/go-sql-driver/mysql
    
