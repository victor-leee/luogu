def str_hash(s):
    res = 0
    for ch in s:
        res = res * 26 + ord(ch) - ord('a')
    
    return res

if __name__ == '__main__':
    s = 'sooptsoooooosssssoopt'
    p = 'soopt'
    p_hash = str_hash(p)
    s_hash = str_hash(s[:len(p)])
    for i in range(len(s) - len(p) + 1):
        if p_hash == s_hash:
            print(f'found same substring in main string, starting from {i}')
        if i == len(s) - len(p):
            break
        s_hash = (s_hash - ((26 ** (len(p) - 1)) * (ord(s[i]) - ord('a')))) * 26 + ord(s[i+len(p)]) - ord('a')
    
