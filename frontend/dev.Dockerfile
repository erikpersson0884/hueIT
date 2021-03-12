FROM node:latest

RUN mkdir -p /usr/src/huego/frontend
RUN chown -R node /usr/src/huego/frontend

USER node

WORKDIR /usr/src/huego/frontend

COPY ./src/ .
COPY package.json .
COPY ./public/ .
COPY .env.development .env

RUN yarn install
RUN yarn global add react-scripts

EXPOSE 3000

CMD yarn start