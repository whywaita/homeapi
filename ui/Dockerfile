FROM node:8-alpine AS build-temp

WORKDIR /app

# needed because some package refers GitHub.
RUN apk add --update --no-cache --virtual .build-dep \
        git

ADD package.json yarn.lock /app/

RUN yarn

COPY . /app
RUN yarn run build

FROM lkwg82/h2o-http2-server

ADD h2o.conf /etc/h2o
COPY --from=build-temp /app/dist /public
