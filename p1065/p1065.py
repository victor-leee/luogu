def get_list():
    return [int(s) for s in input().strip().split(" ")]

if __name__ == '__main__':
    component_order = [[]]
    component_cost = [[]]
    first_line = get_list()
    m, n = first_line[0], first_line[1]
    order = get_list()
    for _ in range(n):
        component_order.append(get_list())
    for _ in range(n):
        component_cost.append(get_list())
    # 每一个零件的最后更新时间，component_last_update[i]和machine_last_update[j]
    # 一起决定了该零件i在工序j上（也是机器j）的起始时间，并且会新增machine_gaps[j]
    component_last_update = [0 for _ in range(n + 1)]
    # 每一个零件的进度
    component_progress = [0 for _ in range(n + 1)]
    # 每一台机器的最后使用时间
    machine_last_update = [0 for _ in range(m + 1)]
    # 机器的可用时间间隔数组
    machine_gaps = [[] for _ in range(m + 1)]
    # 记录每一个gap的开始时间
    gap_start_time = [[] for _ in range(m + 1)]
    
    for component_id in order:
        progress_id = component_order[component_id][component_progress[component_id]]
        cost = component_cost[component_id][component_progress[component_id]]
        if len(machine_gaps[progress_id]) > 0:
            success = False
            for i, gap in enumerate(machine_gaps[progress_id]):
                # 确保这个零件的这个工序能够在这个gap里面完成
                if max(gap_start_time[progress_id][i], component_last_update[component_id]) + cost > gap_start_time[progress_id][i] + gap:
                    continue

                # 
                component_last_update[component_id] = max(component_last_update[component_id], gap_start_time[progress_id][i])
                # 直接插入两个新的gap
                machine_gaps[progress_id][i] = 0
                machine_gaps[progress_id].insert(i, component_last_update[component_id] - gap_start_time[progress_id][i])
                machine_gaps[progress_id].insert(i + 1, gap_start_time[progress_id][i] + gap - (component_last_update[component_id] + cost))
                gap_start_time[progress_id].insert(i, gap_start_time[progress_id][i])
                gap_start_time[progress_id].insert(i + 1, component_last_update[component_id] + cost)

                # 
                component_progress[component_id] += 1
                component_last_update[component_id] += cost
                success = True

                break

            if success:
                continue


        start_time = max(component_last_update[component_id], machine_last_update[progress_id])
        end_time = start_time + cost

        if component_last_update[component_id] > machine_last_update[progress_id]:
            machine_gaps[progress_id].append(component_last_update[component_id] - machine_last_update[progress_id])
            gap_start_time[progress_id].append(machine_last_update[progress_id])
        
        component_last_update[component_id] = end_time
        component_progress[component_id] += 1
        machine_last_update[progress_id] = end_time
    
    print(max(machine_last_update))