docker-compose down
sleep 3
# 删除未使用的镜像
docker image prune -f
echo "清理完成"
git pull
sleep 3
docker-compose build
echo "构建完成"
sleep 3
docker-compose up -d
echo "部署完成"
docker builder prune -f

echo "缓存镜像清理完成"

docker system prune -f
echo "清理完成"