git add .
msg=${1:-"fix: update file"}
git commit -m "${msg}"
git push
echo  "推送成功"

