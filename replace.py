import os

file_name = "serverless.yml"

with open(file_name, encoding="cp932") as f:
    data_lines = f.read()

LINE_CHANNEL_SECRET = os.environ["LINE_CHANNEL_SECRET"]
LINE_ACCESS_TOKEN = os.environ["LINE_ACCESS_TOKEN"]
# 文字列置換
data_lines = data_lines.replace("LINE_CHANNEL_ACCESS_TOKEN_STRING",LINE_CHANNEL_SECRET)
data_lines = data_lines.replace("LINE_ACCESS_TOKEN_STRING",LINE_ACCESS_TOKEN)

# 同じファイル名で保存
with open(file_name, mode="w", encoding="cp932") as f:
    f.write(data_lines)