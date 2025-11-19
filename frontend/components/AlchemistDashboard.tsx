"use client";

import { useEffect, useState } from "react";
import api, { Mission, Transmutation, Material } from "@/lib/api";
import MissionList from "./MissionList";
import TransmutationList from "./TransmutationList";
import CreateMissionModal from "./CreateMissionModal";
import CreateTransmutationModal from "./CreateTransmutationModal";

export default function AlchemistDashboard() {
  const [missions, setMissions] = useState<Mission[]>([]);
  const [transmutations, setTransmutations] = useState<Transmutation[]>([]);
  const [materials, setMaterials] = useState<Material[]>([]);
  const [showMissionModal, setShowMissionModal] = useState(false);
  const [showTransmutationModal, setShowTransmutationModal] = useState(false);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadData();
    setupWebSocket();
  }, []);

  const loadData = async () => {
    try {
      const [missionsRes, transmutationsRes, materialsRes] = await Promise.all([
        api.get("/api/missions"),
        api.get("/api/transmutations"),
        api.get("/api/materials"),
      ]);
      setMissions(missionsRes.data);
      setTransmutations(transmutationsRes.data);
      setMaterials(materialsRes.data);
    } catch (error) {
      console.error("Error loading data:", error);
    } finally {
      setLoading(false);
    }
  };

  const setupWebSocket = () => {
    const token = localStorage.getItem("token");
    if (!token) return;

    const ws = new WebSocket(`ws://localhost:8000/api/ws?token=${token}`);

    ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      if (
        message.type === "mission_status_changed" ||
        message.type === "transmutation_status_changed"
      ) {
        loadData();
      }
    };

    ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };
  };

  if (loading) {
    return (
      <div className="text-center py-12">
        <div className="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-cyan-500"></div>
        <p className="mt-4 text-slate-400">Cargando...</p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header con acciones flotantes */}
      <div className="flex items-start justify-between mb-8">
        <div>
          <h2 className="text-4xl font-bold bg-gradient-to-r from-cyan-400 via-purple-400 to-pink-400 bg-clip-text text-transparent mb-2">
            Panel de Alquimista
          </h2>
          <p className="text-slate-400 text-base">
            Gestiona tus misiones y transmutaciones
          </p>
        </div>
        <div className="flex flex-col gap-2">
          <button
            onClick={() => setShowMissionModal(true)}
            className="bg-gradient-to-r from-blue-600 to-cyan-600 hover:from-blue-500 hover:to-cyan-500 text-white px-6 py-3 rounded-xl transition-all shadow-lg shadow-blue-500/30 font-semibold text-sm"
          >
            + Nueva Misión
          </button>
          <button
            onClick={() => setShowTransmutationModal(true)}
            className="bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-500 hover:to-pink-500 text-white px-6 py-3 rounded-xl transition-all shadow-lg shadow-purple-500/30 font-semibold text-sm"
          >
            + Nueva Transmutación
          </button>
        </div>
      </div>

      {/* Layout vertical con secciones apiladas */}
      <div className="space-y-6">
        {/* Sección de Misiones */}
        <section className="bg-slate-800/50 border border-slate-700/50 rounded-2xl shadow-2xl overflow-hidden">
          <div className="bg-gradient-to-r from-cyan-600/20 to-blue-600/20 border-b border-slate-700/50 px-6 py-4">
            <div className="flex items-center justify-between">
              <h3 className="text-2xl font-bold text-slate-200 flex items-center gap-3">
                <div className="w-2 h-8 bg-gradient-to-b from-cyan-400 to-blue-500 rounded-full"></div>
                Misiones
              </h3>
              <span className="text-slate-400 text-sm bg-slate-700/50 px-3 py-1 rounded-full">
                {missions.length} total
              </span>
            </div>
          </div>
          <div className="p-6">
            <MissionList missions={missions} onUpdate={loadData} />
          </div>
        </section>

        {/* Sección de Transmutaciones */}
        <section className="bg-slate-800/50 border border-slate-700/50 rounded-2xl shadow-2xl overflow-hidden">
          <div className="bg-gradient-to-r from-purple-600/20 to-pink-600/20 border-b border-slate-700/50 px-6 py-4">
            <div className="flex items-center justify-between">
              <h3 className="text-2xl font-bold text-slate-200 flex items-center gap-3">
                <div className="w-2 h-8 bg-gradient-to-b from-purple-400 to-pink-500 rounded-full"></div>
                Transmutaciones
              </h3>
              <span className="text-slate-400 text-sm bg-slate-700/50 px-3 py-1 rounded-full">
                {transmutations.length} total
              </span>
            </div>
          </div>
          <div className="p-6">
            <TransmutationList
              transmutations={transmutations}
              onUpdate={loadData}
            />
          </div>
        </section>
      </div>

      {showMissionModal && (
        <CreateMissionModal
          onClose={() => setShowMissionModal(false)}
          onSuccess={() => {
            setShowMissionModal(false);
            loadData();
          }}
        />
      )}

      {showTransmutationModal && (
        <CreateTransmutationModal
          materials={materials}
          onClose={() => setShowTransmutationModal(false)}
          onSuccess={() => {
            setShowTransmutationModal(false);
            loadData();
          }}
        />
      )}
    </div>
  );
}
