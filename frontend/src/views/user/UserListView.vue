<template>
   <DashboardLayout>
      <h1>User List</h1>
      <p>List of registered users.</p>
      <div class="table-container">
        <table class="custom-table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Email</th>
            <th>Role</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.id">
            <td>{{ user.id }}</td>
            <td>{{ user.email }}</td>
            <td>{{ user.role }}</td>
          </tr>
        </tbody>
      </table>
      </div>
   </DashboardLayout>
</template>

<script>
import DashboardLayout from '../../components/layout/DashboardLayout.vue'
import { useUserStore } from '../../stores/user'

export default {
  components: {
    DashboardLayout,
  },
   computed: {
    users() {
      const userStore = useUserStore()
      return userStore.users
    },
  },
  mounted() {
    const userStore = useUserStore()
    userStore.getUsers()
  }
}

</script>

<style scoped>
h1 {
  margin-bottom: 1rem;
  font-size: 1.5rem;
  color: #333;
}
p {
  margin-bottom: 2rem;
  color: #666;
}
</style>
