package logger

import (
	"os"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	// Настройка формата логов
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Установка уровня логирования
	Log.SetLevel(logrus.InfoLevel)

	// Запись логов в файл
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.Warn("Failed to open log file, using default stderr")
	}
}
