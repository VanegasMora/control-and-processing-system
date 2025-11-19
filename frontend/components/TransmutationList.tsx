'use client'

import { Transmutation } from '@/lib/api'
import api from '@/lib/api'

interface TransmutationListProps {
  transmutations: Transmutation[]
  onUpdate: () => void
  isSupervisor?: boolean
}

export default function TransmutationList({ transmutations, onUpdate, isSupervisor }: TransmutationListProps) {
  const updateStatus = async (id: number, status: string, result?: string) => {
    try {
      await api.put(`/api/transmutations/${id}/status`, { status, result })
      onUpdate()
    } catch (error) {
      console.error('Error updating transmutation:', error)
    }
  }

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'pending':
        return 'bg-yellow-100 text-yellow-800'
      case 'approved':
        return 'bg-blue-100 text-blue-800'
      case 'completed':
        return 'bg-green-100 text-green-800'
      case 'rejected':
        return 'bg-red-100 text-red-800'
      default:
        return 'bg-gray-100 text-gray-800'
    }
  }

  return (
    <div className="bg-white rounded-lg shadow overflow-hidden">
      <ul className="divide-y divide-gray-200">
        {transmutations.map((transmutation) => (
          <li key={transmutation.id} className="p-4 hover:bg-gray-50">
            <div className="flex justify-between items-start">
              <div className="flex-1">
                <h4 className="text-lg font-medium text-gray-900">
                  Transmutación #{transmutation.id}
                </h4>
                <p className="text-sm text-gray-500 mt-1">{transmutation.description}</p>
                <div className="mt-2 space-y-1">
                  <p className="text-xs text-gray-600">
                    Costo: {transmutation.cost.toFixed(2)}
                  </p>
                  <p className="text-xs text-gray-600">
                    Materiales de entrada: {transmutation.input_materials.length}
                  </p>
                  {transmutation.result && (
                    <p className="text-xs text-green-600">{transmutation.result}</p>
                  )}
                </div>
                <span className={`inline-block mt-2 px-2 py-1 rounded text-xs font-medium ${getStatusColor(transmutation.status)}`}>
                  {transmutation.status}
                </span>
              </div>
              {isSupervisor && transmutation.status === 'pending' && (
                <div className="ml-4 space-x-2">
                  <button
                    onClick={() => updateStatus(transmutation.id, 'approved')}
                    className="bg-green-600 text-white px-3 py-1 rounded text-sm hover:bg-green-700"
                  >
                    Aprobar
                  </button>
                  <button
                    onClick={() => updateStatus(transmutation.id, 'rejected', 'Transmutación rechazada')}
                    className="bg-red-600 text-white px-3 py-1 rounded text-sm hover:bg-red-700"
                  >
                    Rechazar
                  </button>
                </div>
              )}
            </div>
          </li>
        ))}
        {transmutations.length === 0 && (
          <li className="p-4 text-center text-gray-500">No hay transmutaciones</li>
        )}
      </ul>
    </div>
  )
}

