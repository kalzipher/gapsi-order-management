import type { OrderFiltersOptions, OrdersFilters } from "@/types/orders";

type FiltersPanelProps = {
  filters: OrdersFilters;
  options: OrderFiltersOptions | null;
  isLoading: boolean;
  onChange: (filters: OrdersFilters) => void;
  onClear: () => void;
};

export function FiltersPanel({
  filters,
  options,
  isLoading,
  onChange,
  onClear,
}: FiltersPanelProps) {
  function updateFilter<K extends keyof OrdersFilters>(
    key: K,
    value: OrdersFilters[K]
  ) {
    onChange({
      ...filters,
      [key]: value,
      page: 1,
    });
  }

  return (
    <div className="mt-6 rounded-2xl border border-slate-200 bg-white p-4 shadow-sm sm:p-6">
      <div className="mb-4 flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
        <div>
          <h3 className="text-base font-semibold text-slate-900">Filtros</h3>

          <p className="mt-1 text-sm text-slate-500">
            Filtra la tabla y las métricas por canal, compañía, entrega y tipo
            de producto.
          </p>
        </div>

        <button
          type="button"
          onClick={onClear}
          disabled={isLoading}
          className="rounded-xl border border-slate-300 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-100 disabled:cursor-not-allowed disabled:opacity-50"
        >
          Limpiar filtros
        </button>
      </div>

      <div className="grid gap-3 sm:grid-cols-2 lg:grid-cols-4">
        <FilterSelect
          label="Canal"
          value={filters.canal ?? ""}
          disabled={isLoading}
          options={options?.channels ?? []}
          onChange={(value) => updateFilter("canal", value || undefined)}
        />

        <FilterSelect
          label="Compañía"
          value={filters.company ?? ""}
          disabled={isLoading}
          options={options?.companies ?? []}
          onChange={(value) => updateFilter("company", value || undefined)}
        />

        <FilterSelect
          label="Tipo de entrega"
          value={filters.fulfillmentType ?? ""}
          disabled={isLoading}
          options={options?.fulfillmentTypes ?? []}
          onChange={(value) =>
            updateFilter("fulfillmentType", value || undefined)
          }
        />

        <FilterSelect
          label="Tipo de producto"
          value={filters.productType ?? ""}
          disabled={isLoading}
          options={options?.productTypes ?? []}
          onChange={(value) => updateFilter("productType", value || undefined)}
        />
      </div>
    </div>
  );
}

type FilterSelectProps = {
  label: string;
  value: string;
  options: string[];
  disabled?: boolean;
  onChange: (value: string) => void;
};

function FilterSelect({
  label,
  value,
  options,
  disabled,
  onChange,
}: FilterSelectProps) {
  return (
    <div>
      <label className="mb-2 block text-sm font-medium text-slate-700">
        {label}
      </label>

      <select
        value={value}
        disabled={disabled}
        onChange={(event) => onChange(event.target.value)}
        className="h-11 w-full rounded-xl border border-slate-300 bg-white px-3 text-sm text-slate-700 outline-none transition focus:border-slate-900 disabled:cursor-not-allowed disabled:bg-slate-50 disabled:text-slate-400"
      >
        <option value="">Todos</option>

        {options.map((option) => (
          <option key={option} value={option}>
            {option}
          </option>
        ))}
      </select>
    </div>
  );
}