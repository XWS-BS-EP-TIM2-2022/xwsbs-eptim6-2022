FROM node:16.13 as node
WORKDIR /app
COPY ./dislinkt-webapp ./
RUN npm install
RUN npm run build --prod

FROM nginx:alpine
COPY --from=node /app/dist/dislinkt-webapp /usr/share/nginx/html
COPY /nginx.conf  /etc/nginx/conf.d/default.conf

EXPOSE 80