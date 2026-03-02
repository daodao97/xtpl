#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "使用方法: $0 <项目路径> <新模块名>"
    echo "例如: $0 . github.com/user/newproject"
    echo "或: $0 /path/to/project github.com/user/newproject"
    exit 1
fi

# 将相对路径转换为绝对路径
PROJECT_PATH=$(cd "$1" && pwd)
NEW_MOD_NAME="$2"
MOD_FILE="${PROJECT_PATH}/go.mod"

# 检查 go.mod 是否存在
if [ ! -f "$MOD_FILE" ]; then
    echo "错误: 在 $PROJECT_PATH 中找不到 go.mod 文件"
    exit 1
fi

# 获取项目目录名
PROJECT_NAME=$(basename "$PROJECT_PATH")
# 获取项目的父目录
PARENT_DIR=$(dirname "$PROJECT_PATH")
# 创建备份目录（使用项目名和时间戳）
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="${PARENT_DIR}/${PROJECT_NAME}_backup_${TIMESTAMP}"

# 检查目录是否已存在
if [ -d "$BACKUP_DIR" ]; then
    echo "错误: 备份目录已存在: $BACKUP_DIR"
    exit 1
fi

# 获取原始模块名
OLD_MOD_NAME=$(grep "^module" "$MOD_FILE" | awk '{print $2}')
if [ -z "$OLD_MOD_NAME" ]; then
    echo "错误: 无法从 go.mod 获取当前模块名"
    exit 1
fi

echo "项目路径: $PROJECT_PATH"
echo "当前模块名: $OLD_MOD_NAME"
echo "新模块名: $NEW_MOD_NAME"
echo "备份目录: $BACKUP_DIR"

# 创建完整的备份
echo "创建备份..."
cp -r "$PROJECT_PATH" "$BACKUP_DIR"

# 更新 go.mod
echo "更新 go.mod..."
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' "s|^module.*|module ${NEW_MOD_NAME}|" "$MOD_FILE"
else
    sed -i "s|^module.*|module ${NEW_MOD_NAME}|" "$MOD_FILE"
fi

# 更新所有 .go 文件中的 import 语句
echo "更新所有 .go 文件中的引用..."
find "$PROJECT_PATH" -name "*.go" -type f | while read -r file; do
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS 版本
        sed -i '' "s|\"${OLD_MOD_NAME}/|\"${NEW_MOD_NAME}/|g" "$file"
    else
        # Linux 版本
        sed -i "s|\"${OLD_MOD_NAME}/|\"${NEW_MOD_NAME}/|g" "$file"
    fi
done

# 运行 go mod tidy 整理依赖
echo "运行 go mod tidy..."
cd "$PROJECT_PATH" && go mod tidy

echo "完成！"
echo "项目已备份到: $BACKUP_DIR"
echo ""
echo "如需恢复，可以运行:"
echo "rm -rf \"$PROJECT_PATH\" && cp -r \"$BACKUP_DIR\" \"$PROJECT_PATH\""