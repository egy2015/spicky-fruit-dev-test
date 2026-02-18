import { defineStore } from 'pinia'
import { apiService } from '../services/api'

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
  },
})
