package folder

import (
	"context"
	"notes/internal/common"
)

func FindFolders(ctx context.Context) ([]Folder, error) {
	folders := make([]Folder, 0)

	rows, err := common.DbConn.QueryContext(ctx, `SELECT id, name FROM folders`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var folder Folder
		rows.Scan(&folder.Id, &folder.Name)
		folders = append(folders, folder)
	}

	return folders, nil
}
