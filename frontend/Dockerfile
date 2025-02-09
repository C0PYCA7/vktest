FROM node:latest AS builder

WORKDIR /app

COPY ../frontend/package.json ../frontend/package-lock.json ./

RUN npm install

COPY ../frontend ./frontend

WORKDIR /app/frontend

RUN npm run build

FROM nginx:alpine

WORKDIR /app/frontend

COPY --from=builder /app/frontend/ /usr/share/nginx/html

COPY ../frontend/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["/usr/sbin/nginx", "-g", "daemon off;"]