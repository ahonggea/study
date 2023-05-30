'''
Author: ahonggea🚧
Date: 2023-05-30 20:52:42
LastEditors: ahonggea🚧
LastEditTime: 2023-05-31 00:23:13
Description: file content
'''
# %%

# pip install requests
import requests
# pip install lxml
from lxml import etree
BASE_URL = 'https://www.mca.gov.cn/article/sj/xzqh/2022/202201xzqh.html'
print(BASE_URL)

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36',
    'Referer': 'https://www.mca.gov.cn/article/sj/xzqh/1980/202304/20230400047341.shtml'
}
# 发起请求
resp = requests.get(BASE_URL, headers=headers)

# %%
e = etree.HTML(resp.text)

codes = e.xpath('//table/tr/td[2]/text()')
names = e.xpath('//table/tr/td[3]/text()')
print(codes, names)
# %%
# 分析
level_1 = 0
level_2 = 0
level_3 = 0
print(len(codes))
for code, name, idx in zip(codes, names, range(len(codes))):
    if (idx == 0):
        continue
    if (code.endswith("0000")):
        level_1 += 1
    elif (code.endswith("00")):
        print(code, name)
        level_2 += 1
    else:
        level_3 += 1
print(f'省：{level_1}\n市：{level_2}\n区：{level_3}')
# %%
