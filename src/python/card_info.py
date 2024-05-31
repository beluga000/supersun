import requests
from bs4 import BeautifulSoup
from pymongo import MongoClient

numbers_with_quotes = [
     '10304', '10303','1530','10228','10333','10312','10141','10105','10305','10306'
]


headers = {'User-Agent' : 'Mozilla/5.0 (Windows NT 10.0; Win64; x64)AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36'}

#   #app > div.cardItem > div:nth-child(7) > div > h3
for cardAdId in numbers_with_quotes:
    data = requests.get(f'https://card-search.naver.com/item?cardAdId={cardAdId}',headers=headers)
    soup = BeautifulSoup(data.text, 'html.parser')
    title = soup.select('#app > div.cardItem > div.card_terms > div > h3')
    infos = soup.select('#app > div.cardItem > div.card_terms > div > div')
    titles = [title.text for title in title]
    infos = [info.text for info in infos]
    card_name = soup.select_one('#app > div > div.BaseInfo > div > div.cardname > b')
    card_name_text = card_name.text if card_name else "Unknown"
    title_list = soup.select('#app > div > div.BaseInfo_benefits > div > div > button > span')
    arr_title = [title.text for title in title_list]
    content_list = soup.select('#app > div.cardItem > div.Benefits > div > details > summary > h5 > i')
    arr_content = [content.text for content in content_list]
    benefits = [{'title': t, 'content': c} for t, c in zip(arr_title, arr_content)]
    info_detail =[{'title': t, 'info': i} for t, i in zip(titles, infos)]

    card_info = {
        'cardAdId': cardAdId,  
        'cardName': card_name_text,
        'benefit': benefits,
        'info_detail': info_detail
}

    client = MongoClient('mongodb://localhost:27017/')
    db = client['local']  # Replace with your database name
    collection = db['card_info']  # Replace with your collection name
    collection.insert_one(card_info)
    print(cardAdId,"Data inserted into MongoDB successfully")


# print(titles)
# print(infos)




# cardAdId = 반복문으로 돌리기

# for cardAdId in numbers_with_quotes:
#     data = requests.get(f'https://card-search.naver.com/item?cardAdId={cardAdId}',headers=headers)
#     soup = BeautifulSoup(data.text, 'html.parser')
#     card_name = soup.select_one('#app > div > div.BaseInfo > div > div.cardname > b')
#     card_name_text = card_name.text if card_name else "Unknown"
#     title_list = soup.select('#app > div > div.BaseInfo_benefits > div > div > button > span')
#     arr_title = [title.text for title in title_list]
#     content_list = soup.select('#app > div.cardItem > div.Benefits > div > details > summary > h5 > i')
#     arr_content = [content.text for content in content_list]
#     benefits = [{'title': t, 'content': c} for t, c in zip(arr_title, arr_content)]
#     card_info = {
#         'cardAdId': cardAdId,  
#         'cardName': card_name_text,
#         'benefit': benefits
#     }
#     client = MongoClient('mongodb://localhost:27017/')
#     db = client['local']  # Replace with your database name
#     collection = db['card_info']  # Replace with your collection name
#     collection.insert_one(card_info)
#     print("Data inserted into MongoDB successfully")

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
