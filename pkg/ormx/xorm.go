package ormx

import "xorm.io/xorm"

type TxSession struct {
	session *xorm.Session
}

func (t *TxSession) Tx(fn func(session any) error) error {
	session := t.session
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}
	err := fn(session)
	if err != nil {
		er := session.Rollback()
		if er != nil {
			return er
		}
		return err
	}
	return session.Commit()
}

func (t *TxSession) EngineTx(fn func(engine *xorm.Engine) error) error {
	f := func(session any) error {
		sess, ok := session.(*xorm.Session)
		if !ok {
			return ErrTypeConvert
		}
		return fn(sess.Engine())
	}
	return t.Tx(f)
}

func NewTxSession(engine *xorm.Session) Transaction {
	return &TxSession{
		session: engine,
	}
}
