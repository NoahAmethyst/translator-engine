package translator_engine

type Scene int

const (
	// Medicine 医疗
	Medicine Scene = iota + 1

	// Computers 计算机
	Computers

	// Finance 金融
	Finance

	// Social 社交
	Social

	// Title 商品标题
	Title

	// Desc 商品描述
	Desc

	// Communicate 商品沟通
	Communicate
)
