package domain

type Task struct {
	title       string
	description string
	status      string
	isLocked    bool
	// todo: add method Set
	color string
}

func (task *Task) SetTitle(title string) {
	task.title = title
}

func (task *Task) SetDescription(description string) {
	task.description = description
}

func (task *Task) Lock(isLocked bool) {
	if !task.isLocked {
		task.isLocked = isLocked
	}

	// todo: error
}

func (task *Task) Unlock(isUnlocked bool) {
	if task.isLocked {
		task.isLocked = isUnlocked
	}

	// todo: error
}

func (task *Task) SetStatus(status string) {
	var canChangeStatus bool = task.canChangeTodoStatus(status) ||
		task.canChangeInProgressStatus(status) ||
		task.canChangeDoneStatus(status) ||
		task.canChangeCanceledStatus(status)

	if canChangeStatus {
		task.status = status
	}
}

func (task *Task) canChangeTodoStatus(status string) bool {
	if task.status == "TODO" {
		if status == "IN_PROGRESS" || status == "CANCELED" {
			return true
		}
	}

	// todo: add error
	return false
}

func (task *Task) canChangeInProgressStatus(status string) bool {
	if task.status == "IN_PROGRESS" {
		if status == "TODO" || status == "DONE" || status == "CANCELED" {
			return true
		}
	}

	// todo: add error
	return false
}

func (task *Task) canChangeDoneStatus(status string) bool {
	if task.status == "DONE" {
		return false
	}

	// todo: add error
	return false
}

func (task *Task) canChangeCanceledStatus(status string) bool {
	if task.status == "CANCELED" {
		return false
	}

	// todo: add error
	return false
}
