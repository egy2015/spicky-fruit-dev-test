// stores/user.js
import { defineStore } from 'pinia'
import { apiService } from '../services/api'

export const useUserStore = defineStore('user', {
  state: () => ({
    users: [],
    loading: false,
  }),
  actions: {
    async getUsers() {
      this.loading = true
      try {
        this.users = await apiService.getUsers()
      } finally {
        this.loading = false
      }
    },
  },
})
