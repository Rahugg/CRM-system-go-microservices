package repository

import (
	"crm_system/internal/crm_core/entity"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
func (r *CRMSystemRepo) GetTasksByDealId(dealId string) ([]entity.Task, error) {
	var tasks []entity.Task
	if err := r.DB.Preload("Votes").Where("associated_deal_id = ?", dealId).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *CRMSystemRepo) CreateTask(ctx *gin.Context, newTask *entity.TaskInput) error {
	task := &entity.Task{
		Model:            gorm.Model{},
		Name:             newTask.Name,
		Description:      newTask.Description,
		DueDate:          newTask.DueDate,
		AssignedTo:       newTask.AssignedTo,
		AssociatedDealID: newTask.AssociatedDealID,
		State:            newTask.State,
	}
	if err := r.DB.Create(&task).Error; err != nil {
		return err
	}

	return nil
}
func (r *CRMSystemRepo) CreateVote(user *entity.User, voteInput *entity.VoteInput) error {
	vote := entity.Vote{
		TaskID:   voteInput.TaskID,
		SenderID: user.ID,
	}

	if err := r.DB.Create(&vote).Error; err != nil {
		return err
	}
	return nil
}
func (r *CRMSystemRepo) GetChangesOfTask(id string) (*[]entity.TaskChanges, error) {
	var taskChanges []entity.TaskChanges
	if err := r.DB.Where("task_id = ?", id).Find(&taskChanges).Error; err != nil {
		return nil, errors.New("could not get task history")
	}
	return &taskChanges, nil
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
