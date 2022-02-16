import os
#Create new folder when each website is crawled
def create_project_dir(directory):
    if not os.path.exists(directory):
        print('Creating project file ' + directory)
        os.makedirs(directory)
    else:
        print('Error: Project created!')
#create queue and crawled files (if not created)
def create_data_files(project_name, base_url):
    queue = project_name + '/queue.txt' #list waiting to be crawled
    crawled = project_name + '/crawled.txt' #files crawled
    if not os.path.isfile(queue): #does the file exist? else create file
        write_file(queue, base_url)
    if not os.path.isfile(crawled): #create in crawled file so file is crawled
        write_file(crawled, ' ')
#create a new files
def write_file(path, data):
    f = open(path, 'w') #write ('w') into file (path)
    f.write(data)
    f.close() #prevent data leaks
#add data into an existing files
def append_to_file(path, data):
     with open(path, 'a') as file:
         file.write(data+'\n')
#delete the contents of a files
def delete_file_contents(path):
    with open(path, 'w'):
        pass
#Read a file and convert each line to set item
def file_to_set(file_name):
    results  = set()
    with open(file_name, 'rt') as f:
        for line in f:
            results.add(line.replace('\n', ''))
    return results
#Iterate through a set, each item will be a new item in the files
def set_to_file(links, file):
    delete_file_contents(file)
    for link in sorted(links):
        append_to_file(file, link)
