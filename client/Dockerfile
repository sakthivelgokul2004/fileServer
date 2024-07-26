FROM node:alpine

WORKDIR /app

COPY package*.json .

RUN npm i

COPY . .

VOLUME /app/dist

CMD npm run build
