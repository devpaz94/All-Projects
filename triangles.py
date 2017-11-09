#Take a triangle and select some coordinate inside that triangle and plot a single point. Then randomly pick one of the three 
#corners of the triange and plot a new point exactly half way between the point and the corner. From this new point, repeat 
#the process picking a corner at random and plotting a point exaclty half way between you and that new randomly selected corner. If this 
#process is repeated many many times, it plots a nice unexpected pattern which is quite cool. Thats what this code does

# shown in this numberphile video https://www.youtube.com/watch?v=kbKtFN71Lfs

import matplotlib.pyplot as plt
import random

x_points = [1]
y_points = [1]
    
yi = 1
xi = 1

points = [[1, 1],[2, 3],[3, 1]] # three corners of the tiangle

def yfinder(yi):
    #takes the y-coordinate for the current position inside the triange
    #and returns the new y-coordinate half way between the one and the randomly selected point
    if point[1] <= yi:
        yi = point[1] + (abs(float(yi - point[1])/2))
    else:
        yi = point[1] - (abs(float(yi - point[1])/2))
    return yi
    
for i in range(100000):

    point = (random.choice(points)) # select a random corner of the triangle

    #theses lines take the x-coordinate for the current position inside the triange
    #and returns the new x-coordinate half way between the one and the randomly selected point
    
    if point[0] <= xi:
        xi = point[0] + (abs(float(xi - point[0])/2))
        yi = yfinder(yi)            
    else:
        xi = point[0] - (abs(float(xi - point[0])/2))
        yi = yfinder(yi)
            
    x_points.append(xi)
    y_points.append(yi)
    
plt.plot(x_points, y_points, 'bs', ms=0.5) #plot all the points
plt.show()

        



    
