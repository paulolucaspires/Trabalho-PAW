package mysql
import (
	"database/sql"
	"github.com/paulolucaspires/Trabalho-PAW/projeto/models"
)
type formularioModel struct{
  DB *sql.DB
}
func(m *formularioModel)Insert(nome, email, mensagem string) (int, error){
  stmt := `INSERT INTO formulario (nome, email, mensagem) 
            VALUES(get nome, get email, get mensagem)`

  result, err := m.DB.Exec(stmt, nome, email, mensagem)
  if err != nil{
    return 0, err
  }

  nome, err := result.LastInsertnome()
  if err != nil{
    return 0, err
  }
  return int(nome),nil
}
func(m *formularioModel) Get(nome int)(*models.formulario, error){
  return nil, nil
}
func(m *formularioModel) mensagem([]*models.formulario, error){
  return nil, nil
}