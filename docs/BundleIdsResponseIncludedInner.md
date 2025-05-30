# BundleIdsResponseIncludedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppAttributes**](AppAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**ProfileRelationships**](ProfileRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewBundleIdsResponseIncludedInner

`func NewBundleIdsResponseIncludedInner(type_ string, id string, ) *BundleIdsResponseIncludedInner`

NewBundleIdsResponseIncludedInner instantiates a new BundleIdsResponseIncludedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleIdsResponseIncludedInnerWithDefaults

`func NewBundleIdsResponseIncludedInnerWithDefaults() *BundleIdsResponseIncludedInner`

NewBundleIdsResponseIncludedInnerWithDefaults instantiates a new BundleIdsResponseIncludedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BundleIdsResponseIncludedInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BundleIdsResponseIncludedInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BundleIdsResponseIncludedInner) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BundleIdsResponseIncludedInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BundleIdsResponseIncludedInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BundleIdsResponseIncludedInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BundleIdsResponseIncludedInner) GetAttributes() AppAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BundleIdsResponseIncludedInner) GetAttributesOk() (*AppAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BundleIdsResponseIncludedInner) SetAttributes(v AppAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BundleIdsResponseIncludedInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *BundleIdsResponseIncludedInner) GetRelationships() ProfileRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BundleIdsResponseIncludedInner) GetRelationshipsOk() (*ProfileRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BundleIdsResponseIncludedInner) SetRelationships(v ProfileRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BundleIdsResponseIncludedInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *BundleIdsResponseIncludedInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BundleIdsResponseIncludedInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BundleIdsResponseIncludedInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BundleIdsResponseIncludedInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


