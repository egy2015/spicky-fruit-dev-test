import { http } from './http'

export const apiService = {
  getPayments(status = null) {
    const query = status ? `?status=${status}` : ''
    return http.request(`/dashboard/v1/payments${query}`)
  },

  login(payload) {
    return http.request('/dashboard/v1/auth/login', {
      method: 'POST',
      body: JSON.stringify(payload),
    })
  },
}
