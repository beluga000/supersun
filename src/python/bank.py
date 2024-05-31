import requests
from pymongo import MongoClient
from bs4 import BeautifulSoup


arr = [
  "ef452122331e3a65214cad7d7b0cb4b5",
  "8c94f9ea922d4f8bfe0a70e9644a3c28",
  "25b1b68afed295d85e699b05accb8299",
  "48d2d4435997c53223b9a42b33db5b8f",
  "40c16f367cc18c2aa9bf377504943035",
  "0d1f3fcf449b3ca792bb9a6dd8eaba80",
  "b90538d902264d7c3352e9be787ca614",
  "2fa3636506668558cd95a4ca6d284af3",
  "4beac9fd03a48a444c5280830375875d",
  "0013e8654e0946fb2d154586a76d5dba",
  "60a0ffc4b419868045d3b453bef19ae5",
  "9aa80041f7102e22eb45713323cb56bb",
  "085018690730ec19f3bca218eaa3315d",
  "02e9bed10fb1ba2e09032ad8f8738222",
  "03c34dc99e672ddf7b56840d4d0c0d73",
  "cd62415ff5e88d77a57fb9a1f3b8fb8b",
  "6b259be5a81013887b9251c8875c583d",
  "4130b4ea02aa3dcfb6fd59ba60bc1b4c",
  "6b31c9e8f869e48360fe904248bdcbb6",
  "7998d0d760d07b42b21e8a49cb70700e",
  "7f34789f963a2c63cb769cc3fc2ca14f",
  "7833e371d2fe5e468aa66ef3c591bff3",
  "2519e8d71028355bf3f86d16e4402b10",
  "2dc0ddb3f28a4aa9a049b603b8795bbc",
  "bf84a65719d1b8b59cf7ed7118d862db",
  "216993ddfc3742d2e7af333153b58b49",
  "3d39023112500f788ea4af0845e9f2cd",
  "99215abb36985cf74450d6f86e2db312",
  "f6b6e9378da703224ae6d3866b18e577",
  "dfb5d565b9e660b54ea633cea5a25980",
  "6ac95af42a38581a27a1717f6fdd79f5",
  "1703e2209e74e68ebded1877720a2e17",
  "3d66dde9d905adf75f69c74e1080acc0",
  "81ffa706dc3d3f187a4ad3c1f64cc3bc",
  "6fffbb362f5cf86fe766929d22538961",
  "794f6780cd9438c82e38de21833d4a20",
  "de3e03536a117c8fb5ca9eb001a6813a",
  "b589ee3de46e25526cc7d4c9e98b16f6",
  "c698396131cf1747d8d5abd42936c70b",
  "f4f5f78d771b5c65d5f6901abcbe6ece",
  "1509627bb67535418175a32cf4169afd",
  "b1b7132ee139e09168ecb16eaf35cf0c",
  "313fb0340d3e571c0b6ebaa3f276d877",
  "ea3a73ab6ca0fba8b6229d29a32514dc",
  "6bb99cf6cf07ece3944961baf38f34a5",
  "ad76c58829040284a93a9d513c01019d",
  "5546134db187cb6594e91de95f863db5",
  "3e869f906d6c8b7e248c7bea34c2e382"
]

# MongoDB setup
client = MongoClient('mongodb://localhost:27017/')
db = client["local"]
collection = db["parking_detail"]

for code in arr :
  headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36'
  }
  data = requests.get(f'https://new-m.pay.naver.com/savings/detail/{code}', headers=headers)
  soup = BeautifulSoup(data.text, 'html.parser')
  product_name = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > h3')
  bank_name = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > span.MainInfo_text__V_Fkd')
  sub_service = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > ul')
  arr_sub_service = [sub.text for sub in sub_service.find_all('li')] if sub_service else []
  max_rate = soup.select_one("#content > div.MainInfo_article__iQqCh > div.MainInfo_area-rates__ZlyVd > dl > div:nth-child(1) > dd")
  min_rate = soup.select_one("#content > div.MainInfo_article__iQqCh > div.MainInfo_area-rates__ZlyVd > dl > div:nth-child(2) > dd")
  maininfo_title = soup.select_one("#content > div.MainInfo_article__iQqCh > div.MainInfo_area-event__REIlH > strong")
  maininfo_content = soup.select_one("#content > div.MainInfo_article__iQqCh > div.MainInfo_area-event__REIlH > div")
  join_amount = soup.select_one("#PRODUCT_GUIDE > dl > div:nth-child(1) > dd")
  join_method = soup.select_one("#PRODUCT_GUIDE > dl > div:nth-child(2) > dd")
  join_target = soup.select_one("#PRODUCT_GUIDE > dl > div:nth-child(3) > dd")
  interest_payment = soup.select_one("#PRODUCT_GUIDE > dl > div:nth-child(4) > dd")
  join_note = soup.select_one("#PRODUCT_GUIDE > dl > div > dd.TextList_description__xhuFz.TextList_color-gray__ye_b3")
  join_protection = soup.select_one("#PRODUCT_GUIDE > dl > div:nth-child(6) > dd > p")
  data_to_insert = {
    "code": code,
    "product_name": product_name.text if product_name else None,
    "bank_name": bank_name.text if bank_name else None,
    "sub_service": arr_sub_service,
    "max_rate": max_rate.text if max_rate else None,
    "min_rate": min_rate.text if min_rate else None,
    "maininfo_title": maininfo_title.text if maininfo_title else None,
    "maininfo_content": maininfo_content.text if maininfo_content else None,
    "join_amount": join_amount.text if join_amount else None,
    "join_method": join_method.text if join_method else None,
    "join_target": join_target.text if join_target else None,
    "interest_payment": interest_payment.text if interest_payment else None,
    "join_note": join_note.text if join_note else None,
    "join_protection": join_protection.text if join_protection else None,
  }
  collection.insert_one(data_to_insert)

  print(code,"Data inserted successfully")




# for i in arr:
#     url = f"https://new-m.pay.naver.com/savings/_next/data/myONoeojN_eppLX19bh4D/detail/{i}.json?productCode={i}"
#     response = requests.get(url, headers=headers)
#     data = response.json()
#     collection.insert_one({'dehydratedState': data['pageProps']['dehydratedState'], 'code': i})
#     print(f"{i} inserted.")

# url = "https://new-m.pay.naver.com/savings/_next/data/myONoeojN_eppLX19bh4D/detail/75f7697bb28a2d35a6cc31be3a18d3da.json?productCode=75f7697bb28a2d35a6cc31be3a18d3da"
# response = requests.get(url, headers=headers)
# data = response.json()
# #   queries
# collection.insert_one({'dehydratedState': data['pageProps']['dehydratedState']})

# print(data)

# for offset in arr:
#     url = f"https://new-m.pay.naver.com/savings/_next/data/myONoeojN_eppLX19bh4D/detail/{offset}.json?productCode={offset}"
#     response = requests.get(url, headers=headers)
#     data = response.json()

#     # Step 2: Extract the 'products' list from the JSON data
#     products = data.get('result', {}).get('dehydratedState', [])

#     # Step 3: Connect to MongoDB
#     client = MongoClient('mongodb://localhost:27017/')
#     db = client['local']  # Database name
#     collection = db['deposit_detail']  # Collection name

#     # Step 4: Insert the data into MongoDB
#     if products:
#         collection.insert_many(products)
#         print(f"{len(products)} records inserted.")
#     else:
#         print("No data to insert.")

#     # Step 5: Close the MongoDB connection
#     client.close()

# url = "https://new-m.pay.naver.com/savings/api/v1/productList?productTypeCode=1003&companyGroupCode=BA&regionCode=00&offset=0&sortType=PRIME_INTEREST_RATE"

# response = requests.get(url, headers=headers)
# data = response.json()

# # Step 2: Extract the 'products' list from the JSON data
# products = data.get('result', {}).get('products', [])

# # Step 3: Connect to MongoDB
# client = MongoClient('mongodb://localhost:27017/')
# db = client['local']  # Database name
# collection = db['deposit']  # Collection name

# # Step 4: Insert the data into MongoDB
# if products:
#     collection.insert_many(products)
#     print(f"{len(products)} records inserted.")
# else:
#     print("No data to insert.")

# # Step 5: Close the MongoDB connection
# client.close()
