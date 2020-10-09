package gomie

type Controller struct {
	baseTopic TopicID
	devices   []TopicID
}

func (c *Controller) Start(client MQTTClient) error {
	if err := client.Subscribe("+/+/$homie", c.deviceFound); err != nil {
		return err
	}

	return nil
}

func (c *Controller) deviceFound(message Message) {

}
