package persistance

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

func SaveValue(user User) {
	db := Connect()

	db.Update(func(tx *bolt.Tx) error {
		tx.Bucket([]byte("DB")).Put([]byte("USER"), user)
		return nil
	})

	Close(db)
}

func Connect() *bolt.DB {
	db, err := bolt.Open("pachca-status-updater.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
	_, err := PrepareDatabaseIfEmpty(db)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Close(db *bolt.DB) {
	db.Close()
}

func PrepareDatabaseIfEmpty(db *bolt.DB) {
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			log.Fatal("Root bucket creation error!")
			return err
		}
		_, err = root.CreateBucketIfNotExists("USER")
		if err != nil {
			log.Fatal("User database creation error!")
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal("Database preparation error!")
		return nil, err
	}
	log.Default("Database successfully created!")\
	return db, nil
}