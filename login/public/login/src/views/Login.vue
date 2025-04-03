<!-- src/views/Login.vue -->
<script setup>
import { ref } from 'vue'
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'

const username = ref('')
const password = ref('')
const authStore = useAuthStore()
const router = useRouter()

async function handleLogin() {
    console.log(username.value,password.value,"aa@aa.com");
    
    const success = await authStore.login(username.value, password.value,"aa@aa.com")
    if (success) {
        router.push('/dashboard')
    } else {
        alert('Credenciales incorrectas')
    }
}
</script>

<template>
    <div>
        <h2>Login</h2>
        <input v-model="username" placeholder="Usuario" />
        <input v-model="password" type="password" placeholder="Contraseña" />
        <button @click="handleLogin">Ingresar</button>
    </div>
</template>
