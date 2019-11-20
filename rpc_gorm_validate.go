package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"git.txuntest.com/meeline/protobuf/backend"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"gopkg.in/go-playground/validator.v9"
)

func TestAddRole(t *testing.T) {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	if err != nil {
		fmt.Println("-- did not connect:", err.Error())
		return
	}

	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	cli := backend.NewBackendClient(conn)

	role := new(backend.Role)
	role.Id = 1
	role.Name = "tom0005"
	role.Level = "1"
	role.State = 1

	// 调用gRPC接口
	rst, err := cli.AddRole(context.Background(), role)
	if err != nil {
		fmt.Println(rst)
		log.Fatalf("[请求失败] 原因:%s 错误码:%d\n", rst.GetErrStr(), rst.GetErrCode())
	}

	fmt.Println(rst)
	fmt.Printf("[请求成功] 结果:%v code:%v\n", rst.GetErrStr(), rst.GetErrCode())
}

// 对应数据库表role的数据结构
type RoleModel struct {
	gorm.Model
	Name        string `validate:"eq=tom22223214234" gorm:"type:varchar”`
	Description string `gorm:"type:varchar”`
	State       int    `validate:"min=12,max=15" gorm:"type:integer” `
	Level       int    `gorm:"type:integer”`
}

func TestGorm(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/backend?charset=utf8&parseTime=true")

	defer db.Close()

	if err != nil {
		log.Fatalf("[打开数据库失败]")
		return
	}

	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxIdleConns(100)
	db.DB().Ping()

	var role RoleModel
	// create

	// update

	// r

	// d
	db.Table("role").Where("name = ?", "tom").First(&role)
	// fmt.Println(role)
	fmt.Println("ID:", role.Model.ID, "CreateAt:", role.Model.CreatedAt, " UpdateAt:", role.Model.UpdatedAt, "DeleteAt:", role.Model.DeletedAt)

	// validate
	validate := validator.New()
	err2 := validate.Struct(&role)
	fmt.Println("error:", err2)

	fmt.Println("Role:", role, " State:", role.State)

}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestCreateTable(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/backend?charset=utf8&parseTime=true")

	defer db.Close()

	if err != nil {
		log.Fatalf("[打开数据库失败]")
		return
	}

	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxIdleConns(100)
	db.DB().Ping()

	// 生成数据库表
	// db.AutoMigrate(&Product{})

	// add
	// 新增一条记录：将Product 对象转换成数据库内一条记录
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// 获取对象：将数据库内一条记录转换成 product 对象
	// query
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	fmt.Println(product)

	// 更新记录
	db.Model(&product).Update("Price", 3000)

	//db.First(&product, 1)
	//db.First(&product, "code = ?", "L1212") // find product with code l1212
	//fmt.Println(product)

	// product.Price = 3000

	var product2 Product
	product2.Price = 3000
	db.Delete(&product2)
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
	Age    int    `validate:"min=12,max=15"`
}

func TestValidate(t *testing.T) {
	address := &Address{
		Street: "",
		City:   "beijing",
		Planet: "Persphone",
		Phone:  "none",
		Age:    0,
	}
	validate := validator.New()
	err := validate.Struct(address)
	fmt.Println(err)
}
