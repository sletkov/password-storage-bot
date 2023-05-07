package models

type Service struct {
	Id          int    `json:"id"`
	UserName    string `json:"user_name"`
	ServiceName string `json:"service_name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	//EncryptedLogin    string `json:"-"`
	//EncryptedPassword string `json:"-"`
}

//func (s *Service) BeforeCreate() error {
//	if len(s.Password) > 0 {
//		enc, err := encryptString(s.Password)
//
//		if err != nil {
//			return err
//		}
//
//		s.EncryptedPassword = enc
//	}
//
//	if len(s.Login) > 0 {
//		enc, err := encryptString(s.Password)
//
//		if err != nil {
//			return err
//		}
//
//		s.EncryptedLogin = enc
//	}
//
//	return nil
//}
//
//func (s *Service) Sanitize() {
//	s.Login = ""
//	s.Password = ""
//}
//
//func (s *Service) CompareLogin(login string) bool {
//	return bcrypt.CompareHashAndPassword([]byte(s.EncryptedLogin), []byte(s.Login)) == nil
//}
//
//func (s *Service) ComparePassword(password string) bool {
//	return bcrypt.CompareHashAndPassword([]byte(s.EncryptedPassword), []byte(s.Password)) == nil
//}
//
//func encryptString(s string) (string, error) {
//	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
//
//	if err != nil {
//		return "", err
//	}
//
//	return string(b), nil
//}
