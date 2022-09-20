# sap-api-integrations-purchasing-source-list-reads
sap-api-integrations-purchasing-source-list-reads は、他のすべての sap-api-integrations-purchasing-source-list-reads 作成更新の際の 参照元となる マスタレポジトリです。  
sap-api-integrations-purchasing-source-list-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 供給元一覧データを取得するマイクロサービスです。    
sap-api-integrations-purchasing-source-list-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-purchasing-source-list-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_PURCHASING_SOURCE_SRV_0001/overview  

## 動作環境  
sap-api-integrations-purchasing-source-list-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-purchasing-source-list-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-purchasing-source-list-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PURCHASING_SOURCE_SRV_0001/overview    
* APIサービス名(=baseURL): API_PURCHASING_SOURCE_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-purchasing-source-list-reads には、次の API をコールするためのリソースが含まれています。  

* A_PurchasingSource（供給元一覧）

## API への 値入力条件 の 初期値
sap-api-integrations-purchasing-source-list-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.List.Material（品目）
* inoutSDC.List.Plant（プラント）
* inoutSDC.List.Supplier（仕入先）
* inoutSDC.List.SupplyingPlant（供給プラント）
* inoutSDC.List.ValidityEndDate（有効終了日）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。

```
	"api_schema": "SAPPurchasingSourceListReads",
	"accepter": ["List"],
	"material_code": "21",
	"plant": "0001",
	"source_list_record": "1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPPurchasingSourceListReads",
	"accepter": ["All"],
	"material_code": "21",
	"plant": "0001",
	"source_list_record": "1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetPurchasingSourceList(material, plant, sourceListRecord, supplier, supplyingPlant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "List":
			func() {
				c.List(material, plant, sourceListRecord)
				wg.Done()
			}()
		case "Supplier":
			func() {
				c.Supplier(material, plant, sourceListRecord, supplier)
				wg.Done()
			}()
		case "SupplyingPlant":
			func() {
				c.SupplyingPlant(material, plant, sourceListRecord, supplyingPlant)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP List の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"Material" ～ "SourceOfSupplyIsFixed" は、/SAP_API_Output_Formatter/type.go 内 の Type List {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-purchasing-source-list-reads/SAP_API_Caller/caller.go#L66",
	"function": "sap-api-integrations-purchasing-source-list-reads/SAP_API_Caller.(*SAPAPICaller).List",
	"level": "INFO",
	"message": [
		{
			"Material": "21",
			"Plant": "0001",
			"SourceListRecord": "1",
			"ValidityStartDate": "1998-01-01",
			"ValidityEndDate": "9999-12-31",
			"Supplier": "100000",
			"PurchasingOrganization": "0001",
			"SupplyingPlant": "",
			"OrderQuantityUnit": "",
			"PurchaseOutlineAgreement": "",
			"SupplierIsFixed": false,
			"SourceOfSupplyIsBlocked": false,
			"MRPSourcingControl": "",
			"LastChangeDateTime": "2022-09-20T20:18:27+09:00",
			"IssgPlantIsFixed": false,
			"PurOutlineAgreementIsFixed": false,
			"SourceOfSupplyIsFixed": false
		}
	],
	"time": "2022-09-20T20:29:09+09:00"
}

```
