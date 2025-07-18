<template>
  <div>
    <h1>Add User</h1>
    <form @submit.prevent="submit">
      <div><label>Name: <input v-model="form.name" required></label></div>
      <div><label>Email: <input type="email" v-model="form.email" required></label></div>
      <div><label>Birth Date: <input type="date" v-model="form.birth_date" required></label></div>
      <div><label>Created At: <input type="datetime-local" v-model="form.created_at" required></label></div>
      <button type="submit">Submit</button>
      <button type="button" @click="back">Back</button>
    </form>
    <div v-if="error">{{ error }}</div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { navigateTo } from '#imports'

const form = ref({
  name: '',
  email: '',
  birth_date: '',
  created_at: ''
})

const error = ref(null)

async function submit() {
  error.value = null
  try {
    await $fetch('http://127.0.0.1:8080/users', {
      method: 'POST',
      body: form.value
    })
    navigateTo('/')   // redirect after success
  } catch (e) {
    error.value = 'Failed to add user'
    console.error(e)
  }
}

function back() {
  navigateTo('/')
}
</script>
