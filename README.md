# Order Management Dashboard

Aplicación fullstack para la administración, monitoreo y análisis de órdenes de venta a partir de un dataset CSV.

El sistema permite autenticarse, visualizar métricas agregadas, consultar órdenes paginadas, aplicar filtros y analizar la distribución de órdenes por canal, tipo de entrega y tipo de producto.

---

## Stack técnico

### Backend

* Go 1.26
* Gin
* GORM
* PostgreSQL
* JWT con `access_token` y `refresh_token`
* bcrypt para hashing de contraseñas
* golang-migrate para migraciones SQL
* Docker

### Frontend

* Next.js
* TypeScript
* Tailwind CSS
* Axios
* Recharts
* React Hook Form
* Docker

### Infraestructura

* Docker Compose
* PostgreSQL 16
* Servicio automático de migraciones
* Servicio automático de seed del CSV

---

## Funcionalidades

### Autenticación

* Login con email y contraseña.
* Generación de `access_token`.
* Generación y persistencia de `refresh_token`.
* Renovación de sesión.
* Logout con revocación de refresh token.
* Rutas protegidas por JWT.

### Dashboard

* Cards con métricas principales.
* Gráficos de órdenes por canal y por tipo de producto.
* Tabla de órdenes paginada.
* Filtros por:

  * Canal
  * Compañía
  * Tipo de entrega
  * Tipo de producto
* Selector de tamaño de página.
* Diseño responsive.

### Data

* Importación automática de `orders_db.csv`.
* Normalización ligera de:

  * Valores vacíos
  * Fechas con prefijo `__Timestamp__`
  * Booleanos
  * Números opcionales
* Carga idempotente usando `ON CONFLICT DO NOTHING`.

---

## Requisitos

Para correr el proyecto solo necesitas:

* Docker
* Docker Compose

No es necesario instalar Go, Node.js ni PostgreSQL localmente si usas Docker.

---

## Cómo correr el proyecto

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

Cuando termine de levantar, abre:

```txt
http://localhost:3000/login
```

---

## Usuario demo

```txt
Email: admin@test.com
Password: admin123
```

---

## URLs principales

### Frontend

```txt
http://localhost:3000
```

### Backend

```txt
http://localhost:8080
```

### Health check

```txt
GET http://localhost:8080/health
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
GET /api/orders?page=1&pageSize=10&canal=WEB&company=LP
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

## Variables de entorno

### Backend

Las variables principales del backend son:

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

### Frontend

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

Nota: aunque backend y frontend corren en Docker, el frontend usa `localhost:8080` porque las peticiones se ejecutan desde el navegador del usuario.

---

## Estructura del proyecto

```txt
.
├── backend/
│   ├── cmd/
│   │   ├── api/
│   │   └── seed/
│   ├── internal/
│   │   ├── app/
│   │   ├── auth/
│   │   ├── config/
│   │   ├── database/
│   │   ├── http/
│   │   ├── orders/
│   │   └── stats/
│   ├── migrations/
│   ├── data/
│   │   └── orders_db.csv
│   └── Dockerfile
│
├── frontend/
│   ├── app/
│   ├── components/
│   ├── lib/
│   ├── types/
│   └── Dockerfile
│
├── docker-compose.yml
└── README.md
```

---

## Arquitectura backend

El backend está organizado por módulos:

```txt
auth
orders
stats
```

Cada módulo separa responsabilidades en capas:

```txt
handler     → HTTP
service     → lógica de negocio
repository  → acceso a datos
entity      → mapeo GORM / base de datos
model       → modelos internos y DTOs
ports       → interfaces del módulo
mapper      → conversiones entre entity, model y DTO
```

La capa `app` se encarga del bootstrap:

* Crear router.
* Configurar CORS.
* Crear conexión a base de datos.
* Inicializar módulos.
* Registrar rutas públicas y protegidas.
* Conectar middlewares.

---

## Arquitectura frontend

El frontend separa:

```txt
app/          → rutas Next.js
components/   → componentes visuales
lib/          → cliente API, auth, config y mappers
types/        → contratos TypeScript
```

Se usan mappers para transformar contratos externos del backend en modelos internos del frontend.

Por ejemplo:

```txt
Backend/API: snake_case
Frontend: camelCase
```

Esto evita acoplar los componentes de React al formato exacto de la API.

---

## Decisiones técnicas

### Go para backend

Se eligió Go porque es eficiente, simple de desplegar y adecuado para construir APIs robustas. Además, la prueba técnica menciona Go como preferencia.

### GORM para acceso a datos

Se usó GORM para reducir código repetitivo en repositories, facilitar filtros dinámicos y mantener una capa de acceso a datos clara.

### golang-migrate para migraciones

Aunque se usa GORM, el schema se administra con migraciones SQL versionadas. Esto permite tener control explícito sobre tablas, índices y extensiones.

### PostgreSQL

Se eligió PostgreSQL por su confiabilidad, soporte para queries agregadas, filtros, índices y datos transaccionales.

### JWT con refresh token

Se implementó un flujo con `access_token` de vida corta y `refresh_token` persistido en base de datos. Los refresh tokens se guardan como hash y pueden revocarse en logout.

### Docker Compose

Se usó Docker Compose para que el proyecto pueda levantarse con un solo comando, incluyendo base de datos, migraciones, seed, backend y frontend.

### Mappers en frontend

El frontend trabaja internamente en `camelCase`, aunque la API responda en `snake_case`. Esto mantiene los componentes desacoplados del contrato externo.

---

## Comandos útiles

### Levantar todo

```bash
docker compose up --build
```

### Levantar en segundo plano

```bash
docker compose up -d --build
```

### Ver logs

```bash
docker compose logs -f
```

### Ver logs del backend

```bash
docker compose logs -f backend
```

### Ver logs del frontend

```bash
docker compose logs -f frontend
```

### Bajar servicios

```bash
docker compose down
```

### Borrar datos y levantar desde cero

```bash
docker compose down -v
docker compose up --build
```

---

## Migraciones

Las migraciones se ejecutan automáticamente en Docker Compose mediante el servicio `migrate`.

Si se desea correr manualmente desde la máquina local:

```bash
migrate \
  -path backend/migrations \
  -database "postgres://postgres:postgres@localhost:5433/orders_db?sslmode=disable" \
  up
```

---

## Seed

El seed se ejecuta automáticamente en Docker Compose mediante el servicio `seed`.

El seed carga:

```txt
backend/data/orders_db.csv
```

La carga es idempotente, por lo que correr el seed más de una vez no duplica órdenes.

---

## Mejoras futuras

* Agregar pruebas unitarias y de integración en backend usando subtests.
* Agregar tests para repositories con base de datos de prueba.
* Agregar tests de componentes frontend.
* Documentar los pacquetes con `doc.go` y las funcionalidades existentes. 
* Mover tokens a cookies `httpOnly` para mayor seguridad.
* Agregar observabilidad con logs estructurados.
* Agregar filtros por rango de fechas.
* Agregar exportación CSV desde el dashboard.
* Agregar ordenamiento por columnas en la tabla.
* Agregar cache o vistas materializadas para métricas si el volumen de datos crece.
* Implementar virtual scrolling si se decide abandonar la paginación tradicional.
* Agregar pipeline CI para ejecutar lint, tests y build.
