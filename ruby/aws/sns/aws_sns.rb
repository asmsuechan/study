require 'yaml'
require 'aws-sdk-v1'
 
AWS.config(YAML.load_file('./aws.yml')['development'])
 
sns = AWS::SNS.new
client = sns.client

# platform_applicationを新規作成
client.create_platform_application({
  name: "cw_user_device_id_1",
  platform: "GCM",
  attributes: {
    'PlatformPrincipal' => '',
    'PlatformCredential' => YAML.load_file('./aws.yml')['development']['platform_credential']
  }
})

# あるplatform_application_arnのendpointを全て取得する
endpoints = client.list_endpoints_by_platform_application(
  platform_application_arn: YAML.load_file('./aws.yml')['development']['platform_application_arn']
)

# データを作成しているがうまく動かない
data = { data: {
  title: "asdf",
  text: "hogehoge", 
  post: {
    id: 1, 
    title: "fuga"
  }
}}.to_json.gsub('"', "\\\"")
message = <<-EOD
{"GCM": "#{data}"}
EOD
endpoints[:endpoints].each do |endpoint|
  # disableのendpointがあるとエラーを吐くため
  if endpoint[:attributes]["Enabled"] == "true"
    puts endpoint
    resp = client.publish(
      target_arn: endpoint[:endpoint_arn],
      message: message,
      message_structure: "json"
    )
  end
end

ep_arn = client.create_platform_endpoint({
  platform_application_arn: YAML.load_file('./aws.yml')['development']['platform_application_arn'],
  token: YAML.load_file('./aws.yml')['development']['token']
})

# token(registration_id)からsns_endpoint_arnを取得する
puts ep_arn[:endpoint_arn]

# IAMロールのユーザーにSNSの権限追加しないとUnauthorizeエラー出ます
#
# 参考
# http://www.absolute-keitarou.net/blog/?p=1206
# http://qiita.com/g08m11/items/463af2665f514b35b9f7
#
# 出たエラー
# aws_sns.rb:4:in `<main>': uninitialized constant AWS (NameError)
# aws-sdkのv1.xはAWSネームスペースを使い、aws-sdkのv2.xはAwsネームスペースを使う。fuck。
# http://qiita.com/azusanakano/items/d34dc3471bfbe97c8a61
# http://stackoverflow.com/questions/22826432/uninitialized-constant-aws-nameerror
