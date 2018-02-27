```uml
@startuml

actor USER
participant Browser as UA
participant BO
participant AirID as ID

UA -> BO: API call
BO -> BO: EntryPointにてissueState
UA <-- BO: 401

UA -> ID: 認証要求(返り値にあるurlを叩く)
opt 初回認証時(PLFセッションがない場合)
  UA <-- ID: ログインページへリダイレクト
  USER -> UA: ID/pass入力
UA -> ID: 認証情報post
end

BO <-- ID: 事前登録しているURLにリダイレクト(/auth/authorized)

group 認証の妥当性検証
  BO -> BO: OAuth2CallbackFilterが処理
  BO -> BO: リダイレクトURLよりcode取得
  BO -> BO: リダイレクトURLよりstate取得
  BO -> BO: state復元(consumeState)
  BO -> BO: state削除
end

group 認証詳細情報取得
  BO -> ID: issueToken
  BO <-- ID: accessToken/refreshToken
  BO -> ID: checkToken
  BO <-- ID: 法人情報(今までは店舗情報だったけど...)
  BO -> ID: サービス利用情報取得・更新API
  BO <-- ID
end

group 認証状態保存
  BO -> BO: Authenticatorで認証後処理
  BO -> BO: 取得した情報を基にAuthedModel作成
  BO -> BO: 作成したAuthedModelを基にAuthentication作成
  BO -> BO: AuthenticationをSecurityContextに保存
  BO -> BO: sessionIdをcookieに保存
end

UA <-- BO: TOPにリダイレクト
@enduml
```
