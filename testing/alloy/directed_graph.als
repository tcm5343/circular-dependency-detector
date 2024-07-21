sig Node {
  edges: set Node
}

fact no_self_loop {
  all n: Node | n not in n.edges
}

pred cyclic {
  some n: Node | n in n.^edges
}

pred acyclic {
  no n: Node | n in n.^edges
}

run cyclic for 4  // 792 satisfiable test cases
run acyclic for 4  // 83 satisfiable test cases
