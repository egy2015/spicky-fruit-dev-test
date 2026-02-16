import { http } from './http'

export const apiService = {
  getPayments(status = null) {
    const query = status ? `?status=${status}` : ''
    return http.request(`/payments${query}`)
  },

  login(payload) {
    return http.request('/auth/login', {
      method: 'POST',
      body: JSON.stringify(payload),
    })
  },
}
