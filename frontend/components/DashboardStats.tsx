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
    <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
      <div className="bg-white p-6 rounded-lg shadow">
        <h3 className="text-sm font-medium text-gray-500">Misiones Pendientes</h3>
        <p className="text-2xl font-bold text-gray-900 mt-2">{pendingMissions}</p>
      </div>
      <div className="bg-white p-6 rounded-lg shadow">
        <h3 className="text-sm font-medium text-gray-500">Transmutaciones Pendientes</h3>
        <p className="text-2xl font-bold text-gray-900 mt-2">{pendingTransmutations}</p>
      </div>
      <div className="bg-white p-6 rounded-lg shadow">
        <h3 className="text-sm font-medium text-gray-500">Auditorías Sin Resolver</h3>
        <p className="text-2xl font-bold text-gray-900 mt-2">{unresolvedAudits}</p>
      </div>
      <div className="bg-white p-6 rounded-lg shadow">
        <h3 className="text-sm font-medium text-gray-500">Costo Total</h3>
        <p className="text-2xl font-bold text-gray-900 mt-2">{totalCost.toFixed(2)}</p>
      </div>
    </div>
  )
}

