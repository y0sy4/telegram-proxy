# TG WS Proxy Go

[![Release](https://img.shields.io/github/v/release/y0sy4/tg-ws-proxy-go)](https://github.com/y0sy4/tg-ws-proxy-go/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go](https://img.shields.io/github/go-mod/go-version/y0sy4/tg-ws-proxy-go)](go.mod)

**SOCKS5-прокси для Telegram Desktop на Go.** Ускоряет Telegram через WebSocket к серверам Telegram.

---

## 📥 Скачать (v2.0.7)

### Full версия (~6.4 MB)

| Платформа | Файл |
|-----------|------|
| **Windows** (amd64) | [⬇️ TgWsProxy_windows_amd64.exe](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.7/TgWsProxy_windows_amd64.exe) |
| **Linux** (amd64) | [⬇️ TgWsProxy_linux_amd64](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.7/TgWsProxy_linux_amd64) |
| **macOS** (Intel) | [⬇️ TgWsProxy_darwin_amd64](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.7/TgWsProxy_darwin_amd64) |
| **macOS** (Apple Silicon) | [⬇️ TgWsProxy_darwin_arm64](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.7/TgWsProxy_darwin_arm64) |

### Lite версия (~5 MB) — для роутеров и серверов

| Платформа | Файл |
|-----------|------|
| **Windows** (amd64) | [⬇️ TgWsProxy_lite_windows_amd64.exe](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.7/TgWsProxy_lite_windows_amd64.exe) |
| **Linux** (amd64) | [⬇️ TgWsProxy_lite_linux_amd64](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.7/TgWsProxy_lite_linux_amd64) |
| **Linux** (ARM64) | [⬇️ TgWsProxy_lite_linux_arm64](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.7/TgWsProxy_lite_linux_arm64) |
| **macOS** (Intel) | [⬇️ TgWsProxy_lite_darwin_amd64](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.7/TgWsProxy_lite_darwin_amd64) |
| **macOS** (Apple Silicon) | [⬇️ TgWsProxy_lite_darwin_arm64](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.7/TgWsProxy_lite_darwin_arm64) |

### Какую версию выбрать?

| Версия | Размер | Функции | Для кого |
|--------|--------|---------|----------|
| **Full** | ~6.4 MB | SOCKS5, HTTP proxy, upstream proxy, конфиги, автозапуск Telegram, `--test-dc`, `--test-dc-media` | Обычные пользователи |
| **Lite** | ~5 MB | Только SOCKS5 прокси, `--test-dc` | Роутеры (OpenWRT), серверы, минималисты |

---

## 🚀 Быстрый старт

### Windows
1. Скачай [TgWsProxy_windows_amd64.exe](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.7/TgWsProxy_windows_amd64.exe)
2. Дважды кликни
3. Telegram откроет настройки прокси → нажми "Включить"

### Linux/macOS
```bash
chmod +x TgWsProxy_linux_amd64
./TgWsProxy_linux_amd64
```

**Всё!** Telegram работает через прокси.

---

## ⚙️ Опции

```bash
TgWsProxy [флаги]
```

| Флаг | Описание | По умолчанию |
|------|----------|--------------|
| `--port` | Порт SOCKS5 | 1080 |
| `--host` | Хост | 127.0.0.1 |
| `--dc-ip` | DC:IP (через запятую), `DCm:IP` для media | авто |
| `--auth` | Логин:пароль для прокси | — |
| `--http-port` | HTTP прокси (для браузеров) | 0 (выкл) |
| `--upstream-proxy` | Цепочка через другой прокси | — |
| `-v` | Подробные логи | false |
| `--test-dc` | Тест DC и выход | — |
| `--test-dc-media` | Тест DC + media/CDN | — |
| `--auto-update` | Авто-обновление | false (безопасность) |

### Примеры

**Просто запустить:**
```bash
TgWsProxy
```

**HTTP прокси для браузеров:**
```bash
TgWsProxy --http-port 8080
```

**Через Tor:**
```bash
TgWsProxy --upstream-proxy "socks5://127.0.0.1:9050"
```

**С паролем:**
```bash
TgWsProxy --auth "user:pass"
```

**Media-specific DC (разные IP для текста и медиа):**
```bash
TgWsProxy --dc-ip "2:149.154.167.220,2m:149.154.167.222,4:149.154.167.91,4m:149.154.167.118"
```

**Диагностика DC:**
```bash
# Тест text DC
TgWsProxy --test-dc

# Тест text + media DC
TgWsProxy --test-dc-media
```

---

## 🔧 Что нового в v2.0.7

| Категория | Изменения |
|-----------|-----------|
| 🌐 **Media CDN** | +13 CDN/media IP, домены pluto/venus/kws-2, media-specific `--dc-ip` |
| 🔧 **Pool** | WS возвращается в пул, blacklist TTL 10 мин, IsClosed() проверка |
| ⚡ **Производительность** | sync.Pool буферы (-80% GC), PatchInitDC in-place |
| 💓 **Стабильность** | Heartbeat ping/pong, rate limiting, graceful shutdown |
| 🧪 **Тесты** | +16 unit тест (pool, config, websocket) |
| 🐛 **Фиксы** | Close frame panic fix, module path → `github.com/y0sy4/...` |
| 📦 **Service** | systemd unit file `tg-ws-proxy.service` |

[📋 Все изменения](https://github.com/y0sy4/tg-ws-proxy-go/compare/v2.0.6...v2.0.7)

---

## 📊 Почему Go?

| | Python | Go |
|--|--------|-----|
| Размер | ~50 MB | **~6 MB** |
| Зависимости | pip | **stdlib** |
| Запуск | ~500 ms | **~50 ms** |
| Память | ~50 MB | **~10 MB** |

---

## 🗂️ Структура

```
tg-ws-proxy-go/
├── cmd/proxy/          # Full CLI
├── cmd/lite/           # Lite CLI (minimal)
├── internal/
│   ├── proxy/          # Ядро прокси
│   ├── socks5/         # SOCKS5 сервер
│   ├── websocket/      # WebSocket клиент
│   ├── mtproto/        # MTProto парсинг
│   ├── pool/           # WebSocket pooling
│   ├── config/         # Конфигурация
│   └── telegram/       # Авто-настройка Telegram
├── mobile/             # Android/iOS bindings
├── tg-ws-proxy.service # systemd unit
├── go.mod
└── Makefile
```

---

## 🛠️ Сборка

```bash
# Все платформы
make all

# Или вручную
go build -o TgWsProxy.exe ./cmd/proxy                     # Windows
GOOS=linux GOARCH=amd64 go build -o TgWsProxy_linux ./cmd/proxy
GOOS=darwin GOARCH=amd64 go build -o TgWsProxy_macos ./cmd/proxy
GOOS=darwin GOARCH=arm64 go build -o TgWsProxy_macos_arm64 ./cmd/proxy
```

---

## 📡 OpenWRT / Роутеры

### Быстрый старт

```bash
# Кросс-компиляция (на ПК)
GOOS=linux GOARCH=arm64 go build -o tg-ws-proxy-go ./cmd/lite/

# Копирование на роутер
scp tg-ws-proxy-go root@192.168.1.1:/usr/bin/
chmod +x /usr/bin/tg-ws-proxy-go

# Запуск
/usr/bin/tg-ws-proxy-go --host 0.0.0.0 --port 1080
```

### Media не грузится?

```bash
# Тест media DC
/usr/bin/tg-ws-proxy-go --test-dc-media

# Запуск с конкретными DC
/usr/bin/tg-ws-proxy-go --dc-ip "2:149.154.167.220,4:149.154.167.220" --host 0.0.0.0 --port 1080
```

### systemd (Linux сервер)

```bash
sudo cp tg-ws-proxy.service /etc/systemd/system/
sudo systemctl enable tg-ws-proxy
sudo systemctl start tg-ws-proxy
```

---

## 🔍 Решение проблем

| Проблема | Решение |
|----------|---------|
| Прокси не подключается | Проверь что запущен, Telegram на `127.0.0.1:1080` |
| Текст грузит, медиа нет | Используй `--test-dc-media`, затем `--dc-ip` с рабочими IP |
| Telegram не открывается | Вручную: `tg://socks?server=127.0.0.1&port=1080` |
| Антивирус блокирует | Ложное срабатывание. Код открытый, добавь в исключения |
| Логи | `%APPDATA%\TgWsProxy\proxy.log` (Win), `~/.config/TgWsProxy/proxy.log` (Linux) |

---

## 🤝 Contributing

1. Fork → branch → PR
2. `go test ./...`
3. `gofmt -w .`
4. Без эмоций. По делу.

---

## 📄 License

MIT License

---

**v2.0.7** | Built with ❤️ using Go 1.21
