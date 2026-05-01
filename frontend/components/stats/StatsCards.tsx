import type { StatsResponse } from "@/types/stats";

type StatsCardsProps = {
  stats: StatsResponse | null;
  isLoading: boolean;
};

export function StatsCards({ stats, isLoading }: StatsCardsProps) {
  const totalChannels = stats?.byChannel?.length ?? 0;
  const totalProductTypes = stats?.byProductType?.length ?? 0;

  return (
    <div className="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
      <MetricCard
        title="Total de órdenes"
        value={isLoading ? "—" : formatNumber(stats?.totalOrders ?? 0)}
      />

      <MetricCard
        title="Órdenes con error"
        value={isLoading ? "—" : `${stats?.errorPercentage ?? 0}%`}
      />

      <MetricCard
        title="Canales activos"
        value={isLoading ? "—" : formatNumber(totalChannels)}
      />

      <MetricCard
        title="Tipos de producto"
        value={isLoading ? "—" : formatNumber(totalProductTypes)}
      />
    </div>
  );
}

function MetricCard({ title, value }: { title: string; value: string }) {
  return (
    <article className="rounded-2xl border border-slate-200 bg-white p-4 shadow-sm sm:p-5">
      <p className="text-sm font-medium text-slate-500">{title}</p>
      <p className="mt-3 text-2xl font-bold tracking-tight text-slate-900">
        {value}
      </p>
    </article>
  );
}

function formatNumber(value: number) {
  return new Intl.NumberFormat("es-MX").format(value);
}