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
    <form onSubmit={handleSubmit} className="space-y-4">
      {isRegister && (
        <>
          <div>
            <label htmlFor="name" className="block text-sm font-medium text-gray-700 mb-1">
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
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div>
            <label htmlFor="rank" className="block text-sm font-medium text-gray-700 mb-1">
              Rango
            </label>
            <select
              id="rank"
              name="rank"
              value={rank}
              onChange={(e) => setRank(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="apprentice">Aprendiz</option>
              <option value="state">Estatal</option>
              <option value="national">Nacional</option>
            </select>
          </div>
          <div>
            <label htmlFor="specialty" className="block text-sm font-medium text-gray-700 mb-1">
              Especialidad
            </label>
            <select
              id="specialty"
              name="specialty"
              value={specialty}
              onChange={(e) => setSpecialty(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="combat">Combate</option>
              <option value="research">Investigación</option>
              <option value="medical">Médica</option>
              <option value="industrial">Industrial</option>
            </select>
          </div>
        </>
      )}
      <div>
        <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-1">
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
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>
      <div>
        <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">
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
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>
      {error && (
        <div className="text-red-600 text-sm text-center">{error}</div>
      )}
      <button
        type="submit"
        disabled={loading}
        className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50"
      >
        {loading ? 'Cargando...' : isRegister ? 'Registrarse' : 'Iniciar Sesión'}
      </button>
      <button
        type="button"
        onClick={() => setIsRegister(!isRegister)}
        className="w-full text-blue-600 text-sm hover:underline"
      >
        {isRegister
          ? '¿Ya tienes cuenta? Inicia sesión'
          : '¿No tienes cuenta? Regístrate'}
      </button>
    </form>
  )
}

