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
        return 'bg-yellow-100 text-yellow-800'
      case 'approved':
        return 'bg-blue-100 text-blue-800'
      case 'in_progress':
        return 'bg-purple-100 text-purple-800'
      case 'completed':
        return 'bg-green-100 text-green-800'
      default:
        return 'bg-gray-100 text-gray-800'
    }
  }

  return (
    <div className="bg-white rounded-lg shadow overflow-hidden">
      <ul className="divide-y divide-gray-200">
        {missions.map((mission) => (
          <li key={mission.id} className="p-4 hover:bg-gray-50">
            <div className="flex justify-between items-start">
              <div className="flex-1">
                <h4 className="text-lg font-medium text-gray-900">{mission.title}</h4>
                <p className="text-sm text-gray-500 mt-1">{mission.description}</p>
                <span className={`inline-block mt-2 px-2 py-1 rounded text-xs font-medium ${getStatusColor(mission.status)}`}>
                  {mission.status}
                </span>
              </div>
              {isSupervisor && mission.status === 'pending' && (
                <div className="ml-4 space-x-2">
                  <button
                    onClick={() => updateStatus(mission.id, 'approved')}
                    className="bg-green-600 text-white px-3 py-1 rounded text-sm hover:bg-green-700"
                  >
                    Aprobar
                  </button>
                  <button
                    onClick={() => updateStatus(mission.id, 'cancelled')}
                    className="bg-red-600 text-white px-3 py-1 rounded text-sm hover:bg-red-700"
                  >
                    Rechazar
                  </button>
                </div>
              )}
            </div>
          </li>
        ))}
        {missions.length === 0 && (
          <li className="p-4 text-center text-gray-500">No hay misiones</li>
        )}
      </ul>
    </div>
  )
}

