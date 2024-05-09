package migrations

import (
	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

type migrations struct {
	id string
}

func checkVersion(id string, tx *xorm.Engine) bool{
	// check version
	has, err:= tx.Exist("migrations", "id = ?", id)
	if err != nil {
		return false
	}
	if has {
		return true
	}
	return false
}

func MakeMigration001() *xormigrate.Migration {
    return &xormigrate.Migration{
        ID:          "001",
        Description: "initialization",
        Migrate: func(tx *xorm.Engine) error {
			has := checkVersion("001", tx)
			if has {
				return nil
			}
            // create new tables
            err := tx.Sync2(new(migrations))
            return err
        },
        Rollback: func(tx *xorm.Engine) error {
            // drop tables
            err := tx.DropTables(new(migrations))
            return err
        },
    }
}
