package main

import "fmt"

type MessageService interface {
	SendNotification(int) bool
}

type SMSService struct{}

func (sms *SMSService) SendNotification(amount int) bool {
	fmt.Printf("Sending a notification about a charging: $%v", amount)
	return true
}

type MobifoneService struct {
	messageService MessageService
}

func (s *MobifoneService) Charge(amount int) error {
	//Charge
	//Send a notification
	s.messageService.SendNotification(amount)
	return nil
}

func main() {
	smsService := SMSService{}
	m := MobifoneService{&smsService}
	m.Charge(100)
}
