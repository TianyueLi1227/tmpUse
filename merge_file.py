import pandas as pd


def merge_csv(part1_path, part2_path, output_file):
    # 读取两部分文件
    part1 = pd.read_csv(part1_path)
    part2 = pd.read_csv(part2_path)

    # 合并数据
    merged_data = pd.concat([part1, part2])

    # 保存合并后的文件
    merged_data.to_csv(output_file, index=False)

    print("Files have been merged and saved.")


merge_csv("837500-ord_transfers_part1.csv","837500-ord_transfers_part2.csv", "837500-ord_transfers.csv")
