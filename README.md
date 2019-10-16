# Jmeter-REST

用RESTful API给Jmeter套上一层壳，使其能通过http请求输入JMX文件和输出测试结果

## 使用方法

### 编译（可选）

```shell
docker run -it --rm -v "$(pwd):/app" yindaheng98/go-iris go build -v -o /app/PressureMeter
```

### 打包

```shell
docker build -t pressure_meter .
```

### 运行

```shell
docker run --rm -v "$(pwd)/Data:/jmeter/Data" -p 8080:8080 pressure_meter
```
