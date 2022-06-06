package mysql
import (
	"database/sql"
	
)
type projetoModel struct{
  DB *sql.DB
}
func(m *projetoModel)Insert(title, content, expires string) (int, error){
  stmt := `INSERT INTO snippets (title, content, created, expires) 
            VALUES(?,?,UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

  result, err := m.DB.Exec(stmt, title, content, expires)
  if err != nil{
    return 0, err
  }

  id, err := result.LastInsertId()
  if err != nil{
    return 0, err
  }
  return int(id),nil
}
func(m *projetoModel) Get(id int)(*models.projeto, error){
  return nil, nil
}
func(m * projetoModel) Latest()([]*models.projeto, error){
  return nil, nil
}