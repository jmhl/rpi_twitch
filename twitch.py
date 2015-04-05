import urllib2
import json
from pprint import pprint

streams_as_json = urllib2.urlopen("https://api.twitch.tv/kraken/streams?game=counter-strike:%20Global%20Offensive")
streams = json.load(streams_as_json)["streams"]
options = {}

for stream, i in enumerate(streams):
    pprint(stream)
    channel = stream["channel"]
    name = channel["name"]
    url = channel["links"]["self"]
    pprint(name, url)
