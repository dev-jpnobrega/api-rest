package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
)

// Create - database model
func main() {
	// BaseModel define
	type BaseModel struct {
		ID        uuid.UUID `gorm:"type:uuid;primary_key"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}

	db, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost/profile?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	mm := []*gormigrate.Migration{
		// MigraCreateTableUser(),
		{
			ID: "20200302080000",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					BaseModel
					Name  string
					Email string
				}

				return tx.AutoMigrate(&User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("user").Error
			},
		},
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, mm)

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")

	/*
		absPath, _ := filepath.Abs("./src/db/migration")
		files, err := ioutil.ReadDir(absPath)

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())

			ff := absPath + "\\" + file.Name()

			arg0 := "-e"
			arg1 := "Migra"
			arg2 := "\n\tfrom"
			arg3 := "golang"

			cmd := exec.Command(ff, arg0, arg1, arg2, arg3)
			stdout, err := cmd.Output()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			//command := exec.Command("go", ff, "Migra")

			//var out bytes.Buffer
			//var sterr bytes.Buffer

			//command.Stderr = &sterr
			//command.Stdout = &out

			fmt.Println("stdout", string(stdout))
		}
	*/

}
