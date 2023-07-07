import requests
import datetime
import lxml.html as lh
import pandas as pd

"""
! Web düzeni düzenlemesi için bir tanıtıcı, sayfa oluşturma
! Web sitesinin içeriğini doc altında saklayın
! HTML'nin <tr>..</tr> arasında depolanan verileri ayrıştırın
"""
url = 'https://www.worldometers.info/coronavirus/'
page = requests.get(url)
doc = lh.fromstring(page.content)
tr_elements = doc.xpath('//tr')

tr_elements = doc.xpath('//tr')
# print([len(T) for T in tr_elements[:12]])

""" 
? Boş liste oluştur
? Her satır için, her bir ilk öğeyi (başlık) ve boş bir listeyi saklayın 
"""
col = []
i = 0
for t in tr_elements[0]:
    i += 1
    name = t.text_content()
    # print ('%d:"%s"'%(i,name))
    col.append((name, []))

"""
! İlk satır başlık olduğu için veriler ikinci satırdan itibaren depolanır
! T bizim j'inci sıramız
! Satır 10 boyutunda değilse, //tr verileri tablomuzdan değildir
! i sütunumuzun indeksidir
! Satırın her bir öğesini yineleyin
! Satır boş mu kontrol et
! Herhangi bir sayısal değeri tam sayılara dönüştür
! i'nci sütunun boş listesine verileri ekle
! Bir sonraki sütun için i'yi artır 
"""
for j in range(1, len(tr_elements)):
    T = tr_elements[j]
    if len(T) != 22:
        break

    i = 0
    for t in T.iterchildren():
        data = t.text_content()
        if i > 0:
            try:
                data = int(data)
            except:
                pass

        col[i][1].append(data)
        i += 1
"""
? Parse edilen veriyi sözlüğe yazdır
? Parse zamanını alık yeni bir sözlüğe aktar
? Sözlüğü birleştir 
? Gereksiz sütünları çıkar
? Sözlüğü csv olarak kaydet
? .sql dosyasına INSORT komutlarını yazıdır
"""
Dict = {title: column for (title, column) in col}
df = pd.DataFrame(Dict)
df = df.iloc[8:231]

new_time = [i for i in range(df.shape[0])]
new_time_2 = [str(i+1) for i in range(df.shape[0])]
for i in range(df.shape[0]):
    an = datetime.datetime.now()
    new_time[i] = datetime.datetime.strftime(an, '%x %X')


df2 = pd.DataFrame({'#': new_time_2, 'DateTime': new_time})
df_final = pd.merge(df, df2, on="#")

for i in df_final.columns:
    if i != '#' and i != 'Country,Other' and i != 'TotalCases' and i != 'TotalDeaths' and i != 'DateTime':
        df_final.drop(i,  axis='columns', inplace=True)

df_final.to_csv('covid-19-data.csv')
with open("deneme.sql", "w") as file:
    for i in range(df_final.shape[0]):
        file.write("INSERT INTO `covid19`(date_time,country,cases,deaths) VALUES ({},{},{},{});\n".format(str(df_final['DateTime'][i]),
                                                                                                        str(df_final['Country,Other'][i]),
                                                                                                        str(df_final['TotalCases'][i]),
                                                                                                        str(df_final['TotalDeaths'][i])))