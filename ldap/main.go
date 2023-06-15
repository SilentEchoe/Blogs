package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func main() {
	l, err := ldap.Dial("tcp", "")
	if err != nil {
		fmt.Println("连接失败", err)

	}
	err = l.Bind("cn=admin,dc=example,dc=org", "adminpassword")
	if err != nil {
		fmt.Println("管理员认证失败", err)
	}
	fmt.Println("连接成功")

	// TODO 创建新用户
	addResponse := ldap.NewAddRequest("cn=users04,ou=users,dc=example,dc=org", []ldap.Control{})
	addResponse.Attribute("ou", []string{"test"})
	addResponse.Attribute("gidNumber", []string{"1004"})
	addResponse.Attribute("homeDirectory", []string{"/home/user04"})
	addResponse.Attribute("sn", []string{"Bar4"})
	addResponse.Attribute("uid", []string{"user04"})
	addResponse.Attribute("uidNumber", []string{"1004"})
	addResponse.Attribute("mail", []string{"user04@example.com"})
	addResponse.Attribute("objectClass", []string{"inetOrgPerson", "posixAccount", "shadowAccount"})
	err = l.Add(addResponse)
	if err != nil {
		fmt.Println("创建用户失败", err.Error())
	} else {
		fmt.Println("创建用户成功")
	}

}
