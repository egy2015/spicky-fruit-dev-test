<template>
  <div class="dashboard-layout">
    <header class="navbar">
      <div class="navbar-content">
        <h2><img src="../../../public/favicon.ico" width="20px" /> Spicky Fruit Monitoring</h2>
        <button @click="handleLogout" class="logout-btn">Logout</button>
      </div>
    </header>

    <div class="dashboard-body">
      <aside class="sidebar">
        <nav class="menu">
          <router-link to="/dashboard" class="menu-item" active-class="active">
            Dashboard Analytics
          </router-link>

          <router-link to="/profile" class="menu-item" active-class="active">
            Profile
          </router-link>
          <div class="mobile-nav">
            <hr style="margin: 1rem 0;" />
            <button @click="handleLogout" class="logout-btn">Logout</button>
          </div>
        </nav>
      </aside>

      <main class="main-content">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '../../stores/auth.js'

const router = useRouter()
const authStore = useAuthStore()

const handleLogout = () => {
  authStore.clearAuth()
  router.push('/login')
}
</script>

<style scoped>
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
