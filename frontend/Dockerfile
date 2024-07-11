FROM node:22-slim

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
RUN pnpm config set store-dir ${PNPM_HOME}/store
COPY ./app /app
WORKDIR /app
EXPOSE 5173

CMD pnpm install && pnpm run dev