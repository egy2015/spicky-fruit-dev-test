import { defineStore } from 'pinia'
import { apiService } from '../services/api'
import { http } from '../services/http'

export const usePaymentStore = defineStore('payment', {
  state: () => ({
    payments: [],
    loading: false,
    error: null,
  }),

  actions: {
    async getPayments() {
      this.loading = true
      this.error = null

      try {
        const data = await apiService.getPayments()
        this.payments = data
      } catch (err) {
        this.error = err
      } finally {
        this.loading = false
      }
    },
     async updatePayment(id, payload) {
      const updated = await http.request(`/dashboard/v1/payments/${id}`, {
        method: 'PUT',
        body: JSON.stringify(payload),
      })

      const idx = this.payments.findIndex(p => p.id === id)
      if (idx !== -1) {
        this.payments[idx] = {
          ...this.payments[idx],
          ...updated,
        }
      }

      return updated
    },
  },
})
