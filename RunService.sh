#!/bin/bash
# 撰寫人員: Neil_Hsieh
# 撰寫日期：2019/01/14
# 說明： 啟動Golang的服務
#
# 備註：
#   
# 取 OS 系統
SYSTEM=$(uname)

# 執行專案的目錄
WORK_PATH=""
# 執行各容器，須掛載的資料夾位置
VOLUME_PATH=""
# 當前用戶名稱
WHOAMI=""
# 用戶專用名稱
USER_PATH=""
# 服務類型
SERVICE_LIST=('http' 'queue')
# 服務 ENV
SERVICE_ENV="local-docker"
# 服務名稱
SERVICE_NAME="all"
# Log 檔案名稱，不指定就吃 code 預設
LOG_NAME=""
# 啟動服務的指令
COMMAND=""

## 選擇環境
ChooseENV(){
    echo -e "\033[1;35m======================\n"

      # 取得全部目錄的字串
    envs=$(ls -l $WORK_PATH/env | awk '/^d/ {print $NF}' | grep -v "common")

    # 執行選項
    printf "\033[1;33m"
    echo -e "請選擇工具代碼：\n"

    # 顯示全部的資料夾清單 + 選擇專案
    select env in $envs
    do
        # 專案名稱
        SERVICE_ENV=$env
        break
    done

}

## 選擇啟動的服務類型
ChooseService(){
    echo -e "\033[1;35m======================\n"

    # 執行選項
    printf "\033[1;33m"
    echo -e "請選擇服務代碼：\n"

    # 顯示全部的資料夾清單 + 選擇專案
    select name in ${SERVICE_LIST[@]}
    do
    
        LOG_NAME=$name
        if [ "$name" = "http" ]
        then
            LOG_NAME=""
        fi
         
        # 專案名稱
        SERVICE_NAME=$name
        
        break
    done
}


# 執行 RunService.sh 的目錄(透過readlink 獲取執行腳本的絕對路徑，再透過dirname取出目錄)
if [ "$SYSTEM" = "Linux" ]
then
    WORK_PATH=$(dirname $(readlink -f $0))
    VOLUME_PATH=$(dirname $(greadlink -f $0))/storage
fi

# For Mac
if [ "$SYSTEM" = "Darwin" ]
then
    # 檢查指令是否存在，不存在直接安裝
    isExist=$(which greadlink)
    if [ -z $isExist ]
    then
        echo "指令不存在，開始安裝 greadlink"
        brew install coreutils
    fi
    WORK_PATH=$(dirname $(greadlink -f $0))
    VOLUME_PATH=$(dirname $(greadlink -f $0))/storage
    WHOAMI=$(whoami)
    USER_PATH="/Users/$WHOAMI"
fi

# 專案名稱(取當前資料夾路徑最後一個資料夾名稱)
PROJECT_NAME=${WORK_PATH##*/}
# 環境變數
ENV="local-docker"
# swagger path
SWAGGER_PATH="$GOPATH/pkg/mod"

# 本機開發須安裝swagger + 初始化文件
if [ ! -d "$SWAGGER_PATH" ]; then
    echo "===== Swagger not exist, prepare to install ===="
    go get -u github.com/swaggo/swag/cmd/swag
fi

cd $WORK_PATH
swag init

## air hot reload command init, only not exist
if ! [ -f "$WORK_PATH/.air.toml" ]; then
echo "$WORK_PATH/.air.toml"
    $GOPATH/bin/air init
fi

#############################
#############################
docker network ls | grep "web_service" >/dev/null 2>&1
    if  [ $? -ne 0 ]; then
        docker network create web_service
    fi


## 選擇 ENV
# ChooseENV

## 選擇啟動的服務類型
# ChooseService

## 自動下載不存在本地端的 package
# go mod download


## 設定 env 參數
echo "ENV=$ENV">.env
echo "PROJECT_NAME=$PROJECT_NAME">>.env
echo "SERVICE_NAME=$SERVICE_NAME">>.env
echo "LOG_NAME=$LOG_NAME">>.env

# 啟動容器服務
if [ "$ENV" = "local" ] ||  [ "$ENV" = "local-docker" ] 
then
    echo "COMMAND"="air -c .air.toml">>.env
else
    echo "COMMAND"="go run main.go">>.env
fi

echo "ENV=$SERVICE_ENV USER_PATH=$USER_PATH VOLUME_PATH=$VOLUME_PATH  docker-compose up -d"
ENV=$SERVICE_ENV USER_PATH=$USER_PATH VOLUME_PATH=$VOLUME_PATH  docker-compose up -d

# echo "ENV=$SERVICE_ENV USER_PATH=$USER_PATH VOLUME_PATH=$VOLUME_PATH  docker-compose up -d golang-order-$SERVICE_NAME"
# ENV=$SERVICE_ENV USER_PATH=$USER_PATH VOLUME_PATH=$VOLUME_PATH  docker-compose up -d golang-order-$SERVICE_NAME
