package config

import (
	"errors"
	"strconv"
	"strings"
	"os"
)

type ConfigParameters struct {
	Connection struct {
			   RabbitmqURL string
			   RetryDelay  int
		   }
	Exchange   struct {
			   Name string
		   }
	Queue      struct {
			   Name          string
			   WaitDelay     int
			   PrefetchCount int
		   }
	Message    struct {
			   DefaultTTL int
		   }
	Http       struct {
			   DefaultMethod string
			   Timeout       int
		   }
	Log        struct {
			   LogFile string
			   ErrFile string
		   }
}

func (config *ConfigParameters) ReadEnvVars() error {
	user := "guest"
	pass := "guest"
	port := "5672"
	host := "localhost"
	exchange := "syntercom_exchnage"
	RetryDelay := 30
	QueueName := "syntercom"
	PrefetchCount := 10
	QueueWaitDelay := 30
	DefaultTtl := 86400
	DefaultMethod := "POST"
	Timeout := 30

	if len(os.Getenv("RABBIT_USER")) > 0 {
		user = os.Getenv("RABBIT_USER")
	}
	if len(os.Getenv("RABBIT_PASS")) > 0 {
		pass = os.Getenv("RABBIT_PASS")
	}
	if len(os.Getenv("RABBIT_PORT")) > 0 {
		port = os.Getenv("RABBIT_PORT")
	}
	if len(os.Getenv("RABBIT_HOST")) > 0 {
		host = os.Getenv("RABBIT_HOST")
	}
	if len(os.Getenv("RABBIT_EXCHANGE")) > 0 {
		exchange = os.Getenv("RABBIT_EXCHANGE")
	}

	config.Connection.RabbitmqURL = "amqp://" + user + ":" + pass + "@" + host + ":" + port + "/"
	config.Connection.RetryDelay = RetryDelay
	config.Queue.Name = QueueName
	config.Queue.PrefetchCount = PrefetchCount
	config.Queue.WaitDelay = QueueWaitDelay
	config.Message.DefaultTTL = DefaultTtl
	config.Http.DefaultMethod = DefaultMethod
	config.Http.Timeout = Timeout
	config.Exchange.Name = exchange

	if config.Connection.RetryDelay < 5 {
		return errors.New("Connection Retry Delay must be at least 5 seconds")
	}

	if len(config.Queue.Name) == 0 {
		return errors.New("Queue Name is empty or missing")
	}

	if config.Queue.WaitDelay < 1 {
		return errors.New("Queue Wait Delay must be at least 1 second")
	}

	if config.Queue.PrefetchCount < 1 {
		return errors.New("PrefetchCount cannot be negative")
	}

	if config.Message.DefaultTTL < 1 {
		return errors.New("Message Default TTL must be at least 1 second")
	}

	var ok bool
	config.Http.DefaultMethod, ok = CheckMethod(config.Http.DefaultMethod)
	if !ok {
		return errors.New("Http Default Method is not recognized: " + config.Http.DefaultMethod)
	}

	if config.Http.Timeout < 5 {
		return errors.New("Http Timeout must be at least 5 seconds")
	}

	config.Log.LogFile = "rabbitmq-worker.log"
	config.Log.ErrFile = "rabbitmq-worker.err"

	return nil
}

func (config *ConfigParameters) String() string {
	cfgDtls := ""
	cfgDtls += "[Connection]\n"
	cfgDtls += "  RabbitmqURL = \"" + config.Connection.RabbitmqURL + "\"\n"
	cfgDtls += "  RetryDelay = " + strconv.Itoa(config.Connection.RetryDelay) + "\n"
	cfgDtls += "[Queue]\n"
	cfgDtls += "  Name = \"" + config.Queue.Name + "\"\n"
	cfgDtls += "  WaitDelay = " + strconv.Itoa(config.Queue.WaitDelay) + "\n"
	cfgDtls += "  PrefetchCount = " + strconv.Itoa(config.Queue.PrefetchCount) + "\n"
	cfgDtls += "[Message]\n"
	cfgDtls += "  DefaultTTL = " + strconv.Itoa(config.Message.DefaultTTL) + "\n"
	cfgDtls += "[Http]\n"
	cfgDtls += "  DefaultMethod = " + config.Http.DefaultMethod + "\n"
	cfgDtls += "  Timeout = " + strconv.Itoa(config.Http.Timeout) + "\n"
	cfgDtls += "[Log]\n"
	cfgDtls += "  LogFile = \"" + config.Log.LogFile + "\"\n"
	cfgDtls += "  ErrFile = \"" + config.Log.ErrFile + "\""

	return cfgDtls
}

func CheckMethod(method string) (upperMethod string, ok bool) {
	methods := make(map[string]bool)
	methods["GET"] = true
	methods["HEAD"] = true
	methods["POST"] = true
	methods["PUT"] = true
	methods["PATCH"] = true
	methods["DELETE"] = true
	methods["CONNECT"] = true
	methods["OPTIONS"] = true
	methods["TRACE"] = true

	upperMethod = strings.ToUpper(method)
	_, ok = methods[upperMethod]

	return
}
