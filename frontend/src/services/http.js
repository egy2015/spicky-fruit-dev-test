const API_BASE_URL =
  import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/dashboard/v1'

export const http = {
  async request(path, options = {}) {
    const response = await fetch(`${API_BASE_URL}${path}`, {
      headers: {
        'Content-Type': 'application/json',
        ...(options.headers || {}),
      },
      ...options,
    })

    if (!response.ok) {
      const errorText = await response.text()
      throw new Error(errorText || `HTTP Error ${response.status}`)
    }

    return response.json()
  },
}
