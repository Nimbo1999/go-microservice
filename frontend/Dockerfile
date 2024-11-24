FROM node:18-alpine as builder

ARG VITE_HOST
ENV VITE_HOST=$VITE_HOST

WORKDIR /app

COPY package.json .
RUN yarn install

COPY public ./public
COPY src ./src
COPY index.html .
COPY tsconfig.json .
COPY tsconfig.node.json .
COPY vite.config.ts .

RUN yarn build

FROM nginx:1.25.3-alpine

COPY --from=builder /app/dist /usr/share/nginx/html
