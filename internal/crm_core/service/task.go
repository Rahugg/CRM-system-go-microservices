package service

import (
	"crm_system/internal/crm_core/entity"
	"github.com/google/uuid"
	"sort"
	"strconv"
)

func (s *Service) GetTasks(dealId, sortBy, sortOrder, stateInput string, user *entity.User) ([]map[string]interface{}, error) {
	tasks, err := s.Repo.GetTasksByDealId(dealId)
	if err != nil {
		return nil, err
	}

	dashboardTodos := map[string][]entity.TaskResult{
		"task":     nil,
		"planning": nil,
		"doing":    nil,
		"done":     nil,
	}

	stateTypes := []string{"task", "planning", "doing", "done"}
	identifier := map[string]string{
		"task":     "📄 Задача",
		"planning": "📝 Планируется",
		"doing":    "⚡ В процессе",
		"done":     "✅ Сделано",
	}

	dashboardTodos = appendTasks(tasks, user, dashboardTodos)
	var columns []map[string]interface{}
	for _, state := range stateTypes {
		column := map[string]interface{}{
			"id":    state,
			"name":  identifier[state],
			"tasks": dashboardTodos[state],
		}
		columns = append(columns, column)
	}

	if sortBy != "" {
		columns, err = s.sortTasks(columns, sortBy, sortOrder, stateInput)
		if err != nil {
			return nil, err
		}
	}

	if stateInput != "" && sortBy == "" && sortOrder == "" {
		column, err := s.filterTasksByStates(columns, stateInput)
		if err != nil {
			return nil, err
		}
		columns = nil
		columns = append(columns, column)
	}

	return columns, nil
}

type CompareFunc[T any] func(i, j T) bool

func SortByField[T any](items []T, less CompareFunc[T]) {
	sort.Slice(items, func(i, j int) bool {
		return less(items[i], items[j])
	})
}

func TaskComparison(i, j entity.TaskResult, sortBy string) bool {
	switch sortBy {
	case "name":
		return i.Task.Name < j.Task.Name
	case "description":
		return i.Task.Description < j.Task.Description
	default:
		return i.Task.DueDate.Before(j.Task.DueDate)
	}
}

func (s *Service) sortTasks(columns []map[string]interface{}, sortBy, sortOrder, state string) ([]map[string]interface{}, error) {
	var tasks []entity.TaskResult
	for _, column := range columns {
		if column["id"] == state {
			tasks = column["tasks"].([]entity.TaskResult)
			break
		}
	}

	less := func(i, j entity.TaskResult) bool {
		return TaskComparison(i, j, sortBy)
	}

	SortByField(tasks, less)

	if sortOrder == "desc" {
		sort.SliceStable(tasks, func(i, j int) bool {
			return !less(tasks[i], tasks[j])
		})
	}

	columns = getColumns(columns, sortBy, tasks)

	return columns, nil
}

func getColumns(columns []map[string]interface{}, sortBy string, tasks []entity.TaskResult) []map[string]interface{} {
	for i, column := range columns {
		if column["id"] == sortBy {
			columns[i]["tasks"] = tasks
			break
		}
	}
	return columns
}

func (s *Service) filterTasksByStates(columns []map[string]interface{}, state string) (map[string]interface{}, error) {
	var response map[string]interface{}
	for _, column := range columns {
		if column["id"] == state {
			response = column
			break
		}
	}
	return response, nil
}

func appendTasks(tasks []entity.Task, user *entity.User, dashboardTodos map[string][]entity.TaskResult) map[string][]entity.TaskResult {
	for i, task := range tasks {
		var taskResult entity.TaskResult
		taskResult.Task = tasks[i]
		taskResult.VoteCount = len(tasks[i].Votes)

		if taskResult.Task.AssignedTo == user.ID {
			taskResult.UserVoted = true
		} else {
			for _, vote := range tasks[i].Votes {
				if vote.SenderID == user.ID {
					taskResult.UserVoted = true
					break
				}
			}
		}

		dashboardTodos[task.State] = append(dashboardTodos[task.State], taskResult)
	}

	for _, task := range tasks {
		sort.Slice(dashboardTodos[task.State], func(i, j int) bool {
			return dashboardTodos[task.State][i].VoteCount > dashboardTodos[task.State][j].VoteCount
		})
	}
	return dashboardTodos
}

func (s *Service) GetTask(id string) (*entity.Task, error) {
	task, err := s.Repo.GetTask(id)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Service) CreateTask(task *entity.TaskInput) error {
	if err := s.Repo.CreateTask(task); err != nil {
		return err
	}

	return nil
}

func (s *Service) Vote(user *entity.User, voteInput *entity.VoteInput) error {
	if err := s.Repo.CreateVote(user, voteInput); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetChangesOfTask(id string) (*[]entity.TaskChanges, error) {
	todoChanges, err := s.Repo.GetChangesOfTask(id)
	if err != nil {
		return nil, err
	}
	return todoChanges, nil
}

func (s *Service) UpdateTask(newTask entity.TaskEditInput, id string, user *entity.User) error {
	task, err := s.Repo.GetTask(id)
	if err != nil {
		return err
	}

	var taskChanges entity.TaskChanges
	taskChanges.ManagerID = user.ID

	if newTask.Name != "" {
		taskChanges.TaskID = task.ID
		taskChanges.ChangedField = "name"
		taskChanges.OldValue = task.Name
		taskChanges.NewValue = newTask.Name
		task.Name = newTask.Name
	}
	if newTask.Description != "" {
		taskChanges.TaskID = task.ID
		taskChanges.ChangedField = "description"
		taskChanges.OldValue = task.Description
		taskChanges.NewValue = newTask.Description
		task.Description = newTask.Description
	}

	if !newTask.DueDate.IsZero() {
		taskChanges.TaskID = task.ID
		taskChanges.ChangedField = "due_date"
		taskChanges.OldValue = task.DueDate.String()
		taskChanges.NewValue = newTask.DueDate.String()
		task.DueDate = newTask.DueDate
	}

	if newTask.AssignedTo != uuid.Nil {
		taskChanges.TaskID = task.ID
		taskChanges.ChangedField = "assigned_to"
		taskChanges.OldValue = task.AssignedTo.String()
		taskChanges.NewValue = newTask.AssignedTo.String()
		task.AssignedTo = newTask.AssignedTo
	}

	if newTask.AssociatedDealID != 0 {
		taskChanges.TaskID = task.ID
		taskChanges.ChangedField = "associated_deal_id"
		taskChanges.OldValue = strconv.Itoa(int(task.AssociatedDealID))
		taskChanges.NewValue = strconv.Itoa(int(newTask.AssociatedDealID))
		task.AssociatedDealID = newTask.AssociatedDealID
	}

	if newTask.State != "" {
		taskChanges.TaskID = task.ID
		taskChanges.ChangedField = "state"
		taskChanges.OldValue = task.State
		taskChanges.NewValue = newTask.State
		task.State = newTask.State
	}

	if err = s.Repo.CreateTaskChanges(&taskChanges); err != nil {
		return err
	}

	if err = s.Repo.SaveTask(task); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTask(id string) error {
	company, err := s.Repo.GetTask(id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteTask(id, company); err != nil {
		return err
	}

	return nil
}
func (s *Service) SearchTask(query string) (*[]entity.Task, error) {
	tasks, err := s.Repo.SearchTask(query)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}
