import threading
from queue import Queue
from spider import Spider
from domain import *
from crawler import *
import sys
import time


class Colors:
    CRED2 = "\33[91m"
    CBLUE2 = "\33[94m"
    ENDC = "\033[0m"


banner = ("""


 ▄▀▀▄    ▄▀▀▄  ▄▀▀█▄▄▄▄  ▄▀▀█▄▄       ▄▀▄▄▄▄   ▄▀▀▄▀▀▀▄  ▄▀▀█▄   ▄▀▀▄    ▄▀▀▄  ▄▀▀▀▀▄     ▄▀▀█▄▄▄▄  ▄▀▀▄▀▀▀▄ 
█   █    ▐  █ ▐  ▄▀   ▐ ▐ ▄▀   █     █ █    ▌ █   █   █ ▐ ▄▀ ▀▄ █   █    ▐  █ █    █     ▐  ▄▀   ▐ █   █   █ 
▐  █        █   █▄▄▄▄▄    █▄▄▄▀      ▐ █      ▐  █▀▀█▀    █▄▄▄█ ▐  █        █ ▐    █       █▄▄▄▄▄  ▐  █▀▀█▀  
  █   ▄    █    █    ▌    █   █        █       ▄▀    █   ▄▀   █   █   ▄    █      █        █    ▌   ▄▀    █  
   ▀▄▀ ▀▄ ▄▀   ▄▀▄▄▄▄    ▄▀▄▄▄▀       ▄▀▄▄▄▄▀ █     █   █   ▄▀     ▀▄▀ ▀▄ ▄▀    ▄▀▄▄▄▄▄▄▀ ▄▀▄▄▄▄   █     █   
         ▀     █    ▐   █    ▐       █     ▐  ▐     ▐   ▐   ▐            ▀      █         █    ▐   ▐     ▐   
               ▐        ▐            ▐                                          ▐         ▐       	  v1.0
                                                                                                                """)

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

y = ("""\n\t\t\t  --*--Rules--*--
        [!]It should start the homepage with http or https. 
        [!]Avoid sites that have lots of links                    \n""")
for col in y:
    print(Colors.CRED2 + col, end="")
    sys.stdout.flush()
    time.sleep(0.0040)

z = "\n"
for col in z:
    print(Colors.ENDC + col, end="")
    sys.stdout.flush()
    time.sleep(0.4)


# Create worker threads (will die when main exits)
def create_threads():
    """
    create 8 threads for 8 spiders,
    based on NUMBER_OF_THREADS = 8
    """
    # just for looping 8 times so _ is used
    for _ in range(NUMBER_OF_THREADS):
        # create a thread, pasing a target, which is work
        t = threading.Thread(target=work)
        # run in daemon process and die once main exits
        t.daemon = True
        t.start()


def work():
    """
    do the next job in the queue
    """
    while True:
        url = queue.get()
        # spider to crawl the url using the current thread created.
        Spider.crawl_page(threading.current_thread().name, url)
        queue.task_done()


# Each queued link is a new job for the spiders
def create_jobs():
    # for all links in the file put into set (memory) queue
    for link in file_to_set(QUEUE_FILE):
        queue.put(link)
    queue.join()
    crawl()


# check if there's items in queue, if so continue to crawl
def crawl():
    # extract links from queue.txt for the links
    queued_links = file_to_set(QUEUE_FILE)
    # if queue file have more than 1 links, carry on to extract the links
    if len(queued_links) > 0:
        print(str(len(queued_links)) + ' links in the queue')
        create_jobs()


if __name__ == '__main__':
    print("\033[1;91m[!] Project Name: \033[0", end="")
    PROJECT_NAME = input()
    print("\033[1;91m[!] Home Page: \033[0", end="")
    HOMEPAGE = input()
    if "http" in HOMEPAGE:
        DOMAIN_NAME = get_domain_name(HOMEPAGE)
        QUEUE_FILE = PROJECT_NAME + '/queue.txt'
        CRAWLED_FILE = PROJECT_NAME + '/crawled.txt'
        NUMBER_OF_THREADS = 8
        queue = Queue()
        Spider(PROJECT_NAME, HOMEPAGE, DOMAIN_NAME)
        create_threads()
        crawl()
        create_threads()
        crawl()
    else:
        print("incorrect home page name ")