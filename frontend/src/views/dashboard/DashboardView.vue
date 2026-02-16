<template>
  <DashboardLayout>
    <div class="dashboard-content">
      <h1>Payments Dashboard</h1>

      <!-- Summary Widget -->
      <div class="summary-widget">
        <div class="summary-card">
          <h3>Total Payments</h3>
          <p class="summary-value">{{ totalPayments }}</p>
        </div>
        <div class="summary-card success">
          <h3>Completed</h3>
          <p class="summary-value">{{ completedCount }}</p>
        </div>
        <div class="summary-card processing">
          <h3>Processing</h3>
          <p class="summary-value">{{ processingCount }}</p>
        </div>
        <div class="summary-card failed">
          <h3>Failed</h3>
          <p class="summary-value">{{ failedCount }}</p>
        </div>
      </div>

      <!-- Filters -->
      <div class="filters">
        <button
          v-for="status in ['all', 'completed', 'processing', 'failed']"
          :key="status"
          @click="filterByStatus(status)"
          :class="{ active: activeFilter === status }"
          class="filter-btn"
        >
          {{ status.charAt(0).toUpperCase() + status.slice(1) }}
        </button>
      </div>

      <!-- Payments Table -->
      <div v-if="isLoading" class="loading">
        <Loading />
      </div>

      <div v-else-if="filteredPayments.length > 0" class="table-container">
        <PaymentsTable :payments="filteredPayments" />
      </div>

      <div v-else class="no-data">
        <p>No payments found</p>
      </div>

      <div v-if="error" class="error-message">
        {{ error }}
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useAuthStore } from '../../stores/auth.js'
import { useRouter } from 'vue-router'
import { apiService } from '../../services/api.js'
import DashboardLayout from '../../components/layout/DashboardLayout.vue'
import PaymentsTable from '../../components/payments/PaymentsTable.vue'
import Loading from '../../components/common/Loading.vue'

const router = useRouter()
const authStore = useAuthStore()

const payments = ref([])
const isLoading = ref(false)
const error = ref(null)
const activeFilter = ref('all')

const filteredPayments = computed(() => {
  if (activeFilter.value === 'all') {
    return payments.value
  }
  return payments.value.filter((p) => p.status === activeFilter.value)
})

const totalPayments = computed(() => payments.value.length)
const completedCount = computed(() =>
  payments.value.filter((p) => p.status === 'completed').length
)
const processingCount = computed(() =>
  payments.value.filter((p) => p.status === 'processing').length
)
const failedCount = computed(() =>
  payments.value.filter((p) => p.status === 'failed').length
)

const fetchPayments = async () => {
  isLoading.value = true
  error.value = null

  try {
    const data = await apiService.getPayments()
    payments.value = data.payments || data || []
  } catch (err) {
    error.value = 'Failed to load payments. Please try again.'
    console.error(err)
  } finally {
    isLoading.value = false
  }
}

const filterByStatus = (status) => {
  activeFilter.value = status
}

const logout = () => {
  authStore.clearAuth()
  router.push('/login')
}

onMounted(() => {
  fetchPayments()
})
</script>

<style scoped>
.dashboard-content {
  padding: 2rem;
}

.dashboard-content h1 {
  margin-bottom: 2rem;
  font-size: 2rem;
  color: #333;
}

.summary-widget {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.summary-card {
  background: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  border-left: 4px solid #0066cc;
}

.summary-card h3 {
  margin: 0;
  font-size: 0.9rem;
  color: #666;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin-bottom: 0.5rem;
}

.summary-value {
  margin: 0;
  font-size: 2rem;
  font-weight: 600;
  color: #333;
}

.summary-card.success {
  border-left-color: #28a745;
}

.summary-card.processing {
  border-left-color: #ffc107;
}

.summary-card.failed {
  border-left-color: #dc3545;
}

.filters {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 2rem;
  flex-wrap: wrap;
}

.filter-btn {
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: white;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.3s;
}

.filter-btn:hover {
  border-color: #0066cc;
  color: #0066cc;
}

.filter-btn.active {
  background-color: #0066cc;
  color: white;
  border-color: #0066cc;
}

.table-container {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.loading,
.no-data {
  text-align: center;
  padding: 3rem 2rem;
  background: white;
  border-radius: 8px;
}

.no-data p {
  color: #666;
  font-size: 1.1rem;
}

.error-message {
  margin-top: 1rem;
  padding: 1rem;
  background-color: #fee;
  color: #c33;
  border-radius: 4px;
  border: 1px solid #fcc;
}
</style>
