

### crud

```

最简单的 crud demo

gin
gorm + mysql
godotenv


后续计划补个 JWT

```



```
2022/09/01

问题 
1. 主键缺失（不规范）
2. 表名复数


gorm.Model 默认添加

type Model struct {
    ID        uint      `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt DeletedAt `gorm:"index"`
}

```


```

参考文档：
1. https://gin-gonic.com/zh-cn/docs/
2. https://gorm.io/zh_CN/docs/
3. https://www.kancloud.cn/shuangdeyu/gin_book

```