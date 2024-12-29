# validator-gen
参数校验插件生成的代码
生成后的 pb/validator.go 文件
```go

    package pb
    
    import (
    	"github.com/go-playground/validator/v10"
    )
    
    var validate = validator.New()
```

生成后的 validator_req.go 文件
```go
    package pb
    
    // Validate req validator.
    func (r *HelloReq) Validate() error {
    	return validate.Struct(r)
    }
```
