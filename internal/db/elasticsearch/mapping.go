package elasticsearch

//VideoMapping ..
const VideoMapping = `{
    "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 0,
        "analysis": {
            "analyzer": {
                "ik_max_synonym": {
                    "type": "custom",
                    "tokenizer": "ik_max_word",
                    "filter": [
                        "my_filter"
                    ]
                },
                "ik_smart_synonym": {
                    "type": "custom",
                    "tokenizer": "ik_smart",
                    "filter": [
                        "my_filter"
                    ]
                }
            },
            "filter": {
                "my_filter": {
                    "type": "synonym",
                    "synonyms_path": "analysis/synonym.txt"
                }
            }
        }
    },
    "mappings": {
        "include_in_all": "false",
        "dynamic": true,
        "properties": {
            "id": {
                "type": "long"
            },
            "title": {
                "type": "text",
                "fields": {
                    "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                    }
                },
                "analyzer": "ik_max_synonym",
                "search_analyzer": "ik_smart_synonym"
            },
            "desc": {
                "type": "text",
                "fields": {
                    "keyword": {
                        "type": "keyword",
                        "ignore_above": 5000
                    }
                },
                "analyzer": "ik_max_synonym",
                "search_analyzer": "ik_smart_synonym"
            },
            "pub_date": {
                "type": "long"
            },
            "cover": {
                "type": "keyword"
            },
            "total_num": {
                "type": "long"
            },
            "tags": {
                "type": "keyword"
            },
            "series_id": {
                "type": "long"
            },
            "series_name": {
                "type": "text",
                "fields": {
                    "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                    }
                },
                "analyzer": "ik_max_synonym",
                "search_analyzer": "ik_smart_synonym"
            },
            "series_alias": {
                "type": "text",
                "fields": {
                    "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                    }
                },
                "analyzer": "ik_max_synonym",
                "search_analyzer": "ik_smart_synonym"
            },
            "series_num": {
                "type": "long"
            }
        }
    }
}`
