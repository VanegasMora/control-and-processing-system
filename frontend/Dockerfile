# ---- Deps ----
FROM node:20-alpine AS deps

WORKDIR /app

# Necesario para compilar dependencias de Next.js en Alpine
RUN apk add --no-cache python3 make g++

COPY package.json package-lock.json* ./
RUN npm ci


# ---- Builder ----
FROM node:20-alpine AS builder

WORKDIR /app

# Copiar node_modules desde fase anterior
COPY --from=deps /app/node_modules ./node_modules

# Copiar todo el proyecto
COPY . .

# Asegurar que public exista
RUN mkdir -p public

# Construir el frontend (standalone)
RUN npm run build


# ---- Runner ----
FROM node:20-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production

# Crear usuario sin privilegios
RUN addgroup --system --gid 1001 nodejs && \
    adduser --system --uid 1001 nextjs

# Copiar archivos necesarios para ejecutar Next.js standalone
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static

USER nextjs

EXPOSE 3000
ENV PORT=3000
ENV HOSTNAME="0.0.0.0"

CMD ["node", "server.js"]