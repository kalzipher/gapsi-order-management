"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";

import { login } from "@/lib/api";
import { saveSession } from "@/lib/auth";
import type { LoginRequest } from "@/types/auth";

export default function LoginPage() {
  const router = useRouter();
  const [error, setError] = useState("");

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<LoginRequest>({
    defaultValues: {
      email: "admin@test.com",
      password: "admin123",
    },
  });

  async function onSubmit(values: LoginRequest) {
    setError("");

    try {
      const session = await login(values);
      saveSession(session);
      router.push("/dashboard");
    } catch {
      setError("Credenciales inválidas o error al iniciar sesión.");
    }
  }

  return (
    <main className="min-h-screen bg-white md:flex md:items-center md:justify-center md:bg-slate-100 md:px-4">
      <section className="min-h-screen w-full px-6 py-10 md:min-h-0 md:max-w-md md:rounded-2xl md:bg-white md:p-8 md:shadow-sm">
        <div className="mb-8">
          <p className="text-sm font-medium text-slate-500">Order Management</p>
          <h1 className="mt-2 text-2xl font-bold tracking-tight text-slate-900 sm:text-3xl">
            Iniciar sesión
          </h1>
          <p className="mt-2 text-sm text-slate-500">
            Accede al dashboard de órdenes.
          </p>
        </div>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-5">
          <div>
            <label className="mb-2 block text-sm font-medium text-slate-700">
              Email
            </label>
            <input
              type="email"
              className="w-full rounded-xl border border-slate-300 px-4 py-3 text-sm outline-none transition focus:border-slate-900"
              placeholder="admin@test.com"
              {...register("email", {
                required: "El email es requerido",
              })}
            />
            {errors.email && (
              <p className="mt-2 text-sm text-red-600">
                {errors.email.message}
              </p>
            )}
          </div>

          <div>
            <label className="mb-2 block text-sm font-medium text-slate-700">
              Contraseña
            </label>
            <input
              type="password"
              className="w-full rounded-xl border border-slate-300 px-4 py-3 text-sm outline-none transition focus:border-slate-900"
              placeholder="admin123"
              {...register("password", {
                required: "La contraseña es requerida",
              })}
            />
            {errors.password && (
              <p className="mt-2 text-sm text-red-600">
                {errors.password.message}
              </p>
            )}
          </div>

          {error && (
            <div className="rounded-xl bg-red-50 px-4 py-3 text-sm text-red-700">
              {error}
            </div>
          )}

          <button
            type="submit"
            disabled={isSubmitting}
            className="w-full rounded-xl bg-slate-900 px-4 py-3 text-sm font-semibold text-white transition hover:bg-slate-800 disabled:cursor-not-allowed disabled:opacity-60"
          >
            {isSubmitting ? "Ingresando..." : "Ingresar"}
          </button>
        </form>

        <div className="mt-6 rounded-xl bg-slate-50 p-4 text-sm text-slate-600">
          <p className="font-medium text-slate-700">Usuario demo</p>
          <p className="mt-1">admin@test.com / admin123</p>
        </div>
      </section>
    </main>
  );
}