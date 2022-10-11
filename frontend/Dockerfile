FROM node:lts-alpine3.15 as builder

WORKDIR '/app'

COPY ./package.json ./

RUN npm install

RUN npm install sass

COPY . .

RUN npm run build

FROM nginx:stable-alpine

COPY --from=builder /app/build /usr/share/nginx/html
