import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { apiService } from '../services/api.js'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token'))
  const role = ref(localStorage.getItem('role'))
  const isLoading = ref(false)
  const error = ref(null)

  const isAuthenticated = computed(() => !!token.value)

  const setAuth = (newToken, newRole) => {
    token.value = newToken
    role.value = newRole
    localStorage.setItem('token', newToken)
    localStorage.setItem('role', newRole)
  }

  const clearAuth = () => {
    token.value = null
    role.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('role')
  }

  const login = async (email, password) => {
    isLoading.value = true
    error.value = null

    try {
      const data = await apiService.login({ email, password })

      setAuth(data.token, data.role)
      return true
    } catch (err) {
      error.value = err.message || 'Login failed'
      return false
    } finally {
      isLoading.value = false
    }
  }

  return {
    token,
    role,
    isAuthenticated,
    isLoading,
    error,
    login,
    clearAuth,
  }
})
