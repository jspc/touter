#!/usr/bin/env ruby
#
# Very, very simple UDP listener and JSON parser

require 'socket'
require 'json'

s = UDPSocket.new
s.bind(nil, 2002)

j, sender = s.recvfrom(10240)
data = JSON.parse j

hostnames = data.keys
puts "Received data about #{hostnames}"

hostnames.each do |host|
  puts "#{host} gives us #{data[host].length} items"
  (0..4).each do |index|
    repo = data[host][index]
    puts "#{repo['Path']} is on branch #{repo['Branch']} (shasum: #{repo['Sha']}"
  end
end
