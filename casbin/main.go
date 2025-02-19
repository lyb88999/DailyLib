package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		log.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		log.Printf("%s CAN NOT %s %s\n", sub, act, obj)
	}
}

func ACL() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("failed to new enforcer: %s", err)
	}
	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")
	check(e, "root", "data3", "read")
	check(e, "root", "data3", "write")
}

func RBAC() {
	e, err := casbin.NewEnforcer("./model_rbac.conf", "./policy_rbac.csv")
	if err != nil {
		log.Fatalf("failed to new enforcer: %s", err)
	}
	check(e, "dajun", "data", "read")
	check(e, "dajun", "data", "write")
	check(e, "lizi", "data", "read")
	check(e, "lizi", "data", "write")
}

func MRBAC() {
	e, err := casbin.NewEnforcer("./model_mrbac.conf", "./policy_mrbac.csv")
	if err != nil {
		log.Fatalf("failed to new enforcer: %s", err)
	}
	check(e, "dajun", "prod.data", "read")
	check(e, "dajun", "prod.data", "write")
	check(e, "lizi", "dev.data", "read")
	check(e, "lizi", "dev.data", "write")
	check(e, "lizi", "prod.data", "write")
}

func MLRBAC() {
	e, err := casbin.NewEnforcer("./model_mrbac.conf", "./policy_mlrbac.csv")
	if err != nil {
		log.Fatalf("failed to new enforcer: %s", err)
	}
	check(e, "dajun", "data", "read")
	check(e, "dajun", "data", "write")
	check(e, "lizi", "data", "read")
	check(e, "lizi", "data", "write")
}

func checkDom(e *casbin.Enforcer, sub, domain, obj, act string) {
	ok, _ := e.Enforce(sub, domain, obj, act)
	if ok {
		log.Printf("%s CAN %s %s in %s\n", sub, act, obj, domain)
	} else {
		log.Printf("%s CAN NOT %s %s in %s\n", sub, act, obj, domain)
	}
}
func RBACDOM() {
	e, err := casbin.NewEnforcer("./model_rbacdom.conf", "./policy_rbacdom.csv")
	if err != nil {
		log.Fatalf("failed to new enforcer: %s", err)
	}
	checkDom(e, "dajun", "tenant1", "data1", "read")
	checkDom(e, "dajun", "tenant2", "data2", "read")
}

type Object struct {
	Name  string
	Owner string
}

type Subject struct {
	Name string
	Hour int
}

func checkABAC(e *casbin.Enforcer, sub Subject, obj Object, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s at %d:00\n", sub.Name, act, obj.Name, sub.Hour)
	} else {
		fmt.Printf("%s CANNOT %s %s at %d:00\n", sub.Name, act, obj.Name, sub.Hour)
	}
}
func ABAC() {
	e, err := casbin.NewEnforcer("./model_abac.conf")
	if err != nil {
		log.Fatalf("failed to new enforcer: %s", err)
	}
	o := Object{
		Name:  "data",
		Owner: "dajun",
	}
	s := Subject{
		Name: "dajun",
		Hour: 10,
	}
	checkABAC(e, s, o, "read")
	s2 := Subject{
		Name: "lizi",
		Hour: 10,
	}
	checkABAC(e, s2, o, "read")
	s3 := Subject{
		Name: "dajun",
		Hour: 20,
	}
	checkABAC(e, s3, o, "read")
	s4 := Subject{
		Name: "lizi",
		Hour: 20,
	}
	checkABAC(e, s4, o, "read")
}

func RBACGorm() {
	a, err := gormadapter.NewAdapter("mysql", "root:Lyb1217..@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Printf("failed to new adapter from mysql: %s", err)
	}
	e, _ := casbin.NewEnforcer("./model.conf", a)
	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")
}

func main() {
	ACL()
	RBAC()
	MRBAC()
	MLRBAC()
	RBACDOM()

	ABAC()

	RBACGorm()

}
