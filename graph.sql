CREATE TABLE graphs (
    graph_id VARCHAR(255), -- id of graph
    graph_name VARCHAR(255), -- name of graph
    PRIMARY KEY (graph_id)
);

CREATE TABLE nodes (
    graph_id VARCHAR(255), -- id of the graph the node belongs to (foreign key)
    node_id VARCHAR(255), -- id of node
    node_name VARCHAR(255), -- name of node
    PRIMARY KEY(node_id),
    CONSTRAINT graph_fk FOREIGN KEY(graph_id) REFERENCES graphs(graph_id)
);

CREATE TABLE edges (
    graph_id VARCHAR(255), -- id of the graph the node belongs to (foreign key)
    edge_id VARCHAR(255), -- id of edge
    edge_name VARCHAR(255), -- name of edge
    from_node VARCHAR(255), -- id of node where edge starts (foreign key)
    to_node VARCHAR(255), -- id of node where edge ends (foreign key)
    PRIMARY KEY(edge_id),
    CONSTRAINT graph_fk FOREIGN KEY(graph_id) REFERENCES graphs(graph_id),
    CONSTRAINT from_node_fk FOREIGN KEY(from_node) REFERENCES nodes(node_id),
    CONSTRAINT to_node_fk FOREIGN KEY(to_node) REFERENCES nodes(node_id)
);