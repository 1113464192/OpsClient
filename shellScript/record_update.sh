#!/bin/bash

##################################
# 该脚本用于更新与MD5校验        #
# 编写时间：2023-11-24 16:01     #
# 修改时间：2023-11-24 16:01     #
# 编写人：fanqihang              #
##################################

if [[ $# -ne 1 ]];then
    echo "文件MD5校验脚本传入参数不正确!"
    exit 1
fi

repo_dir=$1

# 设置并发线程数
THREAD_NUM=30

# 临时svn消息存储文件
svninfo_file=/tmp/svn_update.txt
svnerr_file=/tmp/svn_failure.log


#文件MD5校验
check_file_md5() {
    local file_line=$1
    if [[ $(md5sum ${file_line} | awk '{print $1}') != $(awk '{print $1}' ${file_line}.md5) ]];then
       echo "${file_line}" >> ${svnerr_file}
    fi
}


# 并发执行校验
function thread_check() {
    #并发函数
    check_file_md5() {
        if [[ $(md5sum $line | awk '{print $1}') != $(awk '{print $1}' ${line}.md5) ]];then
           echo "$line" >> ${svnerr_file}
        fi
    }
    
    # 创建一个命名管道，用于线程之间的通信
    tmp_fifofile="./${$}.fifo"
    mkfifo ${tmp_fifofile}
    # 使用 exec 命令以读写模式打开管道，并将描述符 6 与该管道关联
    exec 6<>${tmp_fifofile}
    # 删除临时创建的命名管道文件，因为已经通过描述符打开了它
    /bin/rm ${tmp_fifofile}
    # 初始化描述符 6,空行的作用是占据管道的空间，以控制并发线程的数量
    for ((i=0;i<${THREAD_NUM};i++));do
        echo >&6
    done
    
    while read line
    do
    {
        read -u6
        check_file_md5 "${line}"
        echo >&6
    }&
    done < <(awk 'toupper($1)!~/D|AT/{print $2}' ${svninfo_file} | awk -F '.md5' '{print $1}')
    wait # 等待所有的后台子进程结束
    exec 6>&- # 关闭df6
}

# 判断文件有无问题
function md5_check() {
    path=$1
    cd $path
    revision=$(tail -n 1 ${svninfo_file} | grep -o "[0-9]*")
    update_file_sum=$(awk 'toupper($1)!~/D|AT/{print $2}' ${svninfo_file} | grep -Evc ".*\.md5")
    update_md5_sum=$(awk 'toupper($1)!~/D|AT/{print $2}' ${svninfo_file} | grep -Ec ".*\.md5")
    if [[ $update_file_sum -ne $update_md5_sum ]];then
        echo "md5文件数量与更新文件数量不一致"
        exit 1
    fi
    if [[ $update_file_sum -eq $update_md5_sum && $update_file_sum -eq 0 ]];then
    echo "没有文件需要校验"
        exit 0
    fi
    # 执行多线程校验
    thread_check
}

# SVN UPDATE
function svn_update() {
    > ${svninfo_file}
    > ${svnerr_file}
    cd /data/game_repo/${repo_dir}/
    echo yes | svn update --username 'fanqihang' --password 'test' > ${svninfo_file}
    if [[ $(echo $?) -ne 0 ]];then
        echo "svn更新失败" > $svnerr_file
        echo "svn更新失败"
        exit 1
    fi
}

function main() {
    svn_update 
#    md5_check /data/game_repo/${repo_dir}/ebin
    if [[ -s ${svnerr_file} ]];then
        fail_file=$(uniq ${svnerr_file} | tr '\n' ' ')
        echo svn更新文件校验有误
        exit 1
    else
        echo "更新并校验成功"
        exit 0
    fi
}

main
