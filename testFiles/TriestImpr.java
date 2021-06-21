import java.util.HashMap;
import java.util.ArrayList;
import java.util.HashSet;

public class TriestImpr implements DataStreamAlgo {
	public int M = 1;//storage capacity or sample size
	public int t = 0; //time
	public int D = 0;//triangle count estimate
	public double weight;
	//Store graph connections
	public HashMap<Integer, HashSet<Integer>> G= new HashMap<Integer, HashSet<Integer>>();
	//Store local triangle count for each vertex
	public HashMap<Integer, Integer> deltaV= new HashMap<Integer, Integer>();
	//store edges for faster search
	public ArrayList<Edge> S = new ArrayList<Edge>();
	/*
     * Constructor.
     * The parameter samsize denotes the size of the sample, i.e., the number of
     * edges that the algorithm can store.
     */
	public TriestImpr(int samsize) {
		M = samsize;
		weight = ((double)(M-2)/(double)M);
	}

	public void handleEdge(Edge edge) {
		int u = edge.u;
		int v = edge.v;
		if(u == v){ //invalid edge or it already exists
			return;
		}
		else if(G.containsKey(u)){
			if(G.get(u).contains(v)){
				return;
			}
		}
		t++;
		if(t > M){
			//update weight
			weight = weight / ((double)t-3) * ((double)t-1);
			//do reservoir sampling and update counts
			double prob = (double)M/(double)t;
			if(Math.random()<prob){
				//switch edge in
				swapIn(u,v);
			}
			else{
				accountForEdge(u,v);
			}
		}
		else{
			//add edge
			addEdge(u,v);
		}
	}

	public void swapIn(int u, int v) {
		removeEdge();
		addEdge(u,v);
	}

	public void accountForEdge(int u, int v) {
		//if vertices are disconnected from graph, no triangles will be formed
		if(G.get(u)==null || G.get(v)==null) {
			return;
		}
		//count mutual connections to get number of triangles broken
		HashSet<Integer> uConnections = G.get(u);
		HashSet<Integer> vConnections = G.get(v);
		int mutualCxn = 0;
		for(int i : uConnections){
			if(vConnections.contains(i)){
				mutualCxn++;
			}
		}
		//update total triangle count estimate
		D += weight * mutualCxn;
	}

	public void addEdge(int u, int v) {
		S.add(new Edge(u+" "+v));
		//initialize HashSets for connections and local counts if not present
		if(G.get(u)==null){
			G.put(u, new HashSet<Integer>());
			deltaV.put(u, 0);
		}
		if(G.get(v)==null){
			G.put(v, new HashSet<Integer>());
			deltaV.put(v, 0);
		}
		//count mutual connections to get number of triangles made
		HashSet<Integer> uConnections = G.get(u);
		HashSet<Integer> vConnections = G.get(v);
		int mutualCxn = 0;
		for(int i : uConnections){
			if(vConnections.contains(i)){
				mutualCxn++;
				//increment local triangle count for mutual connection
				deltaV.put(i, deltaV.get(i)+1);
			}
		}

		//now add edge to G
		vConnections.add(u);
		uConnections.add(v);

		//update local triangle counts for vertices
		deltaV.put(u, deltaV.get(u)+mutualCxn);
		deltaV.put(v, deltaV.get(v)+mutualCxn);

		//update total triangle count estimate
		if(t>M){
			D += weight * mutualCxn;
		}else{
			D += mutualCxn;
		}
	}

	public void removeEdge() {
		//remove random edge from S
		Edge removed = S.remove((int)(Math.random()*S.size()));

		//now remove edge from G too
		int u = removed.u;
		int v = removed.v;
		HashSet<Integer> uConnections = G.get(u);
		HashSet<Integer> vConnections = G.get(v);
		uConnections.remove(v);
		vConnections.remove(u);

		//count mutual connections to get number of triangles broken
		int mutualCxn = 0;
		for(int i : uConnections){
			if(vConnections.contains(i)){
				mutualCxn++;
				//decrement local triangle count for mutual connection
				deltaV.put(i, deltaV.get(i)-1);
			}
		}

		//update local triangle counts for vertices
		deltaV.put(u, deltaV.get(u)-mutualCxn);
		deltaV.put(v, deltaV.get(v)-mutualCxn);

	}

	public int getEstimate() {
		return D;
	}
}