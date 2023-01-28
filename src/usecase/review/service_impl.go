package review

func (s *service) Migrate() {
	s.rRepo.AutoMigrate()
}
