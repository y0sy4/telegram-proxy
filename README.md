# TG WS Proxy Go

[![Release](https://img.shields.io/github/v/release/y0sy4/tg-ws-proxy-go)](https://github.com/y0sy4/tg-ws-proxy-go/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**SOCKS5-прокси для Telegram Desktop на Go.** Ускоряет Telegram через WebSocket к серверам Telegram.

---

## 📥 Скачать (v2.0.6)

| Версия | Windows | Linux | macOS |
|--------|---------|-------|-------|
| **Full** | [⬇️ .exe](.../v2.0.6/TgWsProxy.exe) (9.3 MB) | [⬇️ amd64](.../v2.0.6/TgWsProxy_linux_amd64) (8.9 MB) | [⬇️ Intel](.../v2.0.6/TgWsProxy_darwin_amd64) / [⬇️ ARM](.../v2.0.6/TgWsProxy_darwin_arm64) |
| **Lite** | [⬇️ .exe](.../v2.0.6/TgWsProxy_lite.exe) (5.5 MB) | [⬇️ amd64](.../v2.0.6/TgWsProxy_lite_linux) (5.3 MB) | [⬇️ Intel](.../v2.0.6/TgWsProxy_lite_darwin) / [⬇️ ARM](.../v2.0.6/TgWsProxy_lite_darwin_arm) |

**Какую версию выбрать?**

| Версия | Размер | Функции | Для кого |
|--------|--------|---------|----------|
| **Full** | ~9 MB | Авто-обновление, HTTP proxy, upstream proxy, конфиги, автозапуск Telegram | Обычные пользователи |
| **Lite** | ~5.3 MB | Только SOCKS5 прокси, `--test-dc` | Роутеры (OpenWRT), серверы, минималисты |

---

## 🚀 Быстрый старт

### Windows
1. Скачай `TgWsProxy_windows_amd64.exe`
2. Дважды кликни
3. Telegram откроет настройки прокси → нажми "Включить"

### Linux/macOS
```bash
chmod +x TgWsProxy_*
./TgWsProxy_linux_amd64  # или TgWsProxy_darwin_amd64
```

**Всё!** Telegram работает через прокси.

---

## ⚙️ Опции (для профи)

```bash
TgWsProxy.exe [флаги]
```

| Флаг | Описание | По умолчанию |
|------|----------|--------------|
| `--port` | Порт SOCKS5 | 1080 |
| `--host` | Хост | 127.0.0.1 |
| `--dc-ip` | DC:IP (через запятую) | авто |
| `--auth` | Логин:пароль для прокси | — |
| `--http-port` | HTTP прокси (для браузеров) | 0 (выкл) |
| `--upstream-proxy` | Цепочка через другой прокси | — |
| `-v` | Подробные логи | false |
| `--test-dc` | Тест DC и выход | — |
| `--auto-update` | Авто-обновление | false (безопасность) |

### Примеры

**Просто запустить:**
```bash
TgWsProxy.exe
```

**HTTP прокси для браузеров (порт 8080):**
```bash
TgWsProxy.exe --http-port 8080
```
Теперь браузер можно настроить на `127.0.0.1:8080`.

**Через другой прокси (Tor, SSH):**
```bash
TgWsProxy.exe --upstream-proxy "socks5://127.0.0.1:9050"
```

**С паролем:**
```bash
TgWsProxy.exe --auth "user:pass"
```

**Диагностика DC (перед запуском):**
```bash
# Проверить доступность Telegram DC
TgWsProxy.exe --test-dc

# С verbose логами
TgWsProxy.exe --test-dc -v

# Тест конкретных DC
TgWsProxy.exe --test-dc --dc-ip "2:149.154.167.220,4:149.154.167.220"
```

**Авто-обновление (только если доверяете источнику):**
```bash
TgWsProxy.exe --auto-update
```
⚠️ **Внимание:** Авто-обновление отключено по умолчанию из соображений безопасности. Бинарники скачиваются без проверки GPG-подписи.

---

## 🔧 Что нового в v2.0.6

- 🛡️ **Безопасность:** авто-обновление отключено по умолчанию
- 🔍 **--test-dc:** диагностика доступности Telegram DC
- 📡 **OpenWRT:** полная поддержка роутеров + документация
- 🐛 **Issue #5:** решение проблемы с медиа на роутерах
- 📝 **README:** секция для OpenWRT с примерами

[📖 Полные изменения](RELEASE_NOTES_v2.0.6.md)

---

## 🔧 Что было в v2.0.5

- ⚡ **atomic.Int64** для статистики — 0 блокировок
- 🧹 **stdlib вместо велосипедов** — -100 строк
- 🚀 **оптимизация аллокаций** — MTProto быстрее на 50%
- 📱 **Android/iOS** — все оптимизации совместимы

[📖 Полные изменения](RELEASE_NOTES_v2.0.5.md)

---

## 📊 Почему Go?

| | Python | Go |
|--|--------|-----|
| Размер | ~50 MB | **~8 MB** |
| Зависимости | pip | **stdlib** |
| Запуск | ~500 ms | **~50 ms** |
| Память | ~50 MB | **~10 MB** |

---

## 🗂️ Структура

```
tg-ws-proxy-go/
├── cmd/proxy/          # CLI приложение
├── internal/
│   ├── proxy/          # Ядро прокси
│   ├── socks5/         # SOCKS5 сервер
│   ├── websocket/      # WebSocket клиент
│   ├── mtproto/        # MTProto парсинг
│   ├── pool/           # WebSocket pooling
│   ├── config/         # Конфигурация
│   └── telegram/       # Авто-настройка Telegram
├── mobile/             # Android/iOS bindings
├── go.mod
├── Makefile
└── README.md
```

---

## 🛠️ Сборка

```bash
# Windows
go build -o TgWsProxy.exe ./cmd/proxy

# Linux
GOOS=linux GOARCH=amd64 go build -o TgWsProxy_linux ./cmd/proxy

# macOS
GOOS=darwin GOARCH=amd64 go build -o TgWsProxy_macos_amd64 ./cmd/proxy
GOOS=darwin GOARCH=arm64 go build -o TgWsProxy_macos_arm64 ./cmd/proxy

# Все платформы
make all
```

---

## 📱 Android/iOS

```bash
# AAR библиотека
gomobile bind -target android -o android/tgwsproxy.aar ./mobile
```

Все оптимизации совместимы с gomobile (Go 1.21+).

---

## 📡 OpenWRT / Роутеры

### Проблема: текст работает, медиа нет

**Причина:** Telegram использует разные DC для текста и медиа. На роутерах некоторые DC могут быть недоступны.

**Решение 1: Использовать конкретный DC IP**

```bash
# Протестировать DC
./tg-ws-proxy-go --test-dc

# Запустить с рабочими DC (пример для Issue #5)
./tg-ws-proxy-go --dc-ip "2:149.154.167.220,4:149.154.167.220" --host 0.0.0.0 --port 1080
```

**Решение 2: Для веб-версии Telegram**

Добавьте в `/etc/hosts` на роутере:
```
149.154.167.220 web.telegram.org telegram.org t.me telesco.pe
```

**Решение 3: Upstream proxy (для обхода блокировок)**

```bash
# Через Tor
./tg-ws-proxy-go --upstream-proxy "socks5://127.0.0.1:9050"

# Через SSH туннель
./tg-ws-proxy-go --upstream-proxy "socks5://user:pass@proxy-server:port"
```

### Сборка для OpenWRT (ARM64)

```bash
# На ПК (кросс-компиляция)
GOOS=linux GOARCH=arm64 go build -o tg-ws-proxy-go ./cmd/proxy

# Копирование на роутер
scp tg-ws-proxy-go root@192.168.1.1:/usr/bin/
chmod +x /usr/bin/tg-ws-proxy-go

# Запуск
/usr/bin/tg-ws-proxy-go --host 0.0.0.0 --port 1080
```

### Автозапуск (init.d скрипт)

```bash
cat > /etc/init.d/tg-ws-proxy << 'EOF'
#!/bin/sh /etc/rc.common

START=99

start() {
    /usr/bin/tg-ws-proxy-go --dc-ip "2:149.154.167.220,4:149.154.167.220" --host 0.0.0.0 --port 1080 &
}

stop() {
    killall tg-ws-proxy-go
}
EOF

chmod +x /etc/init.d/tg-ws-proxy
/etc/init.d/tg-ws-proxy enable
/etc/init.d/tg-ws-proxy start
```

---

## 🔍 Решение проблем

**Прокси не подключается:**
1. Проверь, запущен ли `TgWsProxy.exe`
2. Убедись, Telegram настроен на `127.0.0.1:1080`
3. Проверь логи: `%APPDATA%\TgWsProxy\proxy.log`

**Telegram не открывается:**
Открой вручную: `tg://socks?server=127.0.0.1&port=1080`

**Антивирус блокирует:**
Ложное срабатывание. Добавь в исключения. Код открытый.

---

## 📖 Документация

- [❓ FAQ](FAQ.md) — частые вопросы
- [📝 Release Notes](RELEASE_NOTES_v2.0.5.md) — изменения v2.0.5
- [👨‍💻 QWEN.md](QWEN.md) — guidelines для разработчиков

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

**v2.0.5** | Built with ❤️ using Go 1.21
