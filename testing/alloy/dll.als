one sig DLL {
header: lone Node
}
sig Node {
prev, link: lone Node,
elem: Int
}
// All nodes should be reachable from the header along the link.
fact Reachable {
Node = DLL.header.*link
}

fact Acyclic {
// The list has no directed cycle along link, i.e., no node is
// reachable from itself following one or more traversals along link.
-- TODO: Your code starts here.
  no n: Node | n in n.^link
}

pred UniqueElem() {
// Unique nodes contain unique elements.
-- TODO: Your code starts here.
  all n1, n2: Node | n1 != n2 implies n1.elem != n2.elem
}

pred Sorted() {
// The list is sorted in ascending order (<=) along link.
-- TODO: Your code starts here.
  all n1, n2: Node | n2 in n1.*link implies n1.elem <= n2.elem
}

pred ConsistentLinkAndPrev() {
// For any node n1 and n2, if n1.link = n2, then n2.prev = n1; and vice versa.
-- TODO: Your code starts here.
  all n1, n2: Node | (n2 = n1.link implies n2.prev = n1) and (n2 = n1.prev implies n2.link = n1)
}

pred AllConditions() {
  UniqueElem and Sorted and ConsistentLinkAndPrev
}

run AllConditions for 3
