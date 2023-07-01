# FundamentalData

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

### NewFundamentalData

`func NewFundamentalData() *FundamentalData`

NewFundamentalData instantiates a new FundamentalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalDataWithDefaults

`func NewFundamentalDataWithDefaults() *FundamentalData`

NewFundamentalDataWithDefaults instantiates a new FundamentalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeta

`func (o *FundamentalData) GetBeta() float32`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *FundamentalData) GetBetaOk() (*float32, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *FundamentalData) SetBeta(v float32)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *FundamentalData) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetBookValuePerShare

`func (o *FundamentalData) GetBookValuePerShare() float32`

GetBookValuePerShare returns the BookValuePerShare field if non-nil, zero value otherwise.

### GetBookValuePerShareOk

`func (o *FundamentalData) GetBookValuePerShareOk() (*float32, bool)`

GetBookValuePerShareOk returns a tuple with the BookValuePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookValuePerShare

`func (o *FundamentalData) SetBookValuePerShare(v float32)`

SetBookValuePerShare sets BookValuePerShare field to given value.

### HasBookValuePerShare

`func (o *FundamentalData) HasBookValuePerShare() bool`

HasBookValuePerShare returns a boolean if a field has been set.

### GetCurrentRatio

`func (o *FundamentalData) GetCurrentRatio() float32`

GetCurrentRatio returns the CurrentRatio field if non-nil, zero value otherwise.

### GetCurrentRatioOk

`func (o *FundamentalData) GetCurrentRatioOk() (*float32, bool)`

GetCurrentRatioOk returns a tuple with the CurrentRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRatio

`func (o *FundamentalData) SetCurrentRatio(v float32)`

SetCurrentRatio sets CurrentRatio field to given value.

### HasCurrentRatio

`func (o *FundamentalData) HasCurrentRatio() bool`

HasCurrentRatio returns a boolean if a field has been set.

### GetDivGrowthRate3Year

`func (o *FundamentalData) GetDivGrowthRate3Year() float32`

GetDivGrowthRate3Year returns the DivGrowthRate3Year field if non-nil, zero value otherwise.

### GetDivGrowthRate3YearOk

`func (o *FundamentalData) GetDivGrowthRate3YearOk() (*float32, bool)`

GetDivGrowthRate3YearOk returns a tuple with the DivGrowthRate3Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivGrowthRate3Year

`func (o *FundamentalData) SetDivGrowthRate3Year(v float32)`

SetDivGrowthRate3Year sets DivGrowthRate3Year field to given value.

### HasDivGrowthRate3Year

`func (o *FundamentalData) HasDivGrowthRate3Year() bool`

HasDivGrowthRate3Year returns a boolean if a field has been set.

### GetDividendAmount

`func (o *FundamentalData) GetDividendAmount() float32`

GetDividendAmount returns the DividendAmount field if non-nil, zero value otherwise.

### GetDividendAmountOk

`func (o *FundamentalData) GetDividendAmountOk() (*float32, bool)`

GetDividendAmountOk returns a tuple with the DividendAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendAmount

`func (o *FundamentalData) SetDividendAmount(v float32)`

SetDividendAmount sets DividendAmount field to given value.

### HasDividendAmount

`func (o *FundamentalData) HasDividendAmount() bool`

HasDividendAmount returns a boolean if a field has been set.

### GetDividendDate

`func (o *FundamentalData) GetDividendDate() string`

GetDividendDate returns the DividendDate field if non-nil, zero value otherwise.

### GetDividendDateOk

`func (o *FundamentalData) GetDividendDateOk() (*string, bool)`

GetDividendDateOk returns a tuple with the DividendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendDate

`func (o *FundamentalData) SetDividendDate(v string)`

SetDividendDate sets DividendDate field to given value.

### HasDividendDate

`func (o *FundamentalData) HasDividendDate() bool`

HasDividendDate returns a boolean if a field has been set.

### GetDividendPayAmount

`func (o *FundamentalData) GetDividendPayAmount() float32`

GetDividendPayAmount returns the DividendPayAmount field if non-nil, zero value otherwise.

### GetDividendPayAmountOk

`func (o *FundamentalData) GetDividendPayAmountOk() (*float32, bool)`

GetDividendPayAmountOk returns a tuple with the DividendPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendPayAmount

`func (o *FundamentalData) SetDividendPayAmount(v float32)`

SetDividendPayAmount sets DividendPayAmount field to given value.

### HasDividendPayAmount

`func (o *FundamentalData) HasDividendPayAmount() bool`

HasDividendPayAmount returns a boolean if a field has been set.

### GetDividendPayDate

`func (o *FundamentalData) GetDividendPayDate() string`

GetDividendPayDate returns the DividendPayDate field if non-nil, zero value otherwise.

### GetDividendPayDateOk

`func (o *FundamentalData) GetDividendPayDateOk() (*string, bool)`

GetDividendPayDateOk returns a tuple with the DividendPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendPayDate

`func (o *FundamentalData) SetDividendPayDate(v string)`

SetDividendPayDate sets DividendPayDate field to given value.

### HasDividendPayDate

`func (o *FundamentalData) HasDividendPayDate() bool`

HasDividendPayDate returns a boolean if a field has been set.

### GetDividendYield

`func (o *FundamentalData) GetDividendYield() float32`

GetDividendYield returns the DividendYield field if non-nil, zero value otherwise.

### GetDividendYieldOk

`func (o *FundamentalData) GetDividendYieldOk() (*float32, bool)`

GetDividendYieldOk returns a tuple with the DividendYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendYield

`func (o *FundamentalData) SetDividendYield(v float32)`

SetDividendYield sets DividendYield field to given value.

### HasDividendYield

`func (o *FundamentalData) HasDividendYield() bool`

HasDividendYield returns a boolean if a field has been set.

### GetEpsChange

`func (o *FundamentalData) GetEpsChange() float32`

GetEpsChange returns the EpsChange field if non-nil, zero value otherwise.

### GetEpsChangeOk

`func (o *FundamentalData) GetEpsChangeOk() (*float32, bool)`

GetEpsChangeOk returns a tuple with the EpsChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsChange

`func (o *FundamentalData) SetEpsChange(v float32)`

SetEpsChange sets EpsChange field to given value.

### HasEpsChange

`func (o *FundamentalData) HasEpsChange() bool`

HasEpsChange returns a boolean if a field has been set.

### GetEpsChangePercentTTM

`func (o *FundamentalData) GetEpsChangePercentTTM() float32`

GetEpsChangePercentTTM returns the EpsChangePercentTTM field if non-nil, zero value otherwise.

### GetEpsChangePercentTTMOk

`func (o *FundamentalData) GetEpsChangePercentTTMOk() (*float32, bool)`

GetEpsChangePercentTTMOk returns a tuple with the EpsChangePercentTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsChangePercentTTM

`func (o *FundamentalData) SetEpsChangePercentTTM(v float32)`

SetEpsChangePercentTTM sets EpsChangePercentTTM field to given value.

### HasEpsChangePercentTTM

`func (o *FundamentalData) HasEpsChangePercentTTM() bool`

HasEpsChangePercentTTM returns a boolean if a field has been set.

### GetEpsChangeYear

`func (o *FundamentalData) GetEpsChangeYear() float32`

GetEpsChangeYear returns the EpsChangeYear field if non-nil, zero value otherwise.

### GetEpsChangeYearOk

`func (o *FundamentalData) GetEpsChangeYearOk() (*float32, bool)`

GetEpsChangeYearOk returns a tuple with the EpsChangeYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsChangeYear

`func (o *FundamentalData) SetEpsChangeYear(v float32)`

SetEpsChangeYear sets EpsChangeYear field to given value.

### HasEpsChangeYear

`func (o *FundamentalData) HasEpsChangeYear() bool`

HasEpsChangeYear returns a boolean if a field has been set.

### GetEpsTTM

`func (o *FundamentalData) GetEpsTTM() float32`

GetEpsTTM returns the EpsTTM field if non-nil, zero value otherwise.

### GetEpsTTMOk

`func (o *FundamentalData) GetEpsTTMOk() (*float32, bool)`

GetEpsTTMOk returns a tuple with the EpsTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsTTM

`func (o *FundamentalData) SetEpsTTM(v float32)`

SetEpsTTM sets EpsTTM field to given value.

### HasEpsTTM

`func (o *FundamentalData) HasEpsTTM() bool`

HasEpsTTM returns a boolean if a field has been set.

### GetGrossMarginMRQ

`func (o *FundamentalData) GetGrossMarginMRQ() float32`

GetGrossMarginMRQ returns the GrossMarginMRQ field if non-nil, zero value otherwise.

### GetGrossMarginMRQOk

`func (o *FundamentalData) GetGrossMarginMRQOk() (*float32, bool)`

GetGrossMarginMRQOk returns a tuple with the GrossMarginMRQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossMarginMRQ

`func (o *FundamentalData) SetGrossMarginMRQ(v float32)`

SetGrossMarginMRQ sets GrossMarginMRQ field to given value.

### HasGrossMarginMRQ

`func (o *FundamentalData) HasGrossMarginMRQ() bool`

HasGrossMarginMRQ returns a boolean if a field has been set.

### GetGrossMarginTTM

`func (o *FundamentalData) GetGrossMarginTTM() float32`

GetGrossMarginTTM returns the GrossMarginTTM field if non-nil, zero value otherwise.

### GetGrossMarginTTMOk

`func (o *FundamentalData) GetGrossMarginTTMOk() (*float32, bool)`

GetGrossMarginTTMOk returns a tuple with the GrossMarginTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossMarginTTM

`func (o *FundamentalData) SetGrossMarginTTM(v float32)`

SetGrossMarginTTM sets GrossMarginTTM field to given value.

### HasGrossMarginTTM

`func (o *FundamentalData) HasGrossMarginTTM() bool`

HasGrossMarginTTM returns a boolean if a field has been set.

### GetHigh52

`func (o *FundamentalData) GetHigh52() float32`

GetHigh52 returns the High52 field if non-nil, zero value otherwise.

### GetHigh52Ok

`func (o *FundamentalData) GetHigh52Ok() (*float32, bool)`

GetHigh52Ok returns a tuple with the High52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh52

`func (o *FundamentalData) SetHigh52(v float32)`

SetHigh52 sets High52 field to given value.

### HasHigh52

`func (o *FundamentalData) HasHigh52() bool`

HasHigh52 returns a boolean if a field has been set.

### GetInterestCoverage

`func (o *FundamentalData) GetInterestCoverage() float32`

GetInterestCoverage returns the InterestCoverage field if non-nil, zero value otherwise.

### GetInterestCoverageOk

`func (o *FundamentalData) GetInterestCoverageOk() (*float32, bool)`

GetInterestCoverageOk returns a tuple with the InterestCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestCoverage

`func (o *FundamentalData) SetInterestCoverage(v float32)`

SetInterestCoverage sets InterestCoverage field to given value.

### HasInterestCoverage

`func (o *FundamentalData) HasInterestCoverage() bool`

HasInterestCoverage returns a boolean if a field has been set.

### GetLow52

`func (o *FundamentalData) GetLow52() float32`

GetLow52 returns the Low52 field if non-nil, zero value otherwise.

### GetLow52Ok

`func (o *FundamentalData) GetLow52Ok() (*float32, bool)`

GetLow52Ok returns a tuple with the Low52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow52

`func (o *FundamentalData) SetLow52(v float32)`

SetLow52 sets Low52 field to given value.

### HasLow52

`func (o *FundamentalData) HasLow52() bool`

HasLow52 returns a boolean if a field has been set.

### GetLtDebtToEquity

`func (o *FundamentalData) GetLtDebtToEquity() float32`

GetLtDebtToEquity returns the LtDebtToEquity field if non-nil, zero value otherwise.

### GetLtDebtToEquityOk

`func (o *FundamentalData) GetLtDebtToEquityOk() (*float32, bool)`

GetLtDebtToEquityOk returns a tuple with the LtDebtToEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtDebtToEquity

`func (o *FundamentalData) SetLtDebtToEquity(v float32)`

SetLtDebtToEquity sets LtDebtToEquity field to given value.

### HasLtDebtToEquity

`func (o *FundamentalData) HasLtDebtToEquity() bool`

HasLtDebtToEquity returns a boolean if a field has been set.

### GetMarketCap

`func (o *FundamentalData) GetMarketCap() float32`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *FundamentalData) GetMarketCapOk() (*float32, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *FundamentalData) SetMarketCap(v float32)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *FundamentalData) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetMarketCapFloat

`func (o *FundamentalData) GetMarketCapFloat() float32`

GetMarketCapFloat returns the MarketCapFloat field if non-nil, zero value otherwise.

### GetMarketCapFloatOk

`func (o *FundamentalData) GetMarketCapFloatOk() (*float32, bool)`

GetMarketCapFloatOk returns a tuple with the MarketCapFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapFloat

`func (o *FundamentalData) SetMarketCapFloat(v float32)`

SetMarketCapFloat sets MarketCapFloat field to given value.

### HasMarketCapFloat

`func (o *FundamentalData) HasMarketCapFloat() bool`

HasMarketCapFloat returns a boolean if a field has been set.

### GetNetProfitMarginMRQ

`func (o *FundamentalData) GetNetProfitMarginMRQ() float32`

GetNetProfitMarginMRQ returns the NetProfitMarginMRQ field if non-nil, zero value otherwise.

### GetNetProfitMarginMRQOk

`func (o *FundamentalData) GetNetProfitMarginMRQOk() (*float32, bool)`

GetNetProfitMarginMRQOk returns a tuple with the NetProfitMarginMRQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfitMarginMRQ

`func (o *FundamentalData) SetNetProfitMarginMRQ(v float32)`

SetNetProfitMarginMRQ sets NetProfitMarginMRQ field to given value.

### HasNetProfitMarginMRQ

`func (o *FundamentalData) HasNetProfitMarginMRQ() bool`

HasNetProfitMarginMRQ returns a boolean if a field has been set.

### GetNetProfitMarginTTM

`func (o *FundamentalData) GetNetProfitMarginTTM() float32`

GetNetProfitMarginTTM returns the NetProfitMarginTTM field if non-nil, zero value otherwise.

### GetNetProfitMarginTTMOk

`func (o *FundamentalData) GetNetProfitMarginTTMOk() (*float32, bool)`

GetNetProfitMarginTTMOk returns a tuple with the NetProfitMarginTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfitMarginTTM

`func (o *FundamentalData) SetNetProfitMarginTTM(v float32)`

SetNetProfitMarginTTM sets NetProfitMarginTTM field to given value.

### HasNetProfitMarginTTM

`func (o *FundamentalData) HasNetProfitMarginTTM() bool`

HasNetProfitMarginTTM returns a boolean if a field has been set.

### GetOperatingMarginMRQ

`func (o *FundamentalData) GetOperatingMarginMRQ() float32`

GetOperatingMarginMRQ returns the OperatingMarginMRQ field if non-nil, zero value otherwise.

### GetOperatingMarginMRQOk

`func (o *FundamentalData) GetOperatingMarginMRQOk() (*float32, bool)`

GetOperatingMarginMRQOk returns a tuple with the OperatingMarginMRQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMarginMRQ

`func (o *FundamentalData) SetOperatingMarginMRQ(v float32)`

SetOperatingMarginMRQ sets OperatingMarginMRQ field to given value.

### HasOperatingMarginMRQ

`func (o *FundamentalData) HasOperatingMarginMRQ() bool`

HasOperatingMarginMRQ returns a boolean if a field has been set.

### GetOperatingMarginTTM

`func (o *FundamentalData) GetOperatingMarginTTM() float32`

GetOperatingMarginTTM returns the OperatingMarginTTM field if non-nil, zero value otherwise.

### GetOperatingMarginTTMOk

`func (o *FundamentalData) GetOperatingMarginTTMOk() (*float32, bool)`

GetOperatingMarginTTMOk returns a tuple with the OperatingMarginTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMarginTTM

`func (o *FundamentalData) SetOperatingMarginTTM(v float32)`

SetOperatingMarginTTM sets OperatingMarginTTM field to given value.

### HasOperatingMarginTTM

`func (o *FundamentalData) HasOperatingMarginTTM() bool`

HasOperatingMarginTTM returns a boolean if a field has been set.

### GetPbRatio

`func (o *FundamentalData) GetPbRatio() float32`

GetPbRatio returns the PbRatio field if non-nil, zero value otherwise.

### GetPbRatioOk

`func (o *FundamentalData) GetPbRatioOk() (*float32, bool)`

GetPbRatioOk returns a tuple with the PbRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbRatio

`func (o *FundamentalData) SetPbRatio(v float32)`

SetPbRatio sets PbRatio field to given value.

### HasPbRatio

`func (o *FundamentalData) HasPbRatio() bool`

HasPbRatio returns a boolean if a field has been set.

### GetPcfRatio

`func (o *FundamentalData) GetPcfRatio() float32`

GetPcfRatio returns the PcfRatio field if non-nil, zero value otherwise.

### GetPcfRatioOk

`func (o *FundamentalData) GetPcfRatioOk() (*float32, bool)`

GetPcfRatioOk returns a tuple with the PcfRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfRatio

`func (o *FundamentalData) SetPcfRatio(v float32)`

SetPcfRatio sets PcfRatio field to given value.

### HasPcfRatio

`func (o *FundamentalData) HasPcfRatio() bool`

HasPcfRatio returns a boolean if a field has been set.

### GetPeRatio

`func (o *FundamentalData) GetPeRatio() float32`

GetPeRatio returns the PeRatio field if non-nil, zero value otherwise.

### GetPeRatioOk

`func (o *FundamentalData) GetPeRatioOk() (*float32, bool)`

GetPeRatioOk returns a tuple with the PeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeRatio

`func (o *FundamentalData) SetPeRatio(v float32)`

SetPeRatio sets PeRatio field to given value.

### HasPeRatio

`func (o *FundamentalData) HasPeRatio() bool`

HasPeRatio returns a boolean if a field has been set.

### GetPegRatio

`func (o *FundamentalData) GetPegRatio() float32`

GetPegRatio returns the PegRatio field if non-nil, zero value otherwise.

### GetPegRatioOk

`func (o *FundamentalData) GetPegRatioOk() (*float32, bool)`

GetPegRatioOk returns a tuple with the PegRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegRatio

`func (o *FundamentalData) SetPegRatio(v float32)`

SetPegRatio sets PegRatio field to given value.

### HasPegRatio

`func (o *FundamentalData) HasPegRatio() bool`

HasPegRatio returns a boolean if a field has been set.

### GetPrRatio

`func (o *FundamentalData) GetPrRatio() float32`

GetPrRatio returns the PrRatio field if non-nil, zero value otherwise.

### GetPrRatioOk

`func (o *FundamentalData) GetPrRatioOk() (*float32, bool)`

GetPrRatioOk returns a tuple with the PrRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrRatio

`func (o *FundamentalData) SetPrRatio(v float32)`

SetPrRatio sets PrRatio field to given value.

### HasPrRatio

`func (o *FundamentalData) HasPrRatio() bool`

HasPrRatio returns a boolean if a field has been set.

### GetQuickRatio

`func (o *FundamentalData) GetQuickRatio() float32`

GetQuickRatio returns the QuickRatio field if non-nil, zero value otherwise.

### GetQuickRatioOk

`func (o *FundamentalData) GetQuickRatioOk() (*float32, bool)`

GetQuickRatioOk returns a tuple with the QuickRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickRatio

`func (o *FundamentalData) SetQuickRatio(v float32)`

SetQuickRatio sets QuickRatio field to given value.

### HasQuickRatio

`func (o *FundamentalData) HasQuickRatio() bool`

HasQuickRatio returns a boolean if a field has been set.

### GetReturnOnAssets

`func (o *FundamentalData) GetReturnOnAssets() float32`

GetReturnOnAssets returns the ReturnOnAssets field if non-nil, zero value otherwise.

### GetReturnOnAssetsOk

`func (o *FundamentalData) GetReturnOnAssetsOk() (*float32, bool)`

GetReturnOnAssetsOk returns a tuple with the ReturnOnAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnAssets

`func (o *FundamentalData) SetReturnOnAssets(v float32)`

SetReturnOnAssets sets ReturnOnAssets field to given value.

### HasReturnOnAssets

`func (o *FundamentalData) HasReturnOnAssets() bool`

HasReturnOnAssets returns a boolean if a field has been set.

### GetReturnOnEquity

`func (o *FundamentalData) GetReturnOnEquity() float32`

GetReturnOnEquity returns the ReturnOnEquity field if non-nil, zero value otherwise.

### GetReturnOnEquityOk

`func (o *FundamentalData) GetReturnOnEquityOk() (*float32, bool)`

GetReturnOnEquityOk returns a tuple with the ReturnOnEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnEquity

`func (o *FundamentalData) SetReturnOnEquity(v float32)`

SetReturnOnEquity sets ReturnOnEquity field to given value.

### HasReturnOnEquity

`func (o *FundamentalData) HasReturnOnEquity() bool`

HasReturnOnEquity returns a boolean if a field has been set.

### GetReturnOnInvestment

`func (o *FundamentalData) GetReturnOnInvestment() float32`

GetReturnOnInvestment returns the ReturnOnInvestment field if non-nil, zero value otherwise.

### GetReturnOnInvestmentOk

`func (o *FundamentalData) GetReturnOnInvestmentOk() (*float32, bool)`

GetReturnOnInvestmentOk returns a tuple with the ReturnOnInvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnInvestment

`func (o *FundamentalData) SetReturnOnInvestment(v float32)`

SetReturnOnInvestment sets ReturnOnInvestment field to given value.

### HasReturnOnInvestment

`func (o *FundamentalData) HasReturnOnInvestment() bool`

HasReturnOnInvestment returns a boolean if a field has been set.

### GetRevChangeIn

`func (o *FundamentalData) GetRevChangeIn() float32`

GetRevChangeIn returns the RevChangeIn field if non-nil, zero value otherwise.

### GetRevChangeInOk

`func (o *FundamentalData) GetRevChangeInOk() (*float32, bool)`

GetRevChangeInOk returns a tuple with the RevChangeIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevChangeIn

`func (o *FundamentalData) SetRevChangeIn(v float32)`

SetRevChangeIn sets RevChangeIn field to given value.

### HasRevChangeIn

`func (o *FundamentalData) HasRevChangeIn() bool`

HasRevChangeIn returns a boolean if a field has been set.

### GetRevChangeTTM

`func (o *FundamentalData) GetRevChangeTTM() float32`

GetRevChangeTTM returns the RevChangeTTM field if non-nil, zero value otherwise.

### GetRevChangeTTMOk

`func (o *FundamentalData) GetRevChangeTTMOk() (*float32, bool)`

GetRevChangeTTMOk returns a tuple with the RevChangeTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevChangeTTM

`func (o *FundamentalData) SetRevChangeTTM(v float32)`

SetRevChangeTTM sets RevChangeTTM field to given value.

### HasRevChangeTTM

`func (o *FundamentalData) HasRevChangeTTM() bool`

HasRevChangeTTM returns a boolean if a field has been set.

### GetRevChangeYear

`func (o *FundamentalData) GetRevChangeYear() float32`

GetRevChangeYear returns the RevChangeYear field if non-nil, zero value otherwise.

### GetRevChangeYearOk

`func (o *FundamentalData) GetRevChangeYearOk() (*float32, bool)`

GetRevChangeYearOk returns a tuple with the RevChangeYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevChangeYear

`func (o *FundamentalData) SetRevChangeYear(v float32)`

SetRevChangeYear sets RevChangeYear field to given value.

### HasRevChangeYear

`func (o *FundamentalData) HasRevChangeYear() bool`

HasRevChangeYear returns a boolean if a field has been set.

### GetSharesOutstanding

`func (o *FundamentalData) GetSharesOutstanding() float32`

GetSharesOutstanding returns the SharesOutstanding field if non-nil, zero value otherwise.

### GetSharesOutstandingOk

`func (o *FundamentalData) GetSharesOutstandingOk() (*float32, bool)`

GetSharesOutstandingOk returns a tuple with the SharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOutstanding

`func (o *FundamentalData) SetSharesOutstanding(v float32)`

SetSharesOutstanding sets SharesOutstanding field to given value.

### HasSharesOutstanding

`func (o *FundamentalData) HasSharesOutstanding() bool`

HasSharesOutstanding returns a boolean if a field has been set.

### GetShortIntDayToCover

`func (o *FundamentalData) GetShortIntDayToCover() float32`

GetShortIntDayToCover returns the ShortIntDayToCover field if non-nil, zero value otherwise.

### GetShortIntDayToCoverOk

`func (o *FundamentalData) GetShortIntDayToCoverOk() (*float32, bool)`

GetShortIntDayToCoverOk returns a tuple with the ShortIntDayToCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortIntDayToCover

`func (o *FundamentalData) SetShortIntDayToCover(v float32)`

SetShortIntDayToCover sets ShortIntDayToCover field to given value.

### HasShortIntDayToCover

`func (o *FundamentalData) HasShortIntDayToCover() bool`

HasShortIntDayToCover returns a boolean if a field has been set.

### GetShortIntToFloat

`func (o *FundamentalData) GetShortIntToFloat() float32`

GetShortIntToFloat returns the ShortIntToFloat field if non-nil, zero value otherwise.

### GetShortIntToFloatOk

`func (o *FundamentalData) GetShortIntToFloatOk() (*float32, bool)`

GetShortIntToFloatOk returns a tuple with the ShortIntToFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortIntToFloat

`func (o *FundamentalData) SetShortIntToFloat(v float32)`

SetShortIntToFloat sets ShortIntToFloat field to given value.

### HasShortIntToFloat

`func (o *FundamentalData) HasShortIntToFloat() bool`

HasShortIntToFloat returns a boolean if a field has been set.

### GetSymbol

`func (o *FundamentalData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FundamentalData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FundamentalData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FundamentalData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTotalDebtToCapital

`func (o *FundamentalData) GetTotalDebtToCapital() float32`

GetTotalDebtToCapital returns the TotalDebtToCapital field if non-nil, zero value otherwise.

### GetTotalDebtToCapitalOk

`func (o *FundamentalData) GetTotalDebtToCapitalOk() (*float32, bool)`

GetTotalDebtToCapitalOk returns a tuple with the TotalDebtToCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebtToCapital

`func (o *FundamentalData) SetTotalDebtToCapital(v float32)`

SetTotalDebtToCapital sets TotalDebtToCapital field to given value.

### HasTotalDebtToCapital

`func (o *FundamentalData) HasTotalDebtToCapital() bool`

HasTotalDebtToCapital returns a boolean if a field has been set.

### GetTotalDebtToEquity

`func (o *FundamentalData) GetTotalDebtToEquity() float32`

GetTotalDebtToEquity returns the TotalDebtToEquity field if non-nil, zero value otherwise.

### GetTotalDebtToEquityOk

`func (o *FundamentalData) GetTotalDebtToEquityOk() (*float32, bool)`

GetTotalDebtToEquityOk returns a tuple with the TotalDebtToEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebtToEquity

`func (o *FundamentalData) SetTotalDebtToEquity(v float32)`

SetTotalDebtToEquity sets TotalDebtToEquity field to given value.

### HasTotalDebtToEquity

`func (o *FundamentalData) HasTotalDebtToEquity() bool`

HasTotalDebtToEquity returns a boolean if a field has been set.

### GetVol10DayAvg

`func (o *FundamentalData) GetVol10DayAvg() float32`

GetVol10DayAvg returns the Vol10DayAvg field if non-nil, zero value otherwise.

### GetVol10DayAvgOk

`func (o *FundamentalData) GetVol10DayAvgOk() (*float32, bool)`

GetVol10DayAvgOk returns a tuple with the Vol10DayAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol10DayAvg

`func (o *FundamentalData) SetVol10DayAvg(v float32)`

SetVol10DayAvg sets Vol10DayAvg field to given value.

### HasVol10DayAvg

`func (o *FundamentalData) HasVol10DayAvg() bool`

HasVol10DayAvg returns a boolean if a field has been set.

### GetVol1DayAvg

`func (o *FundamentalData) GetVol1DayAvg() float32`

GetVol1DayAvg returns the Vol1DayAvg field if non-nil, zero value otherwise.

### GetVol1DayAvgOk

`func (o *FundamentalData) GetVol1DayAvgOk() (*float32, bool)`

GetVol1DayAvgOk returns a tuple with the Vol1DayAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol1DayAvg

`func (o *FundamentalData) SetVol1DayAvg(v float32)`

SetVol1DayAvg sets Vol1DayAvg field to given value.

### HasVol1DayAvg

`func (o *FundamentalData) HasVol1DayAvg() bool`

HasVol1DayAvg returns a boolean if a field has been set.

### GetVol3MonthAvg

`func (o *FundamentalData) GetVol3MonthAvg() float32`

GetVol3MonthAvg returns the Vol3MonthAvg field if non-nil, zero value otherwise.

### GetVol3MonthAvgOk

`func (o *FundamentalData) GetVol3MonthAvgOk() (*float32, bool)`

GetVol3MonthAvgOk returns a tuple with the Vol3MonthAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol3MonthAvg

`func (o *FundamentalData) SetVol3MonthAvg(v float32)`

SetVol3MonthAvg sets Vol3MonthAvg field to given value.

### HasVol3MonthAvg

`func (o *FundamentalData) HasVol3MonthAvg() bool`

HasVol3MonthAvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


