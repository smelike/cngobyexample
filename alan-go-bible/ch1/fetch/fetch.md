fetch.go

这个程序从两个 package 中导入了函数，net/http 和 io/ioutil 包，http.Get 函数是粗昂见 HTTP 请求的函数，如果获取过程没有出错，那么会在 resp 这个结构体中得到访问的请求结果。

resp 的 Body 字段包括一个可读的服务器响应流。ioutil.ReadAll 函数从 response 中读取到全部内容；将其结果保存在变量 b 中。resp.Body.Close 关闭 resp 的 Body 流，防止资源泄露，Printf 函数会将结果 b 写出到标准输出流中。


练习：

1、函数调用 io.Copy(dst, src) 会从 src 中读取内容，并将读到的结果写入到 dst 中，使用这个函数替代掉例子中的 ioutil.ReadAll 来拷贝响应结构体到 os.Stdout，避免申请一个缓冲区（例子中的b）来存储。记得处理 io.Copy 返回结果中的错误。

2、修改 fetch 案例，如果输入的 url 参数没有 http:// 前缀的话，为这个 url 加上前缀。你可能会用到strings.HasPrefix 这个函数。

3、修改 fetch 打印出 HTTP 协议的状态码，可以从 resp.Status 变量得到该状态码。

