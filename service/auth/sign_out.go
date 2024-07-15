package auth

func (s *service) SignOut(userId uint32) (err error) {
	return s.userRepo.Update(&userId, &map[string]any{"token": ""})
}
