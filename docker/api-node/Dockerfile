FROM node:18

WORKDIR /api-node/

ENV PORT=4000
ENV DATABASE_URL=postgres://postgres:foobarbaz@localhost:5432/postgres

COPY package*.json ./

COPY . .

RUN npm install 
RUN npm install -g typescript

RUN tsc -b

CMD ["node", "dist/index.js"]

# docker build .