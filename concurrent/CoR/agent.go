package main

import (
	"context"
	"fmt"
)

// CreateAgent 处理者接口
type CreateAgent interface {
	execute(*AgentRequester) error
	setNext(CreateAgent)
}

// AgentRequester 一个Agent 请求分为三步: 创建PVC, 创建DaemonSet, 创建Service
type AgentRequester struct {
}

// CreateAgentPvc 创建Agent PVC
type CreateAgentPvc struct {
	ctx    context.Context
	client string
	next   CreateAgent
}

func (d *CreateAgentPvc) execute(r *AgentRequester) error {
	// 执行创建pvc的逻辑

	fmt.Println("创建Pvc")
	fmt.Println("如果创建Pvc失败,则直接返回")
	d.next.execute(r)
	return nil
}

func (d *CreateAgentPvc) setNext(next CreateAgent) {
	d.next = next
}

type CreateAgentDaemonSet struct {
	ctx    context.Context
	client string
	next   CreateAgent
}

func (d *CreateAgentDaemonSet) execute(r *AgentRequester) error {
	// 执行创建pvc的逻辑

	fmt.Println("创建DaemonSet")
	fmt.Println("如果创建DaemonSet失败,删除Pvc")
	d.next.execute(r)
	return nil
}

func (d *CreateAgentDaemonSet) setNext(next CreateAgent) {
	d.next = next
}

type CreateAgentSvc struct {
	ctx    context.Context
	client string
	next   CreateAgent
}

func (d *CreateAgentSvc) execute(r *AgentRequester) error {
	// 执行创建pvc的逻辑
	fmt.Println("创建Svc")

	fmt.Println("如果创建svc失败,删除DaemonSet及Pvc")
	return nil
}

func (d *CreateAgentSvc) setNext(next CreateAgent) {
	d.next = next
}

func main() {

	createSvc := &CreateAgentSvc{}

	createDaemonSet := &CreateAgentDaemonSet{}
	createDaemonSet.setNext(createSvc)

	createPvc := &CreateAgentPvc{}
	createPvc.setNext(createDaemonSet)

	r := &AgentRequester{}
	createPvc.execute(r)
}
