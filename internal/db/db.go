package db

import (
	`context`
	
	pgx `github.com/jackc/pgx/v4`
	`github.com/pkg/errors`
)

type DB struct {
	conn *pgx.Conn
}

func New(ctx context.Context, connection string) (*DB, error) {
	conn, err := pgx.Connect(ctx, connection)
	
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to db")
	}
	
	return &DB{conn: conn}, nil
}

func (s * DB) Close(ctx context.Context) error {
	if err := s.conn.Close(ctx); err != nil {
		return err
	}
	
	return nil
}