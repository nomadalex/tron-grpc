#!/usr/bin/env ruby

require "fileutils"

$proto_dir = File.join(Dir.pwd, "proto")

def copy_proto(path)
    content = File.read(path).gsub(/go_package\s*=\s*"github.com\/tronprotocol\/grpc-gateway/, "go_package = \"github.com/fullstackwang/tron-grpc")

    dst = File.join($proto_dir, path)
    FileUtils.mkdir_p(File.dirname(dst))
    File.write(dst, content)
end

FileUtils.remove_dir(File.join($proto_dir, "core"), true)
FileUtils.remove_dir(File.join($proto_dir, "api"), true)

Dir.chdir(ARGV[0]) do
    Dir.glob("**/*.proto") do |fn|
        copy_proto(fn)
    end
end