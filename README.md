#  circular-dependency-detector
A GitHub action to detect and output cycles and topological generations given a directed graph as an adjacency list.

To run:
docker build -t circular-dependency-detector . && docker run circular-dependency-detector


Todo:
Return 2d slice of nodes instead of integers from topological generations

Consider edge cases involving multi-graphs in topological generations

Allow the input file to define nodes as strings instead of just integers (use a set and map)
Input the NetworkX adjacency list format


## Musical Acknowledgements
Bob Dylan - Early Mornin' Rain
Galt MacDermot - Ripped Open By Metal Explosion

