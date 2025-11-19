'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import api from '@/lib/api'

export default function LoginForm() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [isRegister, setIsRegister] = useState(false)
  const [name, setName] = useState('')
  const [rank, setRank] = useState('apprentice')
  const [specialty, setSpecialty] = useState('combat')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)
  const router = useRouter()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    setLoading(true)

    try {
      if (isRegister) {
        const response = await api.post('/api/auth/register', {
          name,
          email,
          password,
          rank,
          specialty,
        })
        localStorage.setItem('token', response.data.token)
        router.push('/dashboard')
      } else {
        const response = await api.post('/api/auth/login', {
          email,
          password,
        })
        localStorage.setItem('token', response.data.token)
        localStorage.setItem('user', JSON.stringify(response.data))
        router.push('/dashboard')
      }
    } catch (err: any) {
      console.error('Error de autenticación:', err)
      if (err.response?.data?.error) {
        setError(err.response.data.error)
      } else if (err.response?.data) {
        setError(JSON.stringify(err.response.data))
      } else if (err.message) {
        setError(err.message)
      } else {
        setError('Error al autenticar. Verifica tus credenciales.')
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-5">
      {isRegister && (
        <>
          <div>
            <label htmlFor="name" className="block text-sm font-medium text-slate-300 mb-2">
              Nombre
            </label>
            <input
              id="name"
              name="name"
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
              autoComplete="name"
              className="w-full px-4 py-3 bg-slate-800/50 border border-slate-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-cyan-500/50 focus:border-cyan-500 text-slate-200 placeholder-slate-500 transition-all"
            />
          </div>
          <div>
            <label htmlFor="rank" className="block text-sm font-medium text-slate-300 mb-2">
              Rango
            </label>
            <select
              id="rank"
              name="rank"
              value={rank}
              onChange={(e) => setRank(e.target.value)}
              className="w-full px-4 py-3 bg-slate-800/50 border border-slate-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-cyan-500/50 focus:border-cyan-500 text-slate-200 transition-all"
            >
              <option value="apprentice" className="bg-slate-800">Aprendiz</option>
              <option value="state" className="bg-slate-800">Estatal</option>
              <option value="national" className="bg-slate-800">Nacional</option>
            </select>
          </div>
          <div>
            <label htmlFor="specialty" className="block text-sm font-medium text-slate-300 mb-2">
              Especialidad
            </label>
            <select
              id="specialty"
              name="specialty"
              value={specialty}
              onChange={(e) => setSpecialty(e.target.value)}
              className="w-full px-4 py-3 bg-slate-800/50 border border-slate-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-cyan-500/50 focus:border-cyan-500 text-slate-200 transition-all"
            >
              <option value="combat" className="bg-slate-800">Combate</option>
              <option value="research" className="bg-slate-800">Investigación</option>
              <option value="medical" className="bg-slate-800">Médica</option>
              <option value="industrial" className="bg-slate-800">Industrial</option>
            </select>
          </div>
        </>
      )}
      <div>
        <label htmlFor="email" className="block text-sm font-medium text-slate-300 mb-2">
          Email
        </label>
        <input
          id="email"
          name="email"
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
          autoComplete="email"
          className="w-full px-4 py-3 bg-slate-800/50 border border-slate-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-cyan-500/50 focus:border-cyan-500 text-slate-200 placeholder-slate-500 transition-all"
        />
      </div>
      <div>
        <label htmlFor="password" className="block text-sm font-medium text-slate-300 mb-2">
          Contraseña
        </label>
        <input
          id="password"
          name="password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          autoComplete={isRegister ? "new-password" : "current-password"}
          className="w-full px-4 py-3 bg-slate-800/50 border border-slate-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-cyan-500/50 focus:border-cyan-500 text-slate-200 placeholder-slate-500 transition-all"
        />
      </div>
      {error && (
        <div className="text-red-400 text-sm text-center bg-red-900/20 border border-red-800/50 rounded-lg p-3">
          {error}
        </div>
      )}
      <button
        type="submit"
        disabled={loading}
        className="w-full bg-gradient-to-r from-cyan-600 to-purple-600 hover:from-cyan-500 hover:to-purple-500 py-3 rounded-lg text-white font-semibold transition-all disabled:opacity-50 disabled:cursor-not-allowed shadow-lg shadow-cyan-500/20 focus:outline-none focus:ring-2 focus:ring-cyan-500"
      >
        {loading ? 'Cargando...' : isRegister ? 'Registrarse' : 'Iniciar Sesión'}
      </button>
      <button
        type="button"
        onClick={() => setIsRegister(!isRegister)}
        className="w-full text-cyan-400 text-sm hover:text-cyan-300 text-center transition-colors"
      >
        {isRegister
          ? '¿Ya tienes cuenta? Inicia sesión'
          : '¿No tienes cuenta? Regístrate'}
      </button>
    </form>
  )
}

