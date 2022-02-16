from urllib.request import urlopen #used in python to connect to webpage
from link_finder import LinkFinder #import link finder in link_finder.py
from crawler import * #import ALL in general.py

class Spider:
    #class variables (shared among all spiders instances)
    project_Name = ''
    base_url = ''
    domain_name = ''
    queue_file = '' #saved inside text file, for resuming later
    crawled_file = '' #saved inside text file, for resuming later
    queue = set()      #using this to stored in ram .
    crawled = set()    #using this to stored in ram.

    def __init__(self, project_name, base_url, domain_name):
        Spider.project_name = project_name
        Spider.base_url = base_url
        Spider.domain_name = domain_name
        Spider.queue_file = Spider.project_name + '/queue.txt'
        Spider.crawled_file = Spider.project_name + '/crawled.txt'
        self.boot()
        self.crawl_page('First Spider', Spider.base_url)

    @staticmethod
    def boot():
        create_project_dir(Spider.project_name)
        create_data_files(Spider.project_name, Spider.base_url)
        #both queue and crawled are retrieved from file and saved in ram for faster operation.
        Spider.queue = file_to_set(Spider.queue_file)
        Spider.crawled = file_to_set(Spider.crawled_file)

    @staticmethod
    def crawl_page(thread_name, page_url):
        if page_url not in Spider.crawled:
            print(thread_name + ' now crawling ' + page_url)
            #print('Queue: ' + str(len(Spider.queue) + ' | Crawled: ' + str(len(Spider.crawled))))
            print('Queue: {}   |   Crawled: {}'.format(len(Spider.queue), len(Spider.crawled)))
            Spider.add_links_to_queue(Spider.gather_links(page_url))
            #remove from queue list after completed.
            Spider.queue.remove(page_url)
            #add links into crawled list after crawled.
            Spider.crawled.add(page_url)
            #call both sets : file_to_set and set_to_file and convert them to file
            Spider.update_files()

    @staticmethod
    def gather_links(page_url):
        html_string = ''
        #using error catching on networking
        try:
            response = urlopen(page_url)

            #make sure its a html page and not some pdf format
            if 'text/html' in response.getheader('Content-Type'):
                #python read in html bytes format
                html_bytes = response.read()
                #convert into human readable character (utf-8)
                html_string = html_bytes.decode('utf-8')
                #create a linkfinder object
            finder = LinkFinder(Spider.base_url, page_url)
            #feed in the html strings
            finder.feed(html_string)
        except:
            print('Error: cannot crawl page!')
            return set()
        return finder.page_links()

    @staticmethod
    def add_links_to_queue(links):
        for url in links:
            #checks if url are being crawled
            if url in Spider.queue:
                continue
            #checks if url was crawled
            if url in Spider.crawled:
                continue
            #checks if spider is in correct site,
            #spiders may accidentally crawled google,FB, Insta,
            #if the site consist them. this is to make sure we are crawling
            #the only site we told spider to crawl. eg: thenewboston.com/******
            if Spider.domain_name not in url:
                continue
            #add into queue
            Spider.queue.add(url)

    @staticmethod
    def update_files():
        set_to_file(Spider.queue, Spider.queue_file)
        set_to_file(Spider.crawled, Spider.crawled_file)
