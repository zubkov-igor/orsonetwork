#!/bin/bash
set -e

DEV="wlp5s0"

# Проверка существования устройства
if ! ip link show "$DEV" >/dev/null 2>&1; then
  echo "Ошибка: интерфейс $DEV не найден." >&2
  exit 1
fi

echo "Текущее состояние $DEV:"
iw dev "$DEV" info

echo "Возвращаем $DEV в режим managed..."
ip link set "$DEV" down
iw dev "$DEV" set type managed
ip link set "$DEV" up
systemctl restart NetworkManager

echo "Финальное состояние:"
iw dev "$DEV" info | grep -i "type"

echo "Готово: $DEV теперь в режиме managed, NetworkManager перезапущен."

