import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'Amestris Alchemy Department',
  description: 'Sistema de gestión de alquimistas estatales',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="es">
      <body>{children}</body>
    </html>
  )
}

