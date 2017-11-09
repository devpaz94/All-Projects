

import itertools
import time

def get_all_subsets(main_set):
    #take an individual set being trialed, return every possible sub set
    #e.g. takes {20, 31, 38, 39, 40, 44, 46}, and returns [[20, 31], [31, 38], ..., [20, 31, 38, 39, 40, 44]]
    #where each array inside the main array constains a sub_set
    all_sub_sets = []
    all_sub_sets.append(main_set)
    for i in range(2, 7):
        sub_set = [''.join(x) for x in itertools.combinations(main_set, i)]
        all_sub_sets.append(sub_set)
    return all_sub_sets


def create_sums_dictionary(all_sub_sets):
    #takes an array of every single sub-set and returns a new array with the size, sum and sub-set
    #e.g. the sub-set [20, 31, 38] will be added to the dictionary as [3, 89, [20, 31, 38]]
    sums_dictionary = []
    set_length = 2
    for sub_set in all_subs_sets[1:]:
        for number in sub_set:
            lst = [number[i:i + 2] for i in range(0, len(number), 2)]
            sum = 0
            for i in lst:
                sum = sum + int(i)
            sums_dictionary.append([set_length, sum, lst])
        set_length = set_length + 1
    return sums_dictionary


def is_disjoint(first_set_values, second_set_values):
    #checks wether two sets are disjoint. Meaning they do not share any of the same values
    result = False
    for number_first_set in first_set_values:
        for number_second_set in second_set_values:
            if number_first_set == number_second_set:
                result = True
    return result


def check_for_equal_disjoint_sets(sums_dictionary):
    #checks every single subset against each other to make sure the problems conditions are met
    #i.e. sums are not equal, they are disjoint and a larger set cannot have a smaller sum than
    #a smaller set (smaller meaning has less elements inside it)
    i = 1
    bool = False
    for first_sum in sums_dictionary:

        first_set_size = first_sum[0]
        first_set_values = first_sum[2]
        first_sum = first_sum[1]

        for second_sum in sums_dictionary[i:]:

            second_set_size = second_sum[0]
            second_set_values = second_sum[2]
            second_sum = second_sum[1]

            result = first_sum - second_sum

            if result == 0:
                bool = True
                break
            elif result > 0:
                if second_set_size > first_set_size:
                    bool = True
                    break
            elif result < 0:
                if first_set_size > second_set_size:
                    bool = True
                    break
            elif is_disjoint(first_set_values, second_set_values):
                bool = True
                break
        i = i + 1

    return bool




start = time.time()

#Method is to start with the set which would be given using the formula given in the problem
#then search for solutions close to that by trying all combinations of sets which are within a given range


current_best = {20, 31, 38, 39, 40, 44, 46} #this is a good guess at the answer
best_sum = sum(current_best)

#try the cuurent best +,- a little bit, creates about 280,000 samples to try
for complete_set in itertools.product(*[list(range(x - 3, x + 4)) for x in current_best]):

    if sum(complete_set) >= best_sum: # first checks that the sum is lower than the current best
        continue

    main_list = []
    for number in complete_set:
        main_list.append(str(number))

    all_subs_sets = get_all_subsets(main_list)
    sums_dictionary = create_sums_dictionary(all_subs_sets)

    if check_for_equal_disjoint_sets(sums_dictionary) == True: #this checks all the other conditions
        continue

    best_sum = sum(complete_set)
    current_best = complete_set

stop = time.time()

print current_best
print best_sum
print "Completed in " +  str(round((stop - start) / 60, 2)) + " minutes."
