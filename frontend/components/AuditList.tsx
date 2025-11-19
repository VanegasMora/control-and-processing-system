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
        return 'bg-blue-900/40 text-blue-300 border border-blue-700/50'
      case 'medium':
        return 'bg-yellow-900/40 text-yellow-300 border border-yellow-700/50'
      case 'high':
        return 'bg-orange-900/40 text-orange-300 border border-orange-700/50'
      case 'critical':
        return 'bg-red-900/40 text-red-300 border border-red-700/50'
      default:
        return 'bg-neutral-800/40 text-neutral-300 border border-neutral-700/50'
    }
  }

  return (
    <div className="space-y-4">
      {audits.length === 0 ? (
        <div className="text-center py-12 text-slate-500">
          <p>No hay auditorías disponibles</p>
        </div>
      ) : (
        audits.map((audit) => (
          <div
            key={audit.id}
            className="bg-slate-700/30 border border-slate-600/50 rounded-xl p-5 hover:bg-slate-700/50 transition-all hover:border-slate-500/50"
          >
            <div className="flex items-start justify-between gap-4">
              <div className="flex-1 min-w-0">
                <div className="flex items-start gap-3 mb-3">
                  <div className="w-1.5 h-12 bg-gradient-to-b from-yellow-400 to-orange-500 rounded-full flex-shrink-0"></div>
                  <div className="flex-1">
                    <div className="flex items-center gap-3 mb-2">
                      <h4 className="text-lg font-semibold text-slate-200">{audit.type}</h4>
                      <span className={`px-3 py-1 rounded-lg text-xs font-semibold backdrop-blur-sm ${getSeverityColor(audit.severity)}`}>
                        {audit.severity}
                      </span>
                    </div>
                    <p className="text-sm text-slate-400 leading-relaxed mb-3">{audit.description}</p>
                    <div className="flex items-center gap-3">
                      {audit.resolved ? (
                        <span className="text-emerald-400 text-sm font-medium bg-emerald-900/20 border border-emerald-700/30 rounded-lg px-3 py-1">
                          ✓ Resuelta
                        </span>
                      ) : (
                        <span className="text-yellow-400 text-sm font-medium bg-yellow-900/20 border border-yellow-700/30 rounded-lg px-3 py-1">
                          ⏳ Pendiente
                        </span>
                      )}
                    </div>
                  </div>
                </div>
              </div>
              {!audit.resolved && (
                <div className="flex-shrink-0">
                  <button
                    onClick={() => resolveAudit(audit.id)}
                    className="bg-emerald-600/80 hover:bg-emerald-600 text-white px-5 py-2.5 rounded-lg transition-all backdrop-blur-sm border border-emerald-500/30 shadow-lg font-medium whitespace-nowrap"
                  >
                    Resolver
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

