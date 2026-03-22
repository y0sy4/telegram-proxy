## 🚀 TG WS Proxy Go v2.0.2

> **Go-переосмысление** [Flowseal/tg-ws-proxy](https://github.com/Flowseal/tg-ws-proxy)  
> Локальный SOCKS5-прокси для Telegram Desktop на Go

---

### ✨ Что нового в v2.0.2

| Функция | Статус |
|---------|--------|
| 🔗 **tg://socks ссылки** | ✅ Работают на Windows |
| 📲 **Авто-конфигурация Telegram** | ✅ Открывает настройки прокси |
| 🔄 **Автообновление** | ✅ Скачивает новую версию |
| 🌐 **IPv6 поддержка** | ✅ Через NAT64 |
| 🔐 **SOCKS5 аутентификация** | ✅ --auth username:password |
| 🛑 **Авто-закрытие дубликатов** | ✅ Завершает старые экземпляры |

---

### 🔧 Исправления v2.0.2

- ✅ **Исправлено:** При запуске второго экземпляра первый автоматически закрывается
- ✅ **Улучшено:** Стабильность работы на Windows
- ✅ **Добавлено:** Проверка и завершение дублирующихся процессов

---

### 📥 Скачать

| Платформа | Файл | Размер | Ссылка |
|-----------|------|--------|--------|
| **Windows x64** | TgWsProxy.exe | 6.6 MB | [⬇️ Скачать](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.2/TgWsProxy_windows_amd64.exe) |
| **Linux x64** | TgWsProxy | 6.5 MB | [⬇️ Скачать](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.2/TgWsProxy_linux_amd64) |
| **macOS Intel** | TgWsProxy | 6.6 MB | [⬇️ Скачать](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.2/TgWsProxy_darwin_amd64) |
| **macOS Apple Silicon** | TgWsProxy | 5.8 MB | [⬇️ Скачать](https://github.com/y0sy4/tg-ws-proxy-go/releases/download/v2.0.2/TgWsProxy_darwin_arm64) |

---

### 🚀 Быстрый старт

**Windows:**
1. Скачай `TgWsProxy.exe`
2. Запусти
3. Telegram автоматически откроет настройки прокси
4. Подтверди добавление

**Linux/macOS:**
```bash
chmod +x TgWsProxy_*
./TgWsProxy_linux_amd64  # или ./TgWsProxy_darwin_amd64
```

---

### 🔧 Командная строка

```bash
./TgWsProxy [опции]

--port int        Порт SOCKS5 (default 1080)
--host string     Хост SOCKS5 (default "127.0.0.1")
--dc-ip string    DC:IP через запятую
--auth string     SOCKS5 аутентификация (username:password)
-v                Подробное логирование
--version         Показать версию
```

---

### 📊 Сравнение с Python

| Метрика | Python | Go |
|---------|--------|-----|
| Размер | ~50 MB | **~6 MB** ⚡ |
| Зависимости | pip | **stdlib** ⚡ |
| Запуск | ~500 ms | **~50 ms** ⚡ |
| Память | ~50 MB | **~10 MB** ⚡ |

---

### 🔗 Ссылки

- 📦 **Релизы:** https://github.com/y0sy4/tg-ws-proxy-go/releases
- 📖 **Документация:** https://github.com/y0sy4/tg-ws-proxy-go#readme
- ❓ **FAQ:** https://github.com/y0sy4/tg-ws-proxy-go/blob/master/FAQ.md
- 🐛 **Баги:** https://github.com/y0sy4/tg-ws-proxy-go/issues

---

**Built with ❤️ using Go 1.21** | **License:** MIT
