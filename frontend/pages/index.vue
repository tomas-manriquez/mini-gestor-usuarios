<template>
  <div>
    <h1>User List</h1>
    <button @click="goAdd">Add New User</button>
    <div v-if="pending">Loading users…</div>
    <div v-else-if="error">Failed to load users.</div>
    <ul v-else>
      <li v-for="user in users" :key="user.id">
        {{ user.name }} — {{ user.email }}
        <button @click="toggle(user.id)">
          {{ expanded[user.id] ? 'Hide' : 'Show' }} details
        </button>
        <div v-if="expanded[user.id]">
          <pre>{{ user }}</pre>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { navigateTo, useFetch } from '#imports'

const { data: users, pending, error } = await useFetch('http://127.0.0.1:8080/users')
const expanded = reactive({})

function toggle(id) {
  expanded[id] = !expanded[id]
}

function goAdd() {
  navigateTo('/add-user')
}
</script>
