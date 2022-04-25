package main

import (
	"eternalChatMacro/config"
	"github.com/MarinX/keylogger"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

var ChatCommandKeys = []string{"/", "s", "p", "e", "e", "c", "h", "SPACE"}

func callback(k *keylogger.KeyLogger, phraseOptions config.PhraseOptions) {
	// Время на то, чтобы человек успел отпустить клавиши
	time.Sleep(1 * time.Second)
	write(k, []string{conf.ChatHotkey})
	time.Sleep(500 * time.Millisecond)
	write(k, ChatCommandKeys)

	subs := strings.Split(phraseOptions.Type, "_")
	ln := len(subs)
	for i, sub := range subs {
		write(k, strings.Split(sub, ""))
		if i != ln-1 {
			writeUnderScore(k)
		}
	}
	write(k, []string{"SPACE"})
	write(k, strings.Split(strconv.Itoa(int(phraseOptions.Index)), ""))
	write(k, []string{"ENTER"})
	wait = false
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

func writeUnderScore(k *keylogger.KeyLogger) {
	writePress(k, "L_SHIFT")
	time.Sleep(600 * time.Millisecond)
	write(k, []string{"-"})
	time.Sleep(100 * time.Millisecond)
	writeRelease(k, "L_SHIFT")
}
