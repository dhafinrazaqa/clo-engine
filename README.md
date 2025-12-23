# CLO Engine - Container Load Optimization (Golang)

Repository ini berisi **engine utama Container Load Optimization (CLO)** yang
diimplementasikan menggunakan bahasa **Golang**

Engine ini bertugas menjalankan algoritma penataan barang ke dalam kontainer
dan menghasilkan posisi penempatan serta metrik pemanfaatan ruang

## Tujuan Engine

CLO Engine dikembangkan untuk:

- Mensimulasikan penataan barang 3D berbentuk balok ke dalam kontainer
- Menghasilkan koordinat posisi barang (x, y, z)
- Menghitung pemanfaatan volume kontainer
- Menjadi komponen logika terpisah dari API (stateless)

## Algoritma Packing

Engine menggunakan algoritma **First Fit Decreasing 3D (FFD-3D)**
dengan pendekatan **Shelf-based Packing**

Langkah umum algoritma:

1. Barang dengan jumlah lebih dari sati diperluas menjadi instance masing-masing
2. Setiap instance dihitung volumenya
3. Barang diurutkan berdasarkan volume (descending)
4. Barang ditempatkan sepanjang sumbu X (panjang kontainer)
5. Jika tidak muat pada shelf yang ada, dibuat shelf baru secara vertikal (Z)

## Sistem Koordinat

Engine menggunakan sistem koordinat kartesian:

- X -> Panjang kontainer (arah masuk barang)
- Y -> Lebar kontainer
- Z -> Tinggi kontainer

Semua koordinat yang dihasilkan bersifat absolut terhadap titik (0,0,0)
di sudut bawah-depan-kiri kontainer

## Struktur Folder

```csharp
clo-engine/
├── cmd/
│ └── main.go # Entry point CLI
├── internal/
│ ├── algorithm/ # Implementasi FFD-3D dan shelf
│ ├── models/ # Struktur data (container, item, shelf, result)
│ ├── util/ # Fungsi utilitas (volume, fit checking)
├── go.mod
├── go.sum
└── README.md
```

## Cara Build

Pastikan Go sudah terinstall.

```bash
go version
```

Build engine:

```bash
go build -o clo-engine.exe ./cmd
```

### Menjalankan Engine

```bash
clo-engine.exe input.json output.json
```

Dengan mode debug:

```bash
clo-engine.exe input.json output.json --debug
```

### Mode Debug

Engine menyediakan flag `--debug` untuk menampilkan log proses algoritma.

Log yang ditampilkan meliputi:

- Urutan penempatan barang
- Percobaan orientasi barang
- Proses pengecekan shelf
- Pembuatan shelf baru
- Penempatan akhir barang

## Format Input

Engine menerima input JSON berisi:

- Dimensi kontainer
- Daftar barang (dengan kuantitas)

Detail format input dijelaskan pada dokumentasi API / umbrella repository

## Format Output

Output berupa JSON yang berisi:

- Status eksekusi
- Metrik pemanfaatan kontainer
- Daftar penempatan barang
- Daftar barang yang tidak dapat ditempatkan (jika ada)

## Determinisme & Batasan

- Engine bersifat deterministik
- Input yang sama akan menghasilkan output yang sama
- Algoritma bersifat heuristik
- Tidak menjamin solusi optimal global
- Tidak mempertimbangkan constraint fisik dunia nyata

## Konteks Penggunaan

Engine ini dikembangkan sebagai:

- Proof-of-concept akademik
- Eksplorasi algoritma heuristik
- Pendukung penyusunan laporan PA

Engine ini tidak digunakan untuk pemuatan fisik kontainer secara operasional.

## Penutup

CLO Engine dirancang agar:

- Mudah diuji
- Mudah dipahami
- Mudah dikembangkan lebih lanjut

Untuk penggunaan melalui HTTP API bisa dilihat di repository CLO API (Laravel).
