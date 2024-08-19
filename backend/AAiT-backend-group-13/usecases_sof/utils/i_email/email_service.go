package iemail

type Service interface {
  Send(mail *Mail) error
}