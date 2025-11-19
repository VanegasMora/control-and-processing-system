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
    return <div className="text-center py-8">Cargando...</div>
  }

  return (
    <div className="space-y-6">
      <h2 className="text-2xl font-bold text-gray-800">Panel de Supervisor</h2>
      
      <DashboardStats 
        missions={missions}
        transmutations={transmutations}
        audits={audits}
      />

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div>
          <h3 className="text-xl font-semibold mb-4">Misiones Pendientes</h3>
          <MissionList missions={missions.filter(m => m.status === 'pending')} onUpdate={loadData} isSupervisor />
        </div>
        <div>
          <h3 className="text-xl font-semibold mb-4">Transmutaciones Pendientes</h3>
          <TransmutationList transmutations={transmutations.filter(t => t.status === 'pending')} onUpdate={loadData} isSupervisor />
        </div>
      </div>

      <div>
        <h3 className="text-xl font-semibold mb-4">Auditorías</h3>
        <AuditList audits={audits} onUpdate={loadData} />
      </div>
    </div>
  )
}

