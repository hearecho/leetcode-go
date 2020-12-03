# [498. 对角线遍历 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/diagonal-traverse/)

> 给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
>
> 输入:
> [
>  [ 1, 2, 3 ],
>  [ 4, 5, 6 ],
>  [ 7, 8, 9 ]
> ]
>
> 输出:  [1,2,4,7,5,3,6,8,9]

### 解题思路

> 沿着对角线向右上方走，沿途记录数据；在最后一个位置开始转向，按照顺时针方向测试右方、下方位置
> 沿着对角线向左下方走，沿途记录数据；在最后一个位置开始转向，按照逆时针方向测试下方、右方位置
>

```go
func findDiagonalOrder(matrix [][]int) []int {
     
    if len(matrix) == 0  {
        return nil
    }

    row, col := len(matrix), len(matrix[0])
    i, j := 0, 0
    res := []int{}

    for {
        // 记录数据
        res = append(res, matrix[i][j])

        // 向右上方走：行标减小，列标增大
        for i - 1 >= 0 && j + 1 < col {
            i--
            j++
            res = append(res, matrix[i][j])
        }
        
        // 顺时针转向：要么选右方，要么选下方；
        if j + 1 < col {
            j++
        } else if i + 1 < row {
            i++   
        } else {
            // 无节点可选
            break
        }
        
        // 记录数据
        res = append(res, matrix[i][j])
        
        //  向左下方走：行标增大，列标减小
        for i + 1 < row && j - 1 >= 0 {
            i++
            j--
            res = append(res, matrix[i][j])
        }

        // 逆时针转向：要么选下方，要么右方;
        if i + 1 < row {
            i++
        } else if j + 1 < col {
            j++ 
        } else {
            // 无节点可选
            break
        }

    }
    return res

}
```

