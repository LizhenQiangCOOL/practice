# 何为工厂？
工厂生产东西，在代码世界里工厂负责如何生产出对象。

# Go 的简单工厂模式
1. go 语言没有构造函数一说，所以一般会定义 NewXXX 函数来初始化相关类。
2. NewXXX 函数返回接口时就是简单工厂模式，也就是 Golang 的一般推荐做法就是简单工厂。

# Go 的 工厂方法模式
1. 定义了一个创建对象的借口，但由子类决定要实力化的类是那一个。工厂方法让类把实例化推迟到子类
2. Go 使用组合，而不是继承
3. 工厂方法模式 是简单工厂的进一步抽象和推广

product                 Creator
                        factoryMethod()
                        anOperation()
   ｜
concreteProduct  <-    concreteCreator
                        factoryMethod()


# 生活实例
1. 假如你有一家 pissa 店，需要创建出不同的 pissa 来满足客户
2. 经营得不错后，加盟 pissa 店，但每家 pissa 店地区不同，想提供不同口味的 pissa，想要 pissa 店的流程，又想加入自己的改良