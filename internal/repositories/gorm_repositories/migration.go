package gorm_migration

type IMigration interface {
	// Migrate fonksiyonu veri tabanı tablolarını oluşturur
	Migrate() error
}

var migration []IMigration

func Add(migrateObj IMigration) {
	migration = append(migration, migrateObj)
}

func StartMigration() error {
	for i := range migration {
		if err := migration[i].Migrate(); err != nil {
			return err
		}
	}
	return nil
}
