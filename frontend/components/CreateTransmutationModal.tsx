'use client'

import { useState } from 'react'
import api, { Material } from '@/lib/api'

interface CreateTransmutationModalProps {
  materials: Material[]
  onClose: () => void
  onSuccess: () => void
}

export default function CreateTransmutationModal({ materials, onClose, onSuccess }: CreateTransmutationModalProps) {
  const [description, setDescription] = useState('')
  const [inputMaterials, setInputMaterials] = useState<Array<{ material_id: number; quantity: number }>>([{ material_id: 0, quantity: 0 }])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const addInputMaterial = () => {
    setInputMaterials([...inputMaterials, { material_id: 0, quantity: 0 }])
  }

  const updateInputMaterial = (index: number, field: 'material_id' | 'quantity', value: number) => {
    const updated = [...inputMaterials]
    updated[index] = { ...updated[index], [field]: value }
    setInputMaterials(updated)
  }

  const removeInputMaterial = (index: number) => {
    setInputMaterials(inputMaterials.filter((_, i) => i !== index))
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    setLoading(true)

    try {
      await api.post('/api/transmutations', {
        description,
        input_materials: inputMaterials.map(im => ({
          material_id: im.material_id,
          quantity: im.quantity,
        })),
      })
      onSuccess()
    } catch (err: any) {
      setError(err.response?.data?.error || 'Error al crear la transmutación')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto">
        <h2 className="text-xl font-bold mb-4">Nueva Transmutación</h2>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Descripción
            </label>
            <textarea
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              required
              rows={4}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Materiales de Entrada
            </label>
            {inputMaterials.map((input, index) => (
              <div key={index} className="flex space-x-2 mb-2">
                <select
                  value={input.material_id}
                  onChange={(e) => updateInputMaterial(index, 'material_id', Number(e.target.value))}
                  required
                  className="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value={0}>Seleccionar material</option>
                  {materials.map((m) => (
                    <option key={m.id} value={m.id}>
                      {m.name} (Stock: {m.stock} {m.unit})
                    </option>
                  ))}
                </select>
                <input
                  type="number"
                  value={input.quantity}
                  onChange={(e) => updateInputMaterial(index, 'quantity', Number(e.target.value))}
                  required
                  min="0.01"
                  step="0.01"
                  placeholder="Cantidad"
                  className="w-32 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
                {inputMaterials.length > 1 && (
                  <button
                    type="button"
                    onClick={() => removeInputMaterial(index)}
                    className="bg-red-600 text-white px-3 py-2 rounded-md hover:bg-red-700"
                  >
                    Eliminar
                  </button>
                )}
              </div>
            ))}
            <button
              type="button"
              onClick={addInputMaterial}
              className="mt-2 text-blue-600 text-sm hover:underline"
            >
              + Agregar material
            </button>
          </div>
          {error && (
            <div className="text-red-600 text-sm">{error}</div>
          )}
          <div className="flex space-x-4">
            <button
              type="submit"
              disabled={loading}
              className="flex-1 bg-green-600 text-white py-2 px-4 rounded-md hover:bg-green-700 disabled:opacity-50"
            >
              {loading ? 'Creando...' : 'Crear'}
            </button>
            <button
              type="button"
              onClick={onClose}
              className="flex-1 bg-gray-300 text-gray-700 py-2 px-4 rounded-md hover:bg-gray-400"
            >
              Cancelar
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

