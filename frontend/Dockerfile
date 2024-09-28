FROM node:16 AS builder

WORKDIR /frontend

COPY package*.json ./

RUN npm install

COPY . .

RUN npm test -- --watchAll=false

RUN npm run build




FROM nginx:alpine

COPY --from=builder /frontend/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
