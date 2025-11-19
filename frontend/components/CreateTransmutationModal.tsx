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
    <div className="fixed inset-0 bg-black/70 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <div className="bg-slate-900 border border-slate-700 rounded-2xl shadow-2xl w-full max-w-3xl max-h-[90vh] flex flex-col">
        {/* Header */}
        <div className="border-b border-slate-700 px-6 py-4 flex-shrink-0">
          <div className="flex items-center justify-between">
            <h2 className="text-2xl font-bold bg-gradient-to-r from-purple-400 to-pink-400 bg-clip-text text-transparent">
              Nueva Transmutación
            </h2>
            <button
              onClick={onClose}
              className="text-slate-400 hover:text-slate-200 transition-colors text-2xl leading-none"
              aria-label="Cerrar"
            >
              ×
            </button>
          </div>
        </div>

        {/* Scrollable Content */}
        <form onSubmit={handleSubmit} className="flex flex-col flex-1 overflow-hidden">
          <div className="overflow-y-auto flex-1 px-6 py-6 space-y-6">
            {/* Description */}
            <div className="space-y-2">
              <label className="block text-sm font-semibold text-slate-200">
                Descripción de la Transmutación
              </label>
              <textarea
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                required
                rows={4}
                placeholder="Describe el proceso de transmutación que deseas realizar..."
                className="w-full px-4 py-3 bg-slate-800 border border-slate-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500/50 focus:border-purple-500 text-slate-100 placeholder-slate-500 text-base resize-none transition-all"
              />
            </div>

            {/* Materials Section */}
            <div className="space-y-3">
              <div className="flex items-center justify-between">
                <label className="block text-sm font-semibold text-slate-200">
                  Materiales de Entrada
                </label>
                {materials.length > 0 && (
                  <span className="text-xs text-slate-400">
                    {materials.length} material{materials.length !== 1 ? 'es' : ''} disponible{materials.length !== 1 ? 's' : ''}
                  </span>
                )}
              </div>

              {materials.length === 0 ? (
                <div className="bg-yellow-900/20 border border-yellow-800/50 rounded-lg p-4">
                  <p className="text-yellow-300 text-sm">
                    No hay materiales disponibles. Contacta al administrador.
                  </p>
                </div>
              ) : (
                <div className="space-y-3">
                  {inputMaterials.map((input, index) => (
                    <div key={index} className="bg-slate-800/50 border border-slate-700 rounded-lg p-4 space-y-3">
                      <div className="flex items-center gap-3">
                        <div className="flex-1">
                          <label className="block text-xs text-slate-400 mb-1.5">
                            Material {index + 1}
                          </label>
                          <select
                            value={input.material_id}
                            onChange={(e) => updateInputMaterial(index, 'material_id', Number(e.target.value))}
                            required
                            className="w-full px-4 py-2.5 bg-slate-700 border border-slate-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500/50 focus:border-purple-500 text-slate-100 text-base transition-all"
                          >
                            <option value={0} className="bg-slate-800">Seleccionar material</option>
                            {materials.map((m) => (
                              <option key={m.id} value={m.id} className="bg-slate-800">
                                {m.name} (Stock: {m.stock} {m.unit})
                              </option>
                            ))}
                          </select>
                        </div>
                        <div className="w-32">
                          <label className="block text-xs text-slate-400 mb-1.5">
                            Cantidad
                          </label>
                          <input
                            type="number"
                            value={input.quantity}
                            onChange={(e) => updateInputMaterial(index, 'quantity', Number(e.target.value))}
                            required
                            min="0.01"
                            step="0.01"
                            placeholder="0.00"
                            className="w-full px-4 py-2.5 bg-slate-700 border border-slate-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500/50 focus:border-purple-500 text-slate-100 placeholder-slate-500 text-base transition-all"
                          />
                        </div>
                        {inputMaterials.length > 1 && (
                          <div className="flex items-end">
                            <button
                              type="button"
                              onClick={() => removeInputMaterial(index)}
                              className="px-4 py-2.5 bg-red-600/80 hover:bg-red-600 text-white rounded-lg transition-all text-sm font-medium"
                            >
                              Eliminar
                            </button>
                          </div>
                        )}
                      </div>
                    </div>
                  ))}
                  <button
                    type="button"
                    onClick={addInputMaterial}
                    className="w-full py-2.5 border-2 border-dashed border-slate-600 hover:border-purple-500 text-slate-400 hover:text-purple-400 rounded-lg transition-all text-sm font-medium"
                  >
                    + Agregar otro material
                  </button>
                </div>
              )}
            </div>

            {error && (
              <div className="bg-red-900/30 border border-red-800/50 rounded-lg p-3">
                <p className="text-red-300 text-sm font-medium">{error}</p>
              </div>
            )}
          </div>

          {/* Footer Actions */}
          <div className="border-t border-slate-700 px-6 py-4 flex-shrink-0">
            <div className="flex gap-3">
              <button
                type="button"
                onClick={onClose}
                className="flex-1 bg-slate-700 hover:bg-slate-600 text-slate-200 py-3 px-4 rounded-lg font-medium transition-all"
              >
                Cancelar
              </button>
              <button
                type="submit"
                disabled={loading || materials.length === 0}
                className="flex-1 bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-500 hover:to-pink-500 text-white py-3 px-4 rounded-lg font-semibold transition-all disabled:opacity-50 disabled:cursor-not-allowed shadow-lg shadow-purple-500/20"
              >
                {loading ? 'Creando...' : 'Crear Transmutación'}
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  )
}

