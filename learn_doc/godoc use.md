# godoc使用教程

## 注释规范
* 注释符//后面要加空格, 例如: // xxx

        在注释符要缩进的话，第二行注释符后面的空格要比上一行的空格多一个
        example:
            // 123
            //  123
* 注释要紧跟package、const、type、func这些关键字，这样子才会被显示
  
        // 功能: 函数H的注释
        //  参  数:
        //     	h1	: 说明参数值
        //     	h2	: 运算的数字
        //  返回值:
        //     	reValue 	: 返回值
        
        //你会发现函数H的功能注释没有显示，这是因为功能函数注释与关键字func之间空了一行
        func H(h1 string, h2 string) (reValue string) {
            fmt.Println(h1, h2)
            return h1
        }
* package的注释最好不好超过3行，不过就算超过3行也没有关系。这里只是做一个规范而已。如果对于pkg描述的注释有很多行的时候，我们可以新建一个doc.go，这个文件用于对package包的描述
* 同一目录下的包可以由很多个文件组成，如果每个文件都有对package进行注释的话，godoc会自动将所有注释"**按照文件名的字母数序**"进行合并
        
        详细等会看实例文件        
* 在无效注释中以BUG(who)开头的注释, 将被识别为已知bug, 显示在bugs区域    

        // BUG(who): 因为前面有BUG(who)这个关键字，所以这句注释就算没有紧跟关键字不会被隐藏掉

## 在终端查看
在终端，进入要查看的包的目录`go doc`显示当前包的文档
这里查看的是包名的文档，以及所有的函数名字跟变量类型名

    example:
        go doc
        
    terminal show:
        package test // import "test/testdoc"
        这个是a.go的包名的描述


        At 20191017 by Baiyu
        
        ---a.go end---
        
        这个是在doc.go 文件中的内容
        
        在Go的源代码中
        
        在Mac系统中
        
        在类Unix系统中
        
        在Windows系统中
        
        ---doc.go end---
        
        这个是在testgodoc.go文件的内容 package Name: test
        
            describe:    程序的入口
        
        ---testgodoc.go end---
        
        这个是z.go 的包名描述
        
        ---z.go end---
        
        const Email ...
        const Baiyu = "baiyu"
        func H(h1 string, h2 string) (reValue string)
        func Test(param1 string, inter int) (k string)
        func Z()
        type Computer struct{ ... }
        type Person struct{ ... }
        
        BUG: 这个注释会生成在文档最后后面，同时因为紧跟着fun Z，所以在上面的func列表里面也有显示
        功能: 生成Z签名
        
        BUG: 因为前面有BUG(who)这个关键字，所以这句注释就算没有紧跟关键字不会被隐藏掉
        
        BUG: BUG(6):格式正确，所以这句注释就算没有紧跟关键字不会被隐藏掉，前面的BUG():、BUG6:、BUG:都是不正确的BUG(who)命名

`go doc <pkg>.<func>`查看pkg包下面的func函数的注释
    
    example:
        go doc test.Test
        
    terminal show:
         package test // import "test/testdoc"
         func Test(param1 string, inter int) (k string)
            功能: 测试函数
       
            参  数:
                param1  : 说明参数值
                inter   : 运算的数字
            返回值:
                k       : 返回值

如果想查看完整的源码`go doc -src <pkg>.<func>`

    example: 
        go doc -src test.Test 
        
    terminal show:
        package test // import "test/testdoc"
        // 功能: 测试函数
        //  参  数:
        //      param1  : 说明参数值
        //      inter   : 运算的数字
        //  返回值:
        //      k       : 返回值
        func Test(param1 string, inter int) (k string) {
                fmt.Println("测试函数")
                return param1
        }

## 在浏览器上查看
使用命令`godoc -http=:6060`
然后在打开浏览器的，在url输入`localhost:6060/pkg`或者`127.0.0.1:6060/pkg`就可以查看到你本地的所有包的信息了

如果你想要找到你特定的包名的话`localhost:6060/pkg/<path>/<packageName>`或者`127.0.0.1:6060/pkg/<path>/<packageName>`

`<path>`是你/src下包所在的路径

`<packageName>`是包名

## 导出godoc文档为HTML
使用命令`godoc -url "http://localhost:6060/pkg/<pkg>/<packageName>/" > <packageName>.html`

`<pkg>`是你包所在的文件夹

`<packageName>`是你的包名

导出来的是纯html文件，没有样式表，不过也有可能可以导出，但是我目前没有找到，如果有找到的朋友也可以交流一下

不过我也有把对应的样式文件给导出来了

下载地址`https://github.com/BaiYu96/golandLearn/tree/master/learn_doc/stylefile`

然后把html文件里面的css跟js指定的路径修改好就可以

## 教程示例文件
下载地址`https://github.com/BaiYu96/golandLearn/tree/master/learn_doc/test`

把test文件夹放到你`$GOPATH/src/`目录下