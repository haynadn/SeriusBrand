<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center min-h-screen">
      <div class="text-center">
        <div class="animate-spin w-16 h-16 border-4 border-blue-500 border-t-transparent rounded-full mx-auto mb-4"></div>
        <p class="text-white text-xl">Memuat halaman...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="flex items-center justify-center min-h-screen px-6">
      <div class="text-center">
        <div class="text-6xl mb-4">😕</div>
        <h1 class="text-4xl font-bold text-white mb-4">Halaman Tidak Ditemukan</h1>
        <p class="text-gray-300 mb-8">Produk yang kamu cari tidak tersedia</p>
        <router-link to="/" class="bg-gradient-to-r from-blue-500 to-purple-600 text-white px-8 py-3 rounded-xl font-semibold hover:shadow-xl transition-all">
          Kembali ke Beranda
        </router-link>
      </div>
    </div>

    <!-- Product Page -->
    <div v-else class="px-6 py-12">
      <div class="container mx-auto max-w-5xl">
        <!-- Back Button -->
        <router-link to="/" class="inline-flex items-center text-white hover:text-blue-400 mb-8 transition-all">
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
          </svg>
          Kembali
        </router-link>

        <div class="grid md:grid-cols-2 gap-12">
          <!-- Video & Media -->
          <div>
            <div class="bg-slate-900/80 backdrop-blur-md rounded-3xl overflow-hidden border border-slate-700 mb-6">
              <video 
                :src="getVideoUrl(umkmPage.video_url)" 
                controls 
                class="w-full"
              ></video>
              <div v-else class="aspect-video bg-slate-800 flex items-center justify-center">
                <p class="text-gray-500">Video tidak tersedia</p>
              </div>
            </div>
            
            <!-- Photo Gallery -->
            <div v-if="umkmPage.photos && umkmPage.photos.length > 0" class="grid grid-cols-3 gap-4">
              <div 
                v-for="(photo, index) in umkmPage.photos" 
                :key="index"
                class="aspect-square bg-slate-800 rounded-xl overflow-hidden border border-slate-700"
              >
                <img :src="'http://localhost:8080' + photo" class="w-full h-full object-cover" :alt="`Foto ${index + 1}`">
              </div>
            </div>
          </div>

          <!-- Product Info -->
          <div>
            <div class="bg-slate-900/80 backdrop-blur-md rounded-3xl p-8 border border-slate-700">
              <h1 class="text-4xl font-bold text-white mb-4">{{ umkmPage.product_name || 'Produk UMKM' }}</h1>
              
              <div v-if="umkmPage.price" class="mb-6">
                <p class="text-gray-400 text-sm mb-1">Harga</p>
                <p class="text-5xl font-extrabold bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
                  {{ umkmPage.price }}
                </p>
              </div>

              <div class="mb-8">
                <h3 class="text-xl font-bold text-white mb-3">Deskripsi Produk</h3>
                <p class="text-gray-300 leading-relaxed">
                  {{ umkmPage.description || 'Produk berkualitas dari UMKM lokal.' }}
                </p>
              </div>

              <!-- CTA Buttons -->
              <div class="space-y-4">
                <a 
                  v-if="umkmPage.whatsapp_link" 
                  :href="umkmPage.whatsapp_link" 
                  target="_blank"
                  class="flex items-center justify-center gap-3 bg-gradient-to-r from-green-500 to-emerald-600 hover:from-green-600 hover:to-emerald-700 text-white py-4 rounded-xl font-bold text-lg transition-all hover:shadow-2xl"
                >
                  <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/>
                  </svg>
                  Pesan via WhatsApp
                </a>

                <a 
                  v-if="umkmPage.marketplace_link" 
                  :href="umkmPage.marketplace_link" 
                  target="_blank"
                  class="flex items-center justify-center gap-3 bg-gradient-to-r from-orange-500 to-red-600 hover:from-orange-600 hover:to-red-700 text-white py-4 rounded-xl font-bold text-lg transition-all hover:shadow-2xl"
                >
                  <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"></path>
                  </svg>
                  Beli di Marketplace
                </a>
              </div>

              <!-- Powered by SeriusBrand -->
              <div class="mt-8 pt-8 border-t border-slate-700">
                <p class="text-gray-400 text-sm text-center">
                  Powered by <router-link to="/" class="text-blue-400 hover:text-blue-300 font-semibold">SeriusBrand</router-link>
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import api from '../api'

const route = useRoute()
const slug = route.params.slug

const umkmPage = ref({})
const loading = ref(true)
const error = ref(false)

function getVideoUrl(path) {
  if (!path) return ''
  const baseUrl = api.defaults.baseURL || 'http://localhost:8080/api'
  // Remove /api from base URL if present to get root URL
  const rootUrl = baseUrl.replace(/\/api$/, '')
  return `${rootUrl}${path}`
}

onMounted(async () => {
  try {
    const response = await api.get(`/umkm/${slug}`)
    umkmPage.value = response.data
    
    // Parse photos if it's a JSON string
    if (typeof umkmPage.value.photos === 'string') {
      try {
        umkmPage.value.photos = JSON.parse(umkmPage.value.photos)
      } catch (e) {
        umkmPage.value.photos = []
      }
    }
  } catch (err) {
    console.error('Failed to load UMKM page:', err)
    error.value = true
  } finally {
    loading.value = false
  }
})
</script>
