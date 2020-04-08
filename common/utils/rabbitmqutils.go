package utils

import (
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMqServer struct {
	dialHost  string
	queueName string
	conn      *amqp.Connection
	channel   *amqp.Channel
}

func NewRabbitMqServer(host,queue string) (mq *RabbitMqServer,err error) {
    conn,err:=amqp.Dial(host)
    if err!=nil{
    	return nil,err
	}
    channel,err:=conn.Channel()
	if err!=nil{
		return nil,err
	}
	return &RabbitMqServer{dialHost: host, queueName:queue,conn:conn,channel:channel},nil
}

func (l *RabbitMqServer) CloseRabbitmqConn() {
	err := l.conn.Close()
	if err != nil {
		fmt.Println("CloseRabbitmqConn Conn Error ", err.Error())
	}
	if l.channel != nil {
		err = l.channel.Close()
		if err != nil {
			fmt.Println("CloseRabbitmqConn Channel Error ", err.Error())
		}
	}
}

func (l *RabbitMqServer) PushMessage(message string) error {
	que,err:=l.channel.QueueDeclare(l.queueName,true,false,false,false,nil)
	if err!=nil{
		return err
	}
	err = l.channel.Publish("",que.Name,false,false,amqp.Publishing{Body:[]byte(message)})
	if err!=nil{
		return err
	}
	return nil
}

//func  传参,业务处理各自逻辑,底层统一处理消费
//string 传参,底层需区分各自的业务逻辑
func (l *RabbitMqServer) ConsumeMessage(consumeFunc func(msg string) error) {
	que, err := l.channel.QueueDeclare(l.queueName, true, false, false, false, nil)
	if err!=nil{
		fmt.Println("ConsumeMessage QueueDeclare Error",err.Error())
	}
	deliveryList,err:=l.channel.Consume(que.Name, "",true,false,false,false,nil)
    go func() {
    	for d:=range deliveryList{
           msgDeli:=string(d.Body)
           fmt.Println("ConsumeMessage Msg -->> ",msgDeli)
           err = consumeFunc(msgDeli)
			if err != nil {
				fmt.Println("ConsumeMessage Error -->> ",err.Error())
				err = l.PushMessage(msgDeli)
				if err!=nil{
					fmt.Println("ConsumeMessage Publish Error -->> ",err.Error())
				}
			} else {
				fmt.Println("ConsumeMessage Success")
			}
		}
	}()
}
