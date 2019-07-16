# goPath和goRoot

- 编辑~/.bash_profile文件，添加以下代码
```
export GOROOT=/usr/local/Cellar/go/1.10.3/libexec

export GOPATH=/Users/chenxingyi/work/go

export GOBIN=

export PATH=$PATH:${GOPATH//://bin:}/bin
```


- 编辑~/.zshrc文件，添加以下代码
```bash
export GOROOT=/usr/local/Cellar/go/1.10.3/libexec

export GOPATH=/Users/chenxingyi/work/go

export GOBIN=

export PATH=$PATH:${GOPATH//://bin:}/bin
```

source ~/.xxx 重启文件



### 参考文章
- [mac下设置gopath环境变量](https://www.jianshu.com/p/5c1873eaf3ca)


