#!/bin/bash

##################################
# 该脚本用于更新执行             #
# 编写时间：2023-12-07 15:50     #
# 修改时间：2023-12-07 15:50     #
# 编写人：fanqihang              #
##################################

if [[ $# -ne 1 ]];then
    echo "更新执行传入参数不正确!未接收到更新类型"
    exit 1
fi

function checkExecResult() {
    status=$1
    if [[ ${status} -ne 0 ]];then
        echo "传参无误，但更新操作执行失败，请联系运维检查"
        exit 1
    fi
}

update_type=$1

cd $(dirname $0)

case "${update_type}" in
1)
    bash stop_server.sh && bash start_server.sh
    checkExecResult $?
    ;;
2)
    bash gamectl hotup
    checkExecResult $?
    ;;
*)
    echo "输入参数类型不对，请找运维确认"
    exit 1
    ;;
esac

echo "更新操作执行成功"
exit 0

