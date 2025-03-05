# there is caching in the the layers
# currently there are 10 layers
# when there is some changes at any layer makes the again run
# from layer that is edited, (rest aove layers are cached)

FROM node:18

# this is what we to containarised
WORKDIR /api-node/

# providing the PORT and DATABASE_URL as a env
ENV PORT=4000
ENV DATABASE_URL=postgres://postgres:foobarbaz@localhost:5432/postgres

# copying the speicifc files like package.json...
COPY package*.json ./

# assigning the owner the code copies in src to be node
# USER node
# COPY --chown=user:node ./src .

# copying the remainder soruce code of our workdir
COPY . .

# installing all the dependecies, typescript for running ts commands
RUN npm install 
RUN npm install -g typescript

RUN tsc -b

# this is exposing port
# EXPOSE 4000

# finally running the docker container
# here there is server so server will start
CMD ["node", "dist/index.js"]

# docker build .