# Python3 required

def enum():
    FILE = '../Package/Shared/opcodes.go'
    opcodes = open(FILE, 'r')
    formatted = list()
    
    AWAIT = -1
    USEIT = 0
    DONE  = 1
    STATE = AWAIT
    COUNT = 0
    
    while STATE != DONE:
        line = opcodes.readline()
        
        if '(' in line:
            STATE = USEIT
            formatted.append(line)
            continue
            
        if ')' in line:
            STATE = DONE
            
        if STATE == AWAIT or STATE == DONE:
            formatted.append(line)
            
        if STATE == USEIT:
            ned = line.rstrip('\n')
            formatted.append( '{}\t// {}\n'.format(ned, COUNT) )
            COUNT += 1
    
    opcodes.close()
    opcodes = open(FILE, 'w')
    for line in formatted:
        opcodes.write(line)

if __name__ == "__main__":
    enum()