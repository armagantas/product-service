# Product Service

Bu servis, e-ticaret mikroservis mimarisi içerisinde ürün taleplerinin yönetiminden sorumludur.

## Teknik Detaylar

- Go programlama dili kullanılmıştır
- Web framework olarak Fiber kullanılmıştır
- Veritabanı olarak PostgreSQL kullanılmıştır
- ORM için GORM kütüphanesi kullanılmıştır

## Kurulum

```bash
# Bağımlılıkları yükleyin
go mod download

# PostgreSQL veritabanını başlatın
docker-compose up -d postgres

# Servisi çalıştırın
go run main.go
```

## API Endpointleri

### Ürün Endpointleri

- `GET /api/products` - Tüm ürün taleplerini listele
- `POST /api/products` - Yeni ürün talebi oluştur
- `GET /api/products/:id` - Tek bir ürün talebinin detaylarını getir
- `PUT /api/products/:id` - Bir ürün talebini güncelle
- `DELETE /api/products/:id` - Bir ürün talebini sil

### Kategori Endpointleri

- `GET /api/categories` - Tüm kategorileri listele
- `POST /api/categories` - Yeni kategori oluştur
- `GET /api/categories/:id` - Tek bir kategorinin detaylarını getir
- `PUT /api/categories/:id` - Bir kategoriyi güncelle
- `DELETE /api/categories/:id` - Bir kategoriyi sil
- `GET /api/categories/:id/products` - Bir kategorideki tüm ürün taleplerini listele

## Veritabanı Şeması

### Category Tablosu

- `id` - Kategori ID (Primary Key)
- `name` - Kategori adı (unique)
- `description` - Kategori açıklaması
- `parent_id` - Üst kategorinin ID'si (hiyerarşik yapı için)
- `is_active` - Kategori aktif mi?
- `created_at` - Oluşturulma tarihi
- `updated_at` - Güncellenme tarihi
- `deleted_at` - Silinme tarihi (soft delete için)

### Product Tablosu

- `id` - Ürün talebi ID (Primary Key)
- `title` - Ürün talebi başlığı
- `description` - Ürün talebi açıklaması
- `category_id` - Kategori ID (Foreign Key)
- `user_id` - Kullanıcı ID
- `quantity` - Talep edilen miktar
- `price` - Fiyat (opsiyonel)
- `min_price` - Minimum fiyat (opsiyonel)
- `max_price` - Maksimum fiyat (opsiyonel)
- `currency` - Para birimi
- `status` - Durum (active, closed, expired)
- `delivery_address` - Teslimat adresi
- `delivery_date` - Teslimat tarihi
- `tags` - Etiketler
- `is_active` - Ürün talebi aktif mi?
- `created_at` - Oluşturulma tarihi
- `updated_at` - Güncellenme tarihi
- `deleted_at` - Silinme tarihi (soft delete için)
- `expires_at` - Son geçerlilik tarihi

## Proje Yapısı

```
product-service/
├── main.go           # Uygulama giriş noktası
├── go.mod           # Go modül dosyası
├── go.sum           # Bağımlılık kontrolü için hash dosyası
├── config/          # Konfigürasyon dosyaları
├── handlers/        # HTTP isteklerini işleyen handler fonksiyonları
├── models/          # Veri modelleri
└── internal/        # İç servis ve paketler
```
