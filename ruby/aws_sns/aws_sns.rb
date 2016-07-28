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
endpoints[:endpoints].each do |endpoint|
  client.publish(
    target_arn: endpoint[:endpoint_arn],
    message: 'push_notification'
  )
end

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
