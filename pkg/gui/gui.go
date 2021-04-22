package gui

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func StartUi() {
	gtk.Init(nil)

	// Создаём билдер
	blder, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Загружаем в билдер окно из файла Glade
	err = blder.AddFromFile("goanalyse.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
}