# QR-CLI

**QR-CLI** — это простое консольное приложение для генерации QR-кодов, которое принимает текст как аргумент и выводит QR-код в виде ASCII символов в консоль.

## Возможности

- Генерация QR-кода на основе текста, переданного в качестве аргумента.
- Вывод QR-кода в текстовом виде в консоль.
  
## Установка

### Требования

- Установленный Go версии 1.16 или выше.
  
### Шаги установки

1. Убедитесь, что Go установлен на вашем компьютере:

   ```bash
   go version
   ```

2. Склонируйте репозиторий проекта (или скачайте его в нужную папку).

   ```bash
   git clone https://github.com/<ваш_пользователь>/qr-cli.git
   cd qr-cli
   ```

3. Установите необходимые зависимости и соберите проект:

   ```bash
   go mod tidy
   go build rsc/* 
   mv main qr-cli
   ```

4. Переместите скомпилированный файл в системную директорию для того, чтобы приложение было доступно глобально:

   ```bash
   sudo cp qr-cli /usr/local/bin/
   ```

Теперь приложение можно запускать командой `qr-cli`.

## Использование

Чтобы сгенерировать QR-код для конкретного текста, выполните команду:

```bash
qr-cli "Your text here"
```

Пример:

```bash
qr-cli "Hello, world!"
```

Это выведет QR-код в виде ASCII символов в консоль.

### Пример вывода

```bash
██████████████████████████████████████████████████████████
██████████████████████████████████████████████████████████
██████████████████████████████████████████████████████████
██████████████████████████████████████████████████████████
████████              ████  ██    ██              ████████
████████  ██████████  ██████  ██████  ██████████  ████████
████████  ██      ██  ██        ████  ██      ██  ████████
████████  ██      ██  ██      ██  ██  ██      ██  ████████
████████  ██      ██  ██  ██  ██  ██  ██      ██  ████████
████████  ██████████  ██  ████  ████  ██████████  ████████
████████              ██  ██  ██  ██              ████████
████████████████████████  ██  ████████████████████████████
████████  ██          ████  ██  ████          ████████████
██████████    ██    ██  ██  ██                ██  ████████
████████  ██  ██        ██    ██      ████      ██████████
████████  ██  ████  ██████  ██      ████      ████████████
██████████████  ██          ████      ██████████  ████████
████████████████████████  ██  ██  ██████    ████  ████████
████████              ████████  ████  ██████    ██████████
████████  ██████████  ██  ████████  ██  ██        ████████
████████  ██      ██  ██  ████  ████    ████████  ████████
████████  ██      ██  ██    ████            ██████████████
████████  ██      ██  ██    ████  ████  ████  ████████████
████████  ██████████  ████    ██    ████      ████████████
████████              ██    ██    ██  ██  ████  ██████████
██████████████████████████████████████████████████████████
██████████████████████████████████████████████████████████
██████████████████████████████████████████████████████████
██████████████████████████████████████████████████████████
```

## Аргументы

| Аргумент     | Описание                        |
|--------------|---------------------------------|
| `<текст>`    | Текст для генерации QR-кода     |

Если текст не передан, программа покажет сообщение об ошибке и завершится.

## Зависимости

Приложение использует стороннюю библиотеку для генерации QR-кодов:

- [go-qrcode](https://github.com/skip2/go-qrcode)

Чтобы установить зависимости, используйте команду:

```bash
go mod tidy
```

## Лицензия

Этот проект распространяется под лицензией GNU GPL v3. Подробности см. в файле [LICENSE](./LICENSE).
