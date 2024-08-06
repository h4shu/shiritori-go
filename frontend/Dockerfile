FROM node:22-slim

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
COPY ./src /app
WORKDIR /app
EXPOSE 5173
RUN corepack enable
RUN pnpm config set store-dir ${PNPM_HOME}/store
CMD pnpm install && pnpm run dev