"use client";

import {
  Bar,
  BarChart,
  CartesianGrid,
  Cell,
  Pie,
  PieChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";

import type { StatsResponse } from "@/types/stats";

type StatsChartsProps = {
  stats: StatsResponse | null;
  isLoading: boolean;
};

const COLORS = ["#3b82f6", "#22c55e", "#f59e0b", "#ef4444", "#8b5cf6"];

export function StatsCharts({ stats, isLoading }: StatsChartsProps) {
  if (isLoading) {
    return (
      <div className="mt-6 grid gap-6 lg:grid-cols-2">
        <ChartSkeleton title="Órdenes por canal" />
        <ChartSkeleton title="Órdenes por tipo de producto" />
      </div>
    );
  }

  return (
    <div className="mt-6 grid gap-6 lg:grid-cols-2">
      <section className="rounded-2xl border border-slate-200 bg-white p-4 shadow-sm sm:p-6">
        <h3 className="text-base font-semibold text-slate-900">
          Órdenes por canal
        </h3>

        <div className="mt-4 min-w-0 min-h-0 h-72">
          <ResponsiveContainer width="100%" aspect={2}>
            <BarChart data={stats?.byChannel ?? []}>
              <CartesianGrid strokeDasharray="3 3" vertical={false} />
              <XAxis dataKey="name" tick={{ fontSize: 12 }} />
              <YAxis tick={{ fontSize: 12 }} />
              <Tooltip />
              <Bar dataKey="total" radius={[8, 8, 0, 0]} fill="#3b82f6" />
            </BarChart>
          </ResponsiveContainer>
        </div>
      </section>

      <section className="rounded-2xl border border-slate-200 bg-white p-4 shadow-sm sm:p-6">
        <h3 className="text-base font-semibold text-slate-900">
          Órdenes por tipo de producto
        </h3>

        <div className="mt-4 h-72 min-w-0 min-h-0">
          <ResponsiveContainer width="100%" aspect={2}>
            <PieChart>
              <Pie
                data={stats?.byProductType ?? []}
                dataKey="total"
                nameKey="name"
                outerRadius={95}
                label
              >
                {(stats?.byProductType ?? []).map((item, index) => (
                  <Cell
                    key={item.name}
                    fill={COLORS[index % COLORS.length]}
                  />
                ))}
              </Pie>
              <Tooltip />
            </PieChart>
          </ResponsiveContainer>
        </div>
      </section>
    </div>
  );
}

function ChartSkeleton({ title }: { title: string }) {
  return (
    <section className="rounded-2xl border border-slate-200 bg-white p-4 shadow-sm sm:p-6">
      <h3 className="text-base font-semibold text-slate-900">{title}</h3>
      <div className="mt-4 flex h-72 animate-pulse items-center justify-center rounded-xl bg-slate-100">
        <p className="text-sm text-slate-400">Cargando gráfico...</p>
      </div>
    </section>
  );
}