from sys import argv

filename = argv[1]
file = open(filename, 'rb').read()
for i in range(0, len(file), 2):
    print( file[i], file[i+1] )
