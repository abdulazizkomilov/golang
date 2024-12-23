def compare_execution_times(go_time_ms, python_time_s):
    go_time_s = go_time_ms / 1000

    speedup_factor = python_time_s / go_time_s

    print(f"Go is approximately {speedup_factor:.2f} times faster than Python")


print(compare_execution_times(55.000208, 2.250244))