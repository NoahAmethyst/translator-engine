# API集成翻译引擎

已接入 有道翻译、腾讯翻译、阿里翻译、百度翻译、火山翻译

#### 有道，阿里支持领域型翻译（专业化），需要开通对应服务

如果指定的翻译模型不支持则会使用普通翻译模型

#### 有道支持专业化列表
```go

        // Medicine 医疗
	Medicine Scene = iota + 1

	// Computers 计算机
	Computers

	// Finance 金融
	Finance
```


#### 阿里支持专业化列表
```go

        // Medicine 医疗
	Medicine Scene = iota + 1

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

```

### 使用

#### 引入依赖
```shell
go get github.com/NoahAmethyst/translator-engine
```

#### 使用示例代码
<details>
<summary>有道</summary>

```go

import (
"github.com/NoahAmethyst/translator-engine"
"os"
"testing"
)

//一般文本翻译
func TestYoudaoTrans(t *testing.T) {
        
    from := translator_engine.AUTO
    to := translator_engine.EN
    src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
    
    apiKey := os.Getenv("API_KEY")
    secretKey := os.Getenv("SECRET_KEY")
    
    cli := translator.EngFactory.BuildYoudaoEng(apiKey, secretKey)
    
    result, err := cli.TransText("This is a test,for testing translation.", translator.AUTO, translator.ZH)
    
    if err != nil {
    panic(err)
    }
    
    t.Logf("%+v", result)

}

//专业化文本翻译
func TestYoudaoTransWithScene(t *testing.T) {
    	
    from := translator_engine.AUTO
    to := translator_engine.EN
    src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
    
    apiKey := os.Getenv("API_KEY")
    secretKey := os.Getenv("SECRET_KEY")
    
    cli := translator_engine.EngFactory.BuildYoudaoEng(appKey, secretKey)
    
    //指定以金融专业模型翻译，如果失败则会使用一般翻译
    result, err := cli.TransText(src, from, to, translator_engine.Finance)
    
    if err != nil {
    panic(err)
    }
    
    t.Logf("youdao translate result:%+v", result)

}

```

</details>


<details>
<summary>腾讯</summary>

```go
/**
腾讯不支持专业化翻译
 */
import (
"github.com/NoahAmethyst/translator-engine"
"os"
"testing"
)

func TestTencentTrans(t *testing.T) {
    	
    from := translator_engine.AUTO
    to := translator_engine.EN
    src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
    
    secretId := os.Getenv("TC_SECRET_ID") 
    secretKey := os.Getenv("TC_SECRET_KEY")
    
    cli, _ := translator_engine.EngFactory.BuildTencentEng(secretId, secretKey)
    
    result, err := cli.TransText(src, from, to)
    
    if err != nil {
    panic(err)
    }
    
    t.Logf("%+v", result)

}

```
</details>


<details>
<summary>阿里</summary>

```go

import (
"github.com/NoahAmethyst/translator-engine"
"os"
"testing"
)

//一般文本翻译
func TestAliTrans(t *testing.T) {
    
    from := translator_engine.AUTO
    to := translator_engine.EN
    src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
    
    accessId := os.Getenv("ALI_ACCESS_ID")
    accessKey := os.Getenv("ALI_ACCESS_SECRET")
    
    cli, err := translator_engine.EngFactory.BuildAliEngine(accessId, accessKey)
    
    if err != nil {
    panic(err)
    }
    
    result, err := cli.TransText(src, from, to)
    
    if err != nil {
    panic(err)
    }
    
    t.Logf("ali translate result:%+v", result)

}

//专业化文本翻译
func TestAliTransWithScene(t *testing.T) {
    
    from := translator_engine.AUTO
    to := translator_engine.EN
    src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
    
    accessId := os.Getenv("ALI_ACCESS_ID")
    accessKey := os.Getenv("ALI_ACCESS_SECRET")
    
    cli, err := translator_engine.EngFactory.BuildAliEngine(id, secret)
    
    if err != nil {
    panic(err)
    }
    
    result, err := cli.TransText(src, from, to, translator_engine.Medicine)
    
    if err != nil {
    panic(err)
    }
    
    t.Logf("ali translate result:%+v", result)

}

```

</details>


<details>
<summary>百度</summary>

```go

/**
百度不支持专业化翻译
*/
import (
"github.com/NoahAmethyst/translator-engine"
"os"
"testing"
)
func TestBaiduTrans(t *testing.T) {

    from := translator_engine.AUTO
    to := translator_engine.EN
    src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
    
    apiKey := os.Getenv("BAIDU_API_KEY")
    secretKey := os.Getenv("BAIDU_SECRET_KEY")
    
    cli := translator_engine.EngFactory.BuildBaiduEng(apiKey, secretKey)
    result, err := cli.TransText(src, from, to)
    
    if err != nil {
    panic(err)
    }
    
    t.Logf("baidu translate result:%+v", result)

}

```

</details>


<details>
<summary>火山</summary>

```go

/**
火山专业化翻译需要联系客服
*/
import (
translator_engine "github.com/NoahAmethyst/translator-engine"
"os"
"testing"
)

func TestVolcTrans(t *testing.T) {
    from := translator_engine.AUTO
    to := translator_engine.EN
    src := "这是一段用来测试的文本，它是中文的，将要翻译成英文"
    
    appKey := os.Getenv("VOLC_ACCESS_KEY")
	secretKey := os.Getenv("VOLC_SECRET_KEY")
    cli := translator_engine.EngFactory.BuildVolcEngine(appKey, secretKey)
    result, err := cli.TransText(src, from, to)
    if err != nil {
    panic(err)
    }
    t.Logf("volc translate result:%+v", result)
    
    to = translator_engine.ZH
    src = result.Dst
	
    result, err = cli.TransText(src, from, to)
	
    if err != nil {
    panic(err)
    }
	
    t.Logf("volc translate result:%+v", result)
    
    }

}

```

</details>


<details>
<summary>通用型</summary>

```go

import (
	"github.com/NoahAmethyst/translator-engine"
	"testing"
)

func TestIBaiduTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
        apiKey := os.Getenv("BAIDU_API_KEY")
        secretKey := os.Getenv("BAIDU_SECRET_KEY")
	baiduEng := translator_engine.EngFactory.BuildBaiduEng(apiKey, secretKey)
	baiduResult, err := translator_engine.TransText(src, from, to, baiduEng)
	if err != nil {
		panic(err)
	}
	t.Logf("baidu trans result:%+v", baiduResult)
}

func TestIYoudaoTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
	
        apiKey := os.Getenv("API_KEY")
        secretKey := os.Getenv("SECRET_KEY")

	youdaoEng := translator_engine.EngFactory.BuildYoudaoEng(appKey, secretKey)
	youdaoResult, err := translator_engine.TransText(src, from, to, youdaoEng, translator_engine.Finance)
	if err != nil {
		panic(err)
	}
	t.Logf("youdao trans result:%+v", youdaoResult)
}

func TestITencentTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
        secretId := os.Getenv("TC_SECRET_ID")
        secretKey := os.Getenv("TC_SECRET_KEY")
	tencentEng, err := translator_engine.EngFactory.BuildTencentEng(secretId, secretKey)
	if err != nil {
		panic(err)
	}
	tencentResult, err := translator_engine.TransText(src, from, to, tencentEng)
	if err != nil {
		panic(err)
	}
	t.Logf("tencent trans result:%+v", tencentResult)
}

func TestIAliTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
	
        accessId := os.Getenv("ALI_ACCESS_ID")
        accessKey := os.Getenv("ALI_ACCESS_SECRET")
	
	aliEng, err := translator_engine.EngFactory.BuildAliEngine(accessId, accessKey)
	if err != nil {
		panic(err)
	}

	aliResult, err := translator_engine.TransText(src, from, to, aliEng, translator_engine.Finance)
	if err != nil {
		panic(err)
	}
	t.Logf("ali trans result:%+v", aliResult)
}

func TestIVolcTrans(t *testing.T) {
        from := translator_engine.AUTO
        to := translator_engine.EN
        src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
        appKey := os.Getenv("VOLC_ACCESS_KEY")
        secretKey := os.Getenv("VOLC_SECRET_KEY")
        volcEng := translator_engine.EngFactory.BuildVolcEngine(accessId, accessKey)
        
        volcResult, err := translator_engine.TransText(src, from, to, volcEng)
        if err != nil {
        panic(err)
        }
        t.Logf("ali trans result:%+v", volcResult)
}

```



</details>