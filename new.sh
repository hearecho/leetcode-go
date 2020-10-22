#! /bin/bash
## new solution make relation doc

# $0 $1 $$2 获取参数
# 三个参数 难度  问题属性 问题中文名称 问题名称

## 先创建problem中对应的文件夹
cd problem
if [ ! -d $2 ]; then
    mkdir $2
fi
cd $2
# 创建文件夹以及文件
name=$4
mkdir ${name^}
cd ${name^}
suffix=".go"
test_suffix="_test.go"
touch $4$suffix
touch $4$test_suffix
# 添加到Readme中
cd ../../..
echo -e "|       |[$3](/problem/$2/$4)|$2|$1|" >> README.md

## 生成docs中的文件
cd docs
echo -e "|       |[$3]($1/$4)|$2|$1|" >> README.md
cd $1
echo -e "|       |[$3]($1/$4)|$2|$1|" >> README.md
touch $4".md"



