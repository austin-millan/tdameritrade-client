# MarginAccountCurrentBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedInterest** | Pointer to **float32** |  | [optional] 
**AvailableFunds** | Pointer to **float32** |  | [optional] 
**AvailableFundsNonMarginableTrade** | Pointer to **float32** |  | [optional] 
**BondValue** | Pointer to **float32** |  | [optional] 
**BuyingPower** | Pointer to **float32** |  | [optional] 
**BuyingPowerNonMarginableTrade** | Pointer to **float32** |  | [optional] 
**CashBalance** | Pointer to **float32** |  | [optional] 
**CashReceipts** | Pointer to **float32** |  | [optional] 
**DayTradingBuyingPower** | Pointer to **float32** |  | [optional] 
**DayTradingBuyingPowerCall** | Pointer to **float32** |  | [optional] 
**Equity** | Pointer to **float32** |  | [optional] 
**EquityPercentage** | Pointer to **float32** |  | [optional] 
**IsInCall** | Pointer to **bool** |  | [optional] 
**LiquidationValue** | Pointer to **float32** |  | [optional] 
**LongMarginValue** | Pointer to **float32** |  | [optional] 
**LongMarketValue** | Pointer to **float32** |  | [optional] 
**LongOptionMarketValue** | Pointer to **float32** |  | [optional] 
**MaintenanceCall** | Pointer to **float32** |  | [optional] 
**MaintenanceRequirement** | Pointer to **float32** |  | [optional] 
**MarginBalance** | Pointer to **float32** |  | [optional] 
**MoneyMarketFund** | Pointer to **float32** |  | [optional] 
**MutualFundValue** | Pointer to **float32** |  | [optional] 
**OptionBuyingPower** | Pointer to **float32** |  | [optional] 
**PendingDeposits** | Pointer to **float32** |  | [optional] 
**RegTCall** | Pointer to **float32** |  | [optional] 
**Savings** | Pointer to **float32** |  | [optional] 
**ShortBalance** | Pointer to **float32** |  | [optional] 
**ShortMarginValue** | Pointer to **float32** |  | [optional] 
**ShortMarketValue** | Pointer to **float32** |  | [optional] 
**ShortOptionMarketValue** | Pointer to **float32** |  | [optional] 
**Sma** | Pointer to **float32** |  | [optional] 
**StockBuyingPower** | Pointer to **float32** |  | [optional] 

## Methods

### NewMarginAccountCurrentBalances

`func NewMarginAccountCurrentBalances() *MarginAccountCurrentBalances`

NewMarginAccountCurrentBalances instantiates a new MarginAccountCurrentBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountCurrentBalancesWithDefaults

`func NewMarginAccountCurrentBalancesWithDefaults() *MarginAccountCurrentBalances`

NewMarginAccountCurrentBalancesWithDefaults instantiates a new MarginAccountCurrentBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedInterest

`func (o *MarginAccountCurrentBalances) GetAccruedInterest() float32`

GetAccruedInterest returns the AccruedInterest field if non-nil, zero value otherwise.

### GetAccruedInterestOk

`func (o *MarginAccountCurrentBalances) GetAccruedInterestOk() (*float32, bool)`

GetAccruedInterestOk returns a tuple with the AccruedInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedInterest

`func (o *MarginAccountCurrentBalances) SetAccruedInterest(v float32)`

SetAccruedInterest sets AccruedInterest field to given value.

### HasAccruedInterest

`func (o *MarginAccountCurrentBalances) HasAccruedInterest() bool`

HasAccruedInterest returns a boolean if a field has been set.

### GetAvailableFunds

`func (o *MarginAccountCurrentBalances) GetAvailableFunds() float32`

GetAvailableFunds returns the AvailableFunds field if non-nil, zero value otherwise.

### GetAvailableFundsOk

`func (o *MarginAccountCurrentBalances) GetAvailableFundsOk() (*float32, bool)`

GetAvailableFundsOk returns a tuple with the AvailableFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFunds

`func (o *MarginAccountCurrentBalances) SetAvailableFunds(v float32)`

SetAvailableFunds sets AvailableFunds field to given value.

### HasAvailableFunds

`func (o *MarginAccountCurrentBalances) HasAvailableFunds() bool`

HasAvailableFunds returns a boolean if a field has been set.

### GetAvailableFundsNonMarginableTrade

`func (o *MarginAccountCurrentBalances) GetAvailableFundsNonMarginableTrade() float32`

GetAvailableFundsNonMarginableTrade returns the AvailableFundsNonMarginableTrade field if non-nil, zero value otherwise.

### GetAvailableFundsNonMarginableTradeOk

`func (o *MarginAccountCurrentBalances) GetAvailableFundsNonMarginableTradeOk() (*float32, bool)`

GetAvailableFundsNonMarginableTradeOk returns a tuple with the AvailableFundsNonMarginableTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFundsNonMarginableTrade

`func (o *MarginAccountCurrentBalances) SetAvailableFundsNonMarginableTrade(v float32)`

SetAvailableFundsNonMarginableTrade sets AvailableFundsNonMarginableTrade field to given value.

### HasAvailableFundsNonMarginableTrade

`func (o *MarginAccountCurrentBalances) HasAvailableFundsNonMarginableTrade() bool`

HasAvailableFundsNonMarginableTrade returns a boolean if a field has been set.

### GetBondValue

`func (o *MarginAccountCurrentBalances) GetBondValue() float32`

GetBondValue returns the BondValue field if non-nil, zero value otherwise.

### GetBondValueOk

`func (o *MarginAccountCurrentBalances) GetBondValueOk() (*float32, bool)`

GetBondValueOk returns a tuple with the BondValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondValue

`func (o *MarginAccountCurrentBalances) SetBondValue(v float32)`

SetBondValue sets BondValue field to given value.

### HasBondValue

`func (o *MarginAccountCurrentBalances) HasBondValue() bool`

HasBondValue returns a boolean if a field has been set.

### GetBuyingPower

`func (o *MarginAccountCurrentBalances) GetBuyingPower() float32`

GetBuyingPower returns the BuyingPower field if non-nil, zero value otherwise.

### GetBuyingPowerOk

`func (o *MarginAccountCurrentBalances) GetBuyingPowerOk() (*float32, bool)`

GetBuyingPowerOk returns a tuple with the BuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingPower

`func (o *MarginAccountCurrentBalances) SetBuyingPower(v float32)`

SetBuyingPower sets BuyingPower field to given value.

### HasBuyingPower

`func (o *MarginAccountCurrentBalances) HasBuyingPower() bool`

HasBuyingPower returns a boolean if a field has been set.

### GetBuyingPowerNonMarginableTrade

`func (o *MarginAccountCurrentBalances) GetBuyingPowerNonMarginableTrade() float32`

GetBuyingPowerNonMarginableTrade returns the BuyingPowerNonMarginableTrade field if non-nil, zero value otherwise.

### GetBuyingPowerNonMarginableTradeOk

`func (o *MarginAccountCurrentBalances) GetBuyingPowerNonMarginableTradeOk() (*float32, bool)`

GetBuyingPowerNonMarginableTradeOk returns a tuple with the BuyingPowerNonMarginableTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingPowerNonMarginableTrade

`func (o *MarginAccountCurrentBalances) SetBuyingPowerNonMarginableTrade(v float32)`

SetBuyingPowerNonMarginableTrade sets BuyingPowerNonMarginableTrade field to given value.

### HasBuyingPowerNonMarginableTrade

`func (o *MarginAccountCurrentBalances) HasBuyingPowerNonMarginableTrade() bool`

HasBuyingPowerNonMarginableTrade returns a boolean if a field has been set.

### GetCashBalance

`func (o *MarginAccountCurrentBalances) GetCashBalance() float32`

GetCashBalance returns the CashBalance field if non-nil, zero value otherwise.

### GetCashBalanceOk

`func (o *MarginAccountCurrentBalances) GetCashBalanceOk() (*float32, bool)`

GetCashBalanceOk returns a tuple with the CashBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBalance

`func (o *MarginAccountCurrentBalances) SetCashBalance(v float32)`

SetCashBalance sets CashBalance field to given value.

### HasCashBalance

`func (o *MarginAccountCurrentBalances) HasCashBalance() bool`

HasCashBalance returns a boolean if a field has been set.

### GetCashReceipts

`func (o *MarginAccountCurrentBalances) GetCashReceipts() float32`

GetCashReceipts returns the CashReceipts field if non-nil, zero value otherwise.

### GetCashReceiptsOk

`func (o *MarginAccountCurrentBalances) GetCashReceiptsOk() (*float32, bool)`

GetCashReceiptsOk returns a tuple with the CashReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashReceipts

`func (o *MarginAccountCurrentBalances) SetCashReceipts(v float32)`

SetCashReceipts sets CashReceipts field to given value.

### HasCashReceipts

`func (o *MarginAccountCurrentBalances) HasCashReceipts() bool`

HasCashReceipts returns a boolean if a field has been set.

### GetDayTradingBuyingPower

`func (o *MarginAccountCurrentBalances) GetDayTradingBuyingPower() float32`

GetDayTradingBuyingPower returns the DayTradingBuyingPower field if non-nil, zero value otherwise.

### GetDayTradingBuyingPowerOk

`func (o *MarginAccountCurrentBalances) GetDayTradingBuyingPowerOk() (*float32, bool)`

GetDayTradingBuyingPowerOk returns a tuple with the DayTradingBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradingBuyingPower

`func (o *MarginAccountCurrentBalances) SetDayTradingBuyingPower(v float32)`

SetDayTradingBuyingPower sets DayTradingBuyingPower field to given value.

### HasDayTradingBuyingPower

`func (o *MarginAccountCurrentBalances) HasDayTradingBuyingPower() bool`

HasDayTradingBuyingPower returns a boolean if a field has been set.

### GetDayTradingBuyingPowerCall

`func (o *MarginAccountCurrentBalances) GetDayTradingBuyingPowerCall() float32`

GetDayTradingBuyingPowerCall returns the DayTradingBuyingPowerCall field if non-nil, zero value otherwise.

### GetDayTradingBuyingPowerCallOk

`func (o *MarginAccountCurrentBalances) GetDayTradingBuyingPowerCallOk() (*float32, bool)`

GetDayTradingBuyingPowerCallOk returns a tuple with the DayTradingBuyingPowerCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradingBuyingPowerCall

`func (o *MarginAccountCurrentBalances) SetDayTradingBuyingPowerCall(v float32)`

SetDayTradingBuyingPowerCall sets DayTradingBuyingPowerCall field to given value.

### HasDayTradingBuyingPowerCall

`func (o *MarginAccountCurrentBalances) HasDayTradingBuyingPowerCall() bool`

HasDayTradingBuyingPowerCall returns a boolean if a field has been set.

### GetEquity

`func (o *MarginAccountCurrentBalances) GetEquity() float32`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *MarginAccountCurrentBalances) GetEquityOk() (*float32, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *MarginAccountCurrentBalances) SetEquity(v float32)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *MarginAccountCurrentBalances) HasEquity() bool`

HasEquity returns a boolean if a field has been set.

### GetEquityPercentage

`func (o *MarginAccountCurrentBalances) GetEquityPercentage() float32`

GetEquityPercentage returns the EquityPercentage field if non-nil, zero value otherwise.

### GetEquityPercentageOk

`func (o *MarginAccountCurrentBalances) GetEquityPercentageOk() (*float32, bool)`

GetEquityPercentageOk returns a tuple with the EquityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityPercentage

`func (o *MarginAccountCurrentBalances) SetEquityPercentage(v float32)`

SetEquityPercentage sets EquityPercentage field to given value.

### HasEquityPercentage

`func (o *MarginAccountCurrentBalances) HasEquityPercentage() bool`

HasEquityPercentage returns a boolean if a field has been set.

### GetIsInCall

`func (o *MarginAccountCurrentBalances) GetIsInCall() bool`

GetIsInCall returns the IsInCall field if non-nil, zero value otherwise.

### GetIsInCallOk

`func (o *MarginAccountCurrentBalances) GetIsInCallOk() (*bool, bool)`

GetIsInCallOk returns a tuple with the IsInCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInCall

`func (o *MarginAccountCurrentBalances) SetIsInCall(v bool)`

SetIsInCall sets IsInCall field to given value.

### HasIsInCall

`func (o *MarginAccountCurrentBalances) HasIsInCall() bool`

HasIsInCall returns a boolean if a field has been set.

### GetLiquidationValue

`func (o *MarginAccountCurrentBalances) GetLiquidationValue() float32`

GetLiquidationValue returns the LiquidationValue field if non-nil, zero value otherwise.

### GetLiquidationValueOk

`func (o *MarginAccountCurrentBalances) GetLiquidationValueOk() (*float32, bool)`

GetLiquidationValueOk returns a tuple with the LiquidationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationValue

`func (o *MarginAccountCurrentBalances) SetLiquidationValue(v float32)`

SetLiquidationValue sets LiquidationValue field to given value.

### HasLiquidationValue

`func (o *MarginAccountCurrentBalances) HasLiquidationValue() bool`

HasLiquidationValue returns a boolean if a field has been set.

### GetLongMarginValue

`func (o *MarginAccountCurrentBalances) GetLongMarginValue() float32`

GetLongMarginValue returns the LongMarginValue field if non-nil, zero value otherwise.

### GetLongMarginValueOk

`func (o *MarginAccountCurrentBalances) GetLongMarginValueOk() (*float32, bool)`

GetLongMarginValueOk returns a tuple with the LongMarginValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMarginValue

`func (o *MarginAccountCurrentBalances) SetLongMarginValue(v float32)`

SetLongMarginValue sets LongMarginValue field to given value.

### HasLongMarginValue

`func (o *MarginAccountCurrentBalances) HasLongMarginValue() bool`

HasLongMarginValue returns a boolean if a field has been set.

### GetLongMarketValue

`func (o *MarginAccountCurrentBalances) GetLongMarketValue() float32`

GetLongMarketValue returns the LongMarketValue field if non-nil, zero value otherwise.

### GetLongMarketValueOk

`func (o *MarginAccountCurrentBalances) GetLongMarketValueOk() (*float32, bool)`

GetLongMarketValueOk returns a tuple with the LongMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMarketValue

`func (o *MarginAccountCurrentBalances) SetLongMarketValue(v float32)`

SetLongMarketValue sets LongMarketValue field to given value.

### HasLongMarketValue

`func (o *MarginAccountCurrentBalances) HasLongMarketValue() bool`

HasLongMarketValue returns a boolean if a field has been set.

### GetLongOptionMarketValue

`func (o *MarginAccountCurrentBalances) GetLongOptionMarketValue() float32`

GetLongOptionMarketValue returns the LongOptionMarketValue field if non-nil, zero value otherwise.

### GetLongOptionMarketValueOk

`func (o *MarginAccountCurrentBalances) GetLongOptionMarketValueOk() (*float32, bool)`

GetLongOptionMarketValueOk returns a tuple with the LongOptionMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongOptionMarketValue

`func (o *MarginAccountCurrentBalances) SetLongOptionMarketValue(v float32)`

SetLongOptionMarketValue sets LongOptionMarketValue field to given value.

### HasLongOptionMarketValue

`func (o *MarginAccountCurrentBalances) HasLongOptionMarketValue() bool`

HasLongOptionMarketValue returns a boolean if a field has been set.

### GetMaintenanceCall

`func (o *MarginAccountCurrentBalances) GetMaintenanceCall() float32`

GetMaintenanceCall returns the MaintenanceCall field if non-nil, zero value otherwise.

### GetMaintenanceCallOk

`func (o *MarginAccountCurrentBalances) GetMaintenanceCallOk() (*float32, bool)`

GetMaintenanceCallOk returns a tuple with the MaintenanceCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCall

`func (o *MarginAccountCurrentBalances) SetMaintenanceCall(v float32)`

SetMaintenanceCall sets MaintenanceCall field to given value.

### HasMaintenanceCall

`func (o *MarginAccountCurrentBalances) HasMaintenanceCall() bool`

HasMaintenanceCall returns a boolean if a field has been set.

### GetMaintenanceRequirement

`func (o *MarginAccountCurrentBalances) GetMaintenanceRequirement() float32`

GetMaintenanceRequirement returns the MaintenanceRequirement field if non-nil, zero value otherwise.

### GetMaintenanceRequirementOk

`func (o *MarginAccountCurrentBalances) GetMaintenanceRequirementOk() (*float32, bool)`

GetMaintenanceRequirementOk returns a tuple with the MaintenanceRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceRequirement

`func (o *MarginAccountCurrentBalances) SetMaintenanceRequirement(v float32)`

SetMaintenanceRequirement sets MaintenanceRequirement field to given value.

### HasMaintenanceRequirement

`func (o *MarginAccountCurrentBalances) HasMaintenanceRequirement() bool`

HasMaintenanceRequirement returns a boolean if a field has been set.

### GetMarginBalance

`func (o *MarginAccountCurrentBalances) GetMarginBalance() float32`

GetMarginBalance returns the MarginBalance field if non-nil, zero value otherwise.

### GetMarginBalanceOk

`func (o *MarginAccountCurrentBalances) GetMarginBalanceOk() (*float32, bool)`

GetMarginBalanceOk returns a tuple with the MarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBalance

`func (o *MarginAccountCurrentBalances) SetMarginBalance(v float32)`

SetMarginBalance sets MarginBalance field to given value.

### HasMarginBalance

`func (o *MarginAccountCurrentBalances) HasMarginBalance() bool`

HasMarginBalance returns a boolean if a field has been set.

### GetMoneyMarketFund

`func (o *MarginAccountCurrentBalances) GetMoneyMarketFund() float32`

GetMoneyMarketFund returns the MoneyMarketFund field if non-nil, zero value otherwise.

### GetMoneyMarketFundOk

`func (o *MarginAccountCurrentBalances) GetMoneyMarketFundOk() (*float32, bool)`

GetMoneyMarketFundOk returns a tuple with the MoneyMarketFund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyMarketFund

`func (o *MarginAccountCurrentBalances) SetMoneyMarketFund(v float32)`

SetMoneyMarketFund sets MoneyMarketFund field to given value.

### HasMoneyMarketFund

`func (o *MarginAccountCurrentBalances) HasMoneyMarketFund() bool`

HasMoneyMarketFund returns a boolean if a field has been set.

### GetMutualFundValue

`func (o *MarginAccountCurrentBalances) GetMutualFundValue() float32`

GetMutualFundValue returns the MutualFundValue field if non-nil, zero value otherwise.

### GetMutualFundValueOk

`func (o *MarginAccountCurrentBalances) GetMutualFundValueOk() (*float32, bool)`

GetMutualFundValueOk returns a tuple with the MutualFundValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualFundValue

`func (o *MarginAccountCurrentBalances) SetMutualFundValue(v float32)`

SetMutualFundValue sets MutualFundValue field to given value.

### HasMutualFundValue

`func (o *MarginAccountCurrentBalances) HasMutualFundValue() bool`

HasMutualFundValue returns a boolean if a field has been set.

### GetOptionBuyingPower

`func (o *MarginAccountCurrentBalances) GetOptionBuyingPower() float32`

GetOptionBuyingPower returns the OptionBuyingPower field if non-nil, zero value otherwise.

### GetOptionBuyingPowerOk

`func (o *MarginAccountCurrentBalances) GetOptionBuyingPowerOk() (*float32, bool)`

GetOptionBuyingPowerOk returns a tuple with the OptionBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionBuyingPower

`func (o *MarginAccountCurrentBalances) SetOptionBuyingPower(v float32)`

SetOptionBuyingPower sets OptionBuyingPower field to given value.

### HasOptionBuyingPower

`func (o *MarginAccountCurrentBalances) HasOptionBuyingPower() bool`

HasOptionBuyingPower returns a boolean if a field has been set.

### GetPendingDeposits

`func (o *MarginAccountCurrentBalances) GetPendingDeposits() float32`

GetPendingDeposits returns the PendingDeposits field if non-nil, zero value otherwise.

### GetPendingDepositsOk

`func (o *MarginAccountCurrentBalances) GetPendingDepositsOk() (*float32, bool)`

GetPendingDepositsOk returns a tuple with the PendingDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDeposits

`func (o *MarginAccountCurrentBalances) SetPendingDeposits(v float32)`

SetPendingDeposits sets PendingDeposits field to given value.

### HasPendingDeposits

`func (o *MarginAccountCurrentBalances) HasPendingDeposits() bool`

HasPendingDeposits returns a boolean if a field has been set.

### GetRegTCall

`func (o *MarginAccountCurrentBalances) GetRegTCall() float32`

GetRegTCall returns the RegTCall field if non-nil, zero value otherwise.

### GetRegTCallOk

`func (o *MarginAccountCurrentBalances) GetRegTCallOk() (*float32, bool)`

GetRegTCallOk returns a tuple with the RegTCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegTCall

`func (o *MarginAccountCurrentBalances) SetRegTCall(v float32)`

SetRegTCall sets RegTCall field to given value.

### HasRegTCall

`func (o *MarginAccountCurrentBalances) HasRegTCall() bool`

HasRegTCall returns a boolean if a field has been set.

### GetSavings

`func (o *MarginAccountCurrentBalances) GetSavings() float32`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *MarginAccountCurrentBalances) GetSavingsOk() (*float32, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *MarginAccountCurrentBalances) SetSavings(v float32)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *MarginAccountCurrentBalances) HasSavings() bool`

HasSavings returns a boolean if a field has been set.

### GetShortBalance

`func (o *MarginAccountCurrentBalances) GetShortBalance() float32`

GetShortBalance returns the ShortBalance field if non-nil, zero value otherwise.

### GetShortBalanceOk

`func (o *MarginAccountCurrentBalances) GetShortBalanceOk() (*float32, bool)`

GetShortBalanceOk returns a tuple with the ShortBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortBalance

`func (o *MarginAccountCurrentBalances) SetShortBalance(v float32)`

SetShortBalance sets ShortBalance field to given value.

### HasShortBalance

`func (o *MarginAccountCurrentBalances) HasShortBalance() bool`

HasShortBalance returns a boolean if a field has been set.

### GetShortMarginValue

`func (o *MarginAccountCurrentBalances) GetShortMarginValue() float32`

GetShortMarginValue returns the ShortMarginValue field if non-nil, zero value otherwise.

### GetShortMarginValueOk

`func (o *MarginAccountCurrentBalances) GetShortMarginValueOk() (*float32, bool)`

GetShortMarginValueOk returns a tuple with the ShortMarginValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMarginValue

`func (o *MarginAccountCurrentBalances) SetShortMarginValue(v float32)`

SetShortMarginValue sets ShortMarginValue field to given value.

### HasShortMarginValue

`func (o *MarginAccountCurrentBalances) HasShortMarginValue() bool`

HasShortMarginValue returns a boolean if a field has been set.

### GetShortMarketValue

`func (o *MarginAccountCurrentBalances) GetShortMarketValue() float32`

GetShortMarketValue returns the ShortMarketValue field if non-nil, zero value otherwise.

### GetShortMarketValueOk

`func (o *MarginAccountCurrentBalances) GetShortMarketValueOk() (*float32, bool)`

GetShortMarketValueOk returns a tuple with the ShortMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMarketValue

`func (o *MarginAccountCurrentBalances) SetShortMarketValue(v float32)`

SetShortMarketValue sets ShortMarketValue field to given value.

### HasShortMarketValue

`func (o *MarginAccountCurrentBalances) HasShortMarketValue() bool`

HasShortMarketValue returns a boolean if a field has been set.

### GetShortOptionMarketValue

`func (o *MarginAccountCurrentBalances) GetShortOptionMarketValue() float32`

GetShortOptionMarketValue returns the ShortOptionMarketValue field if non-nil, zero value otherwise.

### GetShortOptionMarketValueOk

`func (o *MarginAccountCurrentBalances) GetShortOptionMarketValueOk() (*float32, bool)`

GetShortOptionMarketValueOk returns a tuple with the ShortOptionMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortOptionMarketValue

`func (o *MarginAccountCurrentBalances) SetShortOptionMarketValue(v float32)`

SetShortOptionMarketValue sets ShortOptionMarketValue field to given value.

### HasShortOptionMarketValue

`func (o *MarginAccountCurrentBalances) HasShortOptionMarketValue() bool`

HasShortOptionMarketValue returns a boolean if a field has been set.

### GetSma

`func (o *MarginAccountCurrentBalances) GetSma() float32`

GetSma returns the Sma field if non-nil, zero value otherwise.

### GetSmaOk

`func (o *MarginAccountCurrentBalances) GetSmaOk() (*float32, bool)`

GetSmaOk returns a tuple with the Sma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSma

`func (o *MarginAccountCurrentBalances) SetSma(v float32)`

SetSma sets Sma field to given value.

### HasSma

`func (o *MarginAccountCurrentBalances) HasSma() bool`

HasSma returns a boolean if a field has been set.

### GetStockBuyingPower

`func (o *MarginAccountCurrentBalances) GetStockBuyingPower() float32`

GetStockBuyingPower returns the StockBuyingPower field if non-nil, zero value otherwise.

### GetStockBuyingPowerOk

`func (o *MarginAccountCurrentBalances) GetStockBuyingPowerOk() (*float32, bool)`

GetStockBuyingPowerOk returns a tuple with the StockBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockBuyingPower

`func (o *MarginAccountCurrentBalances) SetStockBuyingPower(v float32)`

SetStockBuyingPower sets StockBuyingPower field to given value.

### HasStockBuyingPower

`func (o *MarginAccountCurrentBalances) HasStockBuyingPower() bool`

HasStockBuyingPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


