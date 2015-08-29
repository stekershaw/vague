# Vague

Deserialising JSON to structs when the JSON types are a little vague and loose...

An example fragment of JSON that demonstrates such "looseness", as returned in the response to a [call to the WordPress.org plugins API](https://api.wordpress.org/plugins/info/1.0/saml-20-single-sign-on.json):

```
"ratings" : {
  "5":"2",
  "4":0,
  "3":0,
  "2":"1",
  "1":"1"
}
```

## Types

### vague.Int

A vague.Int is:

```
type Int int
```

Int implements the json.Unmarshaler interface and will attempt to unmarshal JSON strings and numbers to a vague.Int.

## References

+ [Post by Attila Oláh on JSON with loose schema in go](http://attilaolah.eu/2013/11/29/json-decoding-in-go/)
+ [Similar to the above from MAtty Williams](http://mattyjwilliams.blogspot.co.uk/2013/01/using-go-to-unmarshal-json-lists-with.html)
+ [Overview of JSON and go by Zack Bloom](https://eager.io/blog/go-and-json/)
