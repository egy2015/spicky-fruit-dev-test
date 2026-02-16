<template>
  <tr class="payment-row">
    <td>{{ payment.id }}</td>
    <td>{{ payment.merchant_name || payment.merchantName }}</td>
    <td>{{ formatDate(payment.date || payment.created_at) }}</td>
    <td>{{ formatAmount(payment.amount) }}</td>
    <td>
      <StatusBadge :status="payment.status" />
    </td>
  </tr>
</template>

<script setup>
import StatusBadge from './StatusBadge.vue'

defineProps({
  payment: {
    type: Object,
    required: true,
  },
})

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}

const formatAmount = (amount) => {
  if (amount === null || amount === undefined) return 'N/A'
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  }).format(amount)
}
</script>

<style scoped>
.payment-row td {
  padding: 1rem;
  color: #333;
}
</style>
