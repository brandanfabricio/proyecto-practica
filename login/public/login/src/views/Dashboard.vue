<!-- src/views/Dashboard.vue -->
<script setup>
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'

const authStore = useAuthStore()
const router = useRouter()

function logout() {
    authStore.logout()
    router.push('/login')
}
</script>

<template>
    <div>
        <h2>Dashboard</h2>
        <p>Bienvenido al área protegida</p>
        <button @click="logout">Cerrar sesión</button>
    </div>
</template>
