package models

type Workspace struct {
	Id int64
	Name string
	CreatedTime int64
}

func SaveWorkspace(w *Workspace) (int64, error) {
	ret, err := DB.Exec("INSERT INTO workspaces (name, createdTime) VALUES (?, ?)",
		w.Name, w.CreatedTime)
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}

func DeleteWorkspace(id int64) error {
	_, err := DB.Exec("DELETE FROM workspaces WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}