# Frontend

Frontend del **Order Management Dashboard**, construido con Next.js, TypeScript y Tailwind CSS.

La aplicación permite iniciar sesión, visualizar métricas de órdenes, consultar una tabla paginada, aplicar filtros y analizar información mediante gráficos.

---

## Stack

* Next.js
* TypeScript
* Tailwind CSS
* Axios
* Recharts
* React Hook Form

---

## Funcionalidades

* Login responsive.
* Dashboard protegido.
* Consumo de API con JWT.
* Renovación automática de sesión usando `refresh_token`.
* Cards de métricas principales.
* Gráficos con Recharts.
* Tabla de órdenes paginada.
* Filtros por:

  * Canal
  * Compañía
  * Tipo de entrega
  * Tipo de producto
* Selector de tamaño de página.
* Diseño responsive.

---

## Variables de entorno

Crear un archivo `.env.local` en la carpeta `frontend/`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

---

## Ejecutar localmente

Desde la carpeta `frontend/`:

```bash
npm install
npm run dev
```

La aplicación estará disponible en:

```txt
http://localhost:3000
```

---

## Usuario demo

```txt
Email: admin@test.com
Password: admin123
```

---

## Estructura principal

```txt
frontend/
├── app/
│   ├── dashboard/
│   ├── login/
│   ├── layout.tsx
│   └── page.tsx
├── components/
│   ├── orders/
│   └── stats/
├── lib/
│   ├── api.ts
│   ├── auth.ts
│   ├── config.ts
│   └── mappers/
├── types/
└── Dockerfile
```

---

## Decisiones técnicas

### Next.js

Se eligió Next.js porque permite construir una aplicación React moderna con buena estructura de rutas y soporte para TypeScript.

### Tailwind CSS

Se usó Tailwind CSS para crear una interfaz responsive de forma rápida y consistente.

### Axios

Se usó Axios para centralizar el consumo de la API, agregar el `access_token` automáticamente y manejar la renovación de sesión.

### Mappers

El backend responde usando `snake_case`, mientras que el frontend trabaja internamente con `camelCase`.

Por eso se agregaron mappers en `lib/mappers/`, para evitar que los componentes dependan directamente del contrato externo de la API.

### Recharts

Se usó Recharts para visualizar métricas de forma simple y clara dentro del dashboard.

---

## Docker

El frontend puede ejecutarse con Docker usando el `docker-compose.yml` de la raíz del proyecto.

Desde la raíz:

```bash
docker compose up --build
```

Esto levanta backend, frontend, base de datos, migraciones y seed automáticamente.
