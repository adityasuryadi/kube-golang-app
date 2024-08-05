#!/usr/bin/env sh

# build image dari podmanfile dengan nama item-app dan tag v1

echo "build podman"
podman build --platform=linux/amd64 -t golang-app:v2 . 

echo "============================================================"

# lihat semua image yang ada di local
echo "list podman image"
podman images

echo "============================================================"
# mengubah nama image yang sudah di buat ke nama image sesuai dengan ketentuan nama yang akan di push ke repository 
echo "rename podman image"
podman tag golang-app:v2 ghcr.io/adityasuryadi/golang-app:v2
echo "============================================================"
# login ke githu container menggunakan personal access token yang sudah di set sebelum nya
echo "login github"

echo $TOKEN_GITHUB | podman login ghcr.io -u adityasuryadi --password-stdin
echo "============================================================"
# push image yang sudah di rename ke github container
echo "push ke github container registry"
podman push ghcr.io/adityasuryadi/golang-app:v2
echo "============================================================"