#!/usr/bin/env bash
echo "Current Path: `pwd`"
echo "---------Begin----------"
if [[ $# -eq 0 ]]; then
    echo "⚠️WARNING: at least one param !"
    exit
fi

echo "[step 1/5: git pull]"
git pull
echo "[step 2/5: git add .]"
git add .
echo "[step 3/5: git status]"
git status
echo "[step 4/5: git commit -m $1]"
git commit -m "$1"
if [[ $# -eq 2 && ${#2} -le 2 ]]; then
    echo " 🐷 second param's length is too short !!!"
    exit
elif [[ $# -eq 2 && ${#2} -gt 2 ]] ; then
    echo "[step 5/5: git push -u origin $2]"
    git push -u origin $2
else
    echo "[step 5/5: git push]"
    git push
fi

#   $?   上一个命令的退出码
if [[ $? != 0 ]]; then echo " ❌ 提交出错了，请修正！" && exit ; fi


echo "---------End----------👌-"

