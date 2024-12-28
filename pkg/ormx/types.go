package ormx

type DBConn interface {
	Rollback() error
	Commit() error
}

type Transaction interface {
	Tx(fn func(session any) error) error
}
