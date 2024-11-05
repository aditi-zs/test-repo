package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const createQuery = `CREATE TABLE IF NOT EXISTS products 
(
    id int auto_increment primary key,
    name varchar(100) not null,
);`

func create_product_table() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createQuery)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
