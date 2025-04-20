package utils

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/JoelInman-Dev/tickwars/configs"
)

// these are all the loggers currently setup, in their own files
type Loggers struct {
	Player       *os.File
	Event        *os.File
	Error        *os.File
	Ticker       *os.File
	UserActivity *os.File
}

var (
	PlayerLogger       *slog.Logger
	EventLogger        *slog.Logger
	ErrorLogger        *slog.Logger
	TickerLogger       *slog.Logger
	UserActivityLogger *slog.Logger
)

// the following will get the log paths from the .env and create
//
//	or open the files so that we can begin logging to them
func CreateLogFiles() (*Loggers, error) {
	player, err := OpenLogFile(configs.GetEnv("PLAYER_LOG_PATH", "logs/player-bkp.log"))
	if err != nil {
		return nil, err
	}
	event, err := OpenLogFile(configs.GetEnv("EVENT_LOG_PATH", "logs/event-bkp.log"))
	if err != nil {
		return nil, err
	}
	errors, err := OpenLogFile(configs.GetEnv("ERROR_LOG_PATH", "logs/error-bkp.log"))
	if err != nil {
		return nil, err
	}
	ticker, err := OpenLogFile(configs.GetEnv("TICKER_LOG_PATH", "logs/ticker-bkp.log"))
	if err != nil {
		return nil, err
	}
	userActivity, err := OpenLogFile(configs.GetEnv("USER_ACTIVITY_LOG_PATH", "logs/user-activity-bkp.log"))
	if err != nil {
		return nil, err
	}
	// return all the usable log files
	return &Loggers{
		Player:       player,
		Event:        event,
		Error:        errors,
		Ticker:       ticker,
		UserActivity: userActivity,
	}, nil
}

func OpenLogFile(path string) (*os.File, error) {
	// using the path that is passed in, either open the existing file at that
	// path and append new logs to the bottom of it, or create the log file fresh
	// and set it up for logging. I chose this over os.Create() as OpenFile allows
	// persistant logs, whereas os.Create truncates the file to 0 bytes each time this is booted up
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file from: %s", path)
	}
	return file, nil
}

func CloseLogFiles(LogFiles *Loggers) {
	LogFiles.Player.Close()
	LogFiles.Event.Close()
	LogFiles.Error.Close()
	LogFiles.Ticker.Close()
	LogFiles.UserActivity.Close()
}

func InitLoggers(LogFiles *Loggers) {
	// little formatter for the dates and times used in the loggers
	// this converts the time attribute to Y-M-D H:i:s
	dateTimeformatFunc := func(groups []string, attr slog.Attr) slog.Attr {
		// customise the time formatting used in the logs
		if attr.Key == slog.TimeKey {
			time := attr.Value.Time()
			attr.Value = slog.StringValue(time.Format("2006-01-02 15:04:05"))
		}
		return attr
	}

	PlayerLogger = slog.New(slog.NewTextHandler(LogFiles.Player, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: dateTimeformatFunc,
	}))
	EventLogger = slog.New(slog.NewTextHandler(LogFiles.Event, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: dateTimeformatFunc,
	}))
	ErrorLogger = slog.New(slog.NewTextHandler(LogFiles.Error, &slog.HandlerOptions{
		Level:       slog.LevelWarn,
		ReplaceAttr: dateTimeformatFunc,
	}))
	TickerLogger = slog.New(slog.NewTextHandler(LogFiles.Ticker, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: dateTimeformatFunc,
	}))
	UserActivityLogger = slog.New(slog.NewTextHandler(LogFiles.UserActivity, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: dateTimeformatFunc,
	}))
}
