'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import DashboardLayout from '@/components/DashboardLayout'
import AlchemistDashboard from '@/components/AlchemistDashboard'
import SupervisorDashboard from '@/components/SupervisorDashboard'

export default function Dashboard() {
  const [user, setUser] = useState<any>(null)
  const [loading, setLoading] = useState(true)
  const router = useRouter()

  useEffect(() => {
    const token = localStorage.getItem('token')
    const userData = localStorage.getItem('user')
    
    if (!token) {
      router.push('/')
      return
    }

    if (userData) {
      setUser(JSON.parse(userData))
    }
    setLoading(false)
  }, [router])

  if (loading) {
    return <div className="min-h-screen flex items-center justify-center">Cargando...</div>
  }

  return (
    <DashboardLayout user={user}>
      {user?.role === 'supervisor' ? (
        <SupervisorDashboard />
      ) : (
        <AlchemistDashboard />
      )}
    </DashboardLayout>
  )
}

