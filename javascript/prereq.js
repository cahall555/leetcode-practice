const prereqsPossible = (numCourses, prereqs) => {
  const graph = buildGraph(numCourses, prereqs);
  const visited = new Set();
  const visiting = new Set();
  
  for (let node in graph) {
    if (hasCycle(graph, node, visited, visited)) return false;
  }
  return true;
};

const hasCycle = (graph, node, visited, visiting) => {
  if (visited.has(node)) return false;
  if (visiting.has(node)) return true;
  
  visiting.add(node);
  
  for (let neighbor of graph[node]) {
    if (hasCycle(graph, neighbor, visited, visiting)) return true;
  }
  
  visiting.delete(node);
  visited.add(node);
  
  return false;
};

const buildGraph = (numCourses, prereqs) => {
  const graph = {};
  
  for (let i = 0; i < numCourses; i += 1){
    graph[i] = [];
  }
  
  for (let prereq of prereqs) {
    const [a, b] = prereq;
    graph[a].push(String(b));
  }
  return graph;
};

