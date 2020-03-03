// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains simple golden tests for various examples.
// Besides validating the results when the implementation changes,
// it provides a way to look at the generated code without having
// to execute the print statements in one's head.

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Golden represents a test case.
type Golden struct {
	name   string
	input  string // input; the package clause is provided when running the test.
	output string // expected output.
}

var golden = []Golden{
	{"day", dayIn, dayOut},
	{"offset", offsetIn, offsetOut},
	{"gap", gapIn, gapOut},
	{"num", numIn, numOut},
	{"unum", unumIn, unumOut},
	{"prime", primeIn, primeOut},
}

var goldenJSON = []Golden{
	{"primeJson", primeJsonIn, primeJsonOut},
}
var goldenText = []Golden{
	{"primeText", primeTextIn, primeTextOut},
}

var goldenYAML = []Golden{
	{"primeYaml", primeYamlIn, primeYamlOut},
}

var goldenSQL = []Golden{
	{"primeSql", primeSqlIn, primeSqlOut},
}

var goldenCQL = []Golden{
	{"primeCql", primeCqlIn, primeCqlOut},
}

var goldenJSONAndSQL = []Golden{
	{"primeJsonAndSql", primeJsonAndSqlIn, primeJsonAndSqlOut},
}

var goldenTrimPrefix = []Golden{
	{"trimPrefix", trimPrefixIn, trimPrefixOut},
}

var goldenTrimPrefixMultiple = []Golden{
	{"trimPrefixMultiple", trimPrefixMultipleIn, dayNightOut},
}

var goldenWithPrefix = []Golden{
	{"dayWithPrefix", dayIn, prefixedDayOut},
}

var goldenTrimAndAddPrefix = []Golden{
	{"dayTrimAndPrefix", trimPrefixIn, trimmedPrefixedDayOut},
}

// Each example starts with "type XXX [u]int", with a single space separating them.

// Simple test: enumeration of type int starting at 0.
const dayIn = `type Day int
const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
`
const trimPrefixOut = `
const _DayName = "MondayTuesdayWednesdayThursdayFridaySaturdaySunday"

var _DayIndex = [...]uint8{0, 6, 13, 22, 30, 36, 44, 50}

const _DayLowerName = "mondaytuesdaywednesdaythursdayfridaysaturdaysunday"

func (i Day) String() string {
	if i < 0 || i >= Day(len(_DayIndex)-1) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return _DayName[_DayIndex[i]:_DayIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _DayNoOp() {
	var x [1]struct{}
	_ = x[DayMonday-(0)]
	_ = x[DayTuesday-(1)]
	_ = x[DayWednesday-(2)]
	_ = x[DayThursday-(3)]
	_ = x[DayFriday-(4)]
	_ = x[DaySaturday-(5)]
	_ = x[DaySunday-(6)]
}

var _DayValues = []Day{DayMonday, DayTuesday, DayWednesday, DayThursday, DayFriday, DaySaturday, DaySunday}

var _DayNameToValueMap = map[string]Day{
	_DayName[0:6]:        DayMonday,
	_DayLowerName[0:6]:   DayMonday,
	_DayName[6:13]:       DayTuesday,
	_DayLowerName[6:13]:  DayTuesday,
	_DayName[13:22]:      DayWednesday,
	_DayLowerName[13:22]: DayWednesday,
	_DayName[22:30]:      DayThursday,
	_DayLowerName[22:30]: DayThursday,
	_DayName[30:36]:      DayFriday,
	_DayLowerName[30:36]: DayFriday,
	_DayName[36:44]:      DaySaturday,
	_DayLowerName[36:44]: DaySaturday,
	_DayName[44:50]:      DaySunday,
	_DayLowerName[44:50]: DaySunday,
}

var _DayNames = []string{
	_DayName[0:6],
	_DayName[6:13],
	_DayName[13:22],
	_DayName[22:30],
	_DayName[30:36],
	_DayName[36:44],
	_DayName[44:50],
}

// DayString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DayString(s string) (Day, error) {
	if val, ok := _DayNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Day values", s)
}

// DayValues returns all values of the enum
func DayValues() []Day {
	return _DayValues
}

// DayStrings returns a slice of all String values of the enum
func DayStrings() []string {
	strs := make([]string, len(_DayNames))
	copy(strs, _DayNames)
	return strs
}

// IsADay returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Day) IsADay() bool {
	for _, v := range _DayValues {
		if i == v {
			return true
		}
	}
	return false
}
`

const dayOut = `
const _DayName = "MondayTuesdayWednesdayThursdayFridaySaturdaySunday"

var _DayIndex = [...]uint8{0, 6, 13, 22, 30, 36, 44, 50}

const _DayLowerName = "mondaytuesdaywednesdaythursdayfridaysaturdaysunday"

func (i Day) String() string {
	if i < 0 || i >= Day(len(_DayIndex)-1) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return _DayName[_DayIndex[i]:_DayIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _DayNoOp() {
	var x [1]struct{}
	_ = x[Monday-(0)]
	_ = x[Tuesday-(1)]
	_ = x[Wednesday-(2)]
	_ = x[Thursday-(3)]
	_ = x[Friday-(4)]
	_ = x[Saturday-(5)]
	_ = x[Sunday-(6)]
}

var _DayValues = []Day{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}

var _DayNameToValueMap = map[string]Day{
	_DayName[0:6]:        Monday,
	_DayLowerName[0:6]:   Monday,
	_DayName[6:13]:       Tuesday,
	_DayLowerName[6:13]:  Tuesday,
	_DayName[13:22]:      Wednesday,
	_DayLowerName[13:22]: Wednesday,
	_DayName[22:30]:      Thursday,
	_DayLowerName[22:30]: Thursday,
	_DayName[30:36]:      Friday,
	_DayLowerName[30:36]: Friday,
	_DayName[36:44]:      Saturday,
	_DayLowerName[36:44]: Saturday,
	_DayName[44:50]:      Sunday,
	_DayLowerName[44:50]: Sunday,
}

var _DayNames = []string{
	_DayName[0:6],
	_DayName[6:13],
	_DayName[13:22],
	_DayName[22:30],
	_DayName[30:36],
	_DayName[36:44],
	_DayName[44:50],
}

// DayString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DayString(s string) (Day, error) {
	if val, ok := _DayNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Day values", s)
}

// DayValues returns all values of the enum
func DayValues() []Day {
	return _DayValues
}

// DayStrings returns a slice of all String values of the enum
func DayStrings() []string {
	strs := make([]string, len(_DayNames))
	copy(strs, _DayNames)
	return strs
}

// IsADay returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Day) IsADay() bool {
	for _, v := range _DayValues {
		if i == v {
			return true
		}
	}
	return false
}
`

const dayNightOut = `
const _DayName = "MondayTuesdayWednesdayThursdayFridaySaturdaySunday"

var _DayIndex = [...]uint8{0, 6, 13, 22, 30, 36, 44, 50}

const _DayLowerName = "mondaytuesdaywednesdaythursdayfridaysaturdaysunday"

func (i Day) String() string {
	if i < 0 || i >= Day(len(_DayIndex)-1) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return _DayName[_DayIndex[i]:_DayIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _DayNoOp() {
	var x [1]struct{}
	_ = x[DayMonday-(0)]
	_ = x[NightTuesday-(1)]
	_ = x[DayWednesday-(2)]
	_ = x[NightThursday-(3)]
	_ = x[DayFriday-(4)]
	_ = x[NightSaturday-(5)]
	_ = x[DaySunday-(6)]
}

var _DayValues = []Day{DayMonday, NightTuesday, DayWednesday, NightThursday, DayFriday, NightSaturday, DaySunday}

var _DayNameToValueMap = map[string]Day{
	_DayName[0:6]:        DayMonday,
	_DayLowerName[0:6]:   DayMonday,
	_DayName[6:13]:       NightTuesday,
	_DayLowerName[6:13]:  NightTuesday,
	_DayName[13:22]:      DayWednesday,
	_DayLowerName[13:22]: DayWednesday,
	_DayName[22:30]:      NightThursday,
	_DayLowerName[22:30]: NightThursday,
	_DayName[30:36]:      DayFriday,
	_DayLowerName[30:36]: DayFriday,
	_DayName[36:44]:      NightSaturday,
	_DayLowerName[36:44]: NightSaturday,
	_DayName[44:50]:      DaySunday,
	_DayLowerName[44:50]: DaySunday,
}

var _DayNames = []string{
	_DayName[0:6],
	_DayName[6:13],
	_DayName[13:22],
	_DayName[22:30],
	_DayName[30:36],
	_DayName[36:44],
	_DayName[44:50],
}

// DayString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DayString(s string) (Day, error) {
	if val, ok := _DayNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Day values", s)
}

// DayValues returns all values of the enum
func DayValues() []Day {
	return _DayValues
}

// DayStrings returns a slice of all String values of the enum
func DayStrings() []string {
	strs := make([]string, len(_DayNames))
	copy(strs, _DayNames)
	return strs
}

// IsADay returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Day) IsADay() bool {
	for _, v := range _DayValues {
		if i == v {
			return true
		}
	}
	return false
}
`

const prefixedDayOut = `
const _DayName = "DayMondayDayTuesdayDayWednesdayDayThursdayDayFridayDaySaturdayDaySunday"

var _DayIndex = [...]uint8{0, 9, 19, 31, 42, 51, 62, 71}

const _DayLowerName = "daymondaydaytuesdaydaywednesdaydaythursdaydayfridaydaysaturdaydaysunday"

func (i Day) String() string {
	if i < 0 || i >= Day(len(_DayIndex)-1) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return _DayName[_DayIndex[i]:_DayIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _DayNoOp() {
	var x [1]struct{}
	_ = x[Monday-(0)]
	_ = x[Tuesday-(1)]
	_ = x[Wednesday-(2)]
	_ = x[Thursday-(3)]
	_ = x[Friday-(4)]
	_ = x[Saturday-(5)]
	_ = x[Sunday-(6)]
}

var _DayValues = []Day{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}

var _DayNameToValueMap = map[string]Day{
	_DayName[0:9]:        Monday,
	_DayLowerName[0:9]:   Monday,
	_DayName[9:19]:       Tuesday,
	_DayLowerName[9:19]:  Tuesday,
	_DayName[19:31]:      Wednesday,
	_DayLowerName[19:31]: Wednesday,
	_DayName[31:42]:      Thursday,
	_DayLowerName[31:42]: Thursday,
	_DayName[42:51]:      Friday,
	_DayLowerName[42:51]: Friday,
	_DayName[51:62]:      Saturday,
	_DayLowerName[51:62]: Saturday,
	_DayName[62:71]:      Sunday,
	_DayLowerName[62:71]: Sunday,
}

var _DayNames = []string{
	_DayName[0:9],
	_DayName[9:19],
	_DayName[19:31],
	_DayName[31:42],
	_DayName[42:51],
	_DayName[51:62],
	_DayName[62:71],
}

// DayString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DayString(s string) (Day, error) {
	if val, ok := _DayNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Day values", s)
}

// DayValues returns all values of the enum
func DayValues() []Day {
	return _DayValues
}

// DayStrings returns a slice of all String values of the enum
func DayStrings() []string {
	strs := make([]string, len(_DayNames))
	copy(strs, _DayNames)
	return strs
}

// IsADay returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Day) IsADay() bool {
	for _, v := range _DayValues {
		if i == v {
			return true
		}
	}
	return false
}
`

const trimmedPrefixedDayOut = `
const _DayName = "NightMondayNightTuesdayNightWednesdayNightThursdayNightFridayNightSaturdayNightSunday"

var _DayIndex = [...]uint8{0, 11, 23, 37, 50, 61, 74, 85}

const _DayLowerName = "nightmondaynighttuesdaynightwednesdaynightthursdaynightfridaynightsaturdaynightsunday"

func (i Day) String() string {
	if i < 0 || i >= Day(len(_DayIndex)-1) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return _DayName[_DayIndex[i]:_DayIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _DayNoOp() {
	var x [1]struct{}
	_ = x[DayMonday-(0)]
	_ = x[DayTuesday-(1)]
	_ = x[DayWednesday-(2)]
	_ = x[DayThursday-(3)]
	_ = x[DayFriday-(4)]
	_ = x[DaySaturday-(5)]
	_ = x[DaySunday-(6)]
}

var _DayValues = []Day{DayMonday, DayTuesday, DayWednesday, DayThursday, DayFriday, DaySaturday, DaySunday}

var _DayNameToValueMap = map[string]Day{
	_DayName[0:11]:       DayMonday,
	_DayLowerName[0:11]:  DayMonday,
	_DayName[11:23]:      DayTuesday,
	_DayLowerName[11:23]: DayTuesday,
	_DayName[23:37]:      DayWednesday,
	_DayLowerName[23:37]: DayWednesday,
	_DayName[37:50]:      DayThursday,
	_DayLowerName[37:50]: DayThursday,
	_DayName[50:61]:      DayFriday,
	_DayLowerName[50:61]: DayFriday,
	_DayName[61:74]:      DaySaturday,
	_DayLowerName[61:74]: DaySaturday,
	_DayName[74:85]:      DaySunday,
	_DayLowerName[74:85]: DaySunday,
}

var _DayNames = []string{
	_DayName[0:11],
	_DayName[11:23],
	_DayName[23:37],
	_DayName[37:50],
	_DayName[50:61],
	_DayName[61:74],
	_DayName[74:85],
}

// DayString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DayString(s string) (Day, error) {
	if val, ok := _DayNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Day values", s)
}

// DayValues returns all values of the enum
func DayValues() []Day {
	return _DayValues
}

// DayStrings returns a slice of all String values of the enum
func DayStrings() []string {
	strs := make([]string, len(_DayNames))
	copy(strs, _DayNames)
	return strs
}

// IsADay returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Day) IsADay() bool {
	for _, v := range _DayValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Enumeration with an offset.
// Also includes a duplicate.
const offsetIn = `type Number int
const (
	_ Number = iota
	One
	Two
	Three
	AnotherOne = One  // Duplicate; note that AnotherOne doesn't appear below.
)
`

const offsetOut = `
const _NumberName = "OneTwoThree"

var _NumberIndex = [...]uint8{0, 3, 6, 11}

const _NumberLowerName = "onetwothree"

func (i Number) String() string {
	i -= 1
	if i < 0 || i >= Number(len(_NumberIndex)-1) {
		return fmt.Sprintf("Number(%d)", i+1)
	}
	return _NumberName[_NumberIndex[i]:_NumberIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NumberNoOp() {
	var x [1]struct{}
	_ = x[One-(1)]
	_ = x[Two-(2)]
	_ = x[Three-(3)]
}

var _NumberValues = []Number{One, Two, Three}

var _NumberNameToValueMap = map[string]Number{
	_NumberName[0:3]:       One,
	_NumberLowerName[0:3]:  One,
	_NumberName[3:6]:       Two,
	_NumberLowerName[3:6]:  Two,
	_NumberName[6:11]:      Three,
	_NumberLowerName[6:11]: Three,
}

var _NumberNames = []string{
	_NumberName[0:3],
	_NumberName[3:6],
	_NumberName[6:11],
}

// NumberString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NumberString(s string) (Number, error) {
	if val, ok := _NumberNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Number values", s)
}

// NumberValues returns all values of the enum
func NumberValues() []Number {
	return _NumberValues
}

// NumberStrings returns a slice of all String values of the enum
func NumberStrings() []string {
	strs := make([]string, len(_NumberNames))
	copy(strs, _NumberNames)
	return strs
}

// IsANumber returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Number) IsANumber() bool {
	for _, v := range _NumberValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Gaps and an offset.
const gapIn = `type Gap int
const (
	Two Gap = 2
	Three Gap = 3
	Five Gap = 5
	Six Gap = 6
	Seven Gap = 7
	Eight Gap = 8
	Nine Gap = 9
	Eleven Gap = 11
)
`

const gapOut = `
const (
	_GapName_0      = "TwoThree"
	_GapLowerName_0 = "twothree"
	_GapName_1      = "FiveSixSevenEightNine"
	_GapLowerName_1 = "fivesixseveneightnine"
	_GapName_2      = "Eleven"
	_GapLowerName_2 = "eleven"
)

var (
	_GapIndex_0 = [...]uint8{0, 3, 8}
	_GapIndex_1 = [...]uint8{0, 4, 7, 12, 17, 21}
	_GapIndex_2 = [...]uint8{0, 6}
)

func (i Gap) String() string {
	switch {
	case 2 <= i && i <= 3:
		i -= 2
		return _GapName_0[_GapIndex_0[i]:_GapIndex_0[i+1]]
	case 5 <= i && i <= 9:
		i -= 5
		return _GapName_1[_GapIndex_1[i]:_GapIndex_1[i+1]]
	case i == 11:
		return _GapName_2
	default:
		return fmt.Sprintf("Gap(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _GapNoOp() {
	var x [1]struct{}
	_ = x[Two-(2)]
	_ = x[Three-(3)]
	_ = x[Five-(5)]
	_ = x[Six-(6)]
	_ = x[Seven-(7)]
	_ = x[Eight-(8)]
	_ = x[Nine-(9)]
	_ = x[Eleven-(11)]
}

var _GapValues = []Gap{Two, Three, Five, Six, Seven, Eight, Nine, Eleven}

var _GapNameToValueMap = map[string]Gap{
	_GapName_0[0:3]:        Two,
	_GapLowerName_0[0:3]:   Two,
	_GapName_0[3:8]:        Three,
	_GapLowerName_0[3:8]:   Three,
	_GapName_1[0:4]:        Five,
	_GapLowerName_1[0:4]:   Five,
	_GapName_1[4:7]:        Six,
	_GapLowerName_1[4:7]:   Six,
	_GapName_1[7:12]:       Seven,
	_GapLowerName_1[7:12]:  Seven,
	_GapName_1[12:17]:      Eight,
	_GapLowerName_1[12:17]: Eight,
	_GapName_1[17:21]:      Nine,
	_GapLowerName_1[17:21]: Nine,
	_GapName_2[0:6]:        Eleven,
	_GapLowerName_2[0:6]:   Eleven,
}

var _GapNames = []string{
	_GapName_0[0:3],
	_GapName_0[3:8],
	_GapName_1[0:4],
	_GapName_1[4:7],
	_GapName_1[7:12],
	_GapName_1[12:17],
	_GapName_1[17:21],
	_GapName_2[0:6],
}

// GapString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func GapString(s string) (Gap, error) {
	if val, ok := _GapNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Gap values", s)
}

// GapValues returns all values of the enum
func GapValues() []Gap {
	return _GapValues
}

// GapStrings returns a slice of all String values of the enum
func GapStrings() []string {
	strs := make([]string, len(_GapNames))
	copy(strs, _GapNames)
	return strs
}

// IsAGap returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Gap) IsAGap() bool {
	for _, v := range _GapValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Signed integers spanning zero.
const numIn = `type Num int
const (
	m_2 Num = -2 + iota
	m_1
	m0
	m1
	m2
)
`

const numOut = `
const _NumName = "m_2m_1m0m1m2"

var _NumIndex = [...]uint8{0, 3, 6, 8, 10, 12}

const _NumLowerName = "m_2m_1m0m1m2"

func (i Num) String() string {
	i -= -2
	if i < 0 || i >= Num(len(_NumIndex)-1) {
		return fmt.Sprintf("Num(%d)", i+-2)
	}
	return _NumName[_NumIndex[i]:_NumIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NumNoOp() {
	var x [1]struct{}
	_ = x[m_2-(-2)]
	_ = x[m_1-(-1)]
	_ = x[m0-(0)]
	_ = x[m1-(1)]
	_ = x[m2-(2)]
}

var _NumValues = []Num{m_2, m_1, m0, m1, m2}

var _NumNameToValueMap = map[string]Num{
	_NumName[0:3]:        m_2,
	_NumLowerName[0:3]:   m_2,
	_NumName[3:6]:        m_1,
	_NumLowerName[3:6]:   m_1,
	_NumName[6:8]:        m0,
	_NumLowerName[6:8]:   m0,
	_NumName[8:10]:       m1,
	_NumLowerName[8:10]:  m1,
	_NumName[10:12]:      m2,
	_NumLowerName[10:12]: m2,
}

var _NumNames = []string{
	_NumName[0:3],
	_NumName[3:6],
	_NumName[6:8],
	_NumName[8:10],
	_NumName[10:12],
}

// NumString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NumString(s string) (Num, error) {
	if val, ok := _NumNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Num values", s)
}

// NumValues returns all values of the enum
func NumValues() []Num {
	return _NumValues
}

// NumStrings returns a slice of all String values of the enum
func NumStrings() []string {
	strs := make([]string, len(_NumNames))
	copy(strs, _NumNames)
	return strs
}

// IsANum returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Num) IsANum() bool {
	for _, v := range _NumValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Unsigned integers spanning zero.
const unumIn = `type Unum uint
const (
	m_2 Unum = iota + 253
	m_1
)

const (
	m0 Unum = iota
	m1
	m2
)
`

const unumOut = `
const (
	_UnumName_0      = "m0m1m2"
	_UnumLowerName_0 = "m0m1m2"
	_UnumName_1      = "m_2m_1"
	_UnumLowerName_1 = "m_2m_1"
)

var (
	_UnumIndex_0 = [...]uint8{0, 2, 4, 6}
	_UnumIndex_1 = [...]uint8{0, 3, 6}
)

func (i Unum) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _UnumName_0[_UnumIndex_0[i]:_UnumIndex_0[i+1]]
	case 253 <= i && i <= 254:
		i -= 253
		return _UnumName_1[_UnumIndex_1[i]:_UnumIndex_1[i+1]]
	default:
		return fmt.Sprintf("Unum(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _UnumNoOp() {
	var x [1]struct{}
	_ = x[m0-(0)]
	_ = x[m1-(1)]
	_ = x[m2-(2)]
	_ = x[m_2-(253)]
	_ = x[m_1-(254)]
}

var _UnumValues = []Unum{m0, m1, m2, m_2, m_1}

var _UnumNameToValueMap = map[string]Unum{
	_UnumName_0[0:2]:      m0,
	_UnumLowerName_0[0:2]: m0,
	_UnumName_0[2:4]:      m1,
	_UnumLowerName_0[2:4]: m1,
	_UnumName_0[4:6]:      m2,
	_UnumLowerName_0[4:6]: m2,
	_UnumName_1[0:3]:      m_2,
	_UnumLowerName_1[0:3]: m_2,
	_UnumName_1[3:6]:      m_1,
	_UnumLowerName_1[3:6]: m_1,
}

var _UnumNames = []string{
	_UnumName_0[0:2],
	_UnumName_0[2:4],
	_UnumName_0[4:6],
	_UnumName_1[0:3],
	_UnumName_1[3:6],
}

// UnumString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func UnumString(s string) (Unum, error) {
	if val, ok := _UnumNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Unum values", s)
}

// UnumValues returns all values of the enum
func UnumValues() []Unum {
	return _UnumValues
}

// UnumStrings returns a slice of all String values of the enum
func UnumStrings() []string {
	strs := make([]string, len(_UnumNames))
	copy(strs, _UnumNames)
	return strs
}

// IsAUnum returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Unum) IsAUnum() bool {
	for _, v := range _UnumValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Enough gaps to trigger a map implementation of the method.
// Also includes a duplicate to test that it doesn't cause problems
const primeIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"
const _PrimeLowerName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PrimeNoOp() {
	var x [1]struct{}
	_ = x[p2-(2)]
	_ = x[p3-(3)]
	_ = x[p5-(5)]
	_ = x[p7-(7)]
	_ = x[p11-(11)]
	_ = x[p13-(13)]
	_ = x[p17-(17)]
	_ = x[p19-(19)]
	_ = x[p23-(23)]
	_ = x[p29-(29)]
	_ = x[p37-(31)]
	_ = x[p41-(41)]
	_ = x[p43-(43)]
}

var _PrimeValues = []Prime{p2, p3, p5, p7, p11, p13, p17, p19, p23, p29, p37, p41, p43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:        p2,
	_PrimeLowerName[0:2]:   p2,
	_PrimeName[2:4]:        p3,
	_PrimeLowerName[2:4]:   p3,
	_PrimeName[4:6]:        p5,
	_PrimeLowerName[4:6]:   p5,
	_PrimeName[6:8]:        p7,
	_PrimeLowerName[6:8]:   p7,
	_PrimeName[8:11]:       p11,
	_PrimeLowerName[8:11]:  p11,
	_PrimeName[11:14]:      p13,
	_PrimeLowerName[11:14]: p13,
	_PrimeName[14:17]:      p17,
	_PrimeLowerName[14:17]: p17,
	_PrimeName[17:20]:      p19,
	_PrimeLowerName[17:20]: p19,
	_PrimeName[20:23]:      p23,
	_PrimeLowerName[20:23]: p23,
	_PrimeName[23:26]:      p29,
	_PrimeLowerName[23:26]: p29,
	_PrimeName[26:29]:      p37,
	_PrimeLowerName[26:29]: p37,
	_PrimeName[29:32]:      p41,
	_PrimeLowerName[29:32]: p41,
	_PrimeName[32:35]:      p43,
	_PrimeLowerName[32:35]: p43,
}

var _PrimeNames = []string{
	_PrimeName[0:2],
	_PrimeName[2:4],
	_PrimeName[4:6],
	_PrimeName[6:8],
	_PrimeName[8:11],
	_PrimeName[11:14],
	_PrimeName[14:17],
	_PrimeName[17:20],
	_PrimeName[20:23],
	_PrimeName[23:26],
	_PrimeName[26:29],
	_PrimeName[29:32],
	_PrimeName[32:35],
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// PrimeStrings returns a slice of all String values of the enum
func PrimeStrings() []string {
	strs := make([]string, len(_PrimeNames))
	copy(strs, _PrimeNames)
	return strs
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}
`
const primeJsonIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeJsonOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"
const _PrimeLowerName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PrimeNoOp() {
	var x [1]struct{}
	_ = x[p2-(2)]
	_ = x[p3-(3)]
	_ = x[p5-(5)]
	_ = x[p7-(7)]
	_ = x[p11-(11)]
	_ = x[p13-(13)]
	_ = x[p17-(17)]
	_ = x[p19-(19)]
	_ = x[p23-(23)]
	_ = x[p29-(29)]
	_ = x[p37-(31)]
	_ = x[p41-(41)]
	_ = x[p43-(43)]
}

var _PrimeValues = []Prime{p2, p3, p5, p7, p11, p13, p17, p19, p23, p29, p37, p41, p43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:        p2,
	_PrimeLowerName[0:2]:   p2,
	_PrimeName[2:4]:        p3,
	_PrimeLowerName[2:4]:   p3,
	_PrimeName[4:6]:        p5,
	_PrimeLowerName[4:6]:   p5,
	_PrimeName[6:8]:        p7,
	_PrimeLowerName[6:8]:   p7,
	_PrimeName[8:11]:       p11,
	_PrimeLowerName[8:11]:  p11,
	_PrimeName[11:14]:      p13,
	_PrimeLowerName[11:14]: p13,
	_PrimeName[14:17]:      p17,
	_PrimeLowerName[14:17]: p17,
	_PrimeName[17:20]:      p19,
	_PrimeLowerName[17:20]: p19,
	_PrimeName[20:23]:      p23,
	_PrimeLowerName[20:23]: p23,
	_PrimeName[23:26]:      p29,
	_PrimeLowerName[23:26]: p29,
	_PrimeName[26:29]:      p37,
	_PrimeLowerName[26:29]: p37,
	_PrimeName[29:32]:      p41,
	_PrimeLowerName[29:32]: p41,
	_PrimeName[32:35]:      p43,
	_PrimeLowerName[32:35]: p43,
}

var _PrimeNames = []string{
	_PrimeName[0:2],
	_PrimeName[2:4],
	_PrimeName[4:6],
	_PrimeName[6:8],
	_PrimeName[8:11],
	_PrimeName[11:14],
	_PrimeName[14:17],
	_PrimeName[17:20],
	_PrimeName[20:23],
	_PrimeName[23:26],
	_PrimeName[26:29],
	_PrimeName[29:32],
	_PrimeName[32:35],
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// PrimeStrings returns a slice of all String values of the enum
func PrimeStrings() []string {
	strs := make([]string, len(_PrimeNames))
	copy(strs, _PrimeNames)
	return strs
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for Prime
func (i Prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Prime
func (i *Prime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Prime should be a string, got %s", data)
	}

	var err error
	*i, err = PrimeString(s)
	return err
}
`

const primeTextIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeTextOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"
const _PrimeLowerName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PrimeNoOp() {
	var x [1]struct{}
	_ = x[p2-(2)]
	_ = x[p3-(3)]
	_ = x[p5-(5)]
	_ = x[p7-(7)]
	_ = x[p11-(11)]
	_ = x[p13-(13)]
	_ = x[p17-(17)]
	_ = x[p19-(19)]
	_ = x[p23-(23)]
	_ = x[p29-(29)]
	_ = x[p37-(31)]
	_ = x[p41-(41)]
	_ = x[p43-(43)]
}

var _PrimeValues = []Prime{p2, p3, p5, p7, p11, p13, p17, p19, p23, p29, p37, p41, p43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:        p2,
	_PrimeLowerName[0:2]:   p2,
	_PrimeName[2:4]:        p3,
	_PrimeLowerName[2:4]:   p3,
	_PrimeName[4:6]:        p5,
	_PrimeLowerName[4:6]:   p5,
	_PrimeName[6:8]:        p7,
	_PrimeLowerName[6:8]:   p7,
	_PrimeName[8:11]:       p11,
	_PrimeLowerName[8:11]:  p11,
	_PrimeName[11:14]:      p13,
	_PrimeLowerName[11:14]: p13,
	_PrimeName[14:17]:      p17,
	_PrimeLowerName[14:17]: p17,
	_PrimeName[17:20]:      p19,
	_PrimeLowerName[17:20]: p19,
	_PrimeName[20:23]:      p23,
	_PrimeLowerName[20:23]: p23,
	_PrimeName[23:26]:      p29,
	_PrimeLowerName[23:26]: p29,
	_PrimeName[26:29]:      p37,
	_PrimeLowerName[26:29]: p37,
	_PrimeName[29:32]:      p41,
	_PrimeLowerName[29:32]: p41,
	_PrimeName[32:35]:      p43,
	_PrimeLowerName[32:35]: p43,
}

var _PrimeNames = []string{
	_PrimeName[0:2],
	_PrimeName[2:4],
	_PrimeName[4:6],
	_PrimeName[6:8],
	_PrimeName[8:11],
	_PrimeName[11:14],
	_PrimeName[14:17],
	_PrimeName[17:20],
	_PrimeName[20:23],
	_PrimeName[23:26],
	_PrimeName[26:29],
	_PrimeName[29:32],
	_PrimeName[32:35],
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// PrimeStrings returns a slice of all String values of the enum
func PrimeStrings() []string {
	strs := make([]string, len(_PrimeNames))
	copy(strs, _PrimeNames)
	return strs
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalText implements the encoding.TextMarshaler interface for Prime
func (i Prime) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Prime
func (i *Prime) UnmarshalText(text []byte) error {
	var err error
	*i, err = PrimeString(string(text))
	return err
}
`

const primeYamlIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeYamlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"
const _PrimeLowerName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PrimeNoOp() {
	var x [1]struct{}
	_ = x[p2-(2)]
	_ = x[p3-(3)]
	_ = x[p5-(5)]
	_ = x[p7-(7)]
	_ = x[p11-(11)]
	_ = x[p13-(13)]
	_ = x[p17-(17)]
	_ = x[p19-(19)]
	_ = x[p23-(23)]
	_ = x[p29-(29)]
	_ = x[p37-(31)]
	_ = x[p41-(41)]
	_ = x[p43-(43)]
}

var _PrimeValues = []Prime{p2, p3, p5, p7, p11, p13, p17, p19, p23, p29, p37, p41, p43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:        p2,
	_PrimeLowerName[0:2]:   p2,
	_PrimeName[2:4]:        p3,
	_PrimeLowerName[2:4]:   p3,
	_PrimeName[4:6]:        p5,
	_PrimeLowerName[4:6]:   p5,
	_PrimeName[6:8]:        p7,
	_PrimeLowerName[6:8]:   p7,
	_PrimeName[8:11]:       p11,
	_PrimeLowerName[8:11]:  p11,
	_PrimeName[11:14]:      p13,
	_PrimeLowerName[11:14]: p13,
	_PrimeName[14:17]:      p17,
	_PrimeLowerName[14:17]: p17,
	_PrimeName[17:20]:      p19,
	_PrimeLowerName[17:20]: p19,
	_PrimeName[20:23]:      p23,
	_PrimeLowerName[20:23]: p23,
	_PrimeName[23:26]:      p29,
	_PrimeLowerName[23:26]: p29,
	_PrimeName[26:29]:      p37,
	_PrimeLowerName[26:29]: p37,
	_PrimeName[29:32]:      p41,
	_PrimeLowerName[29:32]: p41,
	_PrimeName[32:35]:      p43,
	_PrimeLowerName[32:35]: p43,
}

var _PrimeNames = []string{
	_PrimeName[0:2],
	_PrimeName[2:4],
	_PrimeName[4:6],
	_PrimeName[6:8],
	_PrimeName[8:11],
	_PrimeName[11:14],
	_PrimeName[14:17],
	_PrimeName[17:20],
	_PrimeName[20:23],
	_PrimeName[23:26],
	_PrimeName[26:29],
	_PrimeName[29:32],
	_PrimeName[32:35],
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// PrimeStrings returns a slice of all String values of the enum
func PrimeStrings() []string {
	strs := make([]string, len(_PrimeNames))
	copy(strs, _PrimeNames)
	return strs
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalYAML implements a YAML Marshaler for Prime
func (i Prime) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Prime
func (i *Prime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = PrimeString(s)
	return err
}
`

const primeSqlIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeSqlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"
const _PrimeLowerName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PrimeNoOp() {
	var x [1]struct{}
	_ = x[p2-(2)]
	_ = x[p3-(3)]
	_ = x[p5-(5)]
	_ = x[p7-(7)]
	_ = x[p11-(11)]
	_ = x[p13-(13)]
	_ = x[p17-(17)]
	_ = x[p19-(19)]
	_ = x[p23-(23)]
	_ = x[p29-(29)]
	_ = x[p37-(31)]
	_ = x[p41-(41)]
	_ = x[p43-(43)]
}

var _PrimeValues = []Prime{p2, p3, p5, p7, p11, p13, p17, p19, p23, p29, p37, p41, p43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:        p2,
	_PrimeLowerName[0:2]:   p2,
	_PrimeName[2:4]:        p3,
	_PrimeLowerName[2:4]:   p3,
	_PrimeName[4:6]:        p5,
	_PrimeLowerName[4:6]:   p5,
	_PrimeName[6:8]:        p7,
	_PrimeLowerName[6:8]:   p7,
	_PrimeName[8:11]:       p11,
	_PrimeLowerName[8:11]:  p11,
	_PrimeName[11:14]:      p13,
	_PrimeLowerName[11:14]: p13,
	_PrimeName[14:17]:      p17,
	_PrimeLowerName[14:17]: p17,
	_PrimeName[17:20]:      p19,
	_PrimeLowerName[17:20]: p19,
	_PrimeName[20:23]:      p23,
	_PrimeLowerName[20:23]: p23,
	_PrimeName[23:26]:      p29,
	_PrimeLowerName[23:26]: p29,
	_PrimeName[26:29]:      p37,
	_PrimeLowerName[26:29]: p37,
	_PrimeName[29:32]:      p41,
	_PrimeLowerName[29:32]: p41,
	_PrimeName[32:35]:      p43,
	_PrimeLowerName[32:35]: p43,
}

var _PrimeNames = []string{
	_PrimeName[0:2],
	_PrimeName[2:4],
	_PrimeName[4:6],
	_PrimeName[6:8],
	_PrimeName[8:11],
	_PrimeName[11:14],
	_PrimeName[14:17],
	_PrimeName[17:20],
	_PrimeName[20:23],
	_PrimeName[23:26],
	_PrimeName[26:29],
	_PrimeName[29:32],
	_PrimeName[32:35],
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// PrimeStrings returns a slice of all String values of the enum
func PrimeStrings() []string {
	strs := make([]string, len(_PrimeNames))
	copy(strs, _PrimeNames)
	return strs
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

func (i Prime) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Prime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := PrimeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
`

const primeCqlIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeCqlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"
const _PrimeLowerName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PrimeNoOp() {
	var x [1]struct{}
	_ = x[p2-(2)]
	_ = x[p3-(3)]
	_ = x[p5-(5)]
	_ = x[p7-(7)]
	_ = x[p11-(11)]
	_ = x[p13-(13)]
	_ = x[p17-(17)]
	_ = x[p19-(19)]
	_ = x[p23-(23)]
	_ = x[p29-(29)]
	_ = x[p37-(31)]
	_ = x[p41-(41)]
	_ = x[p43-(43)]
}

var _PrimeValues = []Prime{p2, p3, p5, p7, p11, p13, p17, p19, p23, p29, p37, p41, p43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:        p2,
	_PrimeLowerName[0:2]:   p2,
	_PrimeName[2:4]:        p3,
	_PrimeLowerName[2:4]:   p3,
	_PrimeName[4:6]:        p5,
	_PrimeLowerName[4:6]:   p5,
	_PrimeName[6:8]:        p7,
	_PrimeLowerName[6:8]:   p7,
	_PrimeName[8:11]:       p11,
	_PrimeLowerName[8:11]:  p11,
	_PrimeName[11:14]:      p13,
	_PrimeLowerName[11:14]: p13,
	_PrimeName[14:17]:      p17,
	_PrimeLowerName[14:17]: p17,
	_PrimeName[17:20]:      p19,
	_PrimeLowerName[17:20]: p19,
	_PrimeName[20:23]:      p23,
	_PrimeLowerName[20:23]: p23,
	_PrimeName[23:26]:      p29,
	_PrimeLowerName[23:26]: p29,
	_PrimeName[26:29]:      p37,
	_PrimeLowerName[26:29]: p37,
	_PrimeName[29:32]:      p41,
	_PrimeLowerName[29:32]: p41,
	_PrimeName[32:35]:      p43,
	_PrimeLowerName[32:35]: p43,
}

var _PrimeNames = []string{
	_PrimeName[0:2],
	_PrimeName[2:4],
	_PrimeName[4:6],
	_PrimeName[6:8],
	_PrimeName[8:11],
	_PrimeName[11:14],
	_PrimeName[14:17],
	_PrimeName[17:20],
	_PrimeName[20:23],
	_PrimeName[23:26],
	_PrimeName[26:29],
	_PrimeName[29:32],
	_PrimeName[32:35],
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// PrimeStrings returns a slice of all String values of the enum
func PrimeStrings() []string {
	strs := make([]string, len(_PrimeNames))
	copy(strs, _PrimeNames)
	return strs
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalCQL implements the gocql.Marshaler interface for Prime
func (i Prime) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalCQL implements the gocql.Unmarshaler interface for Prime
func (i *Prime) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	var err error
	*i, err = PrimeString(string(data))
	return err
}
`

const primeJsonAndSqlIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeJsonAndSqlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"
const _PrimeLowerName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PrimeNoOp() {
	var x [1]struct{}
	_ = x[p2-(2)]
	_ = x[p3-(3)]
	_ = x[p5-(5)]
	_ = x[p7-(7)]
	_ = x[p11-(11)]
	_ = x[p13-(13)]
	_ = x[p17-(17)]
	_ = x[p19-(19)]
	_ = x[p23-(23)]
	_ = x[p29-(29)]
	_ = x[p37-(31)]
	_ = x[p41-(41)]
	_ = x[p43-(43)]
}

var _PrimeValues = []Prime{p2, p3, p5, p7, p11, p13, p17, p19, p23, p29, p37, p41, p43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:        p2,
	_PrimeLowerName[0:2]:   p2,
	_PrimeName[2:4]:        p3,
	_PrimeLowerName[2:4]:   p3,
	_PrimeName[4:6]:        p5,
	_PrimeLowerName[4:6]:   p5,
	_PrimeName[6:8]:        p7,
	_PrimeLowerName[6:8]:   p7,
	_PrimeName[8:11]:       p11,
	_PrimeLowerName[8:11]:  p11,
	_PrimeName[11:14]:      p13,
	_PrimeLowerName[11:14]: p13,
	_PrimeName[14:17]:      p17,
	_PrimeLowerName[14:17]: p17,
	_PrimeName[17:20]:      p19,
	_PrimeLowerName[17:20]: p19,
	_PrimeName[20:23]:      p23,
	_PrimeLowerName[20:23]: p23,
	_PrimeName[23:26]:      p29,
	_PrimeLowerName[23:26]: p29,
	_PrimeName[26:29]:      p37,
	_PrimeLowerName[26:29]: p37,
	_PrimeName[29:32]:      p41,
	_PrimeLowerName[29:32]: p41,
	_PrimeName[32:35]:      p43,
	_PrimeLowerName[32:35]: p43,
}

var _PrimeNames = []string{
	_PrimeName[0:2],
	_PrimeName[2:4],
	_PrimeName[4:6],
	_PrimeName[6:8],
	_PrimeName[8:11],
	_PrimeName[11:14],
	_PrimeName[14:17],
	_PrimeName[17:20],
	_PrimeName[20:23],
	_PrimeName[23:26],
	_PrimeName[26:29],
	_PrimeName[29:32],
	_PrimeName[32:35],
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// PrimeStrings returns a slice of all String values of the enum
func PrimeStrings() []string {
	strs := make([]string, len(_PrimeNames))
	copy(strs, _PrimeNames)
	return strs
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for Prime
func (i Prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Prime
func (i *Prime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Prime should be a string, got %s", data)
	}

	var err error
	*i, err = PrimeString(s)
	return err
}

func (i Prime) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Prime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := PrimeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
`

const trimPrefixIn = `type Day int
const (
	DayMonday Day = iota
	DayTuesday
	DayWednesday
	DayThursday
	DayFriday
	DaySaturday
	DaySunday
)
`

const trimPrefixMultipleIn = `type Day int
const (
	DayMonday Day = iota
	NightTuesday
	DayWednesday
	NightThursday
	DayFriday
	NightSaturday
	DaySunday
)
`

func TestGolden(t *testing.T) {
	for _, test := range golden {
		runGoldenTest(t, test, false, false, false, false, false, "", "")
	}
	for _, test := range goldenJSON {
		runGoldenTest(t, test, true, false, false, false, false, "", "")
	}
	for _, test := range goldenText {
		runGoldenTest(t, test, false, false, false, false, true, "", "")
	}
	for _, test := range goldenYAML {
		runGoldenTest(t, test, false, true, false, false, false, "", "")
	}
	for _, test := range goldenSQL {
		runGoldenTest(t, test, false, false, true, false, false, "", "")
	}
	for _, test := range goldenCQL {
		runGoldenTest(t, test, false, false, false, true, false, "", "")
	}
	for _, test := range goldenJSONAndSQL {
		runGoldenTest(t, test, true, false, true, false, false, "", "")
	}
	for _, test := range goldenTrimPrefix {
		runGoldenTest(t, test, false, false, false, false, false, "Day", "")
	}
	for _, test := range goldenTrimPrefixMultiple {
		runGoldenTest(t, test, false, false, false, false, false, "Day,Night", "")
	}
	for _, test := range goldenWithPrefix {
		runGoldenTest(t, test, false, false, false, false, false, "", "Day")
	}
	for _, test := range goldenTrimAndAddPrefix {
		runGoldenTest(t, test, false, false, false, false, false, "Day", "Night")
	}
}

func runGoldenTest(t *testing.T, test Golden, generateJSON, generateYAML, generateSQL, generateCQL, generateText bool, trimPrefix string, prefix string) {
	var g Generator
	file := test.name + ".go"
	input := "package test\n" + test.input

	dir, err := ioutil.TempDir("", "stringer")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err = os.RemoveAll(dir)
		if err != nil {
			t.Error(err)
		}
	}()

	absFile := filepath.Join(dir, file)
	err = ioutil.WriteFile(absFile, []byte(input), 0644)
	if err != nil {
		t.Error(err)
	}
	g.parsePackage([]string{absFile}, nil)
	// Extract the name and type of the constant from the first line.
	tokens := strings.SplitN(test.input, " ", 3)
	if len(tokens) != 3 {
		t.Fatalf("%s: need type declaration on first line", test.name)
	}
	g.generate(tokens[1], generateJSON, generateYAML, generateSQL, generateCQL, generateText, "noop", trimPrefix, prefix, false)
	got := string(g.format())
	if got != test.output {
		// Use this to help build a golden text when changes are needed
		//goldenFile := fmt.Sprintf("./goldendata/%v.golden", test.name)
		//err = ioutil.WriteFile(goldenFile, []byte(got), 0644)
		//if err != nil {
		//	t.Error(err)
		//}
		t.Errorf("%s: got\n====\n%s====\nexpected\n====%s", test.name, got, test.output)
	}
}
