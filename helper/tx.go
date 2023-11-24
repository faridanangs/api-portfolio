package helper

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		err := tx.Rollback()
		PanicIfError(err, "Error Rollback on helper at CommitOrRollback")
		panic(err)
	} else {
		err := tx.Commit()
		PanicIfError(err, "Error Commit on helper at CommitOrRollback")
	}
}
