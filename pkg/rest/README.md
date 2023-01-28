### Sample usage

```
// test http
cli := rest.NewStdHttp(rest.Config{})
// cli := rest.NewFastHttp(rest.Config{})
code, resp, err := cli.Get(context.TODO(), rest.RestParam{
    Url: "https://httpbin.org/get",
    QueryParam: map[string]string{
        "param1": "testparam",
    },
    JsonBodyData: map[string]float64{
        "data1": 123,
        "data2": 33.123,
    },
    Header: map[string][]string{
        "Auth": {"Bearer asdasd", "wwwwaaa"},
    },
})
log.Println("code:", code)
log.Println("resp:", string(resp))
log.Println("err:", err)
code, resp, err = cli.Post(context.TODO(), rest.RestParam{
    Url: "https://httpbin.org/post",
    QueryParam: map[string]string{
        "param1": "testparam",
    },
    JsonBodyData: map[string]float64{
        "data1": 123,
        "data2": 33.123,
    },
    Header: map[string][]string{
        "Auth": {"Bearer asdasd", "wwwwaaa"},
    },
})
log.Println("code:", code)
log.Println("resp:", string(resp))
log.Println("err:", err)
```