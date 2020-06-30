FROM node:erbium-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json ./
COPY yarn.lock ./

RUN apk --no-cache add yarn
RUN yarn install --silent

CMD [ "yarn", "start" ]