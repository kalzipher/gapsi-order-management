import type { OrdersResponse } from "@/types/orders";

type OrdersTableProps = {
  orders: OrdersResponse | null;
  isLoading: boolean;
  onNextPage: () => void;
  onPreviousPage: () => void;
  onPageSizeChange: (pageSize: number) => void;
};

export function OrdersTable({
  orders,
  isLoading,
  onNextPage,
  onPreviousPage,
  onPageSizeChange,
}: OrdersTableProps) {
  const rows = orders?.data ?? [];
  const pagination = orders?.pagination;

  const pageSize = pagination?.pageSize ?? 10;
  const rowHeight = 56;
  const tableHeaderHeight = 44;
  const tableBodyMinHeight = pageSize * rowHeight;
  const tableMinHeight = tableBodyMinHeight + tableHeaderHeight;

  return (
    <div className="mt-6 rounded-2xl border border-slate-200 bg-white shadow-sm">
      <div className="flex flex-col gap-3 border-b border-slate-200 p-4 sm:flex-row sm:items-center sm:justify-between sm:p-6">
        <div>
          <h3 className="text-base font-semibold text-slate-900">
            Listado de órdenes
          </h3>

          <p className="mt-1 text-sm text-slate-500">
            Consulta órdenes paginadas desde el API.
          </p>
        </div>

        <div className="flex flex-col gap-3 sm:items-end">
          {pagination && (
            <p className="text-sm text-slate-500">
              {formatNumber(pagination.total)} órdenes
            </p>
          )}

          <label className="flex items-center gap-2 text-sm text-slate-600">
            Mostrar
            <select
              value={pageSize}
              disabled={isLoading}
              onChange={(event) => onPageSizeChange(Number(event.target.value))}
              className="h-9 rounded-lg border border-slate-300 bg-white px-2 text-sm text-slate-700 outline-none transition focus:border-slate-900 disabled:cursor-not-allowed disabled:bg-slate-50 disabled:text-slate-400"
            >
              <option value={10}>10</option>
              <option value={20}>20</option>
              <option value={50}>50</option>
              <option value={100}>100</option>
            </select>
            filas
          </label>
        </div>
      </div>

      <div className="overflow-x-auto">
        <div style={{ minHeight: tableMinHeight }}>
          <table className="w-full min-w-[980px] border-collapse text-left text-sm">
            <thead className="bg-slate-50 text-xs uppercase tracking-wide text-slate-500">
              <tr className="h-11">
                <th className="px-4 py-3 font-semibold">No. Pedido</th>
                <th className="px-4 py-3 font-semibold">Canal</th>
                <th className="px-4 py-3 font-semibold">SKU</th>
                <th className="px-4 py-3 font-semibold">Fecha estimada</th>
                <th className="px-4 py-3 font-semibold">Entrega</th>
                <th className="px-4 py-3 font-semibold">Producto</th>
                <th className="px-4 py-3 font-semibold">Cantidad</th>
                <th className="px-4 py-3 font-semibold">Fecha compra</th>
                <th className="px-4 py-3 font-semibold">Estado</th>
              </tr>
            </thead>

            <tbody className="divide-y divide-slate-100">
              {isLoading ? (
                <tr>
                  <td
                    colSpan={9}
                    style={{ height: tableBodyMinHeight }}
                    className="px-4 py-10 text-center text-slate-500"
                  >
                    Cargando órdenes...
                  </td>
                </tr>
              ) : rows.length === 0 ? (
                <tr>
                  <td
                    colSpan={9}
                    style={{ height: tableBodyMinHeight }}
                    className="px-4 py-10 text-center text-slate-500"
                  >
                    No hay órdenes para mostrar.
                  </td>
                </tr>
              ) : (
                rows.map((order) => (
                  <tr
                    key={order.id}
                    className="h-14 transition hover:bg-slate-50"
                  >
                    <td className="px-4 py-3 font-medium text-slate-900">
                      {fallback(order.noPedido)}
                    </td>

                    <td className="px-4 py-3 text-slate-700">
                      {fallback(order.canal)}
                    </td>

                    <td className="px-4 py-3 text-slate-700">
                      {fallback(order.sku)}
                    </td>

                    <td className="max-w-[260px] px-4 py-3 text-slate-700">
                      <span className="line-clamp-2">
                        {fallback(order.fechaEstimada)}
                      </span>
                    </td>

                    <td className="px-4 py-3 text-slate-700">
                      {fallback(order.fulfillmentType)}
                    </td>

                    <td className="px-4 py-3 text-slate-700">
                      {fallback(order.productType)}
                    </td>

                    <td className="px-4 py-3 text-slate-700">
                      {order.cantidad ?? "—"}
                    </td>

                    <td className="px-4 py-3 text-slate-700">
                      {formatDate(order.fechaCompra)}
                    </td>

                    <td className="px-4 py-3">
                      {order.hasError ? (
                        <span className="rounded-full bg-red-50 px-3 py-1 text-xs font-medium text-red-700">
                          Error
                        </span>
                      ) : (
                        <span className="rounded-full bg-emerald-50 px-3 py-1 text-xs font-medium text-emerald-700">
                          OK
                        </span>
                      )}
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
      </div>

      <div className="flex flex-col gap-3 border-t border-slate-200 p-4 sm:flex-row sm:items-center sm:justify-between sm:p-6">
        <p className="text-sm text-slate-500">
          Página {pagination?.page ?? 1} de {pagination?.totalPages ?? 1}
        </p>

        <div className="flex gap-2">
          <button
            type="button"
            onClick={onPreviousPage}
            disabled={isLoading || !pagination || pagination.page <= 1}
            className="rounded-xl border border-slate-300 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-100 disabled:cursor-not-allowed disabled:opacity-50"
          >
            Anterior
          </button>

          <button
            type="button"
            onClick={onNextPage}
            disabled={
              isLoading ||
              !pagination ||
              pagination.page >= pagination.totalPages
            }
            className="rounded-xl border border-slate-300 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-100 disabled:cursor-not-allowed disabled:opacity-50"
          >
            Siguiente
          </button>
        </div>
      </div>
    </div>
  );
}

function fallback(value: string | null | undefined) {
  return value && value.trim() !== "" ? value : "—";
}

function formatNumber(value: number) {
  return new Intl.NumberFormat("es-MX").format(value);
}

function formatDate(value: string | null) {
  if (!value) return "—";

  const date = new Date(value);

  if (Number.isNaN(date.getTime())) {
    return "—";
  }

  return new Intl.DateTimeFormat("es-MX", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  }).format(date);
}