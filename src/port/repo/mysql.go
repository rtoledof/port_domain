package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"portdb.io/src/errors"

	"portdb.io/src/models"
)

var _ models.PortDomainRepository = &mysql{}

type sqlPort struct {
	ID          string `db:"id"`
	Code        string `db:"code"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Country     string `db:"country"`
	Alias       string `db:"alias"`
	Regions     string `db:"regions"`
	Coordinates string `db:"coordinates"`
	Province    string `db:"province"`
	Timezone    string `db:"timezone"`
	Unlocs      string `db:"unlocs"`
}

func (sql *sqlPort) Marshal(p *models.Port) error {
	sql.ID = p.ID
	sql.Name = p.Name
	sql.Code = p.Code
	sql.Country = p.Country
	sql.Timezone = p.Timezone
	sql.Province = p.Province
	sql.City = p.City
	data, err := json.Marshal(p.Alias)
	if err != nil {
		return fmt.Errorf("unable to marshal alias: %w", err)
	}
	sql.Alias = string(data)
	data, err = json.Marshal(p.Regions)
	if err != nil {
		return fmt.Errorf("unable to marshal regions: %w", err)
	}
	sql.Regions = string(data)
	data, err = json.Marshal(p.Unlocs)
	if err != nil {
		return fmt.Errorf("unable to marshal unlocs: %w", err)
	}
	sql.Unlocs = string(data)
	data, err = json.Marshal(p.Coordinates)
	if err != nil {
		return fmt.Errorf("unable to marshal coordinates: %w", err)
	}
	sql.Coordinates = string(data)
	return nil
}

func (sql *sqlPort) Unmarshal(p *models.Port) error {
	p.ID = sql.ID
	p.Name = sql.Name
	p.Code = sql.Code
	p.Country = sql.Country
	p.Timezone = sql.Timezone
	p.Province = sql.Province
	p.City = sql.City
	err := json.Unmarshal([]byte(sql.Alias), &p.Alias)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	err = json.Unmarshal([]byte(sql.Regions), &p.Regions)
	if err != nil {
		return fmt.Errorf("unable to unmarshal regions: %w", err)
	}
	err = json.Unmarshal([]byte(sql.Unlocs), &p.Unlocs)
	if err != nil {
		return fmt.Errorf("unable to unmarshal unlocs: %w", err)
	}
	err = json.Unmarshal([]byte(sql.Coordinates), &p.Coordinates)
	if err != nil {
		return fmt.Errorf("unable to unamrshal coordinates: %w", err)
	}
	return nil
}

type mysql struct {
	db *sqlx.DB
}

func (m *mysql) Store(p *models.Port) error {
	tx, err := m.db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("unable to start transaction. :%w", err)
	}
	var sql sqlPort
	if err := sql.Marshal(p); err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Query(portInserSQL,
		sql.ID,
		sql.Name,
		sql.City,
		sql.Country,
		sql.Alias,
		sql.Regions,
		sql.Coordinates,
		sql.Province,
		sql.Timezone,
		sql.Unlocs,
		sql.Code,
	)
	if err != nil {
		return fmt.Errorf("unable to perform insert port query: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("unable to commit transaction to database: %w", err)
	}
	return nil
}

func (m *mysql) Update(p *models.Port) error {
	tx, err := m.db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("unable to start transaction. :%w", err)
	}
	var sql sqlPort
	if err := sql.Marshal(p); err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Query(portUpdateSQL,
		sql.Name,
		sql.City,
		sql.Country,
		sql.Alias,
		sql.Regions,
		sql.Coordinates,
		sql.Province,
		sql.Timezone,
		sql.Unlocs,
		sql.Code,
	)
	if err != nil {
		return fmt.Errorf("unable to perform update port query: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("unable to commit transaction to database: %w", err)
	}
	return nil
}

func (m *mysql) Fetch(code string) (models.Port, error) {
	sqlP := new(sqlPort)
	if err := m.db.Get(sqlP, postSelectSQL, code); err != nil {
		if err == sql.ErrNoRows {
			return models.Port{}, errors.ErrNotFound
		}
		return models.Port{}, fmt.Errorf("unable to execute port select query: %w", err)
	}
	var p models.Port
	if err := sqlP.Unmarshal(&p); err != nil {
		return models.Port{}, err
	}
	return p, nil
}

func NewMysql(db *sqlx.DB) *mysql {
	return &mysql{db: db}
}
