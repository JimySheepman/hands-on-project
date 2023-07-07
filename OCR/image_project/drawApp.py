from tkinter import *
from PIL import ImageGrab
import numpy as np
import scipy.misc
from keras.models import  model_from_json
import os
from keras_cnn import trainModel, getData
import cv2 as cv

class Paint(object):

    def __init__(self):
        self.root = Tk()

        #defining Canvas
        self.c = Canvas(self.root, bg='white', width=280, height=280)
        self.c.grid(row=1, columnspan=5)

        self.classify_button = Button(self.root, text='classify', command=lambda:self.classify(self.c))
        self.classify_button.grid(row=0, column=1)

        self.clear = Button(self.root, text='clear', command=self.clear)
        self.clear.grid(row=0, column=3)

        self.prediction_text = Text(self.root, height=2, width=10)
        self.prediction_text.grid(row=2, column=3)

        self.model = self.loadModel()
        self.setup()
        self.root.mainloop()

    def setup(self):
        self.old_x = None
        self.old_y = None
        self.line_width = 15
        self.color = 'black'
        self.c.bind('<B1-Motion>', self.paint)
        self.c.bind('<ButtonRelease-1>', self.reset)

    def paint(self, event):
        paint_color = self.color
        if self.old_x and self.old_y:
            self.c.create_line(self.old_x, self.old_y, event.x, event.y,
                               width=self.line_width, fill=paint_color,
                               capstyle=ROUND, smooth=TRUE, splinesteps=36)
        self.old_x = event.x
        self.old_y = event.y

    def clear(self):
        """Clear drawing area"""
        self.c.delete("all")

    def reset(self, event):
        """reset old_x and old_y if the left mouse button is released"""
        self.old_x, self.old_y = None, None

    def classify(self, widget):
        #getting pixel information
        x=self.root.winfo_rootx()+widget.winfo_x()
        y=self.root.winfo_rooty()+widget.winfo_y()
        x1=x+widget.winfo_width()
        y1=y+widget.winfo_height()
        #save drawing
        ImageGrab.grab().crop((x,y,x1,y1)).resize((28, 28)).save('classify.png')
        img=cv.imread("split_image_2.jpg")
        #img = scipy.misc.imread('classify.png', flatten=False, mode='P')
        img = np.array(img)
        img = np.reshape(img, (1, 28, 28, 1))
        # Change pixels to work with our classifier
        img[img==0] = 255
        img[img==225] = 0
        # Predict digit
        pred = self.model.predict([img])
        # Get index with highest probability
        pred = np.argmax(pred)
        print(pred)
        self.prediction_text.delete("1.0", END)
        self.prediction_text.insert(END, pred)

    def loadModel(self):
        if(os.path.exists('mnist_model.h5')):
            print('loading model')
            json_file = open('model.json', 'r')
            model_json = json_file.read()
            json_file.close()
            model = model_from_json(model_json)
            model.load_weights("mnist_model.h5")
            return model
        else:
            print('train model')
            X_train, y_train, X_test, y_test = getData()
            model = trainModel(X_train, y_train, X_test, y_test)
            return model

if __name__ == '__main__':
    Paint()