package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB // 已有连接池，无需维持，调用方法查询即可

func initDb() error {
	var err error
	dsn := "root:root@tcp(localhost:3306)/golang_db"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100) // 设置最大连接数
	DB.SetMaxIdleConns(16)  // 设置最大空闲连接
	return nil
}

type User struct {
	Id   int64          `db:"id"`
	Name sql.NullString `db:"string"` //sql.NullString带空字符串的类型
	Age  int            `db:"age"`
}

func testQueryMultilRow() {
	sqlstr := "select id, name, age from user where id > ?"
	rows, err := DB.Query(sqlstr, 0)
	//重点关注， rows对象一定要close掉
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	for rows.Next() { // 遍历所有结果
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age) // 把查询到的结果放入结构体，这里的结构体要注意要用映射，scan也是连接池释放的操作
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}

}

func testQueryData() {
	for i := 0; i < 101; i++ {
		fmt.Printf("query %d times\n", i)
		sqlstr := "select id, name, age from user where id=?"
		row := DB.QueryRow(sqlstr, 2)
		/*if row != nil {
			continue
		}*/
		var user User
		err := row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}

		fmt.Printf("id:%d name:%v age:%d\n", user.Id, user.Name, user.Age)
	}

}

func testInsertData() {
	sqlstr := "insert into user(name, age) values(?, ?)" // 插入语句
	result, err := DB.Exec(sqlstr, "tom", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	id, err := result.LastInsertId() // 对查询到的结果进行提取
	if err != nil {
		fmt.Printf("get last insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

func testUpdateData() { // 更新
	sqlstr := "update user set name=? where id=?"
	result, err := DB.Exec(sqlstr, "jim", 3)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	affected, err := result.RowsAffected() // 更改的行数据
	if err != nil {
		fmt.Printf("get affected rows failed, err:%v\n", err)
	}
	fmt.Printf("update db succ, affected rows:%d\n", affected)
}

func testDeleteData() {
	sqlstr := "delete from user where id=?"
	result, err := DB.Exec(sqlstr, 3)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows failed, err:%v\n", err)
	}
	fmt.Printf("delete db succ, affected rows:%d\n", affected)
}

func testPrepareData() {
	sqlstr := "select id, name, age from user where id > ?"
	stmt, err := DB.Prepare(sqlstr) // 命令传入
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	rows, err := stmt.Query(0) // 数据传入
	//重点关注， rows对象一定要close掉
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}

func testPrepareInsertData() {
	sqlstr := "insert into user(name, age) values(?, ?)"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	result, err := stmt.Exec("jim", 100)
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

func testTrans() {

	conn, err := DB.Begin() // 开启一个事务(连接)
	if err != nil {
		if conn != nil {
			conn.Rollback() // 开启失败进行事务回滚
		}
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}

	sqlstr := "update user set age = 1 where id = ?"
	_, err = conn.Exec(sqlstr, 1) //这里exec方法是来自con（开启的事务对象的）类似于预处理中对象是smt,一般情况是Db对象
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlstr, err)
		return
	}

	sqlstr = "update user set age = 2 where id = ?"
	_, err = conn.Exec(sqlstr, 2)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec second sql:%s failed, err:%v\n", sqlstr, err)
		return
	}
	err = conn.Commit() // 最后提交事务，确定事务结束
	if err != nil {
		fmt.Printf("commit failed, err:%v\n", err)
		conn.Rollback()
		return
	}
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}

	//testQueryData()
	//testQueryMultilRow()
	//testInsertData()
	//testUpdateData()
	//testDeleteData()
	//testPrepareData()
	//testPrepareInsertData()
	testTrans()
}
