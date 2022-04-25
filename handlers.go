package main

import (
	"eternalChatMacro/keyboard"
	"github.com/MarinX/keylogger"
	"time"
)

var CommonPrefix = []string{"BS", "/", "s", "p", "e", "e", "c", "h", "SPACE"}

type Handler struct {
	Keys     []keyboard.Key
	Callback func(k *keylogger.KeyLogger)
	Comment  string
}

func GetHandlers() []Handler {
	return []Handler{
		{
			Keys: []keyboard.Key{"R_ALT", "1"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, []string{"y"})
				time.Sleep(500 * time.Millisecond)
				write(k, CommonPrefix)
				write(k, []string{"h", "e", "l", "l", "o", "SPACE", "1", "7"})
				write(k, []string{"ENTER"})
				wait = false
			},
			Comment: "Мне не смешно",
		},
		{
			Keys: []keyboard.Key{"R_ALT", "2"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, []string{"y"})
				time.Sleep(500 * time.Millisecond)

				write(k, CommonPrefix)
				write(k, []string{"h", "e", "l", "l", "o", "SPACE", "1", "9"})
				write(k, []string{"ENTER"})
				wait = false
			},
			Comment: "Не хочу, чтобы нас видели вместе",
		},
		{
			Keys: []keyboard.Key{"R_ALT", "3"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, []string{"y"})
				time.Sleep(500 * time.Millisecond)

				write(k, CommonPrefix)
				write(k, []string{"h", "e", "l", "l", "o", "SPACE", "2", "7"})
				write(k, []string{"ENTER"})
				wait = false
			},
			Comment: "Ничтожество",
		},
		{
			Keys: []keyboard.Key{"R_ALT", "4"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, []string{"y"})
				time.Sleep(500 * time.Millisecond)

				write(k, CommonPrefix)
				write(k, []string{"h", "e", "l", "l", "o", "SPACE", "3", "1"})
				write(k, []string{"ENTER"})
				wait = false
			},
			Comment: "Не сегодня!",
		},
		{
			Keys: []keyboard.Key{"R_ALT", "5"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, []string{"y"})
				time.Sleep(500 * time.Millisecond)

				write(k, CommonPrefix)
				write(k, []string{"h", "e", "l", "l", "o", "SPACE", "9", "2"})
				write(k, []string{"ENTER"})
				wait = false
			},
			Comment: "Давайте скорее, я не могу целый день тут торчать.",
		},
		{
			Keys: []keyboard.Key{"R_ALT", "6"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, []string{"y"})
				time.Sleep(500 * time.Millisecond)

				write(k, CommonPrefix)
				write(k, []string{"h", "e", "l", "l", "o", "SPACE", "8", "6"})
				write(k, []string{"ENTER"})
				wait = false
			},
			Comment: "Быстрее. У меня мало времени.",
		},
		{
			Keys: []keyboard.Key{"R_ALT", "7"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, []string{"y"})
				time.Sleep(500 * time.Millisecond)

				write(k, CommonPrefix)
				write(k, []string{"o", "r", "d"})
				writePress(k, "L_SHIFT")
				time.Sleep(500 * time.Millisecond)
				write(k, []string{"-"})
				writeRelease(k, "L_SHIFT")
				write(k, []string{"h", "e", "l", "l", "o", "SPACE", "3"})
				write(k, []string{"ENTER"})
				wait = false
			},
			Comment: "Мы следим за тобой. Мерзавец.",
		},
		{
			Keys: []keyboard.Key{"R_ALT", "8"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, []string{"y"})
				time.Sleep(500 * time.Millisecond)

				write(k, CommonPrefix)
				write(k, []string{"h", "e", "l", "l", "o", "SPACE", "2", "1", "7"})
				write(k, []string{"ENTER"})
				wait = false
			},
			Comment: "Я тепло приветствую вас.",
		},
		{
			Keys: []keyboard.Key{"R_ALT", "9"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, []string{"y"})
				time.Sleep(500 * time.Millisecond)

				write(k, CommonPrefix)
				write(k, []string{"h", "e", "l", "l", "o", "SPACE", "1", "3", "8"})
				write(k, []string{"ENTER"})
				wait = false
			},
			Comment: "Как вам Балмора?",
		},
		{
			Keys: []keyboard.Key{"R_ALT", "0"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, []string{"y"})
				time.Sleep(500 * time.Millisecond)

				write(k, CommonPrefix)
				write(k, []string{"h", "e", "l", "l", "o", "SPACE", "1", "1", "9"})
				write(k, []string{"ENTER"})
				wait = false
			},
			Comment: "Мутсера?",
		},
		{
			Keys: []keyboard.Key{"L_ALT", "1"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, append(CommonPrefix, []string{"h", "e", "l", "l", "o", "SPACE", "1", "-", "2", "3", "3"}...))
				wait = false
			},
			Comment: "hello template",
		},
		{
			Keys: []keyboard.Key{"L_ALT", "2"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, append(CommonPrefix, []string{"i", "d", "l", "e", "SPACE", "1", "-", "9"}...))
				wait = false
			},
			Comment: "idle template",
		},
		{
			Keys: []keyboard.Key{"L_ALT", "3"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, append(CommonPrefix, []string{"o", "r", "d"}...))
				writePress(k, "L_SHIFT")
				writePress(k, "-")
				writeRelease(k, "-")
				writeRelease(k, "L_SHIFT")
				write(k, []string{"i", "d", "l", "e", "SPACE", "1", "-", "4"})
				wait = false
			},
			Comment: "ordinator idle template",
		},
		{
			Keys: []keyboard.Key{"L_ALT", "4"},
			Callback: func(k *keylogger.KeyLogger) {
				time.Sleep(1 * time.Second)
				write(k, append(CommonPrefix, []string{"t", "h", "i", "e", "f", "SPACE", "1", "-", "5"}...))
				wait = false
			},
			Comment: "thief template",
		},
	}
}
