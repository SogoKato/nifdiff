# nifdiff

ニフクラリソース間の差分を取るツール

## つかいかた

[バイナリをダウンロード](https://github.com/SogoKato/nifdiff/releases) または Docker で実行できます。

[nifcloud-cli](https://github.com/nifcloud/nifcloud-cli) と同じように環境変数を設定します。

```
$ export NIFCLOUD_ACCESS_KEY_ID=<Your NIFCLOUD Access Key ID>
$ export NIFCLOUD_SECRET_ACCESS_KEY=<Your NIFCLOUD Secret Access Key>
```

このツールは [ニフクラリソースネーム（NRN）](https://docs.nifcloud.com/common/spec/nrn.htm) 形式でリソース名を識別します。`${nifcloud_id}` 部分は見ないので入力不要です。

```
nifdiff nrn:nifcloud:computing:jp-east-1::security_group:FirewallJE1 nrn:nifcloud:computing:jp-west-2::security_group:FirewallJW2
```

Docker で実行する場合:

```
docker run -e NIFCLOUD_ACCESS_KEY_ID=${NIFCLOUD_ACCESS_KEY_ID} -e NIFCLOUD_SECRET_ACCESS_KEY=${NIFCLOUD_SECRET_ACCESS_KEY} ghcr.io/sogokato/nifdiff nrn:nifcloud:computing:jp-east-1::security_group:FirewallJE1 nrn:nifcloud:computing:jp-west-2::security_group:FirewallJW2
```

出力サンプル

```
Mismatch:
  types.SecurityGroupInfo{
-       AvailabilityZone:        &"east-11",
+       AvailabilityZone:        &"jp-west-21",
-       GroupDescription:        &"イーストイチイチ",
+       GroupDescription:        &"",
        GroupLogFilterBroadcast: &true,
        GroupLogFilterNetBios:   &false,
-       GroupLogLimit:           &1000,
+       GroupLogLimit:           &100000,
-       GroupName:               &"FirewallJE1",
+       GroupName:               &"FirewallJW2",
        GroupRuleLimit:          &100,
        GroupStatus:             &"applied",
        ... // 2 ignored fields
        IpPermissions: []types.IpPermissions{
                {
                        ... // 1 ignored and 4 identical fields
                        IpProtocol: &"TCP",
                        IpRanges:   {{CidrIp: &"192.168.0.0/24"}},
-                       ToPort:     &444,
+                       ToPort:     &443,
                        ... // 1 ignored field
                },
                {Description: &"", FromPort: &80, InOut: &"IN", IpProtocol: &"TCP", ...},
        },
        OwnerId:   &"",
        RouterSet: nil,
        ... // 1 ignored and 1 identical fields
  }
```

## 機能

* ニフクラリソース間で差分ないかどうかを確認できます。
* 同一アカウント内であればリージョンをまたいで比較できます。

## サポートされているリソース

* `nrn:nifcloud:computing:*:*:security_group:*` コンピューティング・ファイアウォールグループ
