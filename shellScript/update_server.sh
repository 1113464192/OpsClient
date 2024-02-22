#!/bin/bash

##################################
# 脚本用于推送svn仓库文件到单服  #
# 编写时间：2023-11-24 17:38     #
# 修改时间：2023-11-24 17:38     #
# 编写人：fanqihang              #
##################################

if [ ! $1 ];then
    echo "update脚本未接收到SVN库名"
    exit 1
fi

repo_dir=$1

cd $(dirname $0)


function update() {
    for ((i=1;i<=3;i++))
    do
        rsync -a --delete /data/game_repo/${repo_dir}/ebin $(pwd)/
        if [[ $(echo $?) -ne 0 ]];then
            echo "SVN仓库文件推送到 $(pwd) 单服失败"
            exit 1
        fi
    done
    echo "版本文件成功推送到 $(pwd) 单服"
    exit 0
}



function main() {
    update
}

main
