import java.util.Set;
import java.util.HashSet;

public class IsDag {

    public static void main(String[] args) {
        Node n1 = new Node(1);
        Node n2 = new Node(2);
        Node n3 = new Node(3);

        n1.connect(n2);
        n2.connect(n3);

        System.out.printf("Cycle: %s\n",n1.isCyclic());

        Node m1 = new Node(1);
        Node m2 = new Node(2);
        Node m3 = new Node(3);

        m1.connect(m2);
        m2.connect(m3);
        m3.connect(m1);

        System.out.printf("Cycle: %s\n",m1.isCyclic());
    }

    static class Node {
        private int id;
        private Set<Node> adjacencies;

        Node(int id) {
            this.id = id;
            this.adjacencies = new HashSet<>();
        }

        public void connect(Node node) {
            this.adjacencies.add(node);
        }

        public boolean isCyclic() {
            Set<Node> visited = new HashSet<>();

            return this.DFS(visited);
        }

        private boolean DFS(Set<Node> visited) {
            if(visited == null) {
                return false;
            }

            if(visited.contains(this)) {
                return true;
            }

            visited.add(this);

            for(Node node : this.adjacencies) {
                return node.DFS(visited);
            }

            return false;
        }
    }
}