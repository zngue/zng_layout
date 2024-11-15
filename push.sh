git add .
git commit -m "更新模板文件的默认数据"
git push
echo  "github 推送成功"
sleep 2
git push gitee-origin
echo  "gitee 推送成功"

version=${1:-"v0.0.6"}
git tag -d "${version}"
git push origin :refs/tags/"${version}"
msg=${2:-"Release ${version}"}
git tag -a "${version}" -m "${msg}"
git push origin "${version}"
echo  "推送标签"
git push gitee-origin :refs/tags/"${version}"
git push gitee-origin "${version}"
echo  "推送标签到gitee"
# 删除本地tag
git tag -d "${version}"



