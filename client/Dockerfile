FROM node:alpine

WORKDIR /app

COPY package*.json .

RUN npm i

COPY . .

VOLUME /app/dist

#EXPOSE 5173

#CMD npm run dev
CMD npm run build
