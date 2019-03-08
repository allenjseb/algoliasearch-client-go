// Code generated by go generate. DO NOT EDIT.

package search

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

type Settings struct {
	SearchableAttributes              *opt.SearchableAttributesOption              `json:"searchableAttributes,omitempty"`
	AttributesForFaceting             *opt.AttributesForFacetingOption             `json:"attributesForFaceting,omitempty"`
	UnretrievableAttributes           *opt.UnretrievableAttributesOption           `json:"unretrievableAttributes,omitempty"`
	AttributesToRetrieve              *opt.AttributesToRetrieveOption              `json:"attributesToRetrieve,omitempty"`
	Ranking                           *opt.RankingOption                           `json:"ranking,omitempty"`
	CustomRanking                     *opt.CustomRankingOption                     `json:"customRanking,omitempty"`
	Replicas                          *opt.ReplicasOption                          `json:"replicas,omitempty"`
	MaxValuesPerFacet                 *opt.MaxValuesPerFacetOption                 `json:"maxValuesPerFacet,omitempty"`
	SortFacetValuesBy                 *opt.SortFacetValuesByOption                 `json:"sortFacetValuesBy,omitempty"`
	AttributesToHighlight             *opt.AttributesToHighlightOption             `json:"attributesToHighlight,omitempty"`
	AttributesToSnippet               *opt.AttributesToSnippetOption               `json:"attributesToSnippet,omitempty"`
	HighlightPreTag                   *opt.HighlightPreTagOption                   `json:"highlightPreTag,omitempty"`
	HighlightPostTag                  *opt.HighlightPostTagOption                  `json:"highlightPostTag,omitempty"`
	SnippetEllipsisText               *opt.SnippetEllipsisTextOption               `json:"snippetEllipsisText,omitempty"`
	RestrictHighlightAndSnippetArrays *opt.RestrictHighlightAndSnippetArraysOption `json:"restrictHighlightAndSnippetArrays,omitempty"`
	HitsPerPage                       *opt.HitsPerPageOption                       `json:"hitsPerPage,omitempty"`
	PaginationLimitedTo               *opt.PaginationLimitedToOption               `json:"paginationLimitedTo,omitempty"`
	MinWordSizefor1Typo               *opt.MinWordSizefor1TypoOption               `json:"minWordSizefor1Typo,omitempty"`
	MinWordSizefor2Typos              *opt.MinWordSizefor2TyposOption              `json:"minWordSizefor2Typos,omitempty"`
	TypoTolerance                     *opt.TypoToleranceOption                     `json:"typoTolerance,omitempty"`
	AllowTyposOnNumericTokens         *opt.AllowTyposOnNumericTokensOption         `json:"allowTyposOnNumericTokens,omitempty"`
	DisableTypoToleranceOnAttributes  *opt.DisableTypoToleranceOnAttributesOption  `json:"disableTypoToleranceOnAttributes,omitempty"`
	DisableTypoToleranceOnWords       *opt.DisableTypoToleranceOnWordsOption       `json:"disableTypoToleranceOnWords,omitempty"`
	SeparatorsToIndex                 *opt.SeparatorsToIndexOption                 `json:"separatorsToIndex,omitempty"`
	IgnorePlurals                     *opt.IgnorePluralsOption                     `json:"ignorePlurals,omitempty"`
	RemoveStopWords                   *opt.RemoveStopWordsOption                   `json:"removeStopWords,omitempty"`
	CamelCaseAttributes               *opt.CamelCaseAttributesOption               `json:"camelCaseAttributes,omitempty"`
	DecompoundedAttributes            *opt.DecompoundedAttributesOption            `json:"decompoundedAttributes,omitempty"`
	KeepDiacriticsOnCharacters        *opt.KeepDiacriticsOnCharactersOption        `json:"keepDiacriticsOnCharacters,omitempty"`
	QueryLanguages                    *opt.QueryLanguagesOption                    `json:"queryLanguages,omitempty"`
	QueryType                         *opt.QueryTypeOption                         `json:"queryType,omitempty"`
	RemoveWordsIfNoResults            *opt.RemoveWordsIfNoResultsOption            `json:"removeWordsIfNoResults,omitempty"`
	AdvancedSyntax                    *opt.AdvancedSyntaxOption                    `json:"advancedSyntax,omitempty"`
	OptionalWords                     *opt.OptionalWordsOption                     `json:"optionalWords,omitempty"`
	DisablePrefixOnAttributes         *opt.DisablePrefixOnAttributesOption         `json:"disablePrefixOnAttributes,omitempty"`
	DisableExactOnAttributes          *opt.DisableExactOnAttributesOption          `json:"disableExactOnAttributes,omitempty"`
	ExactOnSingleWordQuery            *opt.ExactOnSingleWordQueryOption            `json:"exactOnSingleWordQuery,omitempty"`
	AlternativesAsExact               *opt.AlternativesAsExactOption               `json:"alternativesAsExact,omitempty"`
	AdvancedSyntaxFeatures            *opt.AdvancedSyntaxFeaturesOption            `json:"advancedSyntaxFeatures,omitempty"`
	EnableRules                       *opt.EnableRulesOption                       `json:"enableRules,omitempty"`
	NumericAttributesForFiltering     *opt.NumericAttributesForFilteringOption     `json:"numericAttributesForFiltering,omitempty"`
	AllowCompressionOfIntegerArray    *opt.AllowCompressionOfIntegerArrayOption    `json:"allowCompressionOfIntegerArray,omitempty"`
	AttributeForDistinct              *opt.AttributeForDistinctOption              `json:"attributeForDistinct,omitempty"`
	Distinct                          *opt.DistinctOption                          `json:"distinct,omitempty"`
	ReplaceSynonymsInHighlight        *opt.ReplaceSynonymsInHighlightOption        `json:"replaceSynonymsInHighlight,omitempty"`
	MinProximity                      *opt.MinProximityOption                      `json:"minProximity,omitempty"`
	ResponseFields                    *opt.ResponseFieldsOption                    `json:"responseFields,omitempty"`
	MaxFacetHits                      *opt.MaxFacetHitsOption                      `json:"maxFacetHits,omitempty"`
}

type settings Settings

type backwardCompatibleSettings struct {
	AttributesToIndex        *opt.SearchableAttributesOption          `json:"attributesToIndex,omitempty"`
	Slaves                   *opt.ReplicasOption                      `json:"slaves,omitempty"`
	NumericAttributesToIndex *opt.NumericAttributesForFilteringOption `json:"numericAttributesToIndex,omitempty"`
}

func (s *Settings) UnmarshalJSON(data []byte) error {
	var bcSettings backwardCompatibleSettings
	err := json.Unmarshal(data, &bcSettings)
	if err != nil {
		return fmt.Errorf("cannot unmarshal backward-compatible settings: %v", err)
	}

	err = json.Unmarshal(data, (*settings)(s))
	if err != nil {
		return fmt.Errorf("cannot unmarshal settings: %v", err)
	}

	if s.SearchableAttributes == nil {
		s.SearchableAttributes = bcSettings.AttributesToIndex
	}

	if s.Replicas == nil {
		s.Replicas = bcSettings.Slaves
	}

	if s.NumericAttributesForFiltering == nil {
		s.NumericAttributesForFiltering = bcSettings.NumericAttributesToIndex
	}

	return nil
}