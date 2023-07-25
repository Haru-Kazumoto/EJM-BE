package cmd

import (
	"EJM/config"
	db "EJM/init"
	"EJM/pkg/models"
	"EJM/utils"
	"log"
	"reflect"
	"time"

	// "github.com/casbin/casbin/v2/model"
	"github.com/spf13/cobra"
)

func initMigration(cfg *config.Config, args []string) {
	start := time.Now()
	dbObject := db.DBInitialize(cfg)
	log.Println("Start migration...")
	isDone := make(chan bool)
	isError := make(chan bool)
	go func() {
		modelsSlices := []interface{}{
			models.User{},
			models.Role{},
			models.Action{},
			models.RolesMenus{},
			models.JWTWhitelist{},
			models.Menu{},
			models.ActionsMenus{},
			models.MappingCode{},
			models.MappingKeywordList{},
			models.MJenisTransaksi{},
			models.ListOpCode{},
			models.BinKartu{},
			models.Term{},
			
		}
		var filtered []interface{}
		if args[0] != utils.ALL {
			for _, modelStruct := range modelsSlices {
				for _, arg := range args {
					if arg == reflect.TypeOf(modelStruct).Name() {
						filtered = append(filtered, modelStruct)
					}
				}
			}
		} else {
			filtered = modelsSlices
		}

		err := dbObject.Debug().AutoMigrate(filtered[:]...)
		if err != nil {
			isError <- true
			defer close(isError)
		}
		isDone <- true

		defer close(isDone)
	}()

	select {
	case <-isDone:
		log.Println("Migration is done!")
	case <-isError:
		log.Println("Migration failed!")
	}
	log.Printf("It takes %f ms", time.Since(start).Seconds())
}

func NewMigration(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "migration",
		Short: "Start Migration",
		Run: func(cmd *cobra.Command, args []string) {
			initMigration(cfg, args)
		},
	}
}

// func init() {
// 	Migrations = append(Migrations, &gormigrate.Migration{
// 		ID: "20230714000001",
// 		Migrate: func(tx *gorm.DB) error {
// 			// GORM raw SQL to create the enum type
// 			rawSQL := "CREATE TYPE status_enum AS ENUM ('success', 'fail');"
// 			if err := tx.Exec(rawSQL).Error; err != nil {
// 				return err
// 			}

// 			// GORM raw SQL to add the column with the enum type
// 			rawSQL = "ALTER TABLE mapping_codes ADD COLUMN status status_enum;"
// 			if err := tx.Exec(rawSQL).Error; err != nil {
// 				return err
// 			}

// 			return nil
// 		},
// 		Rollback: func(tx *gorm.DB) error {
// 			// GORM raw SQL to drop the column
// 			rawSQL := "ALTER TABLE mapping_codes DROP COLUMN status;"
// 			if err := tx.Exec(rawSQL).Error; err != nil {
// 				return err
// 			}

// 			// GORM raw SQL to drop the enum type
// 			rawSQL = "DROP TYPE status_enum;"
// 			if err := tx.Exec(rawSQL).Error; err != nil {
// 				return err
// 			}

// 			return nil
// 		},
// 	})
// }

// var Migrations = []*gormigrate.Migration{
// 	// Definisi migrasi-migrasi lainnya (jika ada)
// 	// ...
// }

