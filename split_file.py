import pandas as pd

def split_csv(file_path, output_prefix):
    # 读取原始CSV文件
    data = pd.read_csv(file_path)

    # 计算分割点
    split_point = len(data) // 2

    # 分割数据为两部分
    first_half = data.iloc[:split_point]
    second_half = data.iloc[split_point:]

    # 保存两个文件
    first_half.to_csv(f"{output_prefix}_part1.csv", index=False)
    second_half.to_csv(f"{output_prefix}_part2.csv", index=False)

    print("Files have been split and saved.")

split_csv("837500-ord_transfers.csv", "837500-ord_transfers")
