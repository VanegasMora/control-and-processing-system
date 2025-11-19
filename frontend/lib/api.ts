import axios from 'axios'

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8000'

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Interceptor para agregar token a las peticiones
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Interceptor para manejar errores de autenticación
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      window.location.href = '/'
    }
    return Promise.reject(error)
  }
)

export default api

export interface Alchemist {
  id: number
  name: string
  email: string
  rank: string
  specialty: string
  role: string
  certified: boolean
  created_at: string
}

export interface Mission {
  id: number
  title: string
  description: string
  status: string
  alchemist_id: number
  requested_at: string
  approved_at?: string
  completed_at?: string
  supervisor_id?: number
}

export interface Material {
  id: number
  name: string
  type: string
  description: string
  stock: number
  unit: string
  price: number
}

export interface Transmutation {
  id: number
  alchemist_id: number
  status: string
  input_materials: TransmutationMaterial[]
  output_materials: TransmutationMaterial[]
  description: string
  cost: number
  result?: string
  supervisor_id?: number
  approved_at?: string
  completed_at?: string
  created_at: string
}

export interface TransmutationMaterial {
  id: number
  material: Material
  quantity: number
  is_input: boolean
}

export interface Audit {
  id: number
  type: string
  severity: string
  description: string
  alchemist_id?: number
  details?: string
  resolved: boolean
  resolved_at?: string
  resolved_by?: number
  created_at: string
}

