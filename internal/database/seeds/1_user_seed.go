package database

func init() {
	//RegisterSeeder(Seeder{
	//	Name: "userSeed",
	//	Run: func(tx *gorm.DB) error {
	//		items := []model.User{
	//			{
	//				Email:    "admin@wtb.id",
	//				Name:     "admin",
	//				Password: "password",
	//				Role:     helper.UserRoleAdmin,
	//			},
	//		}
	//
	//		// You can use tx.Model(&items) or bellow:
	//		if err := tx.Table("users").Create(&items).Error; err != nil {
	//			return err
	//		}
	//
	//		log.Println("✅ user seeded successfully")
	//		return nil
	//	},
	//})
}
