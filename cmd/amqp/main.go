package main

import (
	"fmt"
	"os"

	"git.alibaba.ir/b2b/back/push2b/internal/consume"
	"git.alibaba.ir/b2b/back/push2b/internal/dispatch"
	sender "git.alibaba.ir/b2b/back/push2b/internal/email"
	"git.alibaba.ir/b2b/back/push2b/pkg"
	"gopkg.in/yaml.v2"
)

func main() {

	conf := initConf()
	fillDefaultConfiguratiosn(&conf)

	dispatcher := dispatch.NewInMemoryDispatcher()

	dispatcher.AddListener(pkg.MAIL_GMAIL, sender.NewGoogleSender(conf.GmailConfig))
	dispatcher.AddListener(pkg.MAIL_OUTLOOK, sender.NewOutlookSender(conf.OutlookConfig))

	consumer := consume.NewRabbitConsumer(conf.RabbitConfig, dispatcher)

	//starts the server consuming Rabbitmq for jobs :)
	err := consumer.Consume()

	if nil != err {
		fmt.Println("server stopped with error: ", err)
		return
	}

	fmt.Println("server stopped with error code: 0")
}

func initConf() pkg.Config {
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg pkg.Config

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func fillDefaultConfiguratiosn(cfg *pkg.Config) {

	if cfg.RabbitConfig.RetrySeconds <= 0 || cfg.RabbitConfig.RetrySeconds > 10 {
		cfg.RabbitConfig.RetrySeconds = 1
	}
}
