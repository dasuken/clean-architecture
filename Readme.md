## 参照
[Qiita](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)

## 特徴
外から中へ依存関係は一方通行。
Infrastructure => interface => usecase => domain

### Infrastructure (framework driver)
Database & Frameworkの層

**router.go**
var Routerにinit関数でmux.Routerを代入

**sqlHandler.go**
構造体SqlHandlerにmysqlのコネクションを代入
実際にdataのやりとりはinterfaces/databaseに実装する

### interfaces (interfaces adapter)
usecaseのポートに対して実装を提供する

**controller/user_controller.go**
usecaseのIteratorを呼び出す。iteratorから帰ってきた値をresponse

**database/sql_handler.go**
database/sqlの振る舞いをinterfaceにまとめたもの。database/sqlのsql.DB変数は外側のinfrastructure層で定義されているため、取得する事が出来ない。
そのため、振る舞いのみをinterfaceで表現している

**databse/user_repository.go**
DB操作を定義している。ビジネスロジック

### usecase
interfacesの実装に対するポートを定義

**user_iterator.go**
controllerに対するinput port。domainオブジェクトを返信する。domain層とのやりとりはRepositoryが行う

**user_repository.go**
interfaces/databaseに定義されているメソッドをinterfaceとして定義する。user_iteratorから呼び出される。

## domain
entityを定義。

## 個人的に困惑した部分。
**usecase/user_repositoryの役割**
実際にdatabaseのsqlを実装しているのはinterfaces層のdatabase/user_repository
