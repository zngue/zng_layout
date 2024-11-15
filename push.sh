git add .
git commit -m "更新模板文件的默认数据"
git push
echo  "github 推送成功"
sleep 2
git push gitee-origin
echo  "gitee 推送成功"