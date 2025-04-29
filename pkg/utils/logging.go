package utils

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/JoelInman-Dev/tickwars/configs"
)

// these are all the log files currently setup, paths set in ENV
type Loggers struct {
	Player       *os.File
	Event        *os.File
	Error        *os.File
	Ticker       *os.File
	UserActivity *os.File
}

// these are the actual log handlers which get called
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

func InitLoggers(LogFiles *Loggers) {
	// firstly, a little formatter for the dates and times used in the loggers
	// this converts the time attribute to Y-M-D H:i:s as defaults are FUGLY!
	dateTimeformatFunc := func(groups []string, attr slog.Attr) slog.Attr {
		// customise the time formatting used in the logs
		// the KIND check is basically Slog's version of a type checker, to
		//  ensure your editing the correct type of value
		if attr.Key == slog.TimeKey && attr.Value.Kind() == slog.KindTime {
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

// Don't forget to clean yo shit up afterwards!
func CloseLogFiles(LogFiles *Loggers) {
	loggerFiles := []*os.File{
		LogFiles.Player,
		LogFiles.Event,
		LogFiles.Error,
		LogFiles.Ticker,
		LogFiles.UserActivity,
	}
	// looping the logfiles to close them ensures that the full list
	// of log files is processed in the loop and any errors are logged out
	for _, file := range loggerFiles {
		err := file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to close log file: %v\n", err)
		}
	}
}
