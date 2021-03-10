# statemachine
go 实现的状态机

stateless4j
核心模型
stateless4j是这三款状态机框架中最轻量简单的实现，来源自stateless(C#版本的FSM)

https://segmentfault.com/img/bVbGVMP

一个成熟的状态机引擎架构通常涵盖以上模型，各模型具体分工如下：
StateMachineFactory：状态机的构造工厂，帮助创建StateMachine实例。
TransitionConfigurer & StateConfigurer ：状态机构造相关的配置项，TransitionConfigurer是Transition的配置，StateConfigurer则是State的配置。
StateMachineExecutor： 状态机执行入口，主要作用是接收Event或相对的Trigger，启动状态机，执行状态机，停止状态机等。
StateMachine & Region：状态机模型的核心，最后的操作会落到这个类上面，Region是在StateMachine之上的一层抽象，包含了State和Transition两层概念，大部分的状态流转走的是Region的接口。
State & Event：状态和事件的抽象。通常会使用枚举作为状态和事件的具体实例。Event是整个状态机执行的触发源，几乎所有的状态流转都是通过Event来进行触发的
Transition & Trigger： 从某个状态流转到下一个状态的过程，狭义一点就是一条边，当然，在工程实现上还附带了此次状态流转需要附加的动作，比如Action之类的。Trigger是对Transition的一层封装，起到一个触发器的作用，属于可选的组件。
Action: 执行动作，一般在状态流转时触发。比如spring-sm就定义了多种触发场景，有entry,exit等。
StateMachineAccess：状态机的一些辅助工具，比如可以设置状态机的Interceptor，Monitor，或者是重置状态机等操作。
StateMachinePersist & Reposity：状态机相关的持久化支持。