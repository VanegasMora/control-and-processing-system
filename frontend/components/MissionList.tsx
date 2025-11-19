'use client'

import { Mission } from '@/lib/api'
import api from '@/lib/api'

interface MissionListProps {
  missions: Mission[]
  onUpdate: () => void
  isSupervisor?: boolean
}

export default function MissionList({ missions, onUpdate, isSupervisor }: MissionListProps) {
  const updateStatus = async (id: number, status: string) => {
    try {
      await api.put(`/api/missions/${id}/status`, { status })
      onUpdate()
    } catch (error) {
      console.error('Error updating mission:', error)
    }
  }

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'pending':
        return 'bg-yellow-900/40 text-yellow-300 border border-yellow-700/50'
      case 'approved':
        return 'bg-blue-900/40 text-blue-300 border border-blue-700/50'
      case 'in_progress':
        return 'bg-purple-900/40 text-purple-300 border border-purple-700/50'
      case 'completed':
        return 'bg-emerald-900/40 text-emerald-300 border border-emerald-700/50'
      default:
        return 'bg-neutral-800/40 text-neutral-300 border border-neutral-700/50'
    }
  }

  return (
    <div className="space-y-4">
      {missions.length === 0 ? (
        <div className="text-center py-12 text-slate-500">
          <p>No hay misiones disponibles</p>
        </div>
      ) : (
        missions.map((mission) => (
          <div
            key={mission.id}
            className="bg-slate-700/30 border border-slate-600/50 rounded-xl p-5 hover:bg-slate-700/50 transition-all hover:border-slate-500/50"
          >
            <div className="flex items-start justify-between gap-4">
              <div className="flex-1 min-w-0">
                <div className="flex items-start gap-3 mb-3">
                  <div className="w-1.5 h-12 bg-gradient-to-b from-cyan-400 to-blue-500 rounded-full flex-shrink-0"></div>
                  <div className="flex-1">
                    <h4 className="text-lg font-semibold text-slate-200 mb-1">{mission.title}</h4>
                    <p className="text-sm text-slate-400 leading-relaxed">{mission.description}</p>
                  </div>
                </div>
                <div className="flex items-center gap-3 mt-4">
                  <span className={`px-3 py-1.5 rounded-lg text-xs font-semibold backdrop-blur-sm ${getStatusColor(mission.status)}`}>
                    {mission.status}
                  </span>
                </div>
              </div>
              {isSupervisor && mission.status === 'pending' && (
                <div className="flex flex-col gap-2 flex-shrink-0">
                  <button
                    onClick={() => updateStatus(mission.id, 'approved')}
                    className="bg-emerald-600/80 hover:bg-emerald-600 text-white px-4 py-2 rounded-lg text-sm transition-all backdrop-blur-sm border border-emerald-500/30 shadow-lg font-medium whitespace-nowrap"
                  >
                    Aprobar
                  </button>
                  <button
                    onClick={() => updateStatus(mission.id, 'cancelled')}
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
  )
}

