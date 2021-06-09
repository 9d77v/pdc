FROM 9d77v/pdc-base:0.1.0
COPY . /app

ENV APP_NAME=pdc
RUN cd /app \
    && go build -o $APP_NAME -ldflags "-s -w" cmd/server.go\
    && upx -9 $APP_NAME

FROM 9d77v/pdc-base-deploy:0.0.3
ENV APP_NAME=pdc
COPY --from=0  /app/$APP_NAME /usr/bin/$APP_NAME
COPY ./web/build /app/web/build
COPY ./tpls /app/tpls

WORKDIR /app
CMD [ "pdc" ]
