require 'yaml'
require 'aws-sdk'

creds = YAML.load(File.read('./aws.yml')['development'])

Aws::CloudWatch::Client.new(
  access_key_id: creds['access_key_id'],
  secret_access_key: creds['secret_access_key']
)
