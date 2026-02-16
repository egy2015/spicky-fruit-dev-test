<template>
  <span :class="['status-badge', `status-${status}`]">
    {{ status.charAt(0).toUpperCase() + status.slice(1) }}
  </span>
</template>

<script setup>
defineProps({
  status: {
    type: String,
    required: true,
  },
})
</script>

<style scoped>
.status-badge {
  display: inline-block;
  padding: 0.25rem 0.75rem;
  border-radius: 12px;
  font-size: 0.85rem;
  font-weight: 500;
  text-transform: capitalize;
}

.status-completed {
  background-color: #d4edda;
  color: #155724;
}

.status-processing {
  background-color: #fff3cd;
  color: #856404;
}

.status-failed {
  background-color: #f8d7da;
  color: #721c24;
}
</style>
