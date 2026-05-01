"use client";

import { useEffect, useMemo, useState } from "react";
import { useRouter } from "next/navigation";

import { FiltersPanel } from "@/components/orders/FiltersPanel";
import { getOrderFilters, getOrders, getStats } from "@/lib/api";
import { clearSession, getUser, isAuthenticated } from "@/lib/auth";
import { StatsCards } from "@/components/stats/StatsCards";
import { StatsCharts } from "@/components/stats/StatsCharts";
import type { User } from "@/types/auth";
import type { StatsFilters, StatsResponse } from "@/types/stats";


import { OrdersTable } from "@/components/orders/OrdersTable";
import type {
    OrderFiltersOptions,
    OrdersFilters,
    OrdersResponse,
} from "@/types/orders";

export default function DashboardPage() {
    const router = useRouter();

    const [user, setUser] = useState<User | null>(null);
    const [isCheckingSession, setIsCheckingSession] = useState(true);

    const [ordersFilters, setOrdersFilters] = useState<OrdersFilters>({
        page: 1,
        pageSize: 10,
    });
    const [filterOptions, setFilterOptions] =
        useState<OrderFiltersOptions | null>(null);
    const [isLoadingFilters, setIsLoadingFilters] = useState(true);
    const [filtersError, setFiltersError] = useState("");

    const [stats, setStats] = useState<StatsResponse | null>(null);
    const [isLoadingStats, setIsLoadingStats] = useState(true);
    const [statsError, setStatsError] = useState("");


    const [orders, setOrders] = useState<OrdersResponse | null>(null);
    const [isLoadingOrders, setIsLoadingOrders] = useState(true);
    const [ordersError, setOrdersError] = useState("");

    const statsFilters: StatsFilters = useMemo(() => ({
        canal: ordersFilters.canal,
        company: ordersFilters.company,
        fulfillmentType: ordersFilters.fulfillmentType,
        productType: ordersFilters.productType,
    }), [ordersFilters]);

    useEffect(() => {
        if (!isAuthenticated()) {
            router.replace("/login");
            return;
        }

        setUser(getUser());
        setIsCheckingSession(false);
    }, [router]);

    useEffect(() => {
        if (isCheckingSession) return;

        async function loadStats() {
            setIsLoadingStats(true);
            setStatsError("");

            try {
                const data = await getStats(statsFilters);
                setStats(data);
            } catch {
                setStatsError("No fue posible cargar las métricas.");
            } finally {
                setIsLoadingStats(false);
            }
        }

        loadStats();
    }, [statsFilters, isCheckingSession]);

    useEffect(() => {
        if (isCheckingSession) return;

        async function loadOrders() {
            setIsLoadingOrders(true);
            setOrdersError("");

            try {
                const data = await getOrders(ordersFilters);
                setOrders(data);
            } catch (error) {
                setOrdersError("No fue posible cargar las órdenes.");
            } finally {
                setIsLoadingOrders(false);
            }
        }

        loadOrders();
    }, [ordersFilters, isCheckingSession]);

    useEffect(() => {
        if (isCheckingSession) return;

        async function loadFilters() {
            setIsLoadingFilters(true);
            setFiltersError("");

            try {
                const data = await getOrderFilters();
                setFilterOptions(data);
            } catch (error) {
                setFiltersError("No fue posible cargar los filtros.");
            } finally {
                setIsLoadingFilters(false);
            }
        }

        loadFilters();
    }, [isCheckingSession]);

    function handleFiltersChange(nextFilters: OrdersFilters) {
        setOrdersFilters(nextFilters);
    }

    function handleClearFilters() {
        setOrdersFilters({
            page: 1,
            pageSize: ordersFilters.pageSize,
        });
    }

    function handleNextPage() {
        setOrdersFilters((current) => ({
            ...current,
            page: current.page + 1,
        }));
    }

    function handlePreviousPage() {
        setOrdersFilters((current) => ({
            ...current,
            page: Math.max(1, current.page - 1),
        }));
    }

    function handlePageSizeChange(pageSize: number) {
        setOrdersFilters((current) => ({
            ...current,
            page: 1,
            pageSize,
        }));
    }

    function handleLogout() {
        clearSession();
        router.replace("/login");
    }

    if (isCheckingSession) {
        return (
            <main className="flex min-h-screen items-center justify-center bg-slate-50 px-4">
                <p className="text-sm text-slate-500">Validando sesión...</p>
            </main>
        );
    }

    return (
        <main className="min-h-screen bg-slate-50">
            <header className="sticky top-0 z-10 border-b border-slate-200 bg-white/90 backdrop-blur">
                <div className="mx-auto flex max-w-7xl flex-col gap-4 px-4 py-4 sm:px-6 lg:flex-row lg:items-center lg:justify-between lg:px-8">
                    <div>
                        <p className="text-sm font-medium text-slate-500">
                            Order Management
                        </p>
                        <h1 className="mt-1 text-xl font-bold tracking-tight text-slate-900 sm:text-2xl">
                            Dashboard de órdenes
                        </h1>
                    </div>

                    <div className="flex items-center justify-between gap-3 lg:justify-end">
                        <div className="min-w-0">
                            <p className="truncate text-sm font-medium text-slate-900">
                                {user?.name || "Admin"}
                            </p>
                            <p className="truncate text-xs text-slate-500">
                                {user?.email || "admin@test.com"}
                            </p>
                        </div>

                        <button
                            type="button"
                            onClick={handleLogout}
                            className="rounded-xl border border-slate-300 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-100"
                        >
                            Salir
                        </button>
                    </div>
                </div>
            </header>

            <section className="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
                <div className="mb-6">
                    <h2 className="text-lg font-semibold text-slate-900">
                        Resumen operativo
                    </h2>
                    <p className="mt-1 text-sm text-slate-500">
                        Monitorea órdenes, errores y distribución por canal, entrega y tipo
                        de producto.
                    </p>
                </div>

                {statsError && (
                    <div className="mb-6 rounded-xl bg-red-50 px-4 py-3 text-sm text-red-700">
                        {statsError}
                    </div>
                )}

                <StatsCards stats={stats} isLoading={isLoadingStats} />

                {filtersError && (
                    <div className="mt-6 rounded-xl bg-red-50 px-4 py-3 text-sm text-red-700">
                        {filtersError}
                    </div>
                )}

                <FiltersPanel
                    filters={ordersFilters}
                    options={filterOptions}
                    isLoading={isLoadingFilters}
                    onChange={handleFiltersChange}
                    onClear={handleClearFilters}
                />

                <StatsCharts stats={stats} isLoading={isLoadingStats} />

                {ordersError && (
                    <div className="mt-6 rounded-xl bg-red-50 px-4 py-3 text-sm text-red-700">
                        {ordersError}
                    </div>
                )}

                <OrdersTable
                    orders={orders}
                    isLoading={isLoadingOrders}
                    onNextPage={handleNextPage}
                    onPreviousPage={handlePreviousPage}
                    onPageSizeChange={handlePageSizeChange}
                />
            </section>
        </main>
    );
}