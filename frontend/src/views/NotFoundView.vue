<template>
  <div class="not-found">
    <h1>404</h1>
    <p>Page not found</p>
    <RouterLink to="/dashboard" class="home-link">Go to Dashboard</RouterLink>
  </div>
</template>

<script setup>
import { RouterLink } from 'vue-router'
</script>

<style scoped>
.not-found {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f5f5;
  text-align: center;
}

.not-found h1 {
  font-size: 4rem;
  margin: 0;
  color: #333;
}

.not-found p {
  font-size: 1.2rem;
  color: #666;
  margin: 0.5rem 0 2rem;
}

.home-link {
  padding: 0.75rem 1.5rem;
  background-color: #0066cc;
  color: white;
  text-decoration: none;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.home-link:hover {
  background-color: #0052a3;
}
</style>
