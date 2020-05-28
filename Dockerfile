FROM 9d77v/pdc-base:0.0.2
COPY . /app

ENV APP_NAME=pdc
RUN cd /app \
    && go build -o $APP_NAME -ldflags "-s -w" \
    && upx -9 $APP_NAME

FROM alpine:3.11
ENV APP_NAME=pdc
COPY --from=0  /app/$APP_NAME /usr/bin/$APP_NAME
COPY ./ui/build /app/ui/build

WORKDIR /app
CMD [ "pdc" ]