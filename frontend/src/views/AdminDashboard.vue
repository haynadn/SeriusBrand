<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
    <!-- Header -->
    <header class="bg-slate-900/80 backdrop-blur-md border-b border-slate-700">
      <div class="container mx-auto px-6 py-4">
        <div class="flex justify-between items-center">
          <h1 class="text-2xl font-bold text-white">SeriusBrand Admin</h1>
          <router-link to="/" class="text-gray-300 hover:text-white transition-all">
            ← Kembali ke Beranda
          </router-link>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <div class="container mx-auto px-6 py-12">
      <!-- Stats Cards -->
      <div class="grid md:grid-cols-3 gap-6 mb-8">
        <div class="bg-gradient-to-br from-blue-500/20 to-cyan-500/20 backdrop-blur-md p-6 rounded-2xl border border-blue-500/50">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-gray-300 text-sm mb-1">Total Pesanan</p>
              <p class="text-4xl font-bold text-white">{{ stats.total }}</p>
            </div>
            <div class="text-5xl">📦</div>
          </div>
        </div>
        <div class="bg-gradient-to-br from-yellow-500/20 to-orange-500/20 backdrop-blur-md p-6 rounded-2xl border border-yellow-500/50">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-gray-300 text-sm mb-1">Menunggu Verifikasi</p>
              <p class="text-4xl font-bold text-white">{{ stats.pending }}</p>
            </div>
            <div class="text-5xl">⏳</div>
          </div>
        </div>
        <div class="bg-gradient-to-br from-green-500/20 to-emerald-500/20 backdrop-blur-md p-6 rounded-2xl border border-green-500/50">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-gray-300 text-sm mb-1">Selesai</p>
              <p class="text-4xl font-bold text-white">{{ stats.completed }}</p>
            </div>
            <div class="text-5xl">✅</div>
          </div>
        </div>
      </div>

      <!-- Orders Table -->
      <div class="bg-slate-900/80 backdrop-blur-md rounded-3xl border border-slate-700 overflow-hidden">
        <div class="p-6 border-b border-slate-700">
          <h2 class="text-2xl font-bold text-white">Daftar Pesanan</h2>
        </div>

        <div v-if="loading" class="p-12 text-center">
          <div class="animate-spin w-12 h-12 border-4 border-blue-500 border-t-transparent rounded-full mx-auto mb-4"></div>
          <p class="text-gray-400">Memuat data...</p>
        </div>

        <div v-else-if="orders.length === 0" class="p-12 text-center">
          <div class="text-6xl mb-4">📭</div>
          <p class="text-gray-400">Belum ada pesanan</p>
        </div>

        <div v-else class="overflow-x-auto">
          <table class="w-full">
            <thead class="bg-slate-800/50">
              <tr>
                <th class="px-6 py-4 text-left text-sm font-semibold text-gray-300">ID</th>
                <th class="px-6 py-4 text-left text-sm font-semibold text-gray-300">Customer</th>
                <th class="px-6 py-4 text-left text-sm font-semibold text-gray-300">Produk</th>
                <th class="px-6 py-4 text-left text-sm font-semibold text-gray-300">Paket</th>
                <th class="px-6 py-4 text-left text-sm font-semibold text-gray-300">Harga</th>
                <th class="px-6 py-4 text-left text-sm font-semibold text-gray-300">Status</th>
                <th class="px-6 py-4 text-left text-sm font-semibold text-gray-300">Bukti</th>
                <th class="px-6 py-4 text-left text-sm font-semibold text-gray-300">Aksi</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-700">
              <tr v-for="order in orders" :key="order.id" class="hover:bg-slate-800/30 transition-all">
                <td class="px-6 py-4 text-white font-mono text-sm">#{{ order.id }}</td>
                <td class="px-6 py-4">
                  <div class="text-white font-semibold">{{ order.customer_name }}</div>
                  <div class="text-gray-400 text-sm">{{ order.whatsapp }}</div>
                </td>
                <td class="px-6 py-4 text-white">{{ order.product_name }}</td>
                <td class="px-6 py-4">
                  <span class="px-3 py-1 rounded-full text-xs font-semibold" :class="getPackageBadgeClass(order.package_type)">
                    {{ order.package_type }}
                  </span>
                </td>
                <td class="px-6 py-4 text-white font-semibold">Rp {{ formatPrice(order.price) }}</td>
                <td class="px-6 py-4">
                  <select 
                    v-model="order.status" 
                    @change="updateOrderStatus(order.id, order.status)"
                    class="bg-slate-800 text-white px-3 py-2 rounded-lg border border-slate-600 focus:border-blue-500 focus:outline-none text-sm"
                  >
                    <option value="pending">Pending</option>
                    <option value="paid">Paid</option>
                    <option value="completed">Completed</option>
                  </select>
                </td>
                <td class="px-6 py-4">
                  <button 
                    v-if="order.payment_proof_url" 
                    @click="viewProof(order.payment_proof_url)"
                    class="bg-blue-600 hover:bg-blue-700 text-white px-3 py-1 rounded-lg text-sm transition-all"
                  >
                    Lihat Bukti
                  </button>
                  <span v-else class="text-gray-500 text-sm">Belum ada</span>
                </td>
                <td class="px-6 py-4">
                  <button 
                    @click="createUmkmPage(order)"
                    :disabled="order.status !== 'paid'"
                    class="bg-gradient-to-r from-purple-500 to-pink-600 hover:from-purple-600 hover:to-pink-700 disabled:from-gray-600 disabled:to-gray-700 text-white px-4 py-2 rounded-lg text-sm font-semibold transition-all disabled:cursor-not-allowed"
                  >
                    Buat Halaman
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Proof Modal -->
    <div v-if="showProofModal" @click="showProofModal = false" class="fixed inset-0 bg-black/80 backdrop-blur-sm flex items-center justify-center z-50 p-6">
      <div @click.stop class="bg-slate-900 rounded-3xl p-6 max-w-2xl w-full border border-slate-700">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-bold text-white">Bukti Transfer</h3>
          <button @click="showProofModal = false" class="text-gray-400 hover:text-white text-2xl">×</button>
        </div>
        <img :src="'http://localhost:8081' + currentProof" class="w-full rounded-xl" alt="Bukti Transfer">
      </div>
    </div>

    <!-- Create UMKM Page Modal -->
    <div v-if="showCreateModal" @click="showCreateModal = false" class="fixed inset-0 bg-black/80 backdrop-blur-sm flex items-center justify-center z-50 p-6">
      <div @click.stop class="bg-slate-900 rounded-3xl p-8 max-w-2xl w-full border border-slate-700 max-h-[90vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-2xl font-bold text-white">Buat Halaman UMKM</h3>
          <button @click="showCreateModal = false" class="text-gray-400 hover:text-white text-2xl">×</button>
        </div>

        <form @submit.prevent="submitUmkmPage" class="space-y-4">
          <div>
            <label class="block text-white font-semibold mb-2">Slug (URL) *</label>
            <div class="flex items-center gap-2">
              <span class="text-gray-400 text-sm">seriusbrand.id/umkm/</span>
              <input v-model="umkmForm.slug" required class="flex-1 px-4 py-3 bg-slate-800 text-white rounded-xl border border-slate-700 focus:border-blue-500 focus:outline-none" placeholder="keripik-siti">
            </div>
          </div>

          <div>
            <label class="block text-white font-semibold mb-2">URL Video</label>
            <input v-model="umkmForm.video_url" class="w-full px-4 py-3 bg-slate-800 text-white rounded-xl border border-slate-700 focus:border-blue-500 focus:outline-none" placeholder="/uploads/videos/product.mp4">
          </div>

          <div>
            <label class="block text-white font-semibold mb-2">Deskripsi Produk</label>
            <textarea v-model="umkmForm.description" rows="4" class="w-full px-4 py-3 bg-slate-800 text-white rounded-xl border border-slate-700 focus:border-blue-500 focus:outline-none" placeholder="Keripik pisang original dengan rasa yang renyah..."></textarea>
          </div>

          <div>
            <label class="block text-white font-semibold mb-2">Harga Produk</label>
            <input v-model="umkmForm.price" class="w-full px-4 py-3 bg-slate-800 text-white rounded-xl border border-slate-700 focus:border-blue-500 focus:outline-none" placeholder="Rp 25.000">
          </div>

          <div>
            <label class="block text-white font-semibold mb-2">Link WhatsApp</label>
            <input v-model="umkmForm.whatsapp_link" class="w-full px-4 py-3 bg-slate-800 text-white rounded-xl border border-slate-700 focus:border-blue-500 focus:outline-none" placeholder="https://wa.me/628123456789">
          </div>

          <div>
            <label class="block text-white font-semibold mb-2">Link Marketplace</label>
            <input v-model="umkmForm.marketplace_link" class="w-full px-4 py-3 bg-slate-800 text-white rounded-xl border border-slate-700 focus:border-blue-500 focus:outline-none" placeholder="https://shopee.co.id/...">
          </div>

          <button type="submit" :disabled="submitting" class="w-full bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 disabled:from-gray-600 disabled:to-gray-700 text-white py-4 rounded-xl font-bold text-lg transition-all">
            {{ submitting ? 'Membuat...' : 'Buat Halaman UMKM' }}
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import api from '../api'

const orders = ref([])
const loading = ref(true)
const showProofModal = ref(false)
const currentProof = ref('')
const showCreateModal = ref(false)
const submitting = ref(false)
const selectedOrder = ref(null)

const umkmForm = ref({
  order_id: 0,
  slug: '',
  video_url: '',
  description: '',
  price: '',
  whatsapp_link: '',
  marketplace_link: '',
  photos: '[]'
})

const stats = computed(() => ({
  total: orders.value.length,
  pending: orders.value.filter(o => o.status === 'pending').length,
  completed: orders.value.filter(o => o.status === 'completed').length
}))

onMounted(async () => {
  await fetchOrders()
})

async function fetchOrders() {
  try {
    const response = await api.get('/admin/orders')
    orders.value = response.data || []
  } catch (error) {
    console.error('Failed to fetch orders:', error)
    alert('Gagal memuat data pesanan')
  } finally {
    loading.value = false
  }
}

async function updateOrderStatus(orderId, newStatus) {
  try {
    await api.patch(`/admin/orders/${orderId}`, { status: newStatus })
    alert('Status pesanan berhasil diupdate!')
  } catch (error) {
    console.error('Failed to update status:', error)
    alert('Gagal mengupdate status')
    await fetchOrders()
  }
}

function viewProof(proofUrl) {
  currentProof.value = proofUrl
  showProofModal.value = true
}

function createUmkmPage(order) {
  selectedOrder.value = order
  umkmForm.value.order_id = order.id
  umkmForm.value.slug = order.product_name.toLowerCase().replace(/\s+/g, '-').replace(/[^a-z0-9-]/g, '')
  umkmForm.value.whatsapp_link = `https://wa.me/${order.whatsapp.replace(/^0/, '62')}`
  showCreateModal.value = true
}

async function submitUmkmPage() {
  submitting.value = true
  try {
    await api.post('/admin/umkm', umkmForm.value)
    alert('Halaman UMKM berhasil dibuat!')
    showCreateModal.value = false
    
    // Update order status to completed
    await updateOrderStatus(selectedOrder.value.id, 'completed')
    await fetchOrders()
  } catch (error) {
    console.error('Failed to create UMKM page:', error)
    alert('Gagal membuat halaman UMKM')
  } finally {
    submitting.value = false
  }
}

function getPackageBadgeClass(packageType) {
  const classes = {
    starter: 'bg-blue-500/20 text-blue-400 border border-blue-500/50',
    growth: 'bg-purple-500/20 text-purple-400 border border-purple-500/50',
    pro: 'bg-pink-500/20 text-pink-400 border border-pink-500/50'
  }
  return classes[packageType] || classes.starter
}

function formatPrice(price) {
  return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.')
}
</script>
