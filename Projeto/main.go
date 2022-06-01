package main 

import(
  "database/sql"
  "flag"
  "log"
  
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
  
}

func main () {
  dsn := flag.String("dsn",
    "Wpx1nTJtYB:XomKPIAhyi@tcp(remotemysql.com)/Wpx1nTJtYBparseTime=true",
                    "MySql DSN")
  flag.Parse()
  






  
}
