import requests
from pymongo import MongoClient
from bs4 import BeautifulSoup

# dehydratedState
arr = [
  "d1f2badd0918fcf04d2f06e048699b0f",
  "ad8c6895fdfa41f3b7a384f9abc78ab8",
  "a93398ca648613f2565e973592187f6c",
  "4b357e203eb0ec27195eb774140f4895",
  "7431d4fbe48ab9cc4621ad15413a811a",
  "31321db214427351a6ec87a62c4b9433",
  "e82093bbf079526961fffaaaf43730ef",
  "d3df9e06f739aaca5d0307f000f74284",
  "d8d3745b4ea39ea2b13dcafefe6a8230",
  "66a83dae2c1953b5903ea095ea972717",
  "acef2d4c6153adfd1c49fd36128e2c0c",
  "8e96681b86316f6b10e1d6a4b31a9955",
  "b655c3950677f93daf7fcbf9d39e91f2",
  "63205b6ac851fae90de75937c7af5e77",
  "1f1f097770d025861c130ded276128f1",
  "72b1fbaaacbec77eff5a5fdea7d0704b",
  "3e59ddcbf1aa941dbc0a504e61dbd517",
  "44a7c108de634d7bc99b331a1ba6cd2c",
  "fc7441ea14e510230e7555ea0be5b909",
  "b5262698ed41fc21e044fb8c14e92de9",
  "da804dfacec59071bd22d61aceceb2c6",
  "3aefef34f8b32d7093de80fce490690e",
  "5a293052b02621e7610d492b03c5cf30",
  "00b84f386f64a5809f8d6e3dbbbd87b5",
  "63e9526044f83022c3baebd658be1122",
  "79b9b51575a757fcc46b78aa274af563",
  "971ce82ce064d7a927106eb00291d494",
  "1b0bd3db5b88091055b1003221123a12",
  "ae1d21275f2052ab98f813b062ebb960",
  "b05a63c96864cbfa2afb8a09958a343f",
  "3b3a0f563d840ee01a04a515159d4ea2",
  "7e075ffd48688017b3f3c89bd3434a9d",
  "ea6646f0e3a365beae0a8df97a69ca33",
  "91cda67d721127112a308c3db5e0e07a",
  "7f3fdbddc0850de2bdf3015d54d8a9a0",
  "3ceacd6562dc9f731597a16401b50061",
  "0a029b0f3381d8382a8d06e036d1a07c",
  "1d80c27457d6d909c6567aaca6c23e4b",
  "222079ca8c00875b5b16df535a3dcaa3",
  "aba15bd99f8e4972c6506b0fd2b75d77",
  "9ca8b48ca8bdc3d7472fcf6a643d6b93",
  "9d3e1c051d256b06827ab64e689febe7",
  "60f10da955704ca1349954338c4860b1",
  "3adc7e152e5093d625000086b7c815be",
  "fb1fbbcdbbca475722a530bef5b731f6",
  "e43ec039e2aeaebd0f6f61340914dc9f",
  "f466999ee773d73df33cf99bdf2bbd8a",
  "3278782054a25278f5c78e333fc7b15f",
  "76427cb329c06904ec9fc4582138813a",
  "449f290b22ec00bd90307980cc164dde",
  "7236af06145686bf9885b98e980c8d1a",
  "fdc1c608dd119e6423f0f7e8c505ddfa",
  "a889f848edb145d1b69e375dd09b1744",
  "33f1970a27447b97fcbab97e473ecc33",
  "3da168c55b58df8f82a36e21b6738f29",
  "3db68a56e2045e8a002b7f7c1a37ba55",
  "c71b326f8fa70c7fa2423fa4b3f88f1a",
  "bb5f0da80e925b984c26e3d3d636424d",
  "7821e82547785abf3d2af0b0017538cd",
  "da6c9376a1d28df6591ad896c6693bf1",
  "962859f3fd75d5e405ddcc368d792cd3",
  "4866228bc443cb243f5826b6c8af19aa",
  "c2e03e7f55052eed51157887e8ca6633",
  "289ff3835c8a4a99a1c30f5e9ecb926d",
  "211a313504aeecb546a72ce506523dbe",
  "ad1d9eff03b0ffafaa42600d4a8b9290",
  "f743a4232763e8dea65de7af0ab12361",
  "a897ae3de3789b2a5d4ad77059a1e433",
  "aeb6b09a372aef240f5772edd110491d",
  "b3ecc8aca9d1fe8321b0ad7b49b99a19",
  "e25f216781482e90b05a3295caa68dad",
  "8f3731e6333f560b17d027e4dcbd8cbd",
  "6b7b23dcea4b950ee2132693afbee9a8",
  "9bb72d96705becc827de257da19f8bba",
  "1d863508d970901907c7a1f07a72208a",
  "c64949ce9fd0abf91148aab1c229ae6c",
  "90f544b4fbd8521b59c9da8aa82b8e24",
  "babce2733e63ba7b8ac58786fe2a474a",
  "974f71fbc49acefe7e8a62a57afcf7cd",
  "8838bcf5148c48349d8636bc1f280a6f",
  "0d42e38b898fc6933a7d418793cfe6d0",
  "ea010b547e5fd305aa60d647f86a736e",
  "798731c9bc1ef802733e7930181bfa70",
  "d8a06e426337b710f2e7f488180b2fd1",
  "198cedeaf33eaa4ec0c5ddd5e58fdb56",
  "e090449d40e1635479d2f671a318b9a6",
  "4b021d0864ac64835cf8ae9b4a7e8e32",
  "ce114dfe0686073803e872829d23454d",
  "b22a9097e71102b9e920e2891d35c589",
  "aa75d26df2318da5c60087aaa5a128fe",
  "ef5a789ee045c620c874ed7691a6c0f1",
  "9bd9add33d624ba7abababe732aa024b",
  "26ac6f33ee68f523e114046adaadb00b",
  "5de0503f0464ee7b413c128d5ae6dddd",
  "9693290ab7270db5ed2d3ff3b366e42e",
  "76a8e2e3daffb3cdb1319dc7f6cf04b1",
  "c0e2e013d25daec7f2c1cd258ea717d8",
  "7935883d8fa68a6243adbafa2070c445",
  "d6d1f906422e7f255f50d8a4be491327",
  "ad49af74a87ebd1085136ea0fe9d3890",
  "487a1acd81b7cd9e86bac23684497904",
  "ef8c3c44430c900a4faf65956a9ffe62",
  "fb3324c96df782d8d2b8cc673c022869",
  "8378f99ec5fc2e01c9edc73253f41407",
  "e594cba7899cbd3798cfacda226a6fc2",
  "aa1f0585d8311a3ad76e5233fc2f5476",
  "6c66d3c64e0d61086a694904545dde45",
  "a754fcd7c541a9ebdf97688e1356f7ce",
  "5e733b0bf4daf2ad06adec28627d2e9a",
  "530d2f361dad2129939cb11be1b39369",
  "60cf62a8e3b7ceab25d350b8e364004f",
  "d11e9f48ab67e3ce908f8661630670ed",
  "bd64f6c8bafcc91e4314ce848ad05710",
  "1fc0a781fe705c91da833bc05b8f2e27",
  "af50f5788097d47edcca4c6aaabfc6ae"
]

client = MongoClient('mongodb://localhost:27017/')
db = client["local"]
collection = db["instalment_savings_detail"]

# Step 1: Fetch the JSON data from the API
headers = {'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36'}


data = requests.get(f'https://new-m.pay.naver.com/savings/detail/d1f2badd0918fcf04d2f06e048699b0f',headers=headers)
soup = BeautifulSoup(data.text, 'html.parser')

product_name = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > h3')
bank_name = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > span.MainInfo_text__V_Fkd')
sub_service = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-title__hzVI3 > ul')
arr_sub_service = [sub.text for sub in sub_service]
max = soup.select_one('#content > div.MainInfo_article__iQqCh > div.MainInfo_area-rates__ZlyVd > dl > div:nth-child(1) > dd')
info_year = soup.select_one('#PRODUCT_GUIDE > dl > div:nth-child(1) > dd > p')
info_fee = soup.select_one('#PRODUCT_GUIDE > dl > div:nth-child(2) > dd > p')
info_method = soup.select_one('#PRODUCT_GUIDE > dl > div:nth-child(3) > dd')
info_condition = soup.select_one('#PRODUCT_GUIDE > dl > div:nth-child(5) > dd > p')
info_payment = soup.select_one('#PRODUCT_GUIDE > dl > div:nth-child(6) > dd > p')
info_precautions = soup.select_one('#PRODUCT_GUIDE > dl > div:nth-child(7) > dd > p')
info_protect = soup.select_one('#PRODUCT_GUIDE > dl > div:nth-child(8) > dd > p')
print(product_name.text)
print(bank_name.text)
print(arr_sub_service)
print(max.text)
print(info_year.text)
print(info_fee.text)
print(info_method.text)
print(info_condition)
print(info_payment)
print(info_precautions)
print(info_protect)


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
