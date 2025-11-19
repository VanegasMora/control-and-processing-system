'use client'

import { useEffect } from 'react'
import { useRouter } from 'next/navigation'
import LoginForm from '@/components/LoginForm'

export default function Home() {
  const router = useRouter()

  useEffect(() => {
    const token = localStorage.getItem('token')
    if (token) {
      router.push('/dashboard')
    }
  }, [router])

  return (
    <div className="min-h-screen flex items-center justify-center relative z-10">
      <div className="bg-slate-900/95 border border-slate-700/50 p-8 rounded-2xl shadow-2xl w-full max-w-md backdrop-blur-sm">
        <div className="text-center mb-8">
          <h1 className="text-4xl font-bold mb-3 bg-gradient-to-r from-cyan-400 via-purple-500 to-pink-500 bg-clip-text text-transparent">
            Departamento de Alquimia Estatal
          </h1>
          <div className="h-1 w-24 bg-gradient-to-r from-cyan-500 to-purple-500 mx-auto rounded-full"></div>
          <p className="text-slate-400 mt-6 text-sm">
            Sistema de gestión de alquimistas de Amestris
          </p>
        </div>
        <LoginForm />
      </div>
    </div>
  )
}

