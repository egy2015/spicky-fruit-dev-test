<template>
  <table class="payments-table">
    <thead>
      <tr>
        <th>Payment ID</th>
        <th>Merchant Name</th>
        <th>Date</th>
        <th>Amount</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <PaymentRow v-for="payment in payments" :key="payment.id" :payment="payment" />
    </tbody>
  </table>
</template>

<script setup>
import PaymentRow from './PaymentRow.vue'

defineProps({
  payments: {
    type: Array,
    required: true,
  },
})
</script>

<style scoped>
.payments-table {
  width: 100%;
  border-collapse: collapse;
}

.payments-table thead {
  background-color: #f5f5f5;
}

.payments-table th {
  padding: 1rem;
  text-align: left;
  font-weight: 600;
  color: #333;
  font-size: 0.9rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.payments-table tbody tr:hover {
  background-color: #f9f9f9;
}

.payments-table tbody tr:not(:last-child) {
  border-bottom: 1px solid #eee;
}
</style>
