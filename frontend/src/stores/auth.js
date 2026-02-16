import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || null)
  const role = ref(localStorage.getItem('role') || null)
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
      const response = await fetch('/dashboard/v1/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })

      if (!response.ok) {
        throw new Error('Login failed')
      }

      const data = await response.json()
      setAuth(data.token, data.role)
      return true
    } catch (err) {
      error.value = err.message || 'An error occurred during login'
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
    setAuth,
    clearAuth,
    login,
  }
})
