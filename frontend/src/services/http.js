const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/dashboard/v1'

export const http = {
  async request(path, options = {}) {
    const token = localStorage.getItem('token')
    const headers = {
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(options.headers || {}),
    }

    const response = await fetch(path, {
      ...options,
      headers,
    })

    if (!response.ok) {
      const text = await response.text()
      throw new Error(text || `HTTP error ${response.status}`)
    }

    return response.json()
  },
}
