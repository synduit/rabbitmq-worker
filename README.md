# RabbitMQ Worker

## Overview

In order to reliably send an http request, a client inserts a message containing the request into a RabbitMQ queue. The Go consumer parses messages from this main queue and then attempts the request. Failed requests are placed on a wait queue, where they stay until a configured delay period has expired. They are then moved back to the main queue for retry. Messages are retried until: the http request is successful, the message expiration time has been reached, or a permanent error has been encountered (e.g. message parsing error, 4XX http response code).

## Architecture

[Click here](docs/)

## Development

**Language:** GO Lang

**Platform:** Docker

These steps have been tested with Ubuntu 14.04, but should work for any Linux distro.

- Install RabbitMQ (version 3.6.1 or later): [RabbitMQ Install](http://www.rabbitmq.com/download.html)  
- Enable the following plugins using the 'rabbitmq-plugins' command: [RabbitMQ Plugins Manual Page](https://www.rabbitmq.com/man/rabbitmq-plugins.1.man.html)  
  `rabbitmq_mqtt, mochiweb, webmachine, rabbitmq_web_dispatch, rabbitmq_management_agent, rabbitmq_management, rabbitmq_amqp1_0, amqp_client`
- Install and enable the RabbitMQ Message Timestamp plugin.  
  The Timestamp plugin can be downloaded here: [RabbitMQ Community Plugins](https://www.rabbitmq.com/community-plugins.html)  
  Instructions for installing additional plugins are here: [Additional Plugins](https://www.rabbitmq.com/installing-plugins.html)  
  *NOTE: The Timestamp plugin is not mandatory, but recommended since it is used to create a unique message id if one is not provided.*
- Add the RabbitMQ user, "rmq", and assign permissions:  
`sudo rabbitmqctl add_user rmq rmq`  
`sudo rabbitmqctl set_permissions rmq '.*' '.*' '.*'`  
`sudo rabbitmqctl set_user_tags rmq administrator`
- Download and install Go: [Downloads - The Go Programming Language](https://golang.org/dl/)  
  Follow instructions to:
  - Add the Go bin directory to the PATH environment variable
  - Setup a workspace
  - Set the GOPATH environment variable to point to the workspace.  
  - In addition, add the workspace bin directory, "$GOPATH/bin", to the PATH. This should appear in the shell startup script immediately after GOPATH is set.
- Install Git: [Setup Git](https://help.github.com/articles/set-up-git/)
- Download Go packages:  
  - go get gopkg.in/gcfg.v1
  - go get github.com/streadway/amqp
  - go get github.com/LinioIT/rabbitmq-worker
- Build and install executables:
  - go install github.com/LinioIT/rabbitmq-worker
  - go install github.com/LinioIT/rabbitmq-worker/insertHttpRequest
  - go install github.com/LinioIT/rabbitmq-worker/deleteQueue
  - go install github.com/LinioIT/rabbitmq-worker/webserver

## Project Management

Slack: [#api](https://slack.com/app_redirect?channel=rabbitmq-worker&team=T02F7GWJT)

JIRA: [API](https://brandentity.atlassian.net/secure/RapidBoard.jspa?rapidView=13&projectKey=SCOM)

## Maintainers

Karthik Bodu - @karthik
