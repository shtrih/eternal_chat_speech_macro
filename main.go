package main

import (
	"eternalChatMacro/keyboard"
	"github.com/MarinX/keylogger"
	log "github.com/sirupsen/logrus"
)

var wait bool

func main() {
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

	handlers := GetHandlers()
	printInfo(handlers)

	// range of events
	for e := range events {
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events
		case keylogger.EvKey:
			key := keyboard.Key(e.KeyString())

			// if the state of key is pressed
			if e.KeyPress() {
				log.Debug("[event] press key ", e.KeyString())

				watcher.SetState(key, keyboard.Down)

				for _, h := range handlers {
					combo := true
					for _, comboKey := range h.Keys {
						combo = combo && watcher.Down(comboKey)
					}
					if combo && !wait {
						log.WithField("Keys", h.Keys).Debug("COMBO")
						wait = true
						go h.Callback(k)
					}
				}
			}

			// if the state of key is released
			if e.KeyRelease() {
				log.Debug("[event] release key ", e.KeyString())

				watcher.SetState(key, keyboard.Up)
			}

			break
		}
	}
}

func printInfo(handlers []Handler) {
	for _, h := range handlers {
		log.WithField("Key", h.Keys).Info(h.Comment)
	}
}

func write(k *keylogger.KeyLogger, keys []string) {
	for _, key := range keys {
		// write once will simulate keyboard press/release, for long press or release, lookup at Write
		err := k.WriteOnce(key)
		if err != nil {
			log.Error(err)
		}
	}
}

func writePress(k *keylogger.KeyLogger, key string) {
	err := k.Write(keylogger.KeyPress, key)
	if err != nil {
		log.WithField("key", key).Warn("writePress", err)
	}
}

func writeRelease(k *keylogger.KeyLogger, key string) {
	err := k.Write(keylogger.KeyRelease, key)
	if err != nil {
		log.WithField("key", key).Warn("writeRelease", err)
	}
}
