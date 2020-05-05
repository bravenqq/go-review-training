package domain

//ProcessControlBlock 进程控制块
type ProcessControlBlock struct {
	id    int
	state int
	ProgramCounter
	RegisterInformation
	ScheduleInformation
	AccountingInformation
	Memory
	IOStatusInformation
}

//ProgramCounter 程序计数器
type ProgramCounter struct {
}

//RegisterInformation 寄存器信息
type RegisterInformation struct {
}

//ScheduleInformation 调度信息
type ScheduleInformation struct {
}

//AccountingInformation 记账信息
type AccountingInformation struct {
}

//Memory 内存信息
type Memory struct {
}

//IOStatusInformation 状态信息
type IOStatusInformation struct {
}
