package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	demo "rpc_server/kitex_gen/demo"
	"strconv"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "student"
)

func QueryFromDatabase(id int32, student *demo.Student) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return err
	}

	query, err := database.Query("select * from baseinfo where id = " + strconv.Itoa(int(id)))
	if err != nil {
		log.Fatal(err)
		return err
	}

	for query.Next() {
		err = query.Scan(&student.Id, &student.Name, &student.College, &student.Sex)
	}

	query, err = database.Query("select * from email where id = " + strconv.Itoa(int(id)))

	var emails []string

	for query.Next() {
		var email string
		err = query.Scan(&student.Id, &email)
		emails = append(emails, email)
	}

	student.Email = emails

	err = query.Close()

	err = database.Close()

	return nil
}

func InsertIntoDatabase(student *demo.Student) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return err
	}

	tx, err := database.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}

	stmt, err := tx.Prepare("insert into baseinfo values ($1, $2, $3, $4)")
	if err != nil {
		fmt.Println("Statement Preparation Failed: ", err)
		return err
	}

	_, err = stmt.Exec(student.Id, student.Name, student.College, student.Sex)
	if err != nil {
		fmt.Println("Statement Execution Failed: ", err)
		return err
	}

	for i := 0; i < len(student.Email); i++ {
		stmt, err = tx.Prepare("insert into email values ($1, $2)")
		if err != nil {
			fmt.Println("Statement Preparation Failed: ", err)
			return err
		}

		_, err = stmt.Exec(student.Id, student.Email[i])
		if err != nil {
			fmt.Println("Statement Execution Failed: ", err)
			return err
		}
	}

	tx.Commit()
	return nil
}

var studentMap = make(map[int32]*demo.Student)

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	resp = demo.NewRegisterResp()
	var newStudent demo.Student
	err = QueryFromDatabase(student.Id, &newStudent)
	if err != nil {
		resp.Success = false
		resp.Message = "Internal Exception"
	}
	if newStudent.Id > 0 {
		resp.Success = false
		resp.Message = "Register Failed: Student Information Already Exists"
	} else {
		err = InsertIntoDatabase(student)
		if err != nil {
			resp.Success = false
			resp.Message = "Internal Exception"
		}
		resp.Success = true
		resp.Message = "Register Success"
	}
	fmt.Println(resp)
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	resp = demo.NewStudent()
	var oldStudent demo.Student
	if value, exist := studentMap[req.Id]; exist {
		fmt.Println("Use Cache")
		resp = value
		return
	} else {
		fmt.Println("Query Database")
		err = QueryFromDatabase(req.Id, &oldStudent)
		if err != nil {
			return
		}
		if oldStudent.Id == 0 {
			var student = demo.Student{
				Id:      0,
				Name:    "Student Not Exist",
				College: "Student Not Exist",
				Email:   nil,
			}
			resp = &student
		} else {
			resp = &oldStudent
			studentMap[req.Id] = &oldStudent
		}
		return
	}
}

// Port implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Port(ctx context.Context, portResp *demo.PortReq) (resp *demo.PortResp, err error) {
	resp = new(demo.PortResp)
	resp.Success = true
	resp.Message = strconv.Itoa(bindPort)
	return
}
