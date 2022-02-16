# Keylog Timer


The purpose of the project is to understand whether the person using the computer is the main user from the keyboard usage habits of the people. So, I wrote a keylogger and script to collect data and store data meaningfully for machine learning.

## Working Logic
It is a keylogger that calculates the printing time between letters when you run the code. It writes all keyboard entries in a file. There is a script to organize the data and calculate the times in milliseconds. 
###### File tasks :
- Data collector ``` keylog_timer.py ``` 
- Data split and organizer ``` time_data.py   ``` 

## Quick Start
###### Windows user :
Turn off security settings before downloading, otherwise the program will be detected by the system as a threat and deleted.  
1. ```pip install -r requirements.txt ``` 
2. ```python keylog_timer.py```
3. stop the program
4. Run ```python  time_data.py``` for to process data 
