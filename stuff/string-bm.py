if __name__ == '__main__':
    m = 'aaaaaaaa'
    p = 'aa'
    last_pos = {}
    for i, ch in enumerate(p):
        last_pos[ch] = i
    i = len(p) - 1

    # good suffix rule
    suffix, is_prefix = [-1 for _ in range(len(p)+1)], [False for _ in range(len(p)+1)]
    for i in range(len(p)):
        j, k = i, 0
        while j >= 0 and p[j] == p[len(p)-1-j]:
            j -= 1
            k += 1
            suffix[k] = j + 1
        if j < 0:
            is_prefix[k] = True

    while i < len(m):
        p_i = len(p) - 1
        i_temp = i
        while p_i >= 0 and m[i_temp] == p[p_i]:
            p_i -= 1
            i_temp -= 1
        if p_i < 0:
            # we found one ~
            print(f'found substring in main string that matches pattern string, end with {i}')
            break
            
        # move steps on bad char rule
        bc_move = p_i - (last_pos[m[p_i]] if m[p_i] in last_pos else -1)
        # move steps on good suffix rule
        gs_move = 0
        if p_i < len(p) - 1:
            if suffix[len(p)-p_i-1] >= 0:
                gs_move = p_i - suffix[len(p)-p_i-1] + 1
            else:
                for j in range(p_i + 2, len(p)):
                    if is_prefix[len(p)-j]:
                        gs_move = j
                        break
                gs_move = len(p)
        i += max(bc_move, gs_move)