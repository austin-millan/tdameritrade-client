# MarginAccountInitialBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountValue** | Pointer to **float32** |  | [optional] 
**AccruedInterest** | Pointer to **float32** |  | [optional] 
**AvailableFundsNonMarginableTrade** | Pointer to **float32** |  | [optional] 
**BondValue** | Pointer to **float32** |  | [optional] 
**BuyingPower** | Pointer to **float32** |  | [optional] 
**CashAvailableForTrading** | Pointer to **float32** |  | [optional] 
**CashBalance** | Pointer to **float32** |  | [optional] 
**CashReceipts** | Pointer to **float32** |  | [optional] 
**DayTradingBuyingPower** | Pointer to **float32** |  | [optional] 
**DayTradingBuyingPowerCall** | Pointer to **float32** |  | [optional] 
**DayTradingEquityCall** | Pointer to **float32** |  | [optional] 
**Equity** | Pointer to **float32** |  | [optional] 
**EquityPercentage** | Pointer to **float32** |  | [optional] 
**IsInCall** | Pointer to **bool** |  | [optional] 
**LiquidationValue** | Pointer to **float32** |  | [optional] 
**LongMarginValue** | Pointer to **float32** |  | [optional] 
**LongOptionMarketValue** | Pointer to **float32** |  | [optional] 
**LongStockValue** | Pointer to **float32** |  | [optional] 
**MaintenanceCall** | Pointer to **float32** |  | [optional] 
**MaintenanceRequirement** | Pointer to **float32** |  | [optional] 
**Margin** | Pointer to **float32** |  | [optional] 
**MarginBalance** | Pointer to **float32** |  | [optional] 
**MarginEquity** | Pointer to **float32** |  | [optional] 
**MoneyMarketFund** | Pointer to **float32** |  | [optional] 
**MutualFundValue** | Pointer to **float32** |  | [optional] 
**PendingDeposits** | Pointer to **float32** |  | [optional] 
**RegTCall** | Pointer to **float32** |  | [optional] 
**ShortBalance** | Pointer to **float32** |  | [optional] 
**ShortMarginValue** | Pointer to **float32** |  | [optional] 
**ShortOptionMarketValue** | Pointer to **float32** |  | [optional] 
**ShortStockValue** | Pointer to **float32** |  | [optional] 
**TotalCash** | Pointer to **float32** |  | [optional] 
**UnsettledCash** | Pointer to **float32** |  | [optional] 

## Methods

### NewMarginAccountInitialBalances

`func NewMarginAccountInitialBalances() *MarginAccountInitialBalances`

NewMarginAccountInitialBalances instantiates a new MarginAccountInitialBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountInitialBalancesWithDefaults

`func NewMarginAccountInitialBalancesWithDefaults() *MarginAccountInitialBalances`

NewMarginAccountInitialBalancesWithDefaults instantiates a new MarginAccountInitialBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountValue

`func (o *MarginAccountInitialBalances) GetAccountValue() float32`

GetAccountValue returns the AccountValue field if non-nil, zero value otherwise.

### GetAccountValueOk

`func (o *MarginAccountInitialBalances) GetAccountValueOk() (*float32, bool)`

GetAccountValueOk returns a tuple with the AccountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountValue

`func (o *MarginAccountInitialBalances) SetAccountValue(v float32)`

SetAccountValue sets AccountValue field to given value.

### HasAccountValue

`func (o *MarginAccountInitialBalances) HasAccountValue() bool`

HasAccountValue returns a boolean if a field has been set.

### GetAccruedInterest

`func (o *MarginAccountInitialBalances) GetAccruedInterest() float32`

GetAccruedInterest returns the AccruedInterest field if non-nil, zero value otherwise.

### GetAccruedInterestOk

`func (o *MarginAccountInitialBalances) GetAccruedInterestOk() (*float32, bool)`

GetAccruedInterestOk returns a tuple with the AccruedInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedInterest

`func (o *MarginAccountInitialBalances) SetAccruedInterest(v float32)`

SetAccruedInterest sets AccruedInterest field to given value.

### HasAccruedInterest

`func (o *MarginAccountInitialBalances) HasAccruedInterest() bool`

HasAccruedInterest returns a boolean if a field has been set.

### GetAvailableFundsNonMarginableTrade

`func (o *MarginAccountInitialBalances) GetAvailableFundsNonMarginableTrade() float32`

GetAvailableFundsNonMarginableTrade returns the AvailableFundsNonMarginableTrade field if non-nil, zero value otherwise.

### GetAvailableFundsNonMarginableTradeOk

`func (o *MarginAccountInitialBalances) GetAvailableFundsNonMarginableTradeOk() (*float32, bool)`

GetAvailableFundsNonMarginableTradeOk returns a tuple with the AvailableFundsNonMarginableTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFundsNonMarginableTrade

`func (o *MarginAccountInitialBalances) SetAvailableFundsNonMarginableTrade(v float32)`

SetAvailableFundsNonMarginableTrade sets AvailableFundsNonMarginableTrade field to given value.

### HasAvailableFundsNonMarginableTrade

`func (o *MarginAccountInitialBalances) HasAvailableFundsNonMarginableTrade() bool`

HasAvailableFundsNonMarginableTrade returns a boolean if a field has been set.

### GetBondValue

`func (o *MarginAccountInitialBalances) GetBondValue() float32`

GetBondValue returns the BondValue field if non-nil, zero value otherwise.

### GetBondValueOk

`func (o *MarginAccountInitialBalances) GetBondValueOk() (*float32, bool)`

GetBondValueOk returns a tuple with the BondValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondValue

`func (o *MarginAccountInitialBalances) SetBondValue(v float32)`

SetBondValue sets BondValue field to given value.

### HasBondValue

`func (o *MarginAccountInitialBalances) HasBondValue() bool`

HasBondValue returns a boolean if a field has been set.

### GetBuyingPower

`func (o *MarginAccountInitialBalances) GetBuyingPower() float32`

GetBuyingPower returns the BuyingPower field if non-nil, zero value otherwise.

### GetBuyingPowerOk

`func (o *MarginAccountInitialBalances) GetBuyingPowerOk() (*float32, bool)`

GetBuyingPowerOk returns a tuple with the BuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingPower

`func (o *MarginAccountInitialBalances) SetBuyingPower(v float32)`

SetBuyingPower sets BuyingPower field to given value.

### HasBuyingPower

`func (o *MarginAccountInitialBalances) HasBuyingPower() bool`

HasBuyingPower returns a boolean if a field has been set.

### GetCashAvailableForTrading

`func (o *MarginAccountInitialBalances) GetCashAvailableForTrading() float32`

GetCashAvailableForTrading returns the CashAvailableForTrading field if non-nil, zero value otherwise.

### GetCashAvailableForTradingOk

`func (o *MarginAccountInitialBalances) GetCashAvailableForTradingOk() (*float32, bool)`

GetCashAvailableForTradingOk returns a tuple with the CashAvailableForTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAvailableForTrading

`func (o *MarginAccountInitialBalances) SetCashAvailableForTrading(v float32)`

SetCashAvailableForTrading sets CashAvailableForTrading field to given value.

### HasCashAvailableForTrading

`func (o *MarginAccountInitialBalances) HasCashAvailableForTrading() bool`

HasCashAvailableForTrading returns a boolean if a field has been set.

### GetCashBalance

`func (o *MarginAccountInitialBalances) GetCashBalance() float32`

GetCashBalance returns the CashBalance field if non-nil, zero value otherwise.

### GetCashBalanceOk

`func (o *MarginAccountInitialBalances) GetCashBalanceOk() (*float32, bool)`

GetCashBalanceOk returns a tuple with the CashBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBalance

`func (o *MarginAccountInitialBalances) SetCashBalance(v float32)`

SetCashBalance sets CashBalance field to given value.

### HasCashBalance

`func (o *MarginAccountInitialBalances) HasCashBalance() bool`

HasCashBalance returns a boolean if a field has been set.

### GetCashReceipts

`func (o *MarginAccountInitialBalances) GetCashReceipts() float32`

GetCashReceipts returns the CashReceipts field if non-nil, zero value otherwise.

### GetCashReceiptsOk

`func (o *MarginAccountInitialBalances) GetCashReceiptsOk() (*float32, bool)`

GetCashReceiptsOk returns a tuple with the CashReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashReceipts

`func (o *MarginAccountInitialBalances) SetCashReceipts(v float32)`

SetCashReceipts sets CashReceipts field to given value.

### HasCashReceipts

`func (o *MarginAccountInitialBalances) HasCashReceipts() bool`

HasCashReceipts returns a boolean if a field has been set.

### GetDayTradingBuyingPower

`func (o *MarginAccountInitialBalances) GetDayTradingBuyingPower() float32`

GetDayTradingBuyingPower returns the DayTradingBuyingPower field if non-nil, zero value otherwise.

### GetDayTradingBuyingPowerOk

`func (o *MarginAccountInitialBalances) GetDayTradingBuyingPowerOk() (*float32, bool)`

GetDayTradingBuyingPowerOk returns a tuple with the DayTradingBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradingBuyingPower

`func (o *MarginAccountInitialBalances) SetDayTradingBuyingPower(v float32)`

SetDayTradingBuyingPower sets DayTradingBuyingPower field to given value.

### HasDayTradingBuyingPower

`func (o *MarginAccountInitialBalances) HasDayTradingBuyingPower() bool`

HasDayTradingBuyingPower returns a boolean if a field has been set.

### GetDayTradingBuyingPowerCall

`func (o *MarginAccountInitialBalances) GetDayTradingBuyingPowerCall() float32`

GetDayTradingBuyingPowerCall returns the DayTradingBuyingPowerCall field if non-nil, zero value otherwise.

### GetDayTradingBuyingPowerCallOk

`func (o *MarginAccountInitialBalances) GetDayTradingBuyingPowerCallOk() (*float32, bool)`

GetDayTradingBuyingPowerCallOk returns a tuple with the DayTradingBuyingPowerCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradingBuyingPowerCall

`func (o *MarginAccountInitialBalances) SetDayTradingBuyingPowerCall(v float32)`

SetDayTradingBuyingPowerCall sets DayTradingBuyingPowerCall field to given value.

### HasDayTradingBuyingPowerCall

`func (o *MarginAccountInitialBalances) HasDayTradingBuyingPowerCall() bool`

HasDayTradingBuyingPowerCall returns a boolean if a field has been set.

### GetDayTradingEquityCall

`func (o *MarginAccountInitialBalances) GetDayTradingEquityCall() float32`

GetDayTradingEquityCall returns the DayTradingEquityCall field if non-nil, zero value otherwise.

### GetDayTradingEquityCallOk

`func (o *MarginAccountInitialBalances) GetDayTradingEquityCallOk() (*float32, bool)`

GetDayTradingEquityCallOk returns a tuple with the DayTradingEquityCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradingEquityCall

`func (o *MarginAccountInitialBalances) SetDayTradingEquityCall(v float32)`

SetDayTradingEquityCall sets DayTradingEquityCall field to given value.

### HasDayTradingEquityCall

`func (o *MarginAccountInitialBalances) HasDayTradingEquityCall() bool`

HasDayTradingEquityCall returns a boolean if a field has been set.

### GetEquity

`func (o *MarginAccountInitialBalances) GetEquity() float32`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *MarginAccountInitialBalances) GetEquityOk() (*float32, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *MarginAccountInitialBalances) SetEquity(v float32)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *MarginAccountInitialBalances) HasEquity() bool`

HasEquity returns a boolean if a field has been set.

### GetEquityPercentage

`func (o *MarginAccountInitialBalances) GetEquityPercentage() float32`

GetEquityPercentage returns the EquityPercentage field if non-nil, zero value otherwise.

### GetEquityPercentageOk

`func (o *MarginAccountInitialBalances) GetEquityPercentageOk() (*float32, bool)`

GetEquityPercentageOk returns a tuple with the EquityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityPercentage

`func (o *MarginAccountInitialBalances) SetEquityPercentage(v float32)`

SetEquityPercentage sets EquityPercentage field to given value.

### HasEquityPercentage

`func (o *MarginAccountInitialBalances) HasEquityPercentage() bool`

HasEquityPercentage returns a boolean if a field has been set.

### GetIsInCall

`func (o *MarginAccountInitialBalances) GetIsInCall() bool`

GetIsInCall returns the IsInCall field if non-nil, zero value otherwise.

### GetIsInCallOk

`func (o *MarginAccountInitialBalances) GetIsInCallOk() (*bool, bool)`

GetIsInCallOk returns a tuple with the IsInCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInCall

`func (o *MarginAccountInitialBalances) SetIsInCall(v bool)`

SetIsInCall sets IsInCall field to given value.

### HasIsInCall

`func (o *MarginAccountInitialBalances) HasIsInCall() bool`

HasIsInCall returns a boolean if a field has been set.

### GetLiquidationValue

`func (o *MarginAccountInitialBalances) GetLiquidationValue() float32`

GetLiquidationValue returns the LiquidationValue field if non-nil, zero value otherwise.

### GetLiquidationValueOk

`func (o *MarginAccountInitialBalances) GetLiquidationValueOk() (*float32, bool)`

GetLiquidationValueOk returns a tuple with the LiquidationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationValue

`func (o *MarginAccountInitialBalances) SetLiquidationValue(v float32)`

SetLiquidationValue sets LiquidationValue field to given value.

### HasLiquidationValue

`func (o *MarginAccountInitialBalances) HasLiquidationValue() bool`

HasLiquidationValue returns a boolean if a field has been set.

### GetLongMarginValue

`func (o *MarginAccountInitialBalances) GetLongMarginValue() float32`

GetLongMarginValue returns the LongMarginValue field if non-nil, zero value otherwise.

### GetLongMarginValueOk

`func (o *MarginAccountInitialBalances) GetLongMarginValueOk() (*float32, bool)`

GetLongMarginValueOk returns a tuple with the LongMarginValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMarginValue

`func (o *MarginAccountInitialBalances) SetLongMarginValue(v float32)`

SetLongMarginValue sets LongMarginValue field to given value.

### HasLongMarginValue

`func (o *MarginAccountInitialBalances) HasLongMarginValue() bool`

HasLongMarginValue returns a boolean if a field has been set.

### GetLongOptionMarketValue

`func (o *MarginAccountInitialBalances) GetLongOptionMarketValue() float32`

GetLongOptionMarketValue returns the LongOptionMarketValue field if non-nil, zero value otherwise.

### GetLongOptionMarketValueOk

`func (o *MarginAccountInitialBalances) GetLongOptionMarketValueOk() (*float32, bool)`

GetLongOptionMarketValueOk returns a tuple with the LongOptionMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongOptionMarketValue

`func (o *MarginAccountInitialBalances) SetLongOptionMarketValue(v float32)`

SetLongOptionMarketValue sets LongOptionMarketValue field to given value.

### HasLongOptionMarketValue

`func (o *MarginAccountInitialBalances) HasLongOptionMarketValue() bool`

HasLongOptionMarketValue returns a boolean if a field has been set.

### GetLongStockValue

`func (o *MarginAccountInitialBalances) GetLongStockValue() float32`

GetLongStockValue returns the LongStockValue field if non-nil, zero value otherwise.

### GetLongStockValueOk

`func (o *MarginAccountInitialBalances) GetLongStockValueOk() (*float32, bool)`

GetLongStockValueOk returns a tuple with the LongStockValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongStockValue

`func (o *MarginAccountInitialBalances) SetLongStockValue(v float32)`

SetLongStockValue sets LongStockValue field to given value.

### HasLongStockValue

`func (o *MarginAccountInitialBalances) HasLongStockValue() bool`

HasLongStockValue returns a boolean if a field has been set.

### GetMaintenanceCall

`func (o *MarginAccountInitialBalances) GetMaintenanceCall() float32`

GetMaintenanceCall returns the MaintenanceCall field if non-nil, zero value otherwise.

### GetMaintenanceCallOk

`func (o *MarginAccountInitialBalances) GetMaintenanceCallOk() (*float32, bool)`

GetMaintenanceCallOk returns a tuple with the MaintenanceCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCall

`func (o *MarginAccountInitialBalances) SetMaintenanceCall(v float32)`

SetMaintenanceCall sets MaintenanceCall field to given value.

### HasMaintenanceCall

`func (o *MarginAccountInitialBalances) HasMaintenanceCall() bool`

HasMaintenanceCall returns a boolean if a field has been set.

### GetMaintenanceRequirement

`func (o *MarginAccountInitialBalances) GetMaintenanceRequirement() float32`

GetMaintenanceRequirement returns the MaintenanceRequirement field if non-nil, zero value otherwise.

### GetMaintenanceRequirementOk

`func (o *MarginAccountInitialBalances) GetMaintenanceRequirementOk() (*float32, bool)`

GetMaintenanceRequirementOk returns a tuple with the MaintenanceRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceRequirement

`func (o *MarginAccountInitialBalances) SetMaintenanceRequirement(v float32)`

SetMaintenanceRequirement sets MaintenanceRequirement field to given value.

### HasMaintenanceRequirement

`func (o *MarginAccountInitialBalances) HasMaintenanceRequirement() bool`

HasMaintenanceRequirement returns a boolean if a field has been set.

### GetMargin

`func (o *MarginAccountInitialBalances) GetMargin() float32`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *MarginAccountInitialBalances) GetMarginOk() (*float32, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *MarginAccountInitialBalances) SetMargin(v float32)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *MarginAccountInitialBalances) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetMarginBalance

`func (o *MarginAccountInitialBalances) GetMarginBalance() float32`

GetMarginBalance returns the MarginBalance field if non-nil, zero value otherwise.

### GetMarginBalanceOk

`func (o *MarginAccountInitialBalances) GetMarginBalanceOk() (*float32, bool)`

GetMarginBalanceOk returns a tuple with the MarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBalance

`func (o *MarginAccountInitialBalances) SetMarginBalance(v float32)`

SetMarginBalance sets MarginBalance field to given value.

### HasMarginBalance

`func (o *MarginAccountInitialBalances) HasMarginBalance() bool`

HasMarginBalance returns a boolean if a field has been set.

### GetMarginEquity

`func (o *MarginAccountInitialBalances) GetMarginEquity() float32`

GetMarginEquity returns the MarginEquity field if non-nil, zero value otherwise.

### GetMarginEquityOk

`func (o *MarginAccountInitialBalances) GetMarginEquityOk() (*float32, bool)`

GetMarginEquityOk returns a tuple with the MarginEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginEquity

`func (o *MarginAccountInitialBalances) SetMarginEquity(v float32)`

SetMarginEquity sets MarginEquity field to given value.

### HasMarginEquity

`func (o *MarginAccountInitialBalances) HasMarginEquity() bool`

HasMarginEquity returns a boolean if a field has been set.

### GetMoneyMarketFund

`func (o *MarginAccountInitialBalances) GetMoneyMarketFund() float32`

GetMoneyMarketFund returns the MoneyMarketFund field if non-nil, zero value otherwise.

### GetMoneyMarketFundOk

`func (o *MarginAccountInitialBalances) GetMoneyMarketFundOk() (*float32, bool)`

GetMoneyMarketFundOk returns a tuple with the MoneyMarketFund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyMarketFund

`func (o *MarginAccountInitialBalances) SetMoneyMarketFund(v float32)`

SetMoneyMarketFund sets MoneyMarketFund field to given value.

### HasMoneyMarketFund

`func (o *MarginAccountInitialBalances) HasMoneyMarketFund() bool`

HasMoneyMarketFund returns a boolean if a field has been set.

### GetMutualFundValue

`func (o *MarginAccountInitialBalances) GetMutualFundValue() float32`

GetMutualFundValue returns the MutualFundValue field if non-nil, zero value otherwise.

### GetMutualFundValueOk

`func (o *MarginAccountInitialBalances) GetMutualFundValueOk() (*float32, bool)`

GetMutualFundValueOk returns a tuple with the MutualFundValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualFundValue

`func (o *MarginAccountInitialBalances) SetMutualFundValue(v float32)`

SetMutualFundValue sets MutualFundValue field to given value.

### HasMutualFundValue

`func (o *MarginAccountInitialBalances) HasMutualFundValue() bool`

HasMutualFundValue returns a boolean if a field has been set.

### GetPendingDeposits

`func (o *MarginAccountInitialBalances) GetPendingDeposits() float32`

GetPendingDeposits returns the PendingDeposits field if non-nil, zero value otherwise.

### GetPendingDepositsOk

`func (o *MarginAccountInitialBalances) GetPendingDepositsOk() (*float32, bool)`

GetPendingDepositsOk returns a tuple with the PendingDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDeposits

`func (o *MarginAccountInitialBalances) SetPendingDeposits(v float32)`

SetPendingDeposits sets PendingDeposits field to given value.

### HasPendingDeposits

`func (o *MarginAccountInitialBalances) HasPendingDeposits() bool`

HasPendingDeposits returns a boolean if a field has been set.

### GetRegTCall

`func (o *MarginAccountInitialBalances) GetRegTCall() float32`

GetRegTCall returns the RegTCall field if non-nil, zero value otherwise.

### GetRegTCallOk

`func (o *MarginAccountInitialBalances) GetRegTCallOk() (*float32, bool)`

GetRegTCallOk returns a tuple with the RegTCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegTCall

`func (o *MarginAccountInitialBalances) SetRegTCall(v float32)`

SetRegTCall sets RegTCall field to given value.

### HasRegTCall

`func (o *MarginAccountInitialBalances) HasRegTCall() bool`

HasRegTCall returns a boolean if a field has been set.

### GetShortBalance

`func (o *MarginAccountInitialBalances) GetShortBalance() float32`

GetShortBalance returns the ShortBalance field if non-nil, zero value otherwise.

### GetShortBalanceOk

`func (o *MarginAccountInitialBalances) GetShortBalanceOk() (*float32, bool)`

GetShortBalanceOk returns a tuple with the ShortBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortBalance

`func (o *MarginAccountInitialBalances) SetShortBalance(v float32)`

SetShortBalance sets ShortBalance field to given value.

### HasShortBalance

`func (o *MarginAccountInitialBalances) HasShortBalance() bool`

HasShortBalance returns a boolean if a field has been set.

### GetShortMarginValue

`func (o *MarginAccountInitialBalances) GetShortMarginValue() float32`

GetShortMarginValue returns the ShortMarginValue field if non-nil, zero value otherwise.

### GetShortMarginValueOk

`func (o *MarginAccountInitialBalances) GetShortMarginValueOk() (*float32, bool)`

GetShortMarginValueOk returns a tuple with the ShortMarginValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMarginValue

`func (o *MarginAccountInitialBalances) SetShortMarginValue(v float32)`

SetShortMarginValue sets ShortMarginValue field to given value.

### HasShortMarginValue

`func (o *MarginAccountInitialBalances) HasShortMarginValue() bool`

HasShortMarginValue returns a boolean if a field has been set.

### GetShortOptionMarketValue

`func (o *MarginAccountInitialBalances) GetShortOptionMarketValue() float32`

GetShortOptionMarketValue returns the ShortOptionMarketValue field if non-nil, zero value otherwise.

### GetShortOptionMarketValueOk

`func (o *MarginAccountInitialBalances) GetShortOptionMarketValueOk() (*float32, bool)`

GetShortOptionMarketValueOk returns a tuple with the ShortOptionMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortOptionMarketValue

`func (o *MarginAccountInitialBalances) SetShortOptionMarketValue(v float32)`

SetShortOptionMarketValue sets ShortOptionMarketValue field to given value.

### HasShortOptionMarketValue

`func (o *MarginAccountInitialBalances) HasShortOptionMarketValue() bool`

HasShortOptionMarketValue returns a boolean if a field has been set.

### GetShortStockValue

`func (o *MarginAccountInitialBalances) GetShortStockValue() float32`

GetShortStockValue returns the ShortStockValue field if non-nil, zero value otherwise.

### GetShortStockValueOk

`func (o *MarginAccountInitialBalances) GetShortStockValueOk() (*float32, bool)`

GetShortStockValueOk returns a tuple with the ShortStockValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortStockValue

`func (o *MarginAccountInitialBalances) SetShortStockValue(v float32)`

SetShortStockValue sets ShortStockValue field to given value.

### HasShortStockValue

`func (o *MarginAccountInitialBalances) HasShortStockValue() bool`

HasShortStockValue returns a boolean if a field has been set.

### GetTotalCash

`func (o *MarginAccountInitialBalances) GetTotalCash() float32`

GetTotalCash returns the TotalCash field if non-nil, zero value otherwise.

### GetTotalCashOk

`func (o *MarginAccountInitialBalances) GetTotalCashOk() (*float32, bool)`

GetTotalCashOk returns a tuple with the TotalCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCash

`func (o *MarginAccountInitialBalances) SetTotalCash(v float32)`

SetTotalCash sets TotalCash field to given value.

### HasTotalCash

`func (o *MarginAccountInitialBalances) HasTotalCash() bool`

HasTotalCash returns a boolean if a field has been set.

### GetUnsettledCash

`func (o *MarginAccountInitialBalances) GetUnsettledCash() float32`

GetUnsettledCash returns the UnsettledCash field if non-nil, zero value otherwise.

### GetUnsettledCashOk

`func (o *MarginAccountInitialBalances) GetUnsettledCashOk() (*float32, bool)`

GetUnsettledCashOk returns a tuple with the UnsettledCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsettledCash

`func (o *MarginAccountInitialBalances) SetUnsettledCash(v float32)`

SetUnsettledCash sets UnsettledCash field to given value.

### HasUnsettledCash

`func (o *MarginAccountInitialBalances) HasUnsettledCash() bool`

HasUnsettledCash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


