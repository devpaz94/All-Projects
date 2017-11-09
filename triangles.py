

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

        



    