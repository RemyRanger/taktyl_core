package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"taktyl.com/m/src/models"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var events = []models.Event{
	models.Event{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Event{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

// Load : init database struct
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Event{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Event{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Event{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		events[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Event{}).Create(&events[i]).Error
		if err != nil {
			log.Fatalf("cannot seed events table: %v", err)
		}
	}
}
