def get_nth_number(series, n):
    occurrences_index = {}

    def speak_number(n, round):
        occurrences = occurrences_index.get(n, [])
        if len(occurrences) == 2:
            occurrences[0] = occurrences[1]
            occurrences[1] = round
        else:
            occurrences.append(round)
            if not n in occurrences_index:
                occurrences_index[n] = occurrences

    for i, v in enumerate(series):
        speak_number(v, i)

    last_number = series[-1]
    for i in range(len(series), n):
        last_round = i - 1
        first_occurrence = len(occurrences_index[last_number]) == 1

        if first_occurrence:
            last_number = 0
        else:
            last_number = last_round - occurrences_index[last_number][0]

        speak_number(last_number, i)
        
    return last_number

if __name__ == "__main__":
    series = 1,20,8,12,0,14,
    
    print(get_nth_number(series, 2020))
    print(get_nth_number(series, 30000000))