# SeriusBrand - Platform Branding untuk UMKM

Platform untuk menyediakan branding package (Video + Landing Page) bagi UMKM Indonesia dengan harga terjangkau mulai dari Rp10.000.

## Struktur Project

```
BrandUMKM/
├── backend/           # Go backend dengan Gin framework
│   ├── cmd/server/    # Entry point aplikasi
│   ├── internal/      # Business logic
│   │   ├── db/       # Database connection & migration
│   │   ├── handlers/ # API handlers
│   │   └── models/   # Data models
│   ├── uploads/      # File upload storage
│   └── .env          # Environment variables
└── frontend/          # Vue.js frontend
    ├── src/
    │   ├── api/      # Axios API client
    │   ├── router/   # Vue Router
    │   ├── views/    # Page components
    │   └── components/ # Reusable components
    └── package.json
```

## Tech Stack

### Backend
- **Language**: Go (Golang)
- **Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL
- **File Upload**: Multipart form handling

### Frontend
- **Framework**: Vue 3 (Composition API)
- **Styling**: Tailwind CSS
- **Router**: Vue Router
- **HTTP Client**: Axios
- **Build Tool**: Vite

## Setup & Installation

### Backend

1. Install PostgreSQL dan buat database:
```sql
CREATE DATABASE seriusbrand;
```

2. Setup environment variables di `backend/.env`:
```env
PORT=8080
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=seriusbrand
DB_PORT=5432
```

3. Install dependencies:
```bash
cd backend
go mod download
```

4. Run backend:
```bash
go run cmd/server/main.go
```

Backend akan berjalan di `http://localhost:8080`

### Frontend

1. Install dependencies:
```bash
cd frontend
npm install
```

2. Run development server:
```bash
npm run dev
```

Frontend akan berjalan di `http://localhost:5173`

## API Endpoints

### Public Endpoints
- `POST /api/orders` - Buat pesanan baru
- `POST /api/orders/:id/proof` - Upload bukti transfer
- `GET /api/umkm/:slug` - Get halaman UMKM berdasarkan slug

### Admin Endpoints
- `GET /api/admin/orders` - List semua pesanan
- `PATCH /api/admin/orders/:id` - Update status pesanan
- `POST /api/admin/umkm` - Buat halaman UMKM baru

## Paket Harga

### Starter - Rp10.000
- 1 video packing (15-30 detik)
- 1 landing page produk
- Foto produk dari video
- Tombol WhatsApp & Marketplace
- Hosting 30 hari

### Growth - Rp25.000
- Semua di Starter
- Video 30-60 detik
- Copywriting deskripsi produk
- Tombol Order + Upload Bukti
- Hosting 90 hari

### Pro - Rp50.000
- Semua di Growth
- 3 variasi video (Reels/TikTok/Story)
- Custom subdomain
- Statistik pengunjung
- Hosting 6 bulan

## Flow Pemesanan

1. **Customer** pilih paket di landing page
2. Isi form (Nama UMKM, WhatsApp, Nama Produk, Paket)
3. Submit order → redirect ke halaman sukses
4. Upload bukti transfer
5. **Admin** verifikasi pembayaran
6. **Admin** buat video & landing page
7. Landing page aktif di `/umkm/{slug}`
8. Customer dapat share link ke pembeli

## Database Schema

### Table: orders
- `id` (PK)
- `customer_name`
- `whatsapp`
- `product_name`
- `package_type` (starter/growth/pro)
- `price`
- `status` (pending/paid/completed)
- `payment_proof_url`
- `created_at`, `updated_at`

### Table: umkm_pages
- `id` (PK)
- `order_id` (FK)
- `slug` (unique)
- `video_url`
- `photos` (jsonb)
- `description`
- `price`
- `whatsapp_link`
- `marketplace_link`
- `created_at`, `updated_at`

## Development Notes

- Backend menggunakan CORS untuk allow frontend di localhost:5173
- File uploads disimpan di `backend/uploads/`
- Frontend menggunakan Tailwind dengan custom color palette
- Design menggunakan glassmorphism & gradient untuk tampilan premium
- Semua animasi menggunakan Tailwind transition utilities

## Production Deployment

### Backend
1. Build binary: `go build -o server cmd/server/main.go`
2. Set production environment variables
3. Run: `./server`

### Frontend
1. Build: `npm run build`
2. Deploy folder `dist/` ke static hosting atau serve dengan Nginx

## License

Private project untuk SeriusBrand.
