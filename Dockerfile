####  Build webui distribution ##############
FROM node:latest AS webui

RUN mkdir /app
WORKDIR /app

COPY web-ui/package.json .
RUN npm install 

COPY web-ui .
RUN ls

RUN npm run build.prod

####  build qsave state 
FROM swipl AS saved-state

ENV http_port 80
ENV base_url ''

RUN apt update
RUN apt install -y git

WORKDIR /build

RUN mkdir /build/.packages

COPY controller/*.pl /build/
COPY controller/sources /build/sources/
RUN ls /build/sources

# install custom packages, e.g. 
# RUN swipl -g  "pack_install(prolog_sax, [interactive(false), upgrade(true), url('https://github.com/milung/prolog_sax/archive/v1.0.3.zip'), inquiry(false)])" -t halt

RUN swipl -o bootfile -c run.pl

#### TARGET controller Image
FROM swipl

WORKDIR /app

# COPY --from=saved-state /build/.bin /app
COPY --from=saved-state /build/bootfile /app/bootfile
COPY --from=webui /controller/www /app/www

ENTRYPOINT ["swipl", "-x", "/app/bootfile"]
# remove if only cli is requested
CMD ["server"]

EXPOSE 80