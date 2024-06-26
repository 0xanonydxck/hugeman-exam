package repo

import (
	"fmt"

	"github.com/dxckboi/hugeman-exam/internal/model"
	"github.com/dxckboi/hugeman-exam/pkg/errors"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *todoRepo {
	return &todoRepo{db}
}

func (tr *todoRepo) All(opts ...*AllTodoOption) ([]*model.Todo, error) {
	var opt *AllTodoOption
	if len(opts) == 0 {
		opt = &AllTodoOption{}
	} else {
		opt = opts[0]
	}

	tx := tr.db.Session(&gorm.Session{})
	if opt.Search != nil {
		statement := "title LIKE ? OR description LIKE ?"
		arg := "%" + *opt.Search + "%"
		tx = tx.Where(statement, arg, arg)
	}

	if opt.Sort != nil {
		sort := *opt.Sort
		if opt.Descend {
			sort = fmt.Sprintf("%s %s", sort, "DESC")
		}

		tx = tx.Order(sort)
	}

	var todos []*model.Todo
	if err := tx.Find(&todos).Error; err != nil {
		return nil, errors.ParsePostgresError(err)
	}

	return todos, nil
}

func (tr *todoRepo) Get(id uuid.UUID) (*model.Todo, error) {
	var todo model.Todo
	if err := tr.db.First(&todo, "id = ?", id).Error; err != nil {
		return nil, errors.ParsePostgresError(err)
	}

	return &todo, nil
}

func (tr *todoRepo) Create(todo *model.Todo) error {
	if err := tr.db.Clauses(clause.Returning{}).Omit("ID").Create(todo).Error; err != nil {
		return err
	}

	return nil
}

func (tr *todoRepo) Update(id uuid.UUID, todo *model.Todo) (*model.Todo, error) {
	updatedTodo := &model.Todo{}
	if err := tr.db.Model(updatedTodo).Clauses(clause.Returning{}).Where("id = ?", id).Updates(todo).Error; err != nil {
		return nil, errors.ParsePostgresError(err)
	}

	return updatedTodo, nil
}

func (tr *todoRepo) Delete(id uuid.UUID) error {
	if err := tr.db.Where("id = ?", id).Delete(&model.Todo{}).Error; err != nil {
		return errors.ParsePostgresError(err)
	}

	return nil
}
