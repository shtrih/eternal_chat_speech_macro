package main

import (
	"eternalChatMacro/config"
	"eternalChatMacro/keyboard"
	"flag"
	"github.com/MarinX/keylogger"
	log "github.com/sirupsen/logrus"
)

const Version = "1.0.0"

var wait bool
var conf *config.Config

func main() {
	log.Info("Eternal Chat Speech Macro ver. ", Version)

	configName := flag.String("c", "default", "Название файла конфигурации без расширения")
	logLevel := flag.String("log", "info", "Уровень логирования (debug, info, warn, error)")
	flag.Parse()

	level, err := log.ParseLevel(*logLevel)
	if err != nil {
		log.Error(err)
	}
	log.SetLevel(level)

	// find keyboard device, does not require a root permission
	kbd := keylogger.FindKeyboardDevice()

	// check if we found a path to keyboard
	if len(kbd) <= 0 {
		log.Error("No keyboard found...you will need to provide manual input path")
		return
	}

	log.Println("Found a keyboard at", kbd)
	// init keylogger with keyboard
	k, err := keylogger.New(kbd)
	if err != nil {
		log.Error(err)
		return
	}
	defer k.Close()

	events := k.Read()
	watcher := keyboard.NewWatcher()

	conf = config.Init(*configName)
	log.WithField("Chat Key", conf.ChatHotkey).Info()
	printInfo(conf.Phrases)

	// range of events
	for e := range events {
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events
		case keylogger.EvKey:
			key := keyboard.Key(e.KeyString())

			// if the state of key is pressed
			if e.KeyPress() {
				log.WithField("Key", e.KeyString()).Debug("[event] press key")

				watcher.SetState(key, keyboard.Down)

				for _, h := range conf.Phrases {
					combo := true
					for _, comboKey := range h.Hotkeys {
						combo = combo && watcher.Down(keyboard.Key(comboKey))
					}
					if combo && !wait {
						log.WithFields(log.Fields{"Keys": h.Hotkeys, "Comment": h.Comment}).Debug("Combo")
						wait = true
						go callback(k, h)
					}
				}
			}

			// if the state of key is released
			if e.KeyRelease() {
				log.WithField("Key", e.KeyString()).Debug("[event] release key")

				watcher.SetState(key, keyboard.Up)
			}

			break
		}
	}
}

func printInfo(options []config.PhraseOptions) {
	for _, h := range options {
		log.WithField("Hotkey", h.Hotkeys).Info(h.Comment)
	}
}
