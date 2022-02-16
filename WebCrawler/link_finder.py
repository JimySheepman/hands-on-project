from html.parser import HTMLParser
from urllib import parse

class LinkFinder(HTMLParser):
    def __init__(self, base_url, page_url):
        super().__init__()
        self.base_url = base_url
        self.links = set()

    #if we get any error
    def error(self, message):
        pass

    #handles the start tag in html, <head>
    def handle_starttag(self, tag, attrs):
        if 'a' in tag: #find the tag starts with 'a' do not use '=='
            for (attribute, value) in attrs:
                if 'href' in attribute:
                    url = parse.urljoin(self.base_url, value)
                    #enables the url to be full, most of the time the url would be  /this/this without http://...
                    self.links.add(url)

    def page_links(self):
        return self.links
