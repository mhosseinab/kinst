FROM dockerhub.ir/node:alpine
# Set a working directory
WORKDIR /usr/src/app

RUN apk --no-cache --virtual build-dependencies add \
    python \
    make \
    g++

COPY package.json .
RUN yarn 
RUN yarn install

COPY . .
RUN yarn build

EXPOSE 3000
CMD ["yarn", "start"]
