<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 py-12 px-4 sm:px-6 lg:px-8">
    <!-- Animated Background Elements -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
      <div class="absolute top-1/4 left-1/4 w-96 h-96 bg-purple-500/10 rounded-full blur-3xl animate-pulse"></div>
      <div class="absolute bottom-1/4 right-1/4 w-96 h-96 bg-blue-500/10 rounded-full blur-3xl animate-pulse delay-1000"></div>
    </div>

    <div class="max-w-md w-full space-y-8 relative z-10">
      <!-- Logo/Brand Section -->
      <div class="text-center">
        <div class="inline-block bg-gradient-to-r from-blue-500 to-purple-600 p-4 rounded-2xl mb-4 shadow-2xl">
          <svg class="w-12 h-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
          </svg>
        </div>
        <h2 class="text-4xl font-extrabold text-white mb-2">
          Admin Portal
        </h2>
        <p class="text-gray-400">Sign in to access dashboard</p>
      </div>

      <!-- Login Form -->
      <div class="bg-slate-900/80 backdrop-blur-md rounded-3xl border border-slate-700 p-8 shadow-2xl">
        <form class="space-y-6" @submit.prevent="handleLogin">
          <div class="space-y-4">
            <div>
              <label for="username" class="block text-sm font-semibold text-gray-300 mb-2">Username</label>
              <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="h-5 w-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                  </svg>
                </div>
                <input 
                  id="username" 
                  name="username" 
                  type="text" 
                  v-model="username" 
                  required 
                  class="appearance-none block w-full pl-10 pr-3 py-3 bg-slate-800 border border-slate-700 text-white rounded-xl placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all" 
                  placeholder="Enter your username"
                >
              </div>
            </div>
            
            <div>
              <label for="password" class="block text-sm font-semibold text-gray-300 mb-2">Password</label>
              <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="h-5 w-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
                  </svg>
                </div>
                <input 
                  id="password" 
                  name="password" 
                  type="password" 
                  v-model="password" 
                  required 
                  class="appearance-none block w-full pl-10 pr-3 py-3 bg-slate-800 border border-slate-700 text-white rounded-xl placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all" 
                  placeholder="Enter your password"
                >
              </div>
            </div>
          </div>

          <div v-if="error" class="bg-red-500/20 border border-red-500/50 text-red-400 px-4 py-3 rounded-xl text-sm flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            {{ error }}
          </div>

          <button 
            type="submit" 
            :disabled="isLoading"
            class="group relative w-full flex justify-center items-center gap-2 py-3 px-4 border border-transparent text-base font-bold rounded-xl text-white bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-all disabled:opacity-50 disabled:cursor-not-allowed shadow-lg hover:shadow-2xl hover:scale-105"
          >
            <svg v-if="!isLoading" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"></path>
            </svg>
            <div v-else class="animate-spin w-5 h-5 border-2 border-white border-t-transparent rounded-full"></div>
            {{ isLoading ? 'Signing in...' : 'Sign In' }}
          </button>
        </form>

        <!-- Back to Home Link -->
        <div class="mt-6 text-center">
          <router-link to="/" class="text-sm text-gray-400 hover:text-white transition-colors inline-flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
            </svg>
            Back to Homepage
          </router-link>
        </div>
      </div>

      <!-- Footer -->
      <p class="text-center text-gray-500 text-sm">
        © 2025 SeriusBrand. Secure Admin Access.
      </p>
    </div>

    <!-- Loading Modal -->
    <div v-if="isLoading" class="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center z-50">
      <div class="bg-slate-900/90 backdrop-blur-md p-8 rounded-3xl shadow-2xl border border-slate-700 flex flex-col items-center">
        <div class="animate-spin w-16 h-16 border-4 border-blue-500 border-t-transparent rounded-full mb-4"></div>
        <p class="text-white font-semibold text-lg">Authenticating...</p>
        <p class="text-gray-400 text-sm mt-1">Please wait</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api'

const router = useRouter()
const username = ref('')
const password = ref('')
const error = ref('')
const isLoading = ref(false)

const handleLogin = async () => {
  isLoading.value = true
  error.value = ''
  
  try {
    const response = await api.post('/login', {
      username: username.value,
      password: password.value
    })
    
    // Store token
    localStorage.setItem('admin_token', response.data.token)
    
    // Redirect
    router.push('/admin')
  } catch (err) {
    error.value = err.response?.data?.error || 'Login failed'
  } finally {
    isLoading.value = false
  }
}
</script>
