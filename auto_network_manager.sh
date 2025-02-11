#!/bin/bash

# 启用错误处理
set -e

# 日志文件路径
LOG_FILE="log/docker_auto_network.log"


#判断文件夹是否存在，不存在则创建
if [ ! -d "log" ]; then
  mkdir log
fi

# 需要连接的 overlay 网络名称
OVERLAY_NETWORK="zngue_overlay"

# 判断overlay网络是否存在，不存在则创建
if ! docker network inspect "$OVERLAY_NETWORK" >/dev/null 2>&1; then
  docker network create --driver overlay "$OVERLAY_NETWORK"
  log_message "Created overlay network $OVERLAY_NETWORK"
else
  log_message "Overlay network $OVERLAY_NETWORK already exists"
fi
# 记录日志函数
log_message() {
  local message="$1"
  echo "$(date +'%Y-%m-%d %H:%M:%S') - $message" | tee -a $LOG_FILE
}

# 监听 Docker 事件
# shellcheck disable=SC2162
docker events --filter 'event=start' --filter 'event=destroy' | while read event
do
  # 获取容器 ID
  CONTAINER_ID=$(echo "$event" | awk '{print $4}')

  # 获取事件类型
  EVENT_TYPE=$(echo "$event" | awk '{print $3}')

  # 获取容器名称
  CONTAINER_NAME=$(docker inspect --format '{{.Name}}' "$CONTAINER_ID" 2>/dev/null | sed 's/\///' || echo "Unknown")

  # 获取容器网络信息
  NETWORKS=$(docker inspect --format '{{json .NetworkSettings.Networks}}' "$CONTAINER_ID" 2>/dev/null || echo "{}")

  if [[ "$EVENT_TYPE" == "start" ]]; then
    # 1️⃣ 先检查容器是否已经在 overlay 网络中，防止重复添加
    if echo "$NETWORKS" | grep -q "\"$OVERLAY_NETWORK\""; then
      log_message "Container $CONTAINER_NAME ($CONTAINER_ID) is already in $OVERLAY_NETWORK, skipping..."
    else
      # 2️⃣ 如果容器是 bridge 网络，则加入 overlay 网络
      if echo "$NETWORKS" | grep -q '"bridge"'; then
        log_message "Container $CONTAINER_NAME ($CONTAINER_ID) started, joining overlay network..."
        if docker network connect $OVERLAY_NETWORK "$CONTAINER_ID"; then
          log_message "Successfully connected $CONTAINER_NAME to $OVERLAY_NETWORK"
        else
          log_message "Failed to connect $CONTAINER_NAME to $OVERLAY_NETWORK"
        fi
      fi
    fi
  fi

  if [[ "$EVENT_TYPE" == "destroy" ]]; then
    # 3️⃣ 先检查容器是否在 overlay 网络中，如果在则断开
    if echo "$NETWORKS" | grep -q "\"$OVERLAY_NETWORK\""; then
      log_message "Container $CONTAINER_NAME ($CONTAINER_ID) is being destroyed, disconnecting from overlay network..."
      if docker network disconnect $OVERLAY_NETWORK "$CONTAINER_ID"; then
        log_message "Successfully disconnected $CONTAINER_NAME from $OVERLAY_NETWORK"
      else
        log_message "Failed to disconnect $CONTAINER_NAME from $OVERLAY_NETWORK"
      fi
    else
      log_message "Container $CONTAINER_NAME ($CONTAINER_ID) is not in $OVERLAY_NETWORK, skipping..."
    fi
  fi
done
