import requests
from pymongo import MongoClient
from bs4 import BeautifulSoup


arr = [
  "05bd753b79583beb511b1bbaa1c816e7",
  "ac76adc3c871c900c9563131589d6b11",
  "819b00f22a3a11d56b5fc7a566f99218",
  "93d0589ec14956f87298454daf1b4249",
  "a907fec93cc50630fc48bd75d7dcb2a6",
  "d1e6848366b30a5375fff52d5a28e8b7",
  "a1fb44b590e07ad65fd4a33d750ca7f1",
  "4eebc8f2bc783f77d53c6324461c8929",
  "73544d0d88872f5f4df3d0760ecec6a2",
  "2244945a021250e5b11bac5eaf8d9666",
  "cbe31dd20e8567100ab535f3a464855b",
  "e424394a665f7c554c5a0f8489a74f4b",
  "009c29609d71d5cdd46f9d527c2a78d3",
  "e2874fc5685838a96d82a33f2ba7b814",
  "e024bc0d37996726e3fb7f27651427b5",
  "7b94ae60269852857ba13c870f203ce4",
  "c5b7960d5a5fb6204a35bab4f951e4cb",
  "44d06f1f8cd54bd5dbb6c48bd39bcdf7",
  "aa861dc8d90f681e77e5da8ea0b8e288",
  "6401daf5fc57aba7cfb95e119496d2dd",
  "d697cf26cbde13e35bb41143f26c36d2",
  "2b79aa122c80fe99fc14bebc0ec9bae0",
  "051bd6e3c1f07485085e6fc75bd1f441",
  "ef299fa810fe8bb8d7a22d886e975b08",
  "fedc0082d90db0cc1e873bc9bdd2f5b7",
  "ab31f1daafe919666ccf980887e63f7e",
  "bdf095c16d57c157e95bb0ed70c1efa0",
  "600c634f2a596098ef075eb542de6d69"
]

# MongoDB setup
client = MongoClient('mongodb://localhost:27017/')
db = client["local"]
collection = db["cma_detail"]

# headers = {
#     'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36'
#   }
#   #     https://new-m.pay.naver.com/savings/detail/05bd753b79583beb511b1bbaa1c816e7
# data = requests.get(f'https://new-m.pay.naver.com/savings/detail/05bd753b79583beb511b1bbaa1c816e7', headers=headers)
# soup = BeautifulSoup(data.text, 'html.parser')
# product_name = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > h3')
# bank_name = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > span.MainInfo_text__V_Fkd')
# sub_service = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > ul')
# arr_sub_service = [sub.text for sub in sub_service.find_all('li')] if sub_service else []
# maininfo_rate = soup.select_one("#content > div.MainInfo_article__iQqCh > div.MainInfo_area-rates__ZlyVd > strong")
# maininfo_sub_text = soup.select_one("#content > div.MainInfo_article__iQqCh > div.MainInfo_area-rates__ZlyVd > span")
# maininfo_description = soup.select_one("#content > div.MainInfo_article__iQqCh > p:nth-child(3)")
# join_method = soup.select_one("#PRODUCT_GUIDE > dl > div:nth-child(1) > dd")
# join_protection = soup.select_one("#PRODUCT_GUIDE > dl > div:nth-child(2) > dd")
# price_ratio_type = soup.select_one("#CMA_EARNING_RATE > dl > div > dt")
# price_ratio_content = soup.select_one("#CMA_EARNING_RATE > dl > div > dd")


# print(product_name.text if product_name else None)
# print(bank_name.text if bank_name else None)
# print(arr_sub_service)
# print(maininfo_rate.text if maininfo_rate else None)
# print(maininfo_sub_text.text if maininfo_sub_text else None)
# print(maininfo_description.text if maininfo_description else None)
# print(join_method.text if join_method else None)
# print(join_protection.text if join_protection else None)
# print(price_ratio_type.text if price_ratio_type else None)
# print(price_ratio_content.text if price_ratio_content else None)

for code in arr :
    headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36'
  }
  #     https://new-m.pay.naver.com/savings/detail/05bd753b79583beb511b1bbaa1c816e7
    data = requests.get(f'https://new-m.pay.naver.com/savings/detail/{code}', headers=headers)
    soup = BeautifulSoup(data.text, 'html.parser')
    product_name = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > h3')
    bank_name = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > span.MainInfo_text__V_Fkd')
    sub_service = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > ul')
    arr_sub_service = [sub.text for sub in sub_service.find_all('li')] if sub_service else []
    maininfo_rate = soup.select_one("#content > div.MainInfo_article__iQqCh > div.MainInfo_area-rates__ZlyVd > strong")
    maininfo_sub_text = soup.select_one("#content > div.MainInfo_article__iQqCh > div.MainInfo_area-rates__ZlyVd > span")
    maininfo_description = soup.select_one("#content > div.MainInfo_article__iQqCh > p:nth-child(3)")
    join_method = soup.select_one("#PRODUCT_GUIDE > dl > div:nth-child(1) > dd")
    join_protection = soup.select_one("#PRODUCT_GUIDE > dl > div:nth-child(2) > dd")
    price_ratio_type = soup.select_one("#CMA_EARNING_RATE > dl > div > dt")
    price_ratio_content = soup.select_one("#CMA_EARNING_RATE > dl > div > dd")
    data_to_insert = {
    "code": code,
    "product_name": product_name.text if product_name else None,
    "bank_name": bank_name.text if bank_name else None,
    "sub_service": arr_sub_service,
    "maininfo_rate": maininfo_rate.text if maininfo_rate else None,
    "maininfo_sub_text": maininfo_sub_text.text if maininfo_sub_text else None,
    "maininfo_description": maininfo_description.text if maininfo_description else None,
    "join_method": join_method.text if join_method else None,
    "join_protection": join_protection.text if join_protection else None,
    "price_ratio_type": price_ratio_type.text if price_ratio_type else None,
    "price_ratio_content": price_ratio_content.text if price_ratio_content else None,
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
