"use client";

import { Transmutation } from "@/lib/api";
import api from "@/lib/api";

interface TransmutationListProps {
  transmutations: Transmutation[];
  onUpdate: () => void;
  isSupervisor?: boolean;
}

export default function TransmutationList({
  transmutations,
  onUpdate,
  isSupervisor,
}: TransmutationListProps) {
  const updateStatus = async (id: number, status: string, result?: string) => {
    try {
      await api.put(`/api/transmutations/${id}/status`, { status, result });
      onUpdate();
    } catch (error) {
      console.error("Error updating transmutation:", error);
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case "pending":
        return "bg-yellow-900/40 text-yellow-300 border border-yellow-700/50";
      case "approved":
        return "bg-blue-900/40 text-blue-300 border border-blue-700/50";
      case "completed":
        return "bg-emerald-900/40 text-emerald-300 border border-emerald-700/50";
      case "rejected":
        return "bg-red-900/40 text-red-300 border border-red-700/50";
      default:
        return "bg-neutral-800/40 text-neutral-300 border border-neutral-700/50";
    }
  };

  return (
    <div className="space-y-4">
      {transmutations.length === 0 ? (
        <div className="text-center py-12 text-slate-500">
          <p>No hay transmutaciones disponibles</p>
        </div>
      ) : (
        transmutations.map((transmutation) => (
          <div
            key={transmutation.id}
            className="bg-slate-700/30 border border-slate-600/50 rounded-xl p-5 hover:bg-slate-700/50 transition-all hover:border-slate-500/50"
          >
            <div className="flex items-start justify-between gap-4">
              <div className="flex-1 min-w-0">
                <div className="flex items-start gap-3 mb-3">
                  <div className="w-1.5 h-12 bg-gradient-to-b from-purple-400 to-pink-500 rounded-full flex-shrink-0"></div>
                  <div className="flex-1">
                    <h4 className="text-lg font-semibold text-slate-200 mb-1">
                      Transmutación #{transmutation.id}
                    </h4>
                    <p className="text-sm text-slate-400 leading-relaxed mb-3">
                      {transmutation.description}
                    </p>
                    <div className="flex flex-wrap gap-3 text-xs">
                      <div className="bg-slate-800/50 px-3 py-1.5 rounded-lg border border-slate-600/50">
                        <span className="text-slate-400">Costo: </span>
                        <span className="text-purple-300 font-semibold">
                          {transmutation.cost.toFixed(2)}
                        </span>
                      </div>
                      <div className="bg-slate-800/50 px-3 py-1.5 rounded-lg border border-slate-600/50">
                        <span className="text-slate-400">Materiales: </span>
                        <span className="text-blue-300 font-semibold">
                          {transmutation.input_materials.length}
                        </span>
                      </div>
                    </div>
                    {transmutation.result && (
                      <div className="mt-3 bg-emerald-900/20 border border-emerald-700/30 rounded-lg px-3 py-2">
                        <p className="text-xs text-emerald-300">
                          {transmutation.result}
                        </p>
                      </div>
                    )}
                  </div>
                </div>
                <div className="flex items-center gap-3 mt-4">
                  <span
                    className={`px-3 py-1.5 rounded-lg text-xs font-semibold backdrop-blur-sm ${getStatusColor(
                      transmutation.status
                    )}`}
                  >
                    {transmutation.status}
                  </span>
                </div>
              </div>
              {isSupervisor && transmutation.status === "pending" && (
                <div className="flex flex-col gap-2 flex-shrink-0">
                  <button
                    onClick={() => updateStatus(transmutation.id, "approved")}
                    className="bg-emerald-600/80 hover:bg-emerald-600 text-white px-4 py-2 rounded-lg text-sm transition-all backdrop-blur-sm border border-emerald-500/30 shadow-lg font-medium whitespace-nowrap"
                  >
                    Aprobar
                  </button>
                  <button
                    onClick={() =>
                      updateStatus(
                        transmutation.id,
                        "rejected",
                        "Transmutación rechazada"
                      )
                    }
                    className="bg-red-600/80 hover:bg-red-600 text-white px-4 py-2 rounded-lg text-sm transition-all backdrop-blur-sm border border-red-500/30 shadow-lg font-medium whitespace-nowrap"
                  >
                    Rechazar
                  </button>
                </div>
              )}
            </div>
          </div>
        ))
      )}
    </div>
  );
}
