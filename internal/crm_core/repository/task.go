package repository

import (
	"crm_system/internal/crm_core/entity"
	"github.com/gin-gonic/gin"
)

func (r *CRMSystemRepo) GetTasks(ctx *gin.Context) (*[]entity.Task, error) {
	var tasks *[]entity.Task

	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *CRMSystemRepo) GetTask(ctx *gin.Context, id string) (*entity.Task, error) {
	var task *entity.Task

	if err := r.DB.First(&task, id).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *CRMSystemRepo) CreateTask(ctx *gin.Context, task *entity.Task) error {
	if err := r.DB.Create(&task).Error; err != nil {
		return err
	}

	return nil
}

func (r *CRMSystemRepo) SaveTask(ctx *gin.Context, newTask *entity.Task) error {
	if err := r.DB.Save(&newTask).Error; err != nil {
		return err
	}

	return nil
}

func (r *CRMSystemRepo) DeleteTask(ctx *gin.Context, id string, task *entity.Task) error {
	if err := r.DB.Where("id = ?", id).Delete(task).Error; err != nil {
		return err
	}
	return nil
}
