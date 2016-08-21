# Source this for testing / debugging

function es_wipe_index() {
  curl -XDELETE 'http://192.168.99.100:9200/files' ; echo
}
		
function es_match_all_query() {
  curl -XPOST '192.168.99.100:9200/files/_search?pretty' -d '
{
  "query": { "match_all": {} },
  "size": 10
}' ; echo
}

function es_match_dir_query() {
  curl -XPOST '192.168.99.100:9200/files/_search?pretty' -d '
{
  "query": { "match": { "_type": "directory"} },
  "size": 100
}' ; echo
}
