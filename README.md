# Haley Shopp submission for Tucows technical assessment.

1. Download a text file containing an XML file like the one described
The global variable downloadURL on line 12 of main.go can be updated with the URL for downloading the txt file, then the downloading 
is done at the beginning of the main function.

2. Parse the file to make sure it is syntactically and semantically correct
subject to the following restrictions

This is done in main.go using function parseXML to first read the XML directly into a Graph struct. 
Then the function cleanXML goes through this struct to make sure restrictions are satisfied.

3. Propose a normalized SQL schema to model this graphs in PostgreSQL
using standard SQL data types only. Please explain briefly each attribute
and relationship you propose. You can use SQL documentation facilities
and comments in order to do that.

This can be found in graph.sql and uses PostgreSQL syntax

4. Write an SQL query that finds cycles in a given graph, according to the
data model you proposed on item (3). You can use standard SQL99 or
PL/pgSQL functions.