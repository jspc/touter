#!/usr/bin/env ruby
#
# Very, very simple UDP listener and JSON parser

require 'socket'
require 'json'

s = UDPSocket.new
s.bind(nil, 2002)

j, sender = s.recvfrom(10240)
data = JSON.parse j

puts "Received data about #{data.keys.first}"
puts "Sample repo: #{data[data.keys.first].sample['Path'] }"
