task :build => ["get"]do 
	sh "go build"
end

task :get do 
	sh "go get"
end

namespace :run do
	
	task :local => ["build", "config:local"] do
		sh  "./schnapi"
	end

	task :dev => ["build", "config:dev"] do
		sh  "./schnapi"
	end
end

namespace :config do

	task :local do
		ENV['PORT']="4200"
		ENV['MONGOURI']="127.0.0.1:27017"
		ENV['DBNAME']="simplyactivedb"
		ENV['APIKEY']="1234"
		ENV['CNAME']="users"
	end

	task :mimic do
		ENV['PORT']="3800"
		ENV['MONGOURI']="127.0.0.1:27017"
		ENV['DBNAME']="mongoid"
		ENV['APIKEY']="1234"
		ENV['CNAME']="users"
	end
end