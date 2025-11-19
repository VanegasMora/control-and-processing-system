'use client'

import { useEffect, useState } from 'react'
import api, { Mission, Transmutation, Audit } from '@/lib/api'
import MissionList from './MissionList'
import TransmutationList from './TransmutationList'
import AuditList from './AuditList'
import DashboardStats from './DashboardStats'

export default function SupervisorDashboard() {
  const [missions, setMissions] = useState<Mission[]>([])
  const [transmutations, setTransmutations] = useState<Transmutation[]>([])
  const [audits, setAudits] = useState<Audit[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    loadData()
    setupWebSocket()
  }, [])

  const loadData = async () => {
    try {
      const [missionsRes, transmutationsRes, auditsRes] = await Promise.all([
        api.get('/api/missions'),
        api.get('/api/transmutations'),
        api.get('/api/audits'),
      ])
      setMissions(missionsRes.data)
      setTransmutations(transmutationsRes.data)
      setAudits(auditsRes.data)
    } catch (error) {
      console.error('Error loading data:', error)
    } finally {
      setLoading(false)
    }
  }

  const setupWebSocket = () => {
    const token = localStorage.getItem('token')
    if (!token) return

    const ws = new WebSocket(`ws://localhost:8000/api/ws?token=${token}`)
    
    ws.onmessage = (event) => {
      const message = JSON.parse(event.data)
      if (message.type === 'mission_status_changed' || 
          message.type === 'transmutation_status_changed' ||
          message.type === 'audit_created') {
        loadData()
      }
    }

    ws.onerror = (error) => {
      console.error('WebSocket error:', error)
    }
  }

  if (loading) {
    return (
      <div className="text-center py-12">
        <div className="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-cyan-500"></div>
        <p className="mt-4 text-slate-400">Cargando...</p>
      </div>
    )
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="mb-8">
        <h2 className="text-4xl font-bold bg-gradient-to-r from-cyan-400 via-purple-400 to-pink-400 bg-clip-text text-transparent mb-2">
          Panel de Supervisor
        </h2>
        <p className="text-slate-400 text-base">
          Gestión y supervisión de operaciones alquímicas
        </p>
      </div>
      
      {/* Estadísticas */}
      <DashboardStats 
        missions={missions}
        transmutations={transmutations}
        audits={audits}
      />

      {/* Layout vertical con secciones apiladas */}
      <div className="space-y-6">
        {/* Misiones Pendientes */}
        <section className="bg-slate-800/50 border border-slate-700/50 rounded-2xl shadow-2xl overflow-hidden">
          <div className="bg-gradient-to-r from-cyan-600/20 to-blue-600/20 border-b border-slate-700/50 px-6 py-4">
            <div className="flex items-center justify-between">
              <h3 className="text-2xl font-bold text-slate-200 flex items-center gap-3">
                <div className="w-2 h-8 bg-gradient-to-b from-cyan-400 to-blue-500 rounded-full"></div>
                Misiones Pendientes
              </h3>
              <span className="text-slate-400 text-sm bg-slate-700/50 px-3 py-1 rounded-full">
                {missions.filter(m => m.status === 'pending').length} pendientes
              </span>
            </div>
          </div>
          <div className="p-6">
            <MissionList missions={missions.filter(m => m.status === 'pending')} onUpdate={loadData} isSupervisor />
          </div>
        </section>

        {/* Transmutaciones Pendientes */}
        <section className="bg-slate-800/50 border border-slate-700/50 rounded-2xl shadow-2xl overflow-hidden">
          <div className="bg-gradient-to-r from-purple-600/20 to-pink-600/20 border-b border-slate-700/50 px-6 py-4">
            <div className="flex items-center justify-between">
              <h3 className="text-2xl font-bold text-slate-200 flex items-center gap-3">
                <div className="w-2 h-8 bg-gradient-to-b from-purple-400 to-pink-500 rounded-full"></div>
                Transmutaciones Pendientes
              </h3>
              <span className="text-slate-400 text-sm bg-slate-700/50 px-3 py-1 rounded-full">
                {transmutations.filter(t => t.status === 'pending').length} pendientes
              </span>
            </div>
          </div>
          <div className="p-6">
            <TransmutationList transmutations={transmutations.filter(t => t.status === 'pending')} onUpdate={loadData} isSupervisor />
          </div>
        </section>

        {/* Auditorías */}
        <section className="bg-slate-800/50 border border-slate-700/50 rounded-2xl shadow-2xl overflow-hidden">
          <div className="bg-gradient-to-r from-yellow-600/20 to-orange-600/20 border-b border-slate-700/50 px-6 py-4">
            <div className="flex items-center justify-between">
              <h3 className="text-2xl font-bold text-slate-200 flex items-center gap-3">
                <div className="w-2 h-8 bg-gradient-to-b from-yellow-400 to-orange-500 rounded-full"></div>
                Auditorías
              </h3>
              <span className="text-slate-400 text-sm bg-slate-700/50 px-3 py-1 rounded-full">
                {audits.length} total
              </span>
            </div>
          </div>
          <div className="p-6">
            <AuditList audits={audits} onUpdate={loadData} />
          </div>
        </section>
      </div>
    </div>
  )
}

