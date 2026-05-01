# Backend

Backend del **Order Management Dashboard**, construido con Go, Gin, GORM y PostgreSQL.

La API permite autenticación con JWT, consulta paginada de órdenes, filtros dinámicos, estadísticas agregadas y carga automática del dataset `orders_db.csv`.

---

## Stack

* Go 1.26
* Gin
* GORM
* PostgreSQL
* JWT
* bcrypt
* golang-migrate
* Docker

---

## Funcionalidades

* Login con email y contraseña.
* Generación de `access_token` y `refresh_token`.
* Renovación de sesión.
* Logout con revocación de refresh token.
* Rutas protegidas con middleware JWT.
* Consulta paginada de órdenes.
* Filtros por:

  * Canal
  * Compañía
  * Tipo de entrega
  * Tipo de producto
* Endpoint de opciones para filtros.
* Estadísticas agregadas:

  * Total de órdenes
  * Porcentaje de órdenes con error
  * Órdenes por canal
  * Órdenes por tipo de entrega
  * Órdenes por tipo de producto
* Importación automática del CSV mediante seed.
* Migraciones SQL versionadas.

---

## Variables de entorno

Ejemplo de configuración:

```env
APP_ENV=development
APP_PORT=8080

DATABASE_URL=postgres://postgres:postgres@postgres:5432/orders_db?sslmode=disable

JWT_ACCESS_SECRET=change_me_access_secret
JWT_REFRESH_SECRET=change_me_refresh_secret
JWT_ACCESS_TTL_MINUTES=15
JWT_REFRESH_TTL_HOURS=168

ADMIN_EMAIL=admin@test.com
ADMIN_PASSWORD=admin123
ADMIN_NAME=Admin

CSV_PATH=/app/data/orders_db.csv

CORS_ALLOWED_ORIGINS=http://localhost:3000
```

---

## Ejecutar con Docker

Desde la raíz del proyecto:

```bash
docker compose up --build
```

Esto levanta automáticamente:

1. PostgreSQL
2. Migraciones
3. Seed del CSV
4. Backend
5. Frontend

La API quedará disponible en:

```txt
http://localhost:8080
```

---

## Health check

```txt
GET /health
```

Ejemplo:

```bash
curl http://localhost:8080/health
```

Respuesta esperada:

```json
{
  "status": "ok"
}
```

---

## Usuario demo

El backend crea automáticamente un usuario admin si no existe:

```txt
Email: admin@test.com
Password: admin123
```

---

## Endpoints principales

### Auth

```txt
POST /api/auth/login
POST /api/auth/refresh
POST /api/auth/logout
```

### Orders

```txt
GET /api/orders
GET /api/orders/filters
```

`GET /api/orders` soporta:

```txt
page
pageSize
canal
company
fulfillmentType
productType
```

Ejemplo:

```txt
GET /api/orders?page=1&pageSize=10&canal=WEB
```

### Stats

```txt
GET /api/stats
```

`GET /api/stats` soporta los mismos filtros:

```txt
canal
company
fulfillmentType
productType
```

Ejemplo:

```txt
GET /api/stats?canal=WEB&productType=Soft%20Line
```

---

## Estructura principal

```txt
backend/
├── cmd/
│   ├── api/
│   └── seed/
├── internal/
│   ├── app/
│   ├── auth/
│   ├── config/
│   ├── database/
│   ├── http/
│   ├── orders/
│   └── stats/
├── migrations/
├── data/
│   └── orders_db.csv
└── Dockerfile
```

---

## Arquitectura

El backend está organizado por módulos:

```txt
auth
orders
stats
```

Cada módulo separa responsabilidades:

```txt
handler     → capa HTTP
service     → lógica de negocio
repository  → acceso a datos
entity      → mapeo GORM / base de datos
model       → modelos internos y DTOs
ports       → interfaces
mapper      → conversión entre estructuras
```

La capa `app` se encarga del bootstrap:

* Crear router.
* Configurar CORS.
* Crear conexión a PostgreSQL.
* Inicializar módulos.
* Registrar rutas públicas y protegidas.
* Conectar middlewares.

---

## Migraciones

Las migraciones están en:

```txt
backend/migrations/
```

En Docker Compose se ejecutan automáticamente mediante el servicio `migrate`.

Para correrlas manualmente desde la raíz:

```bash
migrate \
  -path backend/migrations \
  -database "postgres://postgres:postgres@localhost:5433/orders_db?sslmode=disable" \
  up
```

---

## Seed

El seed está en:

```txt
backend/cmd/seed
```

Carga el archivo:

```txt
backend/data/orders_db.csv
```

La carga es idempotente, por lo que puede ejecutarse más de una vez sin duplicar órdenes.

---

## Decisiones técnicas

### Go

Se eligió Go por su rendimiento, simplicidad para construir APIs y facilidad de despliegue.

### Gin

Se usó Gin por ser un framework HTTP ligero, rápido y fácil de estructurar.

### GORM

Se usó GORM para simplificar el acceso a datos, reducir código repetitivo y facilitar filtros dinámicos.

### golang-migrate

Aunque se usa GORM, el esquema de base de datos se controla mediante migraciones SQL versionadas para mantener trazabilidad y control explícito.

### PostgreSQL

Se eligió PostgreSQL por su robustez, soporte para índices, filtros y consultas agregadas.

### JWT con refresh token

Se implementó un flujo con `access_token` de vida corta y `refresh_token` persistido en base de datos. Los refresh tokens se guardan como hash y pueden ser revocados.

### Separación entity/model/DTO

Las entidades con tags de GORM se mantienen separadas de los modelos internos y de los DTOs de respuesta. Esto evita acoplar la API al esquema de base de datos.

---

## Mejoras futuras

* Agregar pruebas unitarias usando subtests.
* Agregar pruebas de integración para repositories.
* Agregar tests para handlers HTTP.
* Agregar logs estructurados.
* Agregar filtros por rango de fechas.
* Agregar ordenamiento dinámico de órdenes.
* Mover secretos a un gestor seguro.
* Agregar rate limiting en endpoints de auth.
* Agregar pipeline CI para lint, tests y build.
