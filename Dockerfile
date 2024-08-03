FROM node:22-alpine as build
WORKDIR /usr/src/app
COPY package.json package-lock.json ./
RUN npm install
COPY ./ ./
RUN npm run build

FROM node:22-alpine as runtime
WORKDIR /usr/src/app
COPY --from=build /usr/src/app/package.json /usr/src/app/package-lock.json ./
COPY --from=build /usr/src/app/node_modules ./node_modules
COPY --from=build /usr/src/app/build ./build

EXPOSE 3000
CMD ["node", "build"]
