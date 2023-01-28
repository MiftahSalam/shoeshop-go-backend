package order

func (s *service) Migrate() {
	s.oRepo.AutoMigrate()
}
