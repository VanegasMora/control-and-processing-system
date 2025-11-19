'use client'

import { useEffect, useState } from 'react'
import api, { Mission, Transmutation, Material } from '@/lib/api'
import MissionList from './MissionList'
import TransmutationList from './TransmutationList'
import CreateMissionModal from './CreateMissionModal'
import CreateTransmutationModal from './CreateTransmutationModal'

export default function AlchemistDashboard() {
  const [missions, setMissions] = useState<Mission[]>([])
  const [transmutations, setTransmutations] = useState<Transmutation[]>([])
  const [materials, setMaterials] = useState<Material[]>([])
  const [showMissionModal, setShowMissionModal] = useState(false)
  const [showTransmutationModal, setShowTransmutationModal] = useState(false)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    loadData()
    setupWebSocket()
  }, [])

  const loadData = async () => {
    try {
      const [missionsRes, transmutationsRes, materialsRes] = await Promise.all([
        api.get('/api/missions'),
        api.get('/api/transmutations'),
        api.get('/api/materials'),
      ])
      setMissions(missionsRes.data)
      setTransmutations(transmutationsRes.data)
      setMaterials(materialsRes.data)
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
      if (message.type === 'mission_status_changed' || message.type === 'transmutation_status_changed') {
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
      <div className="flex justify-between items-center">
        <h2 className="text-2xl font-bold text-gray-800">Panel de Alquimista</h2>
        <div className="space-x-4">
          <button
            onClick={() => setShowMissionModal(true)}
            className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700"
          >
            Nueva Misión
          </button>
          <button
            onClick={() => setShowTransmutationModal(true)}
            className="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700"
          >
            Nueva Transmutación
          </button>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div>
          <h3 className="text-xl font-semibold mb-4">Misiones</h3>
          <MissionList missions={missions} onUpdate={loadData} />
        </div>
        <div>
          <h3 className="text-xl font-semibold mb-4">Transmutaciones</h3>
          <TransmutationList transmutations={transmutations} onUpdate={loadData} />
        </div>
      </div>

      {showMissionModal && (
        <CreateMissionModal
          onClose={() => setShowMissionModal(false)}
          onSuccess={() => {
            setShowMissionModal(false)
            loadData()
          }}
        />
      )}

      {showTransmutationModal && (
        <CreateTransmutationModal
          materials={materials}
          onClose={() => setShowTransmutationModal(false)}
          onSuccess={() => {
            setShowTransmutationModal(false)
            loadData()
          }}
        />
      )}
    </div>
  )
}

