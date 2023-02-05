ARG NODE_VERSION

FROM node:${NODE_VERSION}-alpine

WORKDIR /app

EXPOSE 5173

CMD ["/bin/sh", "-c", "yarn; yarn setup; yarn dev --host"]
