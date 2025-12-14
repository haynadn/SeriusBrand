<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
    <!-- Header -->
    <header class="fixed top-0 left-0 right-0 z-50 bg-white/10 backdrop-blur-md border-b border-white/20">
      <div class="container mx-auto px-6 py-4">
        <div class="flex justify-between items-center">
          <h1 class="text-2xl font-bold text-white">SeriusBrand</h1>
          <button @click="scrollToOrder" class="bg-gradient-to-r from-blue-500 to-purple-600 text-white px-6 py-2 rounded-lg font-semibold hover:shadow-xl hover:scale-105 transition-all">
            Mulai Sekarang
          </button>
        </div>
      </div>
    </header>

    <!-- Hero Section -->
    <section class="pt-32 pb-20 px-6">
      <div class="container mx-auto text-center">
        <h2 class="text-5xl md:text-7xl font-extrabold text-white mb-6 leading-tight">
          Branding Produk UMKM<br />
          <span class="bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
            Mulai 10 Ribu Rupiah
          </span>
        </h2>
        <p class="text-xl md:text-2xl text-gray-300 mb-10 max-w-3xl mx-auto">
          Video profesional + Landing page eksklusif untuk produk UMKM kamu. Tingkatkan kredibilitas dan penjualan dengan branding yang kuat!
        </p>
        <button @click="scrollToOrder" class="bg-gradient-to-r from-blue-500 to-purple-600 text-white px-10 py-4 rounded-full text-lg font-bold hover:shadow-2xl hover:scale-110 transition-all duration-300">
          Buat Branding Sekarang →
        </button>
      </div>
    </section>

    <!-- Problem Section -->
    <section class="py-20 px-6 bg-white/5 backdrop-blur-sm">
      <div class="container mx-auto">
        <h3 class="text-4xl font-bold text-white text-center mb-12">Masalah yang Sering Dialami UMKM</h3>
        <div class="grid md:grid-cols-3 gap-8 max-w-5xl mx-auto">
          <div class="bg-gradient-to-br from-red-500/20 to-orange-500/20 backdrop-blur-md p-8 rounded-2xl border border-red-500/30 hover:scale-105 transition-all">
            <div class="text-5xl mb-4">😞</div>
            <h4 class="text-xl font-bold text-white mb-3">Produk Tidak Terlihat Profesional</h4>
            <p class="text-gray-300">Produk bagus tapi presentasi kurang menarik</p>
          </div>
          <div class="bg-gradient-to-br from-yellow-500/20 to-amber-500/20 backdrop-blur-md p-8 rounded-2xl border border-yellow-500/30 hover:scale-105 transition-all">
            <div class="text-5xl mb-4">🌐</div>
            <h4 class="text-xl font-bold text-white mb-3">Tidak Punya Website</h4>
            <p class="text-gray-300">Sulit bersaing di era digital tanpa presence online</p>
          </div>
          <div class="bg-gradient-to-br from-purple-500/20 to-pink-500/20 backdrop-blur-md p-8 rounded-2xl border border-purple-500/30 hover:scale-105 transition-all">
            <div class="text-5xl mb-4">💸</div>
            <h4 class="text-xl font-bold text-white mb-3">Budget Terbatas</h4>
            <p class="text-gray-300">Biaya branding profesional terlalu mahal</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Solution Section -->
    <section class="py-20 px-6">
      <div class="container mx-auto">
        <h3 class="text-4xl font-bold text-white text-center mb-12">Solusi dari SeriusBrand</h3>
        <div class="grid md:grid-cols-3 gap-8 max-w-5xl mx-auto">
          <div class="bg-gradient-to-br from-blue-500/30 to-cyan-500/30 backdrop-blur-md p-8 rounded-2xl border border-blue-500/50 hover:scale-105 transition-all">
            <div class="text-5xl mb-4">🎥</div>
            <h4 class="text-xl font-bold text-white mb-3">Video Packing Profesional</h4>
            <p class="text-gray-300">Video berkualitas tinggi untuk produk kamu</p>
          </div>
          <div class="bg-gradient-to-br from-green-500/30 to-emerald-500/30 backdrop-blur-md p-8 rounded-2xl border border-green-500/50 hover:scale-105 transition-all">
            <div class="text-5xl mb-4">🚀</div>
            <h4 class="text-xl font-bold text-white mb-3">Landing Page Instan</h4>
            <p class="text-gray-300">Website produk siap dalam hitungan hari</p>
          </div>
          <div class="bg-gradient-to-br from-purple-500/30 to-pink-500/30 backdrop-blur-md p-8 rounded-2xl border border-purple-500/50 hover:scale-105 transition-all">
            <div class="text-5xl mb-4">💰</div>
            <h4 class="text-xl font-bold text-white mb-3">Harga Terjangkau</h4>
            <p class="text-gray-300">Mulai dari 10 ribu per produk!</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Pricing Section -->
    <section id="pricing" class="py-20 px-6 bg-white/5 backdrop-blur-sm">
      <div class="container mx-auto">
        <h3 class="text-4xl font-bold text-white text-center mb-4">Pilih Paket yang Sesuai</h3>
        <p class="text-gray-300 text-center mb-12 text-lg">Investasi kecil untuk hasil maksimal</p>
        
        <div class="grid md:grid-cols-3 gap-8 max-w-6xl mx-auto">
          <!-- Starter Package -->
          <div class="bg-gradient-to-b from-slate-800/80 to-slate-900/80 backdrop-blur-md p-8 rounded-3xl border-2 border-slate-700 hover:border-blue-500 hover:scale-105 transition-all duration-300">
            <div class="text-center mb-6">
              <h4 class="text-2xl font-bold text-white mb-2">Starter</h4>
              <div class="text-5xl font-extrabold text-blue-400 mb-2">Rp10.000</div>
              <p class="text-gray-400">per produk</p>
            </div>
            <ul class="space-y-4 mb-8">
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">1 video packing (15-30 detik)</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">1 landing page produk</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Foto produk dari video</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Tombol WhatsApp & Marketplace</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Hosting aktif 30 hari</span>
              </li>
            </ul>
            <button @click="selectPackage('starter')" class="w-full bg-blue-600 hover:bg-blue-700 text-white py-3 rounded-xl font-semibold transition-all">
              Pilih Paket
            </button>
          </div>

          <!-- Growth Package -->
          <div class="bg-gradient-to-b from-purple-800/80 to-purple-900/80 backdrop-blur-md p-8 rounded-3xl border-2 border-purple-500 transform scale-105 shadow-2xl">
            <div class="bg-gradient-to-r from-yellow-400 to-orange-500 text-black text-xs font-bold px-3 py-1 rounded-full inline-block mb-4">
              TERPOPULER
            </div>
            <div class="text-center mb-6">
              <h4 class="text-2xl font-bold text-white mb-2">Growth</h4>
              <div class="text-5xl font-extrabold text-purple-400 mb-2">Rp25.000</div>
              <p class="text-gray-400">per produk</p>
            </div>
            <ul class="space-y-4 mb-8">
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Semua di Starter</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Video 30-60 detik</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Copywriting deskripsi produk</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Tombol Order + Upload Bukti</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Hosting 90 hari</span>
              </li>
            </ul>
            <button @click="selectPackage('growth')" class="w-full bg-gradient-to-r from-purple-500 to-pink-600 hover:from-purple-600 hover:to-pink-700 text-white py-3 rounded-xl font-semibold transition-all">
              Pilih Paket
            </button>
          </div>

          <!-- Pro Package -->
          <div class="bg-gradient-to-b from-slate-800/80 to-slate-900/80 backdrop-blur-md p-8 rounded-3xl border-2 border-slate-700 hover:border-purple-500 hover:scale-105 transition-all duration-300">
            <div class="text-center mb-6">
              <h4 class="text-2xl font-bold text-white mb-2">Pro</h4>
              <div class="text-5xl font-extrabold text-purple-400 mb-2">Rp50.000</div>
              <p class="text-gray-400">per produk</p>
            </div>
            <ul class="space-y-4 mb-8">
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Semua di Growth</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">3 variasi video (Reels/TikTok/Story)</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Custom subdomain</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Statistik pengunjung</span>
              </li>
              <li class="flex items-start">
                <span class="text-green-400 mr-2">✓</span>
                <span class="text-gray-300">Hosting 6 bulan</span>
              </li>
            </ul>
            <button @click="selectPackage('pro')" class="w-full bg-purple-600 hover:bg-purple-700 text-white py-3 rounded-xl font-semibold transition-all">
              Pilih Paket
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- How It Works -->
    <section class="py-20 px-6">
      <div class="container mx-auto max-w-4xl">
        <h3 class="text-4xl font-bold text-white text-center mb-12">Cara Kerja</h3>
        <div class="space-y-6">
          <div class="flex items-start gap-6 bg-white/5 backdrop-blur-md p-6 rounded-2xl border border-white/10">
            <div class="bg-gradient-to-r from-blue-500 to-purple-600 text-white w-12 h-12 rounded-full flex items-center justify-center font-bold text-xl flex-shrink-0">1</div>
            <div>
              <h4 class="text-xl font-bold text-white mb-2">Pilih Paket</h4>
              <p class="text-gray-300">Pilih paket yang sesuai dengan kebutuhan dan budget kamu</p>
            </div>
          </div>
          <div class="flex items-start gap-6 bg-white/5 backdrop-blur-md p-6 rounded-2xl border border-white/10">
            <div class="bg-gradient-to-r from-blue-500 to-purple-600 text-white w-12 h-12 rounded-full flex items-center justify-center font-bold text-xl flex-shrink-0">2</div>
            <div>
              <h4 class="text-xl font-bold text-white mb-2">Isi Form & Transfer</h4>
              <p class="text-gray-300">Lengkapi form pemesanan dan lakukan pembayaran</p>
            </div>
          </div>
          <div class="flex items-start gap-6 bg-white/5 backdrop-blur-md p-6 rounded-2xl border border-white/10">
            <div class="bg-gradient-to-r from-blue-500 to-purple-600 text-white w-12 h-12 rounded-full flex items-center justify-center font-bold text-xl flex-shrink-0">3</div>
            <div>
              <h4 class="text-xl font-bold text-white mb-2">Kirim Produk/Video Mentah</h4>
              <p class="text-gray-300">Kirim produk atau video mentah untuk proses editing</p>
            </div>
          </div>
          <div class="flex items-start gap-6 bg-white/5 backdrop-blur-md p-6 rounded-2xl border border-white/10">
            <div class="bg-gradient-to-r from-blue-500 to-purple-600 text-white w-12 h-12 rounded-full flex items-center justify-center font-bold text-xl flex-shrink-0">4</div>
            <div>
              <h4 class="text-xl font-bold text-white mb-2">Terima Link Landing Page</h4>
              <p class="text-gray-300">Dalam beberapa hari, landing page produk kamu siap dibagikan!</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- UMKM Showcase Section -->
    <section v-if="umkmPages.length > 0" class="py-20 px-6 bg-white/5 backdrop-blur-sm">
      <div class="container mx-auto">
        <h3 class="text-4xl font-bold text-white text-center mb-12">UMKM yang Sudah Bergabung</h3>
        <div class="grid md:grid-cols-3 gap-8">
          <div v-for="page in umkmPages" :key="page.id" class="bg-slate-900/80 backdrop-blur-md rounded-2xl overflow-hidden border border-slate-700 hover:border-blue-500 transition-all hover:scale-105 group">
            <div class="aspect-video bg-slate-800 relative overflow-hidden">
               <video 
                v-if="page.video_url" 
                :src="'http://localhost:8080' + page.video_url" 
                class="w-full h-full object-cover opacity-80 group-hover:opacity-100 transition-opacity"
                muted
                loop
                
                @mouseover="$event.target.play()"
                @mouseout="$event.target.pause()"
              ></video>
              <div v-else class="w-full h-full flex items-center justify-center text-gray-500">
                <span>No Video</span>
              </div>
            </div>
            <div class="p-6">
              <h4 class="text-xl font-bold text-white mb-2">{{ page.slug.replace(/-/g, ' ').toUpperCase() }}</h4>
              <p class="text-gray-400 text-sm mb-4 line-clamp-2">{{ page.description || 'Produk UMKM berkualitas' }}</p>
              <router-link :to="'/umkm/' + page.slug" class="block w-full text-center bg-white/10 hover:bg-white/20 text-white py-2 rounded-lg transition-colors">
                Lihat Halaman
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Order Form -->
    <section id="order" class="py-20 px-6 bg-white/5 backdrop-blur-sm">
      <div class="container mx-auto max-w-2xl">
        <h3 class="text-4xl font-bold text-white text-center mb-4">Mulai Branding Sekarang!</h3>
        <p class="text-gray-300 text-center mb-10">Isi form di bawah untuk memulai perjalanan branding produk kamu</p>
        
        <form @submit.prevent="submitOrder" class="bg-slate-900/80 backdrop-blur-md p-8 rounded-3xl border border-slate-700">
          <div class="mb-6">
            <label class="block text-white font-semibold mb-2">Nama UMKM / Pemilik *</label>
            <input v-model="orderForm.customerName" type="text" required class="w-full px-4 py-3 bg-slate-800 text-white rounded-xl border border-slate-700 focus:border-blue-500 focus:outline-none transition-all" placeholder="Contoh: Bu Siti">
          </div>
          
          <div class="mb-6">
            <label class="block text-white font-semibold mb-2">Nomor WhatsApp *</label>
            <input v-model="orderForm.whatsapp" type="tel" required class="w-full px-4 py-3 bg-slate-800 text-white rounded-xl border border-slate-700 focus:border-blue-500 focus:outline-none transition-all" placeholder="08xxxxxxxxxx">
          </div>
          
          <div class="mb-6">
            <label class="block text-white font-semibold mb-2">Nama Produk *</label>
            <input v-model="orderForm.productName" type="text" required class="w-full px-4 py-3 bg-slate-800 text-white rounded-xl border border-slate-700 focus:border-blue-500 focus:outline-none transition-all" placeholder="Contoh: Keripik Pisang Original">
          </div>
          
          <div class="mb-8">
            <label class="block text-white font-semibold mb-2">Pilih Paket *</label>
            <select v-model="orderForm.packageType" required class="w-full px-4 py-3 bg-slate-800 text-white rounded-xl border border-slate-700 focus:border-blue-500 focus:outline-none transition-all">
              <option value="">-- Pilih Paket --</option>
              <option value="starter">Starter - Rp10.000</option>
              <option value="growth">Growth - Rp25.000</option>
              <option value="pro">Pro - Rp50.000</option>
            </select>
          </div>
          
          <button type="submit" :disabled="loading" class="w-full bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 disabled:from-gray-600 disabled:to-gray-700 text-white py-4 rounded-xl font-bold text-lg transition-all hover:shadow-2xl">
            {{ loading ? 'Memproses...' : 'Buat Pesanan' }}
          </button>
          
          <p class="text-gray-400 text-sm text-center mt-4">
            Setelah submit, kamu akan diarahkan ke halaman pembayaran
          </p>
        </form>
      </div>
    </section>

    <!-- Footer -->
    <footer class="py-10 px-6 bg-black/40 backdrop-blur-md border-t border-white/10">
      <div class="container mx-auto text-center">
        <h3 class="text-2xl font-bold text-white mb-2">SeriusBrand</h3>
        <p class="text-gray-400 mb-4">Branding Profesional untuk UMKM Indonesia</p>
        <p class="text-gray-500 text-sm">© 2025 SeriusBrand. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api'

const router = useRouter()
const loading = ref(false)
const umkmPages = ref([])

const orderForm = ref({
  customerName: '',
  whatsapp: '',
  productName: '',
  packageType: ''
})

onMounted(async () => {
  try {
    const response = await api.get('/umkm')
    umkmPages.value = response.data || []
  } catch (error) {
    console.error('Failed to fetch UMKM pages:', error)
  }
})

function scrollToOrder() {
  document.getElementById('order').scrollIntoView({ behavior: 'smooth' })
}

function selectPackage(packageType) {
  orderForm.value.packageType = packageType
  scrollToOrder()
}

async function submitOrder() {
  loading.value = true
  
  // Transform to snake_case for backend
  const payload = {
    customer_name: orderForm.value.customerName,
    whatsapp: orderForm.value.whatsapp,
    product_name: orderForm.value.productName,
    package_type: orderForm.value.packageType
  }
  
  try {
    const response = await api.post('/orders', payload)
    
    // Redirect to success page with order ID
    router.push(`/order-success/${response.data.id}`)
  } catch (error) {
    console.error('Order failed:', error)
    // Show more specific error if available
    const msg = error.response?.data?.error || 'Gagal membuat pesanan. Silakan coba lagi.'
    alert(msg)
  } finally {
    loading.value = false
  }
}
</script>
