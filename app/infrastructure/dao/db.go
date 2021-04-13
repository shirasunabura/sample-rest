// Package dao ...
package dao

import (
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connection DB情報
type Connection struct {
	read  *gorm.DB
	write *gorm.DB
}

// location
const location = "Asia/Tokyo"

//Read 書き込み
func (con *Connection) Read() *gorm.DB {
	return con.read
}

//Write 書き込み
func (con *Connection) Write() *gorm.DB {
	return con.write
}

// Connect 接続
func (con *Connection) Connect() {
	var err error

	//TimeZoneをJCTに
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	//DB 接続
	con.read, err = gorm.Open(mysql.Open(os.Getenv("DB_READ_USER")+"@"+os.Getenv("READ_CONNECTION")), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err)
	}
	con.write, err = gorm.Open(mysql.Open(os.Getenv("DB_WRITE_USER")+"@"+os.Getenv("READ_CONNECTION")), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err)
	}
}

// CloseConn クローズ
func (con *Connection) CloseConn() {
	//con.read.Close()
	//con.write.Close()
}