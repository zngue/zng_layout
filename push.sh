git add .
git commit -m "update"
git push
echo  "github 推送成功"
sleep 2
git push gitee-origin
echo  "gitee 推送成功"