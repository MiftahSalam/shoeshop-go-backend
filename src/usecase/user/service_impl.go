package user

func (s *service) Migrate() {
	s.uRepo.AutoMigrate()
}
