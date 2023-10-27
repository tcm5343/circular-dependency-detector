#!/usr/bin/python3 

import matplotlib.pyplot as plt
import networkx as nx
import os

fig = plt.figure(figsize=(10, 10))
ax = plt.subplot(111)
path = os.getcwd()
G = nx.read_adjlist("list/adjacency_list_cycle_17.csv", create_using=nx.DiGraph())
pos = nx.spring_layout(G, 50)
#pos = nx.spectral_layout(G)
nx.draw_networkx(G, with_labels=True, arrows=True)
#nx.draw_networkx_nodes(G, pos, node_color="r", alpha=0.6, node_size=500)
#nx.draw_networkx_edges(G, pos, alpha=0.4, arrows=True)
#edge_labels = nx.get_edge_attributes(G,'weight')
#nx.draw_networkx_edge_labels(G, pos, edge_labels=edge_labels)

plt.axis("off")
plt.show()
