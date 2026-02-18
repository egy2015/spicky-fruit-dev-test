<template>
  <div class="dashboard-layout">
    <header class="navbar">
      <div class="navbar-content">
        <h2><img src="../../../public/favicon.ico" width="20px" /> Spicky Fruit Monitoring</h2>
        <div @mouseenter="mouseOver = true" @mouseleave="mouseOver = false" class="user-info">
          <div class="avatar">
            {{ authStore.role.charAt(0, 1).toUpperCase() }}
          </div>
          <span v-show="mouseOver" class="user-label">
            {{ authStore.role }}
          </span>
          <button @click="handleLogout" class="logout-btn">Logout</button>
        </div>
      </div>
    </header>

    <div class="dashboard-body">
      <template v-if="!isCollapsed">
        <aside class="sidebar">
          <nav class="menu">
            <router-link to="/dashboard" class="menu-item" :class="{ 'active': router.currentRoute.value.path.includes('/dashboard') }">
              Dashboard Analytics
            </router-link>

            <router-link to="/users" class="menu-item" :class="{ 'active': router.currentRoute.value.path.includes('/users') }">
              User List
            </router-link>
            <div class="mobile-nav">
              <hr style="margin: 1rem 0;" />
              <button @click="handleLogout" class="logout-btn">Logout {{ authStore.role }}</button>
            </div>
          </nav>

        </aside>
      </template>
      <button @click="isCollapsed = !isCollapsed" class="collaptor">
        <span class="collaptor-text">{{ !isCollapsed ? 'Hide v' : 'Show ^' }}</span>
      </button>

      <main class="main-content">
        <div class="dashboard-content">
          <slot />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '../../stores/auth.js'
import { ref } from 'vue'

const mouseOver = ref(false)
const isCollapsed = ref(false);
const router = useRouter()
const authStore = useAuthStore()

const handleLogout = () => {
  authStore.clearAuth()
  router.push('/login')
}
</script>

<style scoped>
.collaptor {
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding-top: 2rem;
  background-color: white;
  border: none;
}

.collaptor-text {
  display: inline-block;
  transform: rotate(90deg);
  transform-origin: center;
}

.dashboard-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  background-color: #f5f5f5;
}

.navbar {
  background-color: #333;
  color: white;
  padding: 1rem 0;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.navbar-content {
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.navbar h2 {
  margin: 0;
  font-size: 1.5rem;
}

.logout-btn {
  padding: 0.5rem 1rem;
  background-color: var(--primary-color);
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: background-color 0.3s;
}

.logout-btn:hover {
  background-color: var(--secondary-color);
}

.main-content {
  flex: 1;
  max-width: 1400px;
  width: 100%;
  margin: 0 auto;
  padding: 0 2rem;
}

.dashboard-body {
  display: flex;
  flex: 1;
}

.sidebar {
  width: 220px;
  background-color: #fff;
  border-right: 1px solid #ddd;
  padding: 1.5rem 1rem;
}

.menu {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.menu-item {
  padding: 0.5rem 0.75rem;
  border-radius: 4px;
  color: #333;
  cursor: pointer;
  font-weight: 500;
}

.menu-item.active {
  background-color: var(--primary-color);
  color: #fff;
}

.menu-item:hover {
  background-color: var(--secondary-color);
  color: #fff;
}

.mobile-nav {
  display: none;
}


.dashboard-content {
  padding: 2rem;
}

.dashboard-content h1 {
  margin-bottom: 2rem;
  font-size: 2rem;
  color: #333;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background-color: var(--primary-color);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 0.9rem;
  text-transform: uppercase;
}

.user-label {
  font-size: 0.9rem;
  color: #eee;
  font-weight: 500;
}

@media (max-width: 768px) {
  .navbar {
    display: none;
  }

  .navbar-content {
    flex-direction: column;
    gap: 1rem;
  }

  .navbar-content h2 {
    width: 100%;
    text-align: center;
  }

  .mobile-nav {
    margin-top: 2rem;
    display: block;
  }

  .logout-btn {
    width: 100%;
  }
}
</style>
