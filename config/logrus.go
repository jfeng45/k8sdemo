// package logrus handles creating logrus logger
package config

import (
	"fmt"
	"github.com/jfeng45/k8sdemo/tool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

//type loggerWrapper struct {
//	lw *logrus.Logger
//}
//
//func (logger *loggerWrapper) Errorf(format string, args ...interface{}) {
//	logger.lw.Errorf(format, args)
//}
//func (logger *loggerWrapper) Fatalf(format string, args ...interface{}) {
//	logger.lw.Fatalf(format, args)
//}
//func (logger *loggerWrapper) Fatal(args ...interface{}) {
//	logger.lw.Fatal(args)
//}
//func (logger *loggerWrapper) Infof(format string, args ...interface{}) {
//	logger.lw.Infof(format, args)
//}
//func (logger *loggerWrapper) Warnf(format string, args ...interface{}) {
//	logger.lw.Warnf(format, args)
//}
//func (logger *loggerWrapper) Debugf(format string, args ...interface{}) {
//	logger.lw.Debugf(format, args)
//}
//func (logger *loggerWrapper) Printf(format string, args ...interface{}) {
//	logger.lw.Printf(format, args)
//}
//func (logger *loggerWrapper) Println(args ...interface{}) {
//	logger.lw.Println(args)
//}

func RegisterLogrusLog() error {
	//standard configuration
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetReportCaller(true)
	file, err := os.OpenFile("../logs/demo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : ", err)
		return errors.Wrap(err, "")
	}
	mw := io.MultiWriter(os.Stdout,file)
	log.SetOutput(mw)
	//log.SetOutput(os.Stdout)
	//customize it from configuration file
	err = customizeLogrusFromConfig(log)
	if err != nil {
		return errors.Wrap(err, "")
	}
	//This is for loggerWrapper implementation
	//logger.Logger(&loggerWrapper{log})
	tool.SetLogger(log)
	return nil
}

// customizeLogFromConfig customize log based on parameters from configuration file
func customizeLogrusFromConfig(log *logrus.Logger) error {
	log.SetReportCaller(false)
	//log.SetOutput(os.Stdout)
	l := &log.Level
	err := l.UnmarshalText([]byte("debug"))
	if err != nil {
		return errors.Wrap(err, "")
	}
	log.SetLevel(*l)
	return nil
}
