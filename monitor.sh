#!/bin/bash
set -e

DEV="wlp5s0"

if ! ip link show "$DEV" >/dev/null 2>&1; then
  echo "Ошибка: интерфейс $DEV не найден." >&2
  exit 1
fi

echo "Текущее состояние $DEV:"
iw dev "$DEV" info

echo "Убиваем мешающие процессы (NetworkManager, wpa_supplicant и др.)..."
airmon-ng check kill

echo "Переводим $DEV в режим monitor..."
ip link set "$DEV" down
iw dev "$DEV" set type monitor
ip link set "$DEV" up

echo "Финальное состояние:"
iw dev "$DEV" info | grep -i "type"

echo "Готово: $DEV теперь в режиме monitor."
echo "Можно запускать: sudo airodump-ng -w capture/capture $DEV"



