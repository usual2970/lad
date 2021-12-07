# lad
一个敏感词过滤包，使用它你可以：
* 判断文本中是否包含敏感词。
* 查找文本中包含了哪些敏感词。
* 将查找到的敏感词替换为指定的字符串。

目前支持中英文，后续打算添加对多音词支持。

## 安装
```
go get -u github.com/usual2970/lad  
```

## 快速开始
添加模式串
```golang
machine := lad.New()
machine.Add("hello")
machine.Add("world")
machine.Build()
```

从文件中加载模式串，格式为一行一个词组
```
machine := lad.New()
if err := machine.Load("data path"); err != nil {
    t.Error(err)
}
machine.Build()
```

匹配字符串
```
machine.Match(text)
```

查找字符串
```
machine.Find(text)
```

替换字符串
```
machine.Replacee(text,target)
```

## Todos
- [x] 添加模式串,构建ac自动机
- [x] 从文件中加载模式串，构建ac自动机
- [x] 匹配字符串
- [x] 支持英文
- [ ] 支持多音字
- [ ] 词组支持


<hr>

Released under the [MIT License](LICENSE.txt).  
