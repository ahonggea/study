'''
Author: ahonggea🚧
Date: 2023-05-30 20:52:42
LastEditors: ahonggea🚧
LastEditTime: 2023-06-06 12:03:40
Description: file content
'''
# %%

# pip install requests
from  requests import get
# pip install lxml
from lxml import etree
# 国家网行政规划地址
DEFULT_BASE_URL='https://www.mca.gov.cn/mzsj/xzqh/2022/202201xzqh.html'
BASE_URL = input (f'请输入要抓取的国家网行政区划地址:\n默认:{DEFULT_BASE_URL}：') or  DEFULT_BASE_URL
# 腾讯地图
TENCENT_MAP_URL='https://apis.map.qq.com'
DEFAULT_TENCENT_MAP_KEY='XQCBZ-IG6CI-SHTGB-5CXJN-IVYRJ-WCFSJ'
TENCENT_MAP_KEY=input(f'请输入腾讯地图API的key:\n默认:{DEFAULT_TENCENT_MAP_KEY}') or DEFAULT_TENCENT_MAP_KEY

# 特殊市区编号
SPECIAL_CITY_CODE_PID={
    '4190':"411700",
    '4290':"422800",
    '4690':"460400",
    '5001':"500000",
    '5002':"500000",
    '6590':'654300'
}
headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36',
    # 'Referer': 'https://www.mca.gov.cn/article/sj/xzqh/1980/202304/20230400047341.shtml'
}
# %%
# 请求详细数据
def tencentMapCurl(url,params):
    params['key']=TENCENT_MAP_KEY
    req=get(TENCENT_MAP_URL+url,params)
    res = req.json()
    #print(req)
    m = res["result"]
    # g = f"{m['lng']},{m['lat']}"
    return m

def main():
    # 发起请求
    resp = get(BASE_URL, headers=headers)
    # %%
    # 查找值
    e = etree.HTML(resp.text,etree.HTMLParser())
    codes = e.xpath('/html/body/div/table/tr[*]/td[2]')
    names = e.xpath('/html/body/div/table/tr[*]/td[3]')
    if(len(codes)==len(names) and len(codes)>0):
        print(f'获取国家网数据[{BASE_URL}]成功')
    else:
        print(f'获取国家网数据[{BASE_URL}]失败：\ncode.len:{len(codes)}\nname.len:{len(names)}')
        return
    # 获取详细列表
    cityInfoLists= tencentMapCurl('/ws/district/v1/list',{})
    if(len(cityInfoLists)>0):
        print('获取腾讯详细城市信息成功')
    else:
        print('获取腾讯详细城市信息失败')
    # %% 
    cityDict={}
    for list in cityInfoLists :
        for index,item in enumerate(list):
            id = item['id']
            item['sort'] = index+1
            cityDict[id]=item
    # %%
    # 组装数据
    datas = []
    def pushData(code,pid,level):
        global datas
        if code not  in cityDict:
            print(f'未找到{code}的资料')
            return
        item = cityDict[code]
        fullname =  item['fullname']
        name=fullname
        if 'name' in item:
            name=item['name']
        datas.append({
            'CityCode':code,
            'CityName':name,
            'SimpleName':fullname,
            'Pid':pid,
            'Lng':item['location']['lng'],
            'Lat':item['location']['lat'],
            'Sort':item['sort'],
            'Level':level,
        })
    # %%
    # 处理国家网数据
    level_1_codes = []
    level_2_codes = []
    level_3_codes = []
    for code, name, idx in zip(codes, names, range(len(codes))):
        code =code.xpath("text()")
        if len(code)==0 :
            continue
        name =name.xpath("text()")
        if len(name)==0 :
            continue
        code=code[0]
        name=name[0]
        if not code.isdigit() :
            continue
        level = 1
        pid = 0
        if code.endswith("0000"):
            level = 1
            level_1_codes.append(code)
            pid = 0
        elif code.endswith("00"):
            level = 2
            pid = code[0:2]+"0000"
            level_2_codes.append(code)
        else:
            level = 3
            provinceCode=code[0:4]
            pid = provinceCode+"00"
            if provinceCode in SPECIAL_CITY_CODE_PID:
                pid = SPECIAL_CITY_CODE_PID[provinceCode]
            # 如果不在level2里面，一般为直辖市，为保证数据完整性，加回去
            if pid not in level_2_codes:
                ppid=  pid[0:2]+"0000"
                level_2_codes.append(pid)
                # 字典增加不存在城市
                cityDict[pid]=cityDict[ppid]
                pushData(code=pid,pid=ppid,level=2)
            level_3_codes.append(code)
        pushData(code=code,pid=pid,level=level)

    #%%
    # 数据写入文本
    from  csv import DictWriter
    student_header=['CityCode','CityName','SimpleName','Pid','Lng','Lat','Sort','Level']
    filename = 'data.csv'
    with open(filename, 'w',newline='') as file:
        print(f'数据写入文本[{filename}]')
        writer = DictWriter(file, fieldnames=student_header)
        writer.writeheader()
        writer.writerows(datas)

    # %%
