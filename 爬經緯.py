import csv
import requests
import time
from selenium import webdriver
from selenium.webdriver.support.ui import Select
from bs4 import BeautifulSoup
options = webdriver.ChromeOptions()
options.add_argument("headless")
import csv

def get_coordinate(addr):
    browser = webdriver.Chrome(executable_path=r'C:/Users/mypc\Desktop/chromedriver',options=options)
    browser.get("http://www.map.com.tw/")
    search = browser.find_element_by_id("searchWord")
    search.clear()
    search.send_keys(addr)
    browser.find_element_by_xpath("/html/body/form/div[10]/div[2]/img[2]").click() 
    time.sleep(2)
    iframe = browser.find_elements_by_tag_name("iframe")[1]
    browser.switch_to.frame(iframe)
    coor_btn = browser.find_element_by_xpath("/html/body/form/div[4]/table/tbody/tr[3]/td/table/tbody/tr/td[2]")
    coor_btn.click()
    coor = browser.find_element_by_xpath("/html/body/form/div[5]/table/tbody/tr[2]/td")
    coor = coor.text.strip().split(" ")
    lat = coor[-1].split("：")[-1]
    log = coor[0].split("：")[-1]
    browser.quit()
    return (lat, log)

arr=[]

with open(r'C:/Users/mypc\Desktop/7887.csv','r',newline='',encoding="utf-8") as csvfile:

  # 讀取 CSV 檔內容，將每一列轉成一個 dictionary
  rows = csv.reader(csvfile)
  
  # 以迴圈輸出指定欄位
  for row in rows:
    print(get_coordinate(row))
          

#with open('/Users/julian/Desktop/test2.csv','w',newline='') as csvfile:
  #writer = csv.writer(csvfile)
  #for x in arr :
      #writer.writerow([x])
