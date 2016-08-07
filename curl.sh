# Generic tasks
curl -XDELETE 'http://192.168.99.100:9200/files' ; echo
		
curl -XPOST '192.168.99.100:9200/files/_search?pretty' -d '
{
	  "query": { "match_all": {} },
	    "size": 10
}' ; echo
