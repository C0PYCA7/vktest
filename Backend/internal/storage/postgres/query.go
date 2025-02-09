package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
	"vk/Backend/internal/models"
	"vk/Backend/internal/storage"
)

const UNIQUE_CODE_PG = "23505"

func (d *Database) SelectAllContainersData() ([]models.Container, error) {
	var (
		container  models.Container
		ctx        = context.Background()
		containers = make([]models.Container, 0)
	)

	rows, err := d.db.Query(ctx, "SELECT containerIP, pingTimeMKS, lastSuccessDate FROM containers_stats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&container.ContainerIP, &container.PingTimeMKs, &container.LastSuccessDate); err != nil {
			return nil, err
		}
		containers = append(containers, container)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return containers, nil
}

func (d *Database) CreateContainer(containerIp string) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		tx.Rollback(ctx)
		return storage.ErrBeginTx
	}

	pgErr := &pgconn.PgError{}
	_, err = tx.Exec(ctx, "INSERT INTO containers_stats(containerIP) VALUES($1)", containerIp)
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == UNIQUE_CODE_PG {
			tx.Rollback(ctx)
			return storage.ErrUniqueIP
		}
		tx.Rollback(ctx)
		return err
	}
	tx.Commit(ctx)
	return nil
}

func (d *Database) UpdateContainerInfo(containerIP string, pingTimeMs int, lastSuccessData time.Time) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		tx.Rollback(ctx)
		return storage.ErrBeginTx
	}

	_, err = tx.Exec(ctx, "UPDATE containers_stats SET pingTimeMKS = $1, lastSuccessDate = $2 WHERE containerIP = $3", pingTimeMs, lastSuccessData, containerIP)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	tx.Commit(ctx)
	return nil
}
