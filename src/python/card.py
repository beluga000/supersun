import requests
from bs4 import BeautifulSoup
from pymongo import MongoClient

numbers_with_quotes = [
    '3977', '1684', '1531', '10157', '10071', '3881', '10107', '10277', '3856', '10333', '10108', '10312', '277', '10276', '10102', 
    '10033', '10286', '1408', '10340', '1465', '10295', '10126', '10321', '1260', '10103', '10341', '593', '10342', '10284', '10256', 
    '10345', '10100', '10161', '10328', '10285', '10281', '10282', '10280', '10169', '3896', '3857', '10293', '4047', '10184', '10156', 
    '10162', '10311', '10279', '10307', '10168', '2225', '2348', '10237', '1695', '10166', '10246', '10167', '10229', '10146', '10165', 
    '10329', '10310', '10327', '10357', '10343', '2337', '1322', '10309', '1342', '10195', '10130', '10344', '10134', '1360', '10317', 
    '1340', '10193', '2776', '10360', '3837', '1203', '10035', '3878', '10354', '10268', '10349', '10338', '10298', '115', '3959', '1772', 
    '10267', '10347', '1692', '2005', '111', '10290', '1202', '10046', '3717', '10320', '10330', '1294', '2332', '10216', '1361', '2472', 
    '10302', '1715', '10332', '2612', '10301', '10315', '3836', '3638', '2916', '1768', '1451', '1030', '1316', '99', '10053', '10043', 
    '1362', '10214', '10220', '2857', '10252', '10215', '1470', '2856', '10221', '2716', '10204', '10039', '10294', '10355', '10352', 
    '2345', '2486', '2490', '10070', '1681', '2489', '2185', '10322', '10323', '10334', '10337', '10266', '2226', '1337', '1385', '10056', 
    '10076', '10351', '10336', '710', '10339', '2696', '10225', '2487', '10112', '10335', '10223', '10346', '10200', '10353', '10226', 
    '10135', '2717', '3076', '160', '2888', '10047', '10288', '10350', '10314', '10359', '10203', '10113', '10197', '10264', '10142', 
    '10313', '10219', '2778', '2343', '2392', '3796', '1614', '1698', '2779', '10361', '10362', '10363', '1680', '247', '10348', '10189', 
    '10066', '10227', '1384', '2676', '10160', '10097', '10316', '10271', '10186', '1400', '1388', '3798', '3797', '2418', '3936', '2533', 
    '10261', '1243', '2422', '3055', '3996', '1153', '10250', '10287', '2349', '3616', '10187', '3756', '10262', '10068', '1889', '3196', 
    '10199', '10164', '3258', '2797', '3777', '3776', '2423', '4044', '1591', '10283', '3716', '2419', '3958', '3799', '10233', '10356', 
    '3257', '10128', '10236', '10207', '10118', '652', '10318', '10297', '10289', '10325', '10098', '1571', '2206', '10235', '1570', '10291', 
    '2886', '3378', '3436', '870', '10143', '2615', '2885', '10248', '10151', '3237', '1597', '3556', '10265', '3251', '10057', '1573', '1292', 
    '10211', '1773', '10144', '10234', '10132', '10241', '3957', '10258', '10260', '10212', '10275', '2836', '2511', '10058', '3877', '10201', 
    '10218', '10185', '3636', '10331', '10181', '2532', '10138', '10096', '1158', '2350', '3337', '4048', '10029', '2976', '10049', '3816', 
    '10245', '3338', '4096', '10247', '10232', '3937', '10254', '10299', '1157', '10062', '10059', '205', '10249'
]


headers = {'User-Agent' : 'Mozilla/5.0 (Windows NT 10.0; Win64; x64)AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36'}
# cardAdId = 반복문으로 돌리기

for cardAdId in numbers_with_quotes:
    data = requests.get(f'https://card-search.naver.com/item?cardAdId={cardAdId}',headers=headers)
    soup = BeautifulSoup(data.text, 'html.parser')
    card_name = soup.select_one('#app > div > div.BaseInfo > div > div.cardname > b')
    card_name_text = card_name.text if card_name else "Unknown"
    title_list = soup.select('#app > div > div.BaseInfo_benefits > div > div > button > span')
    arr_title = [title.text for title in title_list]
    content_list = soup.select('#app > div.cardItem > div.Benefits > div > details > summary > h5 > i')
    arr_content = [content.text for content in content_list]
    benefits = [{'title': t, 'content': c} for t, c in zip(arr_title, arr_content)]
    card_info = {
        'cardAdId': cardAdId,  
        'cardName': card_name_text,
        'benefit': benefits
    }
    client = MongoClient('mongodb://localhost:27017/')
    db = client['local']  # Replace with your database name
    collection = db['card_info']  # Replace with your collection name
    collection.insert_one(card_info)
    print("Data inserted into MongoDB successfully")

# data = requests.get('https://card-search.naver.com/item?cardAdId=10304',headers=headers)


# soup = BeautifulSoup(data.text, 'html.parser')
 
# #   title = soup.select_one('#lst50 > td:nth-child(6) > div > div > div.ellipsis.rank01 > span > a ')

# card_name = soup.select_one('#app > div > div.BaseInfo > div > div.cardname > b')
# card_name_text = card_name.text if card_name else "Unknown"

# # #app > div > div.BaseInfo_benefits > div > div > button:nth-child(1) > span

# # print(title.text)

# title_list = soup.select('#app > div > div.BaseInfo_benefits > div > div > button > span')
# arr_title = [title.text for title in title_list]

# # for title in title_list:
# #     arr_title.append(title.text)
# #     print(title.text)

# content = soup.select_one('#app > div.cardItem > div.Benefits > div > details:nth-child(2) > summary > h5 > i ')

# # content_list = soup.select('#app > div.cardItem > div.Benefits > div > details > summary > h5 > i')

# # arr_content = []

# # for content in content_list:
# #     arr_content.append(content.text)
# #     print(content.text)

# content_list = soup.select('#app > div.cardItem > div.Benefits > div > details > summary > h5 > i')
# arr_content = [content.text for content in content_list]



# benefits = [{'title': t, 'content': c} for t, c in zip(arr_title, arr_content)]
# card_info = {
#     # cardAdId 반복문 돌리기
#     'cardAdId': 10304,  
#     'cardName': card_name_text,
#     'benefit': benefits
# }



# client = MongoClient('mongodb://localhost:27017/')
# db = client['local']  # Replace with your database name
# collection = db['card_info']  # Replace with your collection name
# collection.insert_one(card_info)

# print("Data inserted into MongoDB successfully")
