# Gunakan base image Alpine dengan Go yang telah di-compile
FROM golang:1.22rc1-alpine3.19 AS builder

# Set jalur kerja (working directory) ke direktori aplikasi
WORKDIR /app

# Copy kode sumber ke dalam kontainer
COPY . .
COPY .env .env
# Unduh dependensi Go dan build aplikasi menjadi binary
RUN go mod tidy
RUN go build -o app-binary

# Gunakan base image yang lebih ringan untuk menjalankan aplikasi
FROM alpine:3.19

# Set jalur kerja
WORKDIR /app

# Copy binary dari kontainer build
COPY --from=builder /app/app-binary .

# Copy file konfigurasi atau environment
COPY .env.example .env

# Menjalankan aplikasi binary
CMD ["./app-binary"]
