<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 flex items-center justify-center px-6">
    <div class="max-w-2xl w-full bg-slate-900/80 backdrop-blur-md p-10 rounded-3xl border border-slate-700 text-center">
      <!-- Success Icon -->
      <div class="w-20 h-20 bg-gradient-to-r from-green-400 to-emerald-500 rounded-full flex items-center justify-center mx-auto mb-6">
        <svg class="w-12 h-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"></path>
        </svg>
      </div>
      
      <h1 class="text-4xl font-bold text-white mb-4">Pesanan Berhasil Dibuat! 🎉</h1>
      <p class="text-gray-300 text-lg mb-8">
        Terima kasih! Pesanan kamu telah kami terima.
      </p>

      <!-- Order ID -->
      <div class="bg-slate-800/50 p-6 rounded-2xl mb-8">
        <p class="text-gray-400 text-sm mb-2">ID Pesanan</p>
        <p class="text-3xl font-bold text-blue-400">#{{ orderId }}</p>
      </div>

      <!-- Payment Info -->
      <div class="bg-blue-500/10 border border-blue-500/30 p-6 rounded-2xl mb-8 text-left">
        <h3 class="text-xl font-bold text-white mb-4">Langkah Selanjutnya:</h3>
        <ol class="space-y-3 text-gray-300">
          <li class="flex items-start">
            <span class="bg-blue-500 text-white w-6 h-6 rounded-full flex items-center justify-center text-sm font-bold mr-3 flex-shrink-0 mt-0.5">1</span>
            <span>Transfer sesuai paket yang dipilih ke rekening berikut:<br>
              <strong class="text-white">BCA 1234567890 a.n. SeriusBrand</strong>
            </span>
          </li>
          <li class="flex items-start">
            <span class="bg-blue-500 text-white w-6 h-6 rounded-full flex items-center justify-center text-sm font-bold mr-3 flex-shrink-0 mt-0.5">2</span>
            <span>Upload bukti transfer di bawah ini</span>
          </li>
          <li class="flex items-start">
            <span class="bg-blue-500 text-white w-6 h-6 rounded-full flex items-center justify-center text-sm font-bold mr-3 flex-shrink-0 mt-0.5">3</span>
            <span>Tim kami akan memverifikasi dan menghubungi kamu via WhatsApp</span>
          </li>
        </ol>
      </div>

      <!-- Upload Payment Proof -->
      <div class="mb-8">
        <label class="block text-white font-semibold mb-4 text-lg">Upload Bukti Transfer</label>
        <div class="border-2 border-dashed border-slate-600 rounded-2xl p-8 hover:border-blue-500 transition-all">
          <input 
            type="file" 
            @change="handleFileChange" 
            accept="image/*"
            class="hidden"
            ref="fileInput"
          >
          <button 
            @click="$refs.fileInput.click()"
            class="bg-slate-800 hover:bg-slate-700 text-white px-6 py-3 rounded-xl font-semibold transition-all"
          >
            {{ fileName || 'Pilih File' }}
          </button>
          <p class="text-gray-400 text-sm mt-3">Format: JPG, PNG (Max. 5MB)</p>
        </div>
        
        <button 
          @click="uploadProof" 
          :disabled="!selectedFile || uploading"
          class="w-full mt-4 bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 disabled:from-gray-600 disabled:to-gray-700 text-white py-4 rounded-xl font-bold text-lg transition-all"
        >
          {{ uploading ? 'Mengupload...' : 'Upload Bukti Transfer' }}
        </button>
        
        <p v-if="uploadSuccess" class="text-green-400 mt-3 font-semibold">✓ Bukti transfer berhasil diupload!</p>
      </div>

      <!-- Back to Home -->
      <router-link 
        to="/" 
        class="inline-block bg-slate-800 hover:bg-slate-700 text-white px-8 py-3 rounded-xl font-semibold transition-all"
      >
        ← Kembali ke Beranda
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

const route = useRoute()
const orderId = route.params.id

const selectedFile = ref(null)
const fileName = ref('')
const uploading = ref(false)
const uploadSuccess = ref(false)
const fileInput = ref(null)

function handleFileChange(event) {
  const file = event.target.files[0]
  if (file) {
    selectedFile.value = file
    fileName.value = file.name
  }
}

async function uploadProof() {
  if (!selectedFile.value) return
  
  uploading.value = true
  
  try {
    const formData = new FormData()
    formData.append('payment_proof', selectedFile.value)
    
    await axios.post(`http://localhost:8080/api/orders/${orderId}/proof`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    
    uploadSuccess.value = true
    alert('Bukti transfer berhasil diupload! Tim kami akan segera memverifikasi.')
  } catch (error) {
    console.error('Upload failed:', error)
    alert('Gagal mengupload bukti transfer. Silakan coba lagi.')
  } finally {
    uploading.value = false
  }
}
</script>
