
#================== PROBLEM OUTLINE ====================================
#Find three, 4-digit prime numbers in an arithmetic sequence where each number is a different permutation of the same 4 digits

import itertools
import time


def perms(number):
    """returns every permutation of a 4-digit number"""
    num = str(number)
    return list(map("".join, itertools.permutations(num)))



def is_prime(p):
    """returns True if p is prime, else False"""
    return all(p % i for i in xrange(2, p))




def arithmetic_sequence(primes):

    #checks a list to see whether it contains an arithmetic sequence
    #and prints the sequance if found

    #'primes' is a list of prime numbers in numerical order where each number is a different permutation of the same 4 digits


    for p in range(0, len(primes)):
        #loops through each prime in the list and creates a new list 'primes_list' which
        #removes the values which have already been checked for a squence

        primes_list = primes[p:]
        counter = 1

        gap = 1
        #these two following loops go through every possible combination of three numbers in the list starting with prime 'p'
        #and checks whether they are arithmetic
        for x in range(1, len(primes_list) - 1):

            diff = int(primes_list[counter]) - int(primes_list[counter - gap]) #get the difference between prime 'p' and a second prime

            for i in primes_list[counter :]:
                if int(i) == diff + int(primes_list[counter]): #check wether the difference to the third prime is the same as between the first two
                    print "Success!"
                    print "The primes are " + str(int(primes_list[counter]) - diff) + ", " + primes_list[counter] + " and " + i


            counter = counter + 1
            gap = gap + 1



start = time.time()

list_all = [''.join(x) for x in itertools.combinations('123456789',4)] #returns a list of all numbers from doing 9 Choose 4. i.e (1234, 1235, 1236, ... , 6789)

primes_list = []

for number in list_all: #for each number in the list of all possible numbers.

    perms_list = perms(number) #get every permutation of that 4-digit number. for example 'perms(1234)' returns '[1234, 1243, 1324, ... , 4321]'

    for permutation in perms_list: #for each permutation of the 4-digit number

        if is_prime(int(permutation)):
            primes_list.append(str(permutation)) #check if prime and if so append to list

    if len(primes_list) > 2:
        arithmetic_sequence(primes_list) #in the new list of primes, check if it contains an arithmetic sequence

    primes_list = []

end = time.time()

print "Completed in " + str(round(end - start, 2)) + " seconds."