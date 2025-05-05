FROM node:23-alpine3.20

WORKDIR /run

COPY . .

RUN apk add --no-cache git
RUN npm install -g @angular/cli@16.2.1
RUN npm install

EXPOSE 4200

CMD [ "ng", "serve", "--host", "0.0.0.0" ]
