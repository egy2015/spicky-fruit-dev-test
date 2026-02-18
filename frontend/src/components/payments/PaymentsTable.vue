<template>
  <table class="custom-table">
    <thead>
      <tr>
        <th>Action</th>
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

