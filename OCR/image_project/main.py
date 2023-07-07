import numpy as np
import cv2 as cv
from matplotlib import pyplot as plt
from split_image import splitImage
from split_char_from_image import split_image_char


# load the image
def convertImage():
    img = cv.imread("kapak.jpg")
    img = cv.cvtColor(img, cv.COLOR_BGR2RGB)
    # detect circles
    gray = cv.medianBlur(cv.cvtColor(img, cv.COLOR_RGB2GRAY), 5)
    circles = cv.HoughCircles(gray, cv.HOUGH_GRADIENT, 1, 20, param1=170, param2=50, minRadius=25, maxRadius=0)
    circles = np.uint16(np.around(circles))
    # draw mask
    mask = np.full((img.shape[0], img.shape[1]), 0, dtype=np.uint8)  # mask is only
    for i in circles[0, :]:
        cv.circle(mask, (i[0], i[1]), i[2], (255, 255, 255), -1)
    # get first masked value (foreground)
    fg = cv.bitwise_or(img, img, mask=mask)
    # get second masked value (background) mask must be inverted
    mask = cv.bitwise_not(mask)
    background = np.full(img.shape, 255, dtype=np.uint8)
    bk = cv.bitwise_or(background, background, mask=mask)
    # combine foreground+background
    final = cv.bitwise_or(fg, bk)
    cv.imshow("1", final)
    cv.waitKey()
    img = final
    edges = cv.Canny(final, 100, 75)
    plt.imshow(edges, cmap='gray')
    x, y, w, h = cv.boundingRect(edges)
    image = cv.rectangle(edges, (x, y), (x + w, y + h), (0, 0, 12), 10)
    cv.imwrite("2.jpg", image)
    plt.imshow(image)
    plt.show()
if __name__ == '__main__':
    convertImage()
    splitImage("2.jpg")
    split_image_char("ROI_1.jpg")
