import cv2 as cv


def splitImage(img):
    image = cv.imread(img)
    gray = cv.cvtColor(image, cv.COLOR_BGR2GRAY)
    blur = cv.GaussianBlur(gray, (5, 5), 0)
    thresh = cv.threshold(blur, 0, 255, cv.THRESH_OTSU + cv.THRESH_BINARY_INV)[1]

    kernel = cv.getStructuringElement(cv.MORPH_RECT, (3, 3))
    opening = cv.morphologyEx(thresh, cv.MORPH_OPEN, kernel, iterations=1)
    cnts = cv.findContours(opening, cv.RETR_TREE, cv.CHAIN_APPROX_SIMPLE)
    cnts = cnts[0] if len(cnts) == 2 else cnts[1]

    ROI_number = 0
    for c in cnts:
        area = cv.contourArea(c)
        peri = cv.arcLength(c, True)
        approx = cv.approxPolyDP(c, 0.02 * peri, True)
        x, y, w, h = cv.boundingRect(approx)
        if len(approx) == 4 and (area > 1000) and (area < 80000):
            ROI = image[y:y + h, x:x + w]
            cv.imwrite('ROI_{}.jpg'.format(ROI_number), ROI)
            ROI_number += 1

    cv.imshow('thresh', thresh)
    cv.imshow('opening', opening)
    cv.waitKey()

