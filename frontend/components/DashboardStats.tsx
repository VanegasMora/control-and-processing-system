'use client'

import { Mission, Transmutation, Audit } from '@/lib/api'

interface DashboardStatsProps {
  missions: Mission[]
  transmutations: Transmutation[]
  audits: Audit[]
}

export default function DashboardStats({ missions, transmutations, audits }: DashboardStatsProps) {
  const pendingMissions = missions.filter(m => m.status === 'pending').length
  const pendingTransmutations = transmutations.filter(t => t.status === 'pending').length
  const unresolvedAudits = audits.filter(a => !a.resolved).length
  const totalCost = transmutations.reduce((sum, t) => sum + t.cost, 0)

  return (
    <div className="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
      <div className="bg-gradient-to-br from-slate-800 to-slate-900 border border-slate-700/50 rounded-2xl p-6 shadow-xl relative overflow-hidden">
        <div className="absolute top-0 right-0 w-32 h-32 bg-cyan-500/10 rounded-full -mr-16 -mt-16"></div>
        <div className="relative">
          <p className="text-slate-400 text-xs uppercase tracking-wider mb-2">Misiones Pendientes</p>
          <p className="text-4xl font-bold bg-gradient-to-r from-cyan-400 to-blue-400 bg-clip-text text-transparent">
            {pendingMissions}
          </p>
        </div>
      </div>
      <div className="bg-gradient-to-br from-slate-800 to-slate-900 border border-slate-700/50 rounded-2xl p-6 shadow-xl relative overflow-hidden">
        <div className="absolute top-0 right-0 w-32 h-32 bg-purple-500/10 rounded-full -mr-16 -mt-16"></div>
        <div className="relative">
          <p className="text-slate-400 text-xs uppercase tracking-wider mb-2">Transmutaciones Pendientes</p>
          <p className="text-4xl font-bold bg-gradient-to-r from-purple-400 to-pink-400 bg-clip-text text-transparent">
            {pendingTransmutations}
          </p>
        </div>
      </div>
      <div className="bg-gradient-to-br from-slate-800 to-slate-900 border border-slate-700/50 rounded-2xl p-6 shadow-xl relative overflow-hidden">
        <div className="absolute top-0 right-0 w-32 h-32 bg-yellow-500/10 rounded-full -mr-16 -mt-16"></div>
        <div className="relative">
          <p className="text-slate-400 text-xs uppercase tracking-wider mb-2">Auditorías Sin Resolver</p>
          <p className="text-4xl font-bold bg-gradient-to-r from-yellow-400 to-orange-400 bg-clip-text text-transparent">
            {unresolvedAudits}
          </p>
        </div>
      </div>
      <div className="bg-gradient-to-br from-slate-800 to-slate-900 border border-slate-700/50 rounded-2xl p-6 shadow-xl relative overflow-hidden">
        <div className="absolute top-0 right-0 w-32 h-32 bg-pink-500/10 rounded-full -mr-16 -mt-16"></div>
        <div className="relative">
          <p className="text-slate-400 text-xs uppercase tracking-wider mb-2">Costo Total</p>
          <p className="text-4xl font-bold bg-gradient-to-r from-pink-400 to-rose-400 bg-clip-text text-transparent">
            {totalCost.toFixed(2)}
          </p>
        </div>
      </div>
    </div>
  )
}

