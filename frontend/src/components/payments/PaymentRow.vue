<template>
  <tr>
    <td><a @click="showDetail(payment.id)">View</a></td>
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
import router from '../../router'
import StatusBadge from './StatusBadge.vue'
import { formatDate } from '../../utils/format'

defineProps({
  payment: {
    type: Object,
    required: true,
  },
})

const showDetail = (paymentId) => {
  router.push('dashboard/payment/' + paymentId)
}

const formatAmount = (amount) => {
  if (amount === null || amount === undefined) return 'N/A'
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  }).format(amount)
}
</script>

