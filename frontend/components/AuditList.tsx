'use client'

import { Audit } from '@/lib/api'
import api from '@/lib/api'

interface AuditListProps {
  audits: Audit[]
  onUpdate: () => void
}

export default function AuditList({ audits, onUpdate }: AuditListProps) {
  const resolveAudit = async (id: number) => {
    try {
      await api.put(`/api/audits/${id}/resolve`)
      onUpdate()
    } catch (error) {
      console.error('Error resolving audit:', error)
    }
  }

  const getSeverityColor = (severity: string) => {
    switch (severity) {
      case 'low':
        return 'bg-blue-100 text-blue-800'
      case 'medium':
        return 'bg-yellow-100 text-yellow-800'
      case 'high':
        return 'bg-orange-100 text-orange-800'
      case 'critical':
        return 'bg-red-100 text-red-800'
      default:
        return 'bg-gray-100 text-gray-800'
    }
  }

  return (
    <div className="bg-white rounded-lg shadow overflow-hidden">
      <table className="min-w-full divide-y divide-gray-200">
        <thead className="bg-gray-50">
          <tr>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Tipo</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Severidad</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Descripción</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Estado</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Acciones</th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {audits.map((audit) => (
            <tr key={audit.id} className="hover:bg-gray-50">
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{audit.type}</td>
              <td className="px-6 py-4 whitespace-nowrap">
                <span className={`px-2 py-1 rounded text-xs font-medium ${getSeverityColor(audit.severity)}`}>
                  {audit.severity}
                </span>
              </td>
              <td className="px-6 py-4 text-sm text-gray-500">{audit.description}</td>
              <td className="px-6 py-4 whitespace-nowrap">
                {audit.resolved ? (
                  <span className="text-green-600 text-sm">Resuelta</span>
                ) : (
                  <span className="text-yellow-600 text-sm">Pendiente</span>
                )}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm">
                {!audit.resolved && (
                  <button
                    onClick={() => resolveAudit(audit.id)}
                    className="bg-green-600 text-white px-3 py-1 rounded hover:bg-green-700"
                  >
                    Resolver
                  </button>
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      {audits.length === 0 && (
        <div className="p-4 text-center text-gray-500">No hay auditorías</div>
      )}
    </div>
  )
}

