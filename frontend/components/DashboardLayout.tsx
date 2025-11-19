"use client";

import { useRouter } from "next/navigation";
import { ReactNode } from "react";

interface DashboardLayoutProps {
  children: ReactNode;
  user: any;
}

export default function DashboardLayout({
  children,
  user,
}: DashboardLayoutProps) {
  const router = useRouter();

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    router.push("/");
  };

  return (
    <div className="min-h-screen flex relative z-10">
      {/* Sidebar Lateral */}
      <aside className="w-72 bg-slate-900/95 border-r border-slate-700/50 backdrop-blur-sm flex flex-col fixed h-screen left-0 top-0">
        <div className="p-6 border-b border-slate-700/50">
          <h1 className="text-2xl font-bold bg-gradient-to-r from-cyan-400 via-purple-400 to-pink-400 bg-clip-text text-transparent mb-2">
            Amestris
          </h1>
          <p className="text-xs text-slate-500 uppercase tracking-wider">
            Alquimia Estatal
          </p>
        </div>
        
        <div className="flex-1 p-6">
          <div className="mb-8">
            <div className="flex items-center space-x-3 mb-2">
              <div className="w-10 h-10 rounded-full bg-gradient-to-br from-cyan-500 to-purple-500 flex items-center justify-center text-white font-bold">
                {user?.name?.charAt(0) || "U"}
              </div>
              <div>
                <p className="text-slate-200 font-medium text-sm">{user?.name}</p>
                <p className="text-slate-500 text-xs">
                  {user?.role === "supervisor" ? "Supervisor" : "Alquimista"}
                </p>
              </div>
            </div>
          </div>
        </div>

        <div className="p-6 border-t border-slate-700/50">
          <button
            onClick={handleLogout}
            className="w-full bg-gradient-to-r from-red-600 to-pink-600 hover:from-red-500 hover:to-pink-500 text-white px-4 py-3 rounded-xl transition-all shadow-lg shadow-red-500/20 font-medium"
          >
            Cerrar Sesión
          </button>
        </div>
      </aside>

      {/* Contenido Principal */}
      <main className="flex-1 ml-72 min-h-screen">
        <div className="p-8">
          {children}
        </div>
      </main>
    </div>
  );
}
