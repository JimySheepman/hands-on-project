import cv2 as cv

def split_image_char(img):
    img = cv.imread(img)

    ROI_number = 0

    a, b, _ = img.shape
    print(a)
    print(b)
    c = 0
    d = 55
    for i in range(0, b, 45):
        if i < 181:
            char = img[:78, c:d + i]
            char2 = img[78:, c:d + i]
            c = d + i

            cv.imwrite('split_image_{}.jpg'.format(ROI_number), char)
            cv.imwrite('split_image_{}.jpg'.format(ROI_number + 1), char2)
            ROI_number += 2