package postgresql

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/user/internal/app/query"
	"github.com/evgsrkn/go-ddd-example/user/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type UserModel struct {
	id           string
	email        string
	username     string
	passwordHash string
	active       bool
	role         string
}

type PostgresqlRepository struct {
	db *pgxpool.Pool
}

func NewPostgresqlRepository(db *pgxpool.Pool) PostgresqlRepository {
	return PostgresqlRepository{
		db: db,
	}
}

func (r PostgresqlRepository) UserById(ctx context.Context, id string) (*query.User, error) {
	row := r.db.QueryRow(ctx, "SELECT * FROM users WHERE id=$1", id)

	return r.rowToUserQueryModel(row)
}

func (r PostgresqlRepository) AllUsers(ctx context.Context) ([]*query.User, error) {
	var users []*query.User

	rows, _ := r.db.Query(ctx, "SELECT * FROM users")
	defer rows.Close()

	for rows.Next() {
		usr, err := r.rowToUserQueryModel(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, usr)
	}

	return users, nil
}

func (r PostgresqlRepository) CreateUser(ctx context.Context, user domain.User) error {
	_, err := r.db.Exec(
		ctx,
		"INSERT INTO users(id, username, email, passwordHash, active, role) values($1, $2, $3, $4, $5, $6)",
	)
	if err != nil {
		return errors.Wrap(err, "unable to insert into table")
	}

	return nil
}

func (r PostgresqlRepository) UpdateUser(
	ctx context.Context,
	id string,
	updateFn func(ctx context.Context, user *domain.User) (*domain.User, error),
) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "cannot begin db transaction")
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	row := tx.QueryRow(ctx, "SELECT * FROM users WHERE id=$1", id)
	user, err := r.rowToUserDomain(row)
	if err != nil {
		return err
	}

	updatedUser, err := updateFn(ctx, user)
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		ctx,
		`UPDATE users
        SET username=$2, email=$3, passwordHash=$4, active=$5, role=$6
        WHERE id=$1`,
		id,
		updatedUser.Username(),
		updatedUser.Email(),
		updatedUser.PasswordHash(),
		updatedUser.IsActive(),
		updatedUser.Role(),
	); err != nil {
		return err
	}

	return nil
}

func (r PostgresqlRepository) rowToUserDomain(row pgx.Row) (*domain.User, error) {
	var user UserModel

	if err := row.Scan(
		&user.id,
		&user.email,
		&user.username,
		&user.passwordHash,
		&user.active,
		&user.role,
	); err != nil {
		return nil, errors.Wrap(err, "unable to scan db row")
	}

	return domain.UnmarshalUserFromDatabase(
		user.id,
		user.email,
		user.username,
		user.passwordHash,
		user.active,
		user.role,
	)
}

func (r PostgresqlRepository) rowToUserQueryModel(row pgx.Row) (*query.User, error) {
	var user query.User

	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.PasswordHash,
		&user.Active,
		&user.Role,
	)

	if err != nil {
		return nil, errors.Wrap(err, "unable to scan db row")
	}

	return &user, nil
}
