package products

const (
	getAllQuery = `SELECT id,name from products`
	createQuery = `INSERT INTO products (name) VALUES (?)`
)
