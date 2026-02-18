<template>
  <DashboardLayout>
    <div class="payment-detail-view">
      <h1>Payment Detail</h1>
      <p>Detailed information about a specific payment.</p>
      <p v-if="isEditMode" class="edit-mode-note">You are in edit mode. You can changes the amount and merchant name.
      </p>

      <form class="payment-detail-form" v-if="payment">
        <div class="form-group">
          <label for="payment-id">Payment ID</label>
          <input id="payment-id" type="text" :value="payment.id" disabled />
        </div>

        <div class="form-group">
          <label for="merchant-name">Merchant Name</label>
          <input id="merchant-name" type="text" v-model="form.merchant_name" :disabled="!isEditMode" />
        </div>

        <div class="form-group">
          <label for="amount">Amount</label>
          <input id="amount" type="number" v-model.number="form.amount" :disabled="!isEditMode" />
        </div>

        <div class="form-group">
          <label for="status">Status</label>
          <input id="status" type="text" :value="payment.status" disabled />
        </div>

        <div class="form-group">
          <label for="date">Date</label>
          <input id="date" type="text" :value="formatDate(payment.created_at)" disabled />
        </div>

        <div class="form-group full">
          <button v-if="!isEditMode" @click="isEditMode = true" class="submit-btn">
            Edit
          </button>

          <button v-else @click="handleSave" :disabled="isSaving" class="submit-btn">
            {{ isSaving ? 'Saving...' : 'Save' }}
          </button>
        </div>
      </form>

      <p v-else-if="isLoading">Loading payment detail...</p>
      <p v-else>No payment data available.</p>
    </div>
  </DashboardLayout>
</template>

<script setup>
import DashboardLayout from "../../../components/layout/DashboardLayout.vue"
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { usePaymentStore } from "../../../stores/payment"
import { formatDate } from '../../../utils/format'

const route = useRoute()
const paymentStore = usePaymentStore()

const payment = ref(null)
const isLoading = ref(false)

onMounted(async () => {
  isLoading.value = true

  try {
    if (!paymentStore.payments.length) {
      await paymentStore.getPayments()
    }

    const id = Number(route.params.id)
    const found = paymentStore.payments.find(p => p.id === id)

    payment.value = found || null
    if (found) {
      payment.value = found
      initForm(found)
    }
  } finally {
    isLoading.value = false
  }
})


const isEditMode = ref(false)
const isSaving = ref(false)

const form = ref({
  merchant_name: '',
  amount: 0,
})

const initForm = (payment) => {
  form.value.merchant_name = payment.merchant_name
  form.value.amount = payment.amount
}

const handleSave = async () => {
  isSaving.value = true
  try {
    await paymentStore.updatePayment(
      Number(route.params.id),
      form.value
    )

    payment.value = {
      ...payment.value,
      ...form.value,
    }

    isEditMode.value = false
  } catch (e) {
    alert('Failed to update payment')
  } finally {
    isSaving.value = false
  }
}
</script>


<style scoped>
.payment-detail-view {
  max-width: 720px;
  background: #fff;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.payment-detail-view h1 {
  margin-bottom: 0.5rem;
  font-size: 1.6rem;
  color: #222;
}

.payment-detail-view p {
  margin-bottom: 2rem;
  color: #666;
  font-size: 0.95rem;
}

.payment-detail-form {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.25rem 1.5rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  font-size: 0.8rem;
  font-weight: 600;
  color: #555;
  margin-bottom: 0.4rem;
}

.form-group input {
  padding: 0.6rem 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 0.9rem;
  background-color: #f9f9f9;
  color: #333;
}

.form-group input:disabled {
  cursor: not-allowed;
}

/* Full width rows if needed */
.form-group.full {
  grid-column: span 2;
}

.edit-mode-note {
  background-color: #fff3cd;
  color: #856404;
  padding: 0.75rem 1rem;
  border-radius: 4px;
  margin-bottom: 1.5rem;
}

/* Mobile */
@media (max-width: 640px) {
  .payment-detail-form {
    grid-template-columns: 1fr;
  }
}
</style>
