FROM node:alpine

WORKDIR '/app'

COPY ./package.json ./

RUN npm install

RUN npm install sass

COPY . .

EXPOSE 3000

CMD ["npm", "run", "start"]
