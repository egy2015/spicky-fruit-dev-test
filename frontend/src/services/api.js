import { useAuthStore } from '../stores/auth.js'

const API_BASE_URL = '/dashboard/v1'

const createHeaders = () => {
  const authStore = useAuthStore()
  const headers = {
    'Content-Type': 'application/json',
  }
  if (authStore.token) {
    headers['Authorization'] = `Bearer ${authStore.token}`
  }
  return headers
}

export const apiService = {
  async getPayments(status = null) {
    try {
      const url = status
        ? `${API_BASE_URL}/payments?status=${status}`
        : `${API_BASE_URL}/payments`

      const response = await fetch(url, {
        method: 'GET',
        headers: createHeaders(),
      })

      if (!response.ok) {
        throw new Error(`Failed to fetch payments: ${response.status}`)
      }

      return await response.json()
    } catch (error) {
      console.error('API Error:', error)
      throw error
    }
  },

  async getPaymentsByStatus(status) {
    return this.getPayments(status)
  },
}
