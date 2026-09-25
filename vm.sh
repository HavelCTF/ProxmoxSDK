#!/usr/bin/env bash
# Proxmox VE QEMU VM.
#   ./vm.sh setup     download the ISO (resumable), verify it, create the disk
#   ./vm.sh install   boot the ISO to install Proxmox onto disk.qcow2
#   ./vm.sh           boot the installed system from disk.qcow2
set -euo pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

PVE_VERSION=9.2-1
ISO="$DIR/proxmox-ve_${PVE_VERSION}.iso"
ISO_URL="https://enterprise.proxmox.com/iso/proxmox-ve_${PVE_VERSION}.iso"
SUMS_URL="https://enterprise.proxmox.com/iso/SHA256SUMS"
SUMS="$DIR/SHA256SUMS"

DISK="$DIR/disk.qcow2"
DISK_SIZE=50G
CPUS=2
MEM=4G
SSH_PORT=2222   # host -> guest 22
GUI_PORT=8006   # host -> guest 8006 (Proxmox web UI)

setup() {
  echo "==> Fetch checksums"
  curl -fsSL -o "$SUMS" "$SUMS_URL"

  if [[ -f "$ISO" ]] && (cd "$DIR" && sha256sum --ignore-missing -c "$SUMS" 2>/dev/null | grep -q "$(basename "$ISO"): OK"); then
    echo "==> ISO already present and verified"
  else
    echo "==> Downloading $(basename "$ISO")"
    curl -fL -C - -o "$ISO" "$ISO_URL"
    echo "==> Verifying ISO"
    (cd "$DIR" && sha256sum --ignore-missing -c "$SUMS" | grep "$(basename "$ISO")")
  fi

  if [[ -f "$DISK" ]]; then
    echo "==> Disk already exists: $DISK ($(qemu-img info "$DISK" | awk '/^virtual size:/ {print $3, $4}'))"
  else
    echo "==> Creating $DISK_SIZE disk"
    qemu-img create -f qcow2 "$DISK" "$DISK_SIZE"
  fi

  echo "==> Setup done, now run: ./vm.sh install"
}

boot() {
  [[ -f "$DISK" ]] || { echo "no disk.qcow2 - run ./vm.sh setup first" >&2; exit 1; }

  args=(
    -name proxmox
    -machine q35,accel=kvm
    -cpu host                        # exposes vmx/svm so Proxmox can nest guests
    -smp "$CPUS"
    -m "$MEM"
    -drive "file=$DISK,if=virtio,format=qcow2,cache=writeback,discard=unmap"
    -netdev "user,id=net0,hostfwd=tcp::${SSH_PORT}-:22,hostfwd=tcp::${GUI_PORT}-:8006"
    -device virtio-net-pci,netdev=net0
    -device virtio-balloon
    -object rng-random,filename=/dev/urandom,id=rng0
    -device virtio-rng-pci,rng=rng0
    -vga std
    -display gtk
    -usb -device usb-tablet
  )

  if [[ "${1:-}" == "install" ]]; then
    [[ -f "$ISO" ]] || { echo "no ISO - run ./vm.sh setup first" >&2; exit 1; }
    args+=( -cdrom "$ISO" -boot d )
  fi

  exec qemu-system-x86_64 "${args[@]}"
}

case "${1:-run}" in
  setup)   setup ;;
  install) boot install ;;
  run)     boot ;;
  *)       echo "usage: $0 [setup|install|run]" >&2; exit 1 ;;
esac