
import csv
#open and read the file after the appending:
response = []
with open('Point_Of_Interest.csv', newline='') as csvfile:
    spamreader = csv.reader(csvfile, delimiter=',')
    for row in spamreader:
        elem = []
        #print(row[0].split(' '))
        separated = row[0].split(' ')
        elem.append(separated[1])
        elem.append(separated[0])
        elem.append(row[15])
        
        #print(row[0],row[15])
        response.append(elem)
print(response)

with open('Point_Of_Interest.txt', 'w', newline="") as x:
   csv.writer(x,delimiter=",").writerows(response)