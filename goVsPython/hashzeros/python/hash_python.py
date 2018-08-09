from hashlib import sha256
from joblib import Parallel, delayed 
import os
import sys 
import argparse

parser = argparse.ArgumentParser(description='this script to give best hashs of text')

parser.add_argument("-t", "--text", dest="text", help="text to hash")
parser.add_argument("-c", "--count", dest="count", help="count of tries")

args = parser.parse_args()

text = ""
if args.text:
    text = args.text
else:
    raise Exception("you should set the text to hash ex.:'-t yourtext'")  

count = 10000
if args.count:
    count = int(args.count)


def printtest(start_tryies,end_tries):  
    expect = '0' 
    zeros_count = 0 
    result = ''
    obj = []
    for i in range(start_tryies,end_tries): 
        data = text + str(i)
        text_hash = sha256(data.encode()).hexdigest()
        if text_hash.startswith(expect):
            expect += '0'
            zeros_count = 0
            for x in text_hash:
                if x == '0':
                    zeros_count += 1
                else:
                    break
            result = str(zeros_count) + ',' + data + "," + text_hash 
            obj = [text,i,data,expect,text_hash] 
    return obj

ranges = range(1000,count,1000) 
best_zeros_at_all = Parallel(n_jobs=20)(delayed(printtest)((rng - ranges.step),rng) for rng in ranges)   

zeros_count = 0
finalRes = []
for best in best_zeros_at_all: 
    if zeros_count < len(best[3]):
        zeros_count = len(best[3])
        finalRes = best 

print(finalRes) 


#how to use
#cd Learning/test
#source activate tweepy36
#python hash_python.py -t foo -c 100000
