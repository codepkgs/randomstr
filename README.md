# randomstr

> 产生指定长度的指定字符集的随机字符串

* 下载
    ```bash
    go get github.com/codepkgs/randomstr
    ```

* 导入
    ```go
    import "github.com/codepkgs/randomstr"
    ```

* 使用
  ```go
  package main
  
  import (
      "fmt"
      "github.com/codepkgs/randomstr"
  )
  
  func main() {
      rs := randomstr.New(randomstr.WithDigit(), randomstr.WithLetter())
      s1 := rs.Random(6)
      s2 := rs.Random(12)
      fmt.Println(s1)
      fmt.Println(s2)
  }
  ```