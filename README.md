# AppEngine Datastore

## App Engineにデプロイ
```sh
gcloud app deploy
```

## Usage

`/users/` でランダムなUUIDをキーにして,
Datastore(Kind `User` )に検索。なければPutする.

### use Bench Tool
[tsenart/vegeta](https://github.com/tsenart/vegeta)
vegetaを使う例
```sh
echo "GET http://<YOUR_APPLICATION>.appspot.com/users/" | vegeta attack -rate=100 -duration=5s | vegeta report
```