#!/bin/bash

set -euo pipefail

NEW_MODULE="$1"

if [ -z "$NEW_MODULE" ]; then
  echo "❌ NEW_MODULE belum diberikan."
  exit 1
fi

# Pastikan file go.mod ada
if [ ! -f go.mod ]; then
  echo "❌ File go.mod tidak ditemukan di direktori saat ini."
  exit 1
fi

# Baca module lama dari go.mod
OLD_MODULE=$(grep '^module ' go.mod | awk '{print $2}')

if [ -z "$OLD_MODULE" ]; then
  echo "❌ Gagal membaca module lama dari go.mod"
  exit 1
fi

echo "📦 Modul lama: $OLD_MODULE"
echo "📦 Modul baru: $NEW_MODULE"

# Update module di go.mod
go mod edit -module "$NEW_MODULE"

# Ganti semua import path dalam file apapun (bukan hanya .go)
echo "🔍 Mengganti semua import path di semua file..."
# Cari semua file reguler, kecuali direktori .git dan file biner
find . -type f ! -path "./.git/*" -print0 | \
  xargs -0 sed -i.bak "s|$OLD_MODULE|$NEW_MODULE|g"

# Hapus file backup dari sed
find . -type f -name '*.bak' -delete

echo "✅ Selesai mengganti module dan semua import path."
