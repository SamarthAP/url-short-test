# url-shortener

Given a shortened form or a url, this program redirects to the full url.
For example, if running locally, localhost:9000/goog will redirect to google.ca. Shortened urls are based on the shorurls.csv file, which 
maps the shortened url to the full url.

```
$ go build .
$ ./url-shortener -h
Usage of ./url-shortener:
  -csv string
        a csv file with the format shortened-url,long-url (default "shorturls.csv")
```

