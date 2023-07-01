# FundamentalFundamental

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beta** | Pointer to **float32** |  | [optional] 
**BookValuePerShare** | Pointer to **float32** |  | [optional] 
**CurrentRatio** | Pointer to **float32** |  | [optional] 
**DivGrowthRate3Year** | Pointer to **float32** |  | [optional] 
**DividendAmount** | Pointer to **float32** |  | [optional] 
**DividendDate** | Pointer to **string** |  | [optional] 
**DividendPayAmount** | Pointer to **float32** |  | [optional] 
**DividendPayDate** | Pointer to **string** |  | [optional] 
**DividendYield** | Pointer to **float32** |  | [optional] 
**EpsChange** | Pointer to **float32** |  | [optional] 
**EpsChangePercentTTM** | Pointer to **float32** |  | [optional] 
**EpsChangeYear** | Pointer to **float32** |  | [optional] 
**EpsTTM** | Pointer to **float32** |  | [optional] 
**GrossMarginMRQ** | Pointer to **float32** |  | [optional] 
**GrossMarginTTM** | Pointer to **float32** |  | [optional] 
**High52** | Pointer to **float32** |  | [optional] 
**InterestCoverage** | Pointer to **float32** |  | [optional] 
**Low52** | Pointer to **float32** |  | [optional] 
**LtDebtToEquity** | Pointer to **float32** |  | [optional] 
**MarketCap** | Pointer to **float32** |  | [optional] 
**MarketCapFloat** | Pointer to **float32** |  | [optional] 
**NetProfitMarginMRQ** | Pointer to **float32** |  | [optional] 
**NetProfitMarginTTM** | Pointer to **float32** |  | [optional] 
**OperatingMarginMRQ** | Pointer to **float32** |  | [optional] 
**OperatingMarginTTM** | Pointer to **float32** |  | [optional] 
**PbRatio** | Pointer to **float32** |  | [optional] 
**PcfRatio** | Pointer to **float32** |  | [optional] 
**PeRatio** | Pointer to **float32** |  | [optional] 
**PegRatio** | Pointer to **float32** |  | [optional] 
**PrRatio** | Pointer to **float32** |  | [optional] 
**QuickRatio** | Pointer to **float32** |  | [optional] 
**ReturnOnAssets** | Pointer to **float32** |  | [optional] 
**ReturnOnEquity** | Pointer to **float32** |  | [optional] 
**ReturnOnInvestment** | Pointer to **float32** |  | [optional] 
**RevChangeIn** | Pointer to **float32** |  | [optional] 
**RevChangeTTM** | Pointer to **float32** |  | [optional] 
**RevChangeYear** | Pointer to **float32** |  | [optional] 
**SharesOutstanding** | Pointer to **float32** |  | [optional] 
**ShortIntDayToCover** | Pointer to **float32** |  | [optional] 
**ShortIntToFloat** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TotalDebtToCapital** | Pointer to **float32** |  | [optional] 
**TotalDebtToEquity** | Pointer to **float32** |  | [optional] 
**Vol10DayAvg** | Pointer to **float32** |  | [optional] 
**Vol1DayAvg** | Pointer to **float32** |  | [optional] 
**Vol3MonthAvg** | Pointer to **float32** |  | [optional] 

## Methods

### NewFundamentalFundamental

`func NewFundamentalFundamental() *FundamentalFundamental`

NewFundamentalFundamental instantiates a new FundamentalFundamental object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalFundamentalWithDefaults

`func NewFundamentalFundamentalWithDefaults() *FundamentalFundamental`

NewFundamentalFundamentalWithDefaults instantiates a new FundamentalFundamental object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeta

`func (o *FundamentalFundamental) GetBeta() float32`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *FundamentalFundamental) GetBetaOk() (*float32, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *FundamentalFundamental) SetBeta(v float32)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *FundamentalFundamental) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetBookValuePerShare

`func (o *FundamentalFundamental) GetBookValuePerShare() float32`

GetBookValuePerShare returns the BookValuePerShare field if non-nil, zero value otherwise.

### GetBookValuePerShareOk

`func (o *FundamentalFundamental) GetBookValuePerShareOk() (*float32, bool)`

GetBookValuePerShareOk returns a tuple with the BookValuePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookValuePerShare

`func (o *FundamentalFundamental) SetBookValuePerShare(v float32)`

SetBookValuePerShare sets BookValuePerShare field to given value.

### HasBookValuePerShare

`func (o *FundamentalFundamental) HasBookValuePerShare() bool`

HasBookValuePerShare returns a boolean if a field has been set.

### GetCurrentRatio

`func (o *FundamentalFundamental) GetCurrentRatio() float32`

GetCurrentRatio returns the CurrentRatio field if non-nil, zero value otherwise.

### GetCurrentRatioOk

`func (o *FundamentalFundamental) GetCurrentRatioOk() (*float32, bool)`

GetCurrentRatioOk returns a tuple with the CurrentRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRatio

`func (o *FundamentalFundamental) SetCurrentRatio(v float32)`

SetCurrentRatio sets CurrentRatio field to given value.

### HasCurrentRatio

`func (o *FundamentalFundamental) HasCurrentRatio() bool`

HasCurrentRatio returns a boolean if a field has been set.

### GetDivGrowthRate3Year

`func (o *FundamentalFundamental) GetDivGrowthRate3Year() float32`

GetDivGrowthRate3Year returns the DivGrowthRate3Year field if non-nil, zero value otherwise.

### GetDivGrowthRate3YearOk

`func (o *FundamentalFundamental) GetDivGrowthRate3YearOk() (*float32, bool)`

GetDivGrowthRate3YearOk returns a tuple with the DivGrowthRate3Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivGrowthRate3Year

`func (o *FundamentalFundamental) SetDivGrowthRate3Year(v float32)`

SetDivGrowthRate3Year sets DivGrowthRate3Year field to given value.

### HasDivGrowthRate3Year

`func (o *FundamentalFundamental) HasDivGrowthRate3Year() bool`

HasDivGrowthRate3Year returns a boolean if a field has been set.

### GetDividendAmount

`func (o *FundamentalFundamental) GetDividendAmount() float32`

GetDividendAmount returns the DividendAmount field if non-nil, zero value otherwise.

### GetDividendAmountOk

`func (o *FundamentalFundamental) GetDividendAmountOk() (*float32, bool)`

GetDividendAmountOk returns a tuple with the DividendAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendAmount

`func (o *FundamentalFundamental) SetDividendAmount(v float32)`

SetDividendAmount sets DividendAmount field to given value.

### HasDividendAmount

`func (o *FundamentalFundamental) HasDividendAmount() bool`

HasDividendAmount returns a boolean if a field has been set.

### GetDividendDate

`func (o *FundamentalFundamental) GetDividendDate() string`

GetDividendDate returns the DividendDate field if non-nil, zero value otherwise.

### GetDividendDateOk

`func (o *FundamentalFundamental) GetDividendDateOk() (*string, bool)`

GetDividendDateOk returns a tuple with the DividendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendDate

`func (o *FundamentalFundamental) SetDividendDate(v string)`

SetDividendDate sets DividendDate field to given value.

### HasDividendDate

`func (o *FundamentalFundamental) HasDividendDate() bool`

HasDividendDate returns a boolean if a field has been set.

### GetDividendPayAmount

`func (o *FundamentalFundamental) GetDividendPayAmount() float32`

GetDividendPayAmount returns the DividendPayAmount field if non-nil, zero value otherwise.

### GetDividendPayAmountOk

`func (o *FundamentalFundamental) GetDividendPayAmountOk() (*float32, bool)`

GetDividendPayAmountOk returns a tuple with the DividendPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendPayAmount

`func (o *FundamentalFundamental) SetDividendPayAmount(v float32)`

SetDividendPayAmount sets DividendPayAmount field to given value.

### HasDividendPayAmount

`func (o *FundamentalFundamental) HasDividendPayAmount() bool`

HasDividendPayAmount returns a boolean if a field has been set.

### GetDividendPayDate

`func (o *FundamentalFundamental) GetDividendPayDate() string`

GetDividendPayDate returns the DividendPayDate field if non-nil, zero value otherwise.

### GetDividendPayDateOk

`func (o *FundamentalFundamental) GetDividendPayDateOk() (*string, bool)`

GetDividendPayDateOk returns a tuple with the DividendPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendPayDate

`func (o *FundamentalFundamental) SetDividendPayDate(v string)`

SetDividendPayDate sets DividendPayDate field to given value.

### HasDividendPayDate

`func (o *FundamentalFundamental) HasDividendPayDate() bool`

HasDividendPayDate returns a boolean if a field has been set.

### GetDividendYield

`func (o *FundamentalFundamental) GetDividendYield() float32`

GetDividendYield returns the DividendYield field if non-nil, zero value otherwise.

### GetDividendYieldOk

`func (o *FundamentalFundamental) GetDividendYieldOk() (*float32, bool)`

GetDividendYieldOk returns a tuple with the DividendYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendYield

`func (o *FundamentalFundamental) SetDividendYield(v float32)`

SetDividendYield sets DividendYield field to given value.

### HasDividendYield

`func (o *FundamentalFundamental) HasDividendYield() bool`

HasDividendYield returns a boolean if a field has been set.

### GetEpsChange

`func (o *FundamentalFundamental) GetEpsChange() float32`

GetEpsChange returns the EpsChange field if non-nil, zero value otherwise.

### GetEpsChangeOk

`func (o *FundamentalFundamental) GetEpsChangeOk() (*float32, bool)`

GetEpsChangeOk returns a tuple with the EpsChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsChange

`func (o *FundamentalFundamental) SetEpsChange(v float32)`

SetEpsChange sets EpsChange field to given value.

### HasEpsChange

`func (o *FundamentalFundamental) HasEpsChange() bool`

HasEpsChange returns a boolean if a field has been set.

### GetEpsChangePercentTTM

`func (o *FundamentalFundamental) GetEpsChangePercentTTM() float32`

GetEpsChangePercentTTM returns the EpsChangePercentTTM field if non-nil, zero value otherwise.

### GetEpsChangePercentTTMOk

`func (o *FundamentalFundamental) GetEpsChangePercentTTMOk() (*float32, bool)`

GetEpsChangePercentTTMOk returns a tuple with the EpsChangePercentTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsChangePercentTTM

`func (o *FundamentalFundamental) SetEpsChangePercentTTM(v float32)`

SetEpsChangePercentTTM sets EpsChangePercentTTM field to given value.

### HasEpsChangePercentTTM

`func (o *FundamentalFundamental) HasEpsChangePercentTTM() bool`

HasEpsChangePercentTTM returns a boolean if a field has been set.

### GetEpsChangeYear

`func (o *FundamentalFundamental) GetEpsChangeYear() float32`

GetEpsChangeYear returns the EpsChangeYear field if non-nil, zero value otherwise.

### GetEpsChangeYearOk

`func (o *FundamentalFundamental) GetEpsChangeYearOk() (*float32, bool)`

GetEpsChangeYearOk returns a tuple with the EpsChangeYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsChangeYear

`func (o *FundamentalFundamental) SetEpsChangeYear(v float32)`

SetEpsChangeYear sets EpsChangeYear field to given value.

### HasEpsChangeYear

`func (o *FundamentalFundamental) HasEpsChangeYear() bool`

HasEpsChangeYear returns a boolean if a field has been set.

### GetEpsTTM

`func (o *FundamentalFundamental) GetEpsTTM() float32`

GetEpsTTM returns the EpsTTM field if non-nil, zero value otherwise.

### GetEpsTTMOk

`func (o *FundamentalFundamental) GetEpsTTMOk() (*float32, bool)`

GetEpsTTMOk returns a tuple with the EpsTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsTTM

`func (o *FundamentalFundamental) SetEpsTTM(v float32)`

SetEpsTTM sets EpsTTM field to given value.

### HasEpsTTM

`func (o *FundamentalFundamental) HasEpsTTM() bool`

HasEpsTTM returns a boolean if a field has been set.

### GetGrossMarginMRQ

`func (o *FundamentalFundamental) GetGrossMarginMRQ() float32`

GetGrossMarginMRQ returns the GrossMarginMRQ field if non-nil, zero value otherwise.

### GetGrossMarginMRQOk

`func (o *FundamentalFundamental) GetGrossMarginMRQOk() (*float32, bool)`

GetGrossMarginMRQOk returns a tuple with the GrossMarginMRQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossMarginMRQ

`func (o *FundamentalFundamental) SetGrossMarginMRQ(v float32)`

SetGrossMarginMRQ sets GrossMarginMRQ field to given value.

### HasGrossMarginMRQ

`func (o *FundamentalFundamental) HasGrossMarginMRQ() bool`

HasGrossMarginMRQ returns a boolean if a field has been set.

### GetGrossMarginTTM

`func (o *FundamentalFundamental) GetGrossMarginTTM() float32`

GetGrossMarginTTM returns the GrossMarginTTM field if non-nil, zero value otherwise.

### GetGrossMarginTTMOk

`func (o *FundamentalFundamental) GetGrossMarginTTMOk() (*float32, bool)`

GetGrossMarginTTMOk returns a tuple with the GrossMarginTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossMarginTTM

`func (o *FundamentalFundamental) SetGrossMarginTTM(v float32)`

SetGrossMarginTTM sets GrossMarginTTM field to given value.

### HasGrossMarginTTM

`func (o *FundamentalFundamental) HasGrossMarginTTM() bool`

HasGrossMarginTTM returns a boolean if a field has been set.

### GetHigh52

`func (o *FundamentalFundamental) GetHigh52() float32`

GetHigh52 returns the High52 field if non-nil, zero value otherwise.

### GetHigh52Ok

`func (o *FundamentalFundamental) GetHigh52Ok() (*float32, bool)`

GetHigh52Ok returns a tuple with the High52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh52

`func (o *FundamentalFundamental) SetHigh52(v float32)`

SetHigh52 sets High52 field to given value.

### HasHigh52

`func (o *FundamentalFundamental) HasHigh52() bool`

HasHigh52 returns a boolean if a field has been set.

### GetInterestCoverage

`func (o *FundamentalFundamental) GetInterestCoverage() float32`

GetInterestCoverage returns the InterestCoverage field if non-nil, zero value otherwise.

### GetInterestCoverageOk

`func (o *FundamentalFundamental) GetInterestCoverageOk() (*float32, bool)`

GetInterestCoverageOk returns a tuple with the InterestCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestCoverage

`func (o *FundamentalFundamental) SetInterestCoverage(v float32)`

SetInterestCoverage sets InterestCoverage field to given value.

### HasInterestCoverage

`func (o *FundamentalFundamental) HasInterestCoverage() bool`

HasInterestCoverage returns a boolean if a field has been set.

### GetLow52

`func (o *FundamentalFundamental) GetLow52() float32`

GetLow52 returns the Low52 field if non-nil, zero value otherwise.

### GetLow52Ok

`func (o *FundamentalFundamental) GetLow52Ok() (*float32, bool)`

GetLow52Ok returns a tuple with the Low52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow52

`func (o *FundamentalFundamental) SetLow52(v float32)`

SetLow52 sets Low52 field to given value.

### HasLow52

`func (o *FundamentalFundamental) HasLow52() bool`

HasLow52 returns a boolean if a field has been set.

### GetLtDebtToEquity

`func (o *FundamentalFundamental) GetLtDebtToEquity() float32`

GetLtDebtToEquity returns the LtDebtToEquity field if non-nil, zero value otherwise.

### GetLtDebtToEquityOk

`func (o *FundamentalFundamental) GetLtDebtToEquityOk() (*float32, bool)`

GetLtDebtToEquityOk returns a tuple with the LtDebtToEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtDebtToEquity

`func (o *FundamentalFundamental) SetLtDebtToEquity(v float32)`

SetLtDebtToEquity sets LtDebtToEquity field to given value.

### HasLtDebtToEquity

`func (o *FundamentalFundamental) HasLtDebtToEquity() bool`

HasLtDebtToEquity returns a boolean if a field has been set.

### GetMarketCap

`func (o *FundamentalFundamental) GetMarketCap() float32`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *FundamentalFundamental) GetMarketCapOk() (*float32, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *FundamentalFundamental) SetMarketCap(v float32)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *FundamentalFundamental) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetMarketCapFloat

`func (o *FundamentalFundamental) GetMarketCapFloat() float32`

GetMarketCapFloat returns the MarketCapFloat field if non-nil, zero value otherwise.

### GetMarketCapFloatOk

`func (o *FundamentalFundamental) GetMarketCapFloatOk() (*float32, bool)`

GetMarketCapFloatOk returns a tuple with the MarketCapFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapFloat

`func (o *FundamentalFundamental) SetMarketCapFloat(v float32)`

SetMarketCapFloat sets MarketCapFloat field to given value.

### HasMarketCapFloat

`func (o *FundamentalFundamental) HasMarketCapFloat() bool`

HasMarketCapFloat returns a boolean if a field has been set.

### GetNetProfitMarginMRQ

`func (o *FundamentalFundamental) GetNetProfitMarginMRQ() float32`

GetNetProfitMarginMRQ returns the NetProfitMarginMRQ field if non-nil, zero value otherwise.

### GetNetProfitMarginMRQOk

`func (o *FundamentalFundamental) GetNetProfitMarginMRQOk() (*float32, bool)`

GetNetProfitMarginMRQOk returns a tuple with the NetProfitMarginMRQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfitMarginMRQ

`func (o *FundamentalFundamental) SetNetProfitMarginMRQ(v float32)`

SetNetProfitMarginMRQ sets NetProfitMarginMRQ field to given value.

### HasNetProfitMarginMRQ

`func (o *FundamentalFundamental) HasNetProfitMarginMRQ() bool`

HasNetProfitMarginMRQ returns a boolean if a field has been set.

### GetNetProfitMarginTTM

`func (o *FundamentalFundamental) GetNetProfitMarginTTM() float32`

GetNetProfitMarginTTM returns the NetProfitMarginTTM field if non-nil, zero value otherwise.

### GetNetProfitMarginTTMOk

`func (o *FundamentalFundamental) GetNetProfitMarginTTMOk() (*float32, bool)`

GetNetProfitMarginTTMOk returns a tuple with the NetProfitMarginTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfitMarginTTM

`func (o *FundamentalFundamental) SetNetProfitMarginTTM(v float32)`

SetNetProfitMarginTTM sets NetProfitMarginTTM field to given value.

### HasNetProfitMarginTTM

`func (o *FundamentalFundamental) HasNetProfitMarginTTM() bool`

HasNetProfitMarginTTM returns a boolean if a field has been set.

### GetOperatingMarginMRQ

`func (o *FundamentalFundamental) GetOperatingMarginMRQ() float32`

GetOperatingMarginMRQ returns the OperatingMarginMRQ field if non-nil, zero value otherwise.

### GetOperatingMarginMRQOk

`func (o *FundamentalFundamental) GetOperatingMarginMRQOk() (*float32, bool)`

GetOperatingMarginMRQOk returns a tuple with the OperatingMarginMRQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMarginMRQ

`func (o *FundamentalFundamental) SetOperatingMarginMRQ(v float32)`

SetOperatingMarginMRQ sets OperatingMarginMRQ field to given value.

### HasOperatingMarginMRQ

`func (o *FundamentalFundamental) HasOperatingMarginMRQ() bool`

HasOperatingMarginMRQ returns a boolean if a field has been set.

### GetOperatingMarginTTM

`func (o *FundamentalFundamental) GetOperatingMarginTTM() float32`

GetOperatingMarginTTM returns the OperatingMarginTTM field if non-nil, zero value otherwise.

### GetOperatingMarginTTMOk

`func (o *FundamentalFundamental) GetOperatingMarginTTMOk() (*float32, bool)`

GetOperatingMarginTTMOk returns a tuple with the OperatingMarginTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMarginTTM

`func (o *FundamentalFundamental) SetOperatingMarginTTM(v float32)`

SetOperatingMarginTTM sets OperatingMarginTTM field to given value.

### HasOperatingMarginTTM

`func (o *FundamentalFundamental) HasOperatingMarginTTM() bool`

HasOperatingMarginTTM returns a boolean if a field has been set.

### GetPbRatio

`func (o *FundamentalFundamental) GetPbRatio() float32`

GetPbRatio returns the PbRatio field if non-nil, zero value otherwise.

### GetPbRatioOk

`func (o *FundamentalFundamental) GetPbRatioOk() (*float32, bool)`

GetPbRatioOk returns a tuple with the PbRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbRatio

`func (o *FundamentalFundamental) SetPbRatio(v float32)`

SetPbRatio sets PbRatio field to given value.

### HasPbRatio

`func (o *FundamentalFundamental) HasPbRatio() bool`

HasPbRatio returns a boolean if a field has been set.

### GetPcfRatio

`func (o *FundamentalFundamental) GetPcfRatio() float32`

GetPcfRatio returns the PcfRatio field if non-nil, zero value otherwise.

### GetPcfRatioOk

`func (o *FundamentalFundamental) GetPcfRatioOk() (*float32, bool)`

GetPcfRatioOk returns a tuple with the PcfRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfRatio

`func (o *FundamentalFundamental) SetPcfRatio(v float32)`

SetPcfRatio sets PcfRatio field to given value.

### HasPcfRatio

`func (o *FundamentalFundamental) HasPcfRatio() bool`

HasPcfRatio returns a boolean if a field has been set.

### GetPeRatio

`func (o *FundamentalFundamental) GetPeRatio() float32`

GetPeRatio returns the PeRatio field if non-nil, zero value otherwise.

### GetPeRatioOk

`func (o *FundamentalFundamental) GetPeRatioOk() (*float32, bool)`

GetPeRatioOk returns a tuple with the PeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeRatio

`func (o *FundamentalFundamental) SetPeRatio(v float32)`

SetPeRatio sets PeRatio field to given value.

### HasPeRatio

`func (o *FundamentalFundamental) HasPeRatio() bool`

HasPeRatio returns a boolean if a field has been set.

### GetPegRatio

`func (o *FundamentalFundamental) GetPegRatio() float32`

GetPegRatio returns the PegRatio field if non-nil, zero value otherwise.

### GetPegRatioOk

`func (o *FundamentalFundamental) GetPegRatioOk() (*float32, bool)`

GetPegRatioOk returns a tuple with the PegRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegRatio

`func (o *FundamentalFundamental) SetPegRatio(v float32)`

SetPegRatio sets PegRatio field to given value.

### HasPegRatio

`func (o *FundamentalFundamental) HasPegRatio() bool`

HasPegRatio returns a boolean if a field has been set.

### GetPrRatio

`func (o *FundamentalFundamental) GetPrRatio() float32`

GetPrRatio returns the PrRatio field if non-nil, zero value otherwise.

### GetPrRatioOk

`func (o *FundamentalFundamental) GetPrRatioOk() (*float32, bool)`

GetPrRatioOk returns a tuple with the PrRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrRatio

`func (o *FundamentalFundamental) SetPrRatio(v float32)`

SetPrRatio sets PrRatio field to given value.

### HasPrRatio

`func (o *FundamentalFundamental) HasPrRatio() bool`

HasPrRatio returns a boolean if a field has been set.

### GetQuickRatio

`func (o *FundamentalFundamental) GetQuickRatio() float32`

GetQuickRatio returns the QuickRatio field if non-nil, zero value otherwise.

### GetQuickRatioOk

`func (o *FundamentalFundamental) GetQuickRatioOk() (*float32, bool)`

GetQuickRatioOk returns a tuple with the QuickRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickRatio

`func (o *FundamentalFundamental) SetQuickRatio(v float32)`

SetQuickRatio sets QuickRatio field to given value.

### HasQuickRatio

`func (o *FundamentalFundamental) HasQuickRatio() bool`

HasQuickRatio returns a boolean if a field has been set.

### GetReturnOnAssets

`func (o *FundamentalFundamental) GetReturnOnAssets() float32`

GetReturnOnAssets returns the ReturnOnAssets field if non-nil, zero value otherwise.

### GetReturnOnAssetsOk

`func (o *FundamentalFundamental) GetReturnOnAssetsOk() (*float32, bool)`

GetReturnOnAssetsOk returns a tuple with the ReturnOnAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnAssets

`func (o *FundamentalFundamental) SetReturnOnAssets(v float32)`

SetReturnOnAssets sets ReturnOnAssets field to given value.

### HasReturnOnAssets

`func (o *FundamentalFundamental) HasReturnOnAssets() bool`

HasReturnOnAssets returns a boolean if a field has been set.

### GetReturnOnEquity

`func (o *FundamentalFundamental) GetReturnOnEquity() float32`

GetReturnOnEquity returns the ReturnOnEquity field if non-nil, zero value otherwise.

### GetReturnOnEquityOk

`func (o *FundamentalFundamental) GetReturnOnEquityOk() (*float32, bool)`

GetReturnOnEquityOk returns a tuple with the ReturnOnEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnEquity

`func (o *FundamentalFundamental) SetReturnOnEquity(v float32)`

SetReturnOnEquity sets ReturnOnEquity field to given value.

### HasReturnOnEquity

`func (o *FundamentalFundamental) HasReturnOnEquity() bool`

HasReturnOnEquity returns a boolean if a field has been set.

### GetReturnOnInvestment

`func (o *FundamentalFundamental) GetReturnOnInvestment() float32`

GetReturnOnInvestment returns the ReturnOnInvestment field if non-nil, zero value otherwise.

### GetReturnOnInvestmentOk

`func (o *FundamentalFundamental) GetReturnOnInvestmentOk() (*float32, bool)`

GetReturnOnInvestmentOk returns a tuple with the ReturnOnInvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnInvestment

`func (o *FundamentalFundamental) SetReturnOnInvestment(v float32)`

SetReturnOnInvestment sets ReturnOnInvestment field to given value.

### HasReturnOnInvestment

`func (o *FundamentalFundamental) HasReturnOnInvestment() bool`

HasReturnOnInvestment returns a boolean if a field has been set.

### GetRevChangeIn

`func (o *FundamentalFundamental) GetRevChangeIn() float32`

GetRevChangeIn returns the RevChangeIn field if non-nil, zero value otherwise.

### GetRevChangeInOk

`func (o *FundamentalFundamental) GetRevChangeInOk() (*float32, bool)`

GetRevChangeInOk returns a tuple with the RevChangeIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevChangeIn

`func (o *FundamentalFundamental) SetRevChangeIn(v float32)`

SetRevChangeIn sets RevChangeIn field to given value.

### HasRevChangeIn

`func (o *FundamentalFundamental) HasRevChangeIn() bool`

HasRevChangeIn returns a boolean if a field has been set.

### GetRevChangeTTM

`func (o *FundamentalFundamental) GetRevChangeTTM() float32`

GetRevChangeTTM returns the RevChangeTTM field if non-nil, zero value otherwise.

### GetRevChangeTTMOk

`func (o *FundamentalFundamental) GetRevChangeTTMOk() (*float32, bool)`

GetRevChangeTTMOk returns a tuple with the RevChangeTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevChangeTTM

`func (o *FundamentalFundamental) SetRevChangeTTM(v float32)`

SetRevChangeTTM sets RevChangeTTM field to given value.

### HasRevChangeTTM

`func (o *FundamentalFundamental) HasRevChangeTTM() bool`

HasRevChangeTTM returns a boolean if a field has been set.

### GetRevChangeYear

`func (o *FundamentalFundamental) GetRevChangeYear() float32`

GetRevChangeYear returns the RevChangeYear field if non-nil, zero value otherwise.

### GetRevChangeYearOk

`func (o *FundamentalFundamental) GetRevChangeYearOk() (*float32, bool)`

GetRevChangeYearOk returns a tuple with the RevChangeYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevChangeYear

`func (o *FundamentalFundamental) SetRevChangeYear(v float32)`

SetRevChangeYear sets RevChangeYear field to given value.

### HasRevChangeYear

`func (o *FundamentalFundamental) HasRevChangeYear() bool`

HasRevChangeYear returns a boolean if a field has been set.

### GetSharesOutstanding

`func (o *FundamentalFundamental) GetSharesOutstanding() float32`

GetSharesOutstanding returns the SharesOutstanding field if non-nil, zero value otherwise.

### GetSharesOutstandingOk

`func (o *FundamentalFundamental) GetSharesOutstandingOk() (*float32, bool)`

GetSharesOutstandingOk returns a tuple with the SharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOutstanding

`func (o *FundamentalFundamental) SetSharesOutstanding(v float32)`

SetSharesOutstanding sets SharesOutstanding field to given value.

### HasSharesOutstanding

`func (o *FundamentalFundamental) HasSharesOutstanding() bool`

HasSharesOutstanding returns a boolean if a field has been set.

### GetShortIntDayToCover

`func (o *FundamentalFundamental) GetShortIntDayToCover() float32`

GetShortIntDayToCover returns the ShortIntDayToCover field if non-nil, zero value otherwise.

### GetShortIntDayToCoverOk

`func (o *FundamentalFundamental) GetShortIntDayToCoverOk() (*float32, bool)`

GetShortIntDayToCoverOk returns a tuple with the ShortIntDayToCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortIntDayToCover

`func (o *FundamentalFundamental) SetShortIntDayToCover(v float32)`

SetShortIntDayToCover sets ShortIntDayToCover field to given value.

### HasShortIntDayToCover

`func (o *FundamentalFundamental) HasShortIntDayToCover() bool`

HasShortIntDayToCover returns a boolean if a field has been set.

### GetShortIntToFloat

`func (o *FundamentalFundamental) GetShortIntToFloat() float32`

GetShortIntToFloat returns the ShortIntToFloat field if non-nil, zero value otherwise.

### GetShortIntToFloatOk

`func (o *FundamentalFundamental) GetShortIntToFloatOk() (*float32, bool)`

GetShortIntToFloatOk returns a tuple with the ShortIntToFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortIntToFloat

`func (o *FundamentalFundamental) SetShortIntToFloat(v float32)`

SetShortIntToFloat sets ShortIntToFloat field to given value.

### HasShortIntToFloat

`func (o *FundamentalFundamental) HasShortIntToFloat() bool`

HasShortIntToFloat returns a boolean if a field has been set.

### GetSymbol

`func (o *FundamentalFundamental) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FundamentalFundamental) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FundamentalFundamental) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FundamentalFundamental) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTotalDebtToCapital

`func (o *FundamentalFundamental) GetTotalDebtToCapital() float32`

GetTotalDebtToCapital returns the TotalDebtToCapital field if non-nil, zero value otherwise.

### GetTotalDebtToCapitalOk

`func (o *FundamentalFundamental) GetTotalDebtToCapitalOk() (*float32, bool)`

GetTotalDebtToCapitalOk returns a tuple with the TotalDebtToCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebtToCapital

`func (o *FundamentalFundamental) SetTotalDebtToCapital(v float32)`

SetTotalDebtToCapital sets TotalDebtToCapital field to given value.

### HasTotalDebtToCapital

`func (o *FundamentalFundamental) HasTotalDebtToCapital() bool`

HasTotalDebtToCapital returns a boolean if a field has been set.

### GetTotalDebtToEquity

`func (o *FundamentalFundamental) GetTotalDebtToEquity() float32`

GetTotalDebtToEquity returns the TotalDebtToEquity field if non-nil, zero value otherwise.

### GetTotalDebtToEquityOk

`func (o *FundamentalFundamental) GetTotalDebtToEquityOk() (*float32, bool)`

GetTotalDebtToEquityOk returns a tuple with the TotalDebtToEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebtToEquity

`func (o *FundamentalFundamental) SetTotalDebtToEquity(v float32)`

SetTotalDebtToEquity sets TotalDebtToEquity field to given value.

### HasTotalDebtToEquity

`func (o *FundamentalFundamental) HasTotalDebtToEquity() bool`

HasTotalDebtToEquity returns a boolean if a field has been set.

### GetVol10DayAvg

`func (o *FundamentalFundamental) GetVol10DayAvg() float32`

GetVol10DayAvg returns the Vol10DayAvg field if non-nil, zero value otherwise.

### GetVol10DayAvgOk

`func (o *FundamentalFundamental) GetVol10DayAvgOk() (*float32, bool)`

GetVol10DayAvgOk returns a tuple with the Vol10DayAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol10DayAvg

`func (o *FundamentalFundamental) SetVol10DayAvg(v float32)`

SetVol10DayAvg sets Vol10DayAvg field to given value.

### HasVol10DayAvg

`func (o *FundamentalFundamental) HasVol10DayAvg() bool`

HasVol10DayAvg returns a boolean if a field has been set.

### GetVol1DayAvg

`func (o *FundamentalFundamental) GetVol1DayAvg() float32`

GetVol1DayAvg returns the Vol1DayAvg field if non-nil, zero value otherwise.

### GetVol1DayAvgOk

`func (o *FundamentalFundamental) GetVol1DayAvgOk() (*float32, bool)`

GetVol1DayAvgOk returns a tuple with the Vol1DayAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol1DayAvg

`func (o *FundamentalFundamental) SetVol1DayAvg(v float32)`

SetVol1DayAvg sets Vol1DayAvg field to given value.

### HasVol1DayAvg

`func (o *FundamentalFundamental) HasVol1DayAvg() bool`

HasVol1DayAvg returns a boolean if a field has been set.

### GetVol3MonthAvg

`func (o *FundamentalFundamental) GetVol3MonthAvg() float32`

GetVol3MonthAvg returns the Vol3MonthAvg field if non-nil, zero value otherwise.

### GetVol3MonthAvgOk

`func (o *FundamentalFundamental) GetVol3MonthAvgOk() (*float32, bool)`

GetVol3MonthAvgOk returns a tuple with the Vol3MonthAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol3MonthAvg

`func (o *FundamentalFundamental) SetVol3MonthAvg(v float32)`

SetVol3MonthAvg sets Vol3MonthAvg field to given value.

### HasVol3MonthAvg

`func (o *FundamentalFundamental) HasVol3MonthAvg() bool`

HasVol3MonthAvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


