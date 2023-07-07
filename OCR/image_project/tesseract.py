import pytesseract
import cv2 as cv
from keras.models import  model_from_json
import os
from keras_cnn import trainModel, getData
import numpy as np

pytesseract.pytesseract.tesseract_cmd = r'C:\Users\yusuf\AppData\Local\Programs\Tesseract-OCR\tesseract.exe'

for i in range(10):
    string_image = "split_image_{}.jpg".format(i)
    img = cv.imread(string_image)
    text = pytesseract.image_to_string(img, lang='eng', config='--psm 9 --oem 3 -c tessedit_char_whitelist=ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789')
    print(i,text,end=" ")
