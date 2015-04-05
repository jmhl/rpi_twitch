require 'net/http'
require 'json'

def get_streams
  cs_url = URI('https://api.twitch.tv/kraken/streams?game=counter-strike:%20Global%20Offensive')
  resp = Net::HTTP.get(cs_url)
  streams = JSON.parse(resp)['streams']
end

def parse_response(streams)
  options = []

  streams.each_with_index do |stream, i|
    channel = stream['channel']
    name = channel['name']
    uri = channel['_links']['self']
    options.push({ name: name })
  end

  options
end

def open_stream(stream)
  system("livestreamer http://twitch.tv/#{stream[:name]} best 'omxplayer -o hdmi'")
  exit
end

def menu(streams, limit)
  puts "OPTIONS: 'q' to quit, number to pick a stream"
  puts '+++++ STREAMS +++++'

  streams.each_with_index do |stream, key|
    spaces = '     '
    spaces = spaces[(key + 1).to_s.length..spaces.length]
    puts "#{key + 1} #{spaces}|| #{stream[:name]}"
  end

  puts 'Pick a stream:'
  choice = gets.chomp

  if choice === 'q'
    exit
  elsif choice === 'm'
  elsif streams[choice.to_i - 1]
    open_stream(streams[choice.to_i - 1])
  end
end

def init
  resp = get_streams
  streams = parse_response(resp)
  limit = 10

  loop do
    menu(streams.first(limit), limit)
    limit += 10
  end
end

# init
puts get_streams
