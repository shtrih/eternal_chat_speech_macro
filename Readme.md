# Eternal Chat Speech Macro

Макрос для голосовых команд чата для [tes3mp](https://github.com/TES3MP) сервера https://morrowindonline.ru/.

В данный момент работает только на Linux. Нужно запускать с `sudo` (требование от https://github.com/MarinX/keylogger#notes).

### Как это работает?

Программа симулирует нажатия клавиатуры: вводит в чат команду `/speech <type> <index>` и жмет `ENTER`. 
Никаких проверок, открыт ли чат, не осуществляется, это отдается на откуп пользователю.

### Профили

Поддерживается загрузка профилей конфигурации. Для каждого персонажа можно добавить отдельный профиль и запускать программу с нужным профилем:
```shell
# Загружаем профиль ./profiles/default.yml
> sudo ./eternal_chat_speech_macro_linux -c default
```
Инструкция по заполнению профиля находится внутри [default.yml](./build/profiles/default.yml).

[//]: # (### Компиляция)
[//]: # (```shell)
[//]: # (> go build -o ./build/eternal_chat_speech_macro_linux main.go)
[//]: # (```)