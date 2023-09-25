# nifdiff

ニフクラリソース間の差分を取るツール

## つかいかた

[nifcloud-cli](https://github.com/nifcloud/nifcloud-cli) と同じように環境変数を設定します。

```
$ export NIFCLOUD_ACCESS_KEY_ID=<Your NIFCLOUD Access Key ID>
$ export NIFCLOUD_SECRET_ACCESS_KEY=<Your NIFCLOUD Secret Access Key>
```

このツールは [ニフクラリソースネーム（NRN）](https://docs.nifcloud.com/common/spec/nrn.htm) 形式でリソース名を識別します。`${nifcloud_id}` 部分は見ないので入力不要です。

```
nifdiff nrn:nifcloud:computing:jp-east-1::security_group:FirewallJE1 nrn:nifcloud:computing:jp-west-2::security_group:FirewallJW2
```

## 機能

* ニフクラリソース間で差分ないかどうかを確認できます。
* 同一アカウント内であればリージョンをまたいで比較できます。

## サポートされているリソース

* `nrn:nifcloud:computing:*:*:security_group:*` コンピューティング・ファイアウォールグループ
