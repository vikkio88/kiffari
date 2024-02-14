package db

func (d *Db) GetHashedKey() (*Passkey, bool) {
	var p Passkey
	trx := d.g.Model(&p).Find(&p)

	return &p, trx.RowsAffected == 1
}

func (d *Db) StorePassword(p Passkey) bool {
	trx := d.g.Create(&p)

	return trx.RowsAffected == 1
}
