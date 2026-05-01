CREATE TABLE IF NOT EXISTS orders (
  id TEXT PRIMARY KEY,
  canal TEXT,
  cantidad INTEGER,
  company TEXT,
  cp TEXT,
  created_at TIMESTAMPTZ,
  days_to_delivery INTEGER,
  error_code TEXT,
  error_message TEXT,
  fecha_compra TIMESTAMPTZ,
  fecha_estimada TEXT,
  fulfillment_type TEXT,
  is_flash BOOLEAN,
  is_marketplace BOOLEAN,
  no_pedido TEXT,
  plan TEXT,
  product_type TEXT,
  sku TEXT,
  store_selected TEXT,
  tipo_pago TEXT,
  edd1 TEXT,
  edd2 TEXT
);

CREATE INDEX IF NOT EXISTS idx_orders_canal
ON orders(canal);

CREATE INDEX IF NOT EXISTS idx_orders_company
ON orders(company);

CREATE INDEX IF NOT EXISTS idx_orders_fulfillment_type
ON orders(fulfillment_type);

CREATE INDEX IF NOT EXISTS idx_orders_product_type
ON orders(product_type);

CREATE INDEX IF NOT EXISTS idx_orders_created_at
ON orders(created_at);

CREATE INDEX IF NOT EXISTS idx_orders_no_pedido
ON orders(no_pedido);