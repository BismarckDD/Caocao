package c2asset

import (
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/BismarckDD/Caocao/c2common/c2file/c2txt"
	"github.com/BismarckDD/Caocao/c2common/c2font"
	"github.com/BismarckDD/Caocao/c2common/c2util"

	"github.com/BismarckDD/Caocao/c2common/c2enum"
	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2loader"
	"github.com/BismarckDD/Caocao/c2common/c2loader/asset/types"
)

const (
	defaultCacheEntryWeight = 1
)

const (
	unitAnimationBudget    = 1024 * 1024 * 128 // 128M
	magicAnimationBudget   = 1024 * 1024 * 128 // 128M
	fontBudget             = 128
	paletteBudget          = 64
	paletteTransformBudget = 64
)

const (
	defaultLanguage    = "ENG"
	logPrefix          = "Asset Manager"
	fmtLoadAsset       = "could not load file stream %s (%v)"
	fmtLoadAnimation   = "loading animation %s, draw effect %d"
	fmtLoadComposite   = "loading composite: type %d, token %s, palette %s"
	fmtLoadFont        = "loading font: table %s, sprite %s, palette %s"
	fmtLoadPalette     = "loading palette %s"
	fmtLoadStringTable = "loading string table: %s"
	fmtLoadTransform   = "loading palette transform: %s"
	fmtLoadDict        = "loading data dictionary: %s"
)

// AssetManager loads files and game objects
type AssetManager struct {
	*c2util.Logger
	*c2loader.Loader
	animations       c2interface.Cache
	fonts            c2interface.Cache
	palettes         c2interface.Cache
	transforms       c2interface.Cache
	Records          *c2records.RecordManager
	language         string
	languageModifier int
}

// SetLogLevel sets the log level for the asset manager,  record manager, and file loader
func (am *AssetManager) SetLogLevel(level c2util.LogLevel) {
	am.Logger.SetLevel(level)
	am.Records.Logger.SetLevel(level)
	am.Loader.Logger.SetLevel(level)
}

// LoadAsset loads an asset
func (am *AssetManager) LoadAsset(filePath string) (io.ReadSeeker, error) {
	data, err := am.Loader.Load(filePath)
	if err != nil {
		errStr := fmt.Sprintf(fmtLoadAsset, filePath, err.Error())
		am.Error(errStr)
	}
	return data, err
}

// LoadFileStream streams an MPQ file from a source file path, 这个函数有啥用？
func (am *AssetManager) LoadFileStream(filePath string) (io.ReadSeeker, error) {
	am.Logger.Debugf("Loading FileStream: %s", filePath)
	return am.LoadAsset(filePath)
}

// LoadFile loads an entire file from a source file path as a []byte
func (am *AssetManager) LoadFile(filePath string) ([]byte, error) { // I DO NOT LIKE THIS! - Essial
	fileAsset, err := am.LoadAsset(filePath)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(fileAsset)
	if err != nil {
		return nil, err
	}

	return data, err
}

// FileExists checks if a file exists on the underlying file system at the given file path.
func (am *AssetManager) FileExists(filePath string) (bool, error) {
	filePath = filepath.Clean(filePath)
	am.Logger.Debugf("Checking if file exists %s", filePath)
	return am.Loader.Exists(filePath), nil
}

// LoadLanguage loads language from resource path
func (am *AssetManager) LoadLanguage(languagePath string) string {
	languageByte, err := am.LoadFile(languagePath)
	if err != nil {
		am.Debugf("Unable to load language file: %s", err)
		return defaultLanguage
	}

	languageCode := languageByte[0]
	am.Debugf("Language code: %#02x", languageCode)

	language := c2resource.GetLanguageLiteral(languageCode)
	am.Infof("Language: %s", language)

	am.language = language
	am.languageModifier = c2resource.GetLabelModifier(language)

	return language
}

// Delete all pallete param. here no longer need this.
// LoadAnimation loads an Animation by its resource path and its palette path
func (am *AssetManager) LoadAnimation(animationPath string) (c2interface.Animation, error) {
	return am.LoadAnimationWithEffect(animationPath, c2enum.DrawEffectNone)
}

// LoadAnimationWithEffect loads an Animation by its resource path and its palette path with a given transparency value
func (am *AssetManager) LoadAnimationWithEffect(animationPath string,
	effect c2enum.DrawEffect) (c2interface.Animation, error) {

	cachePath := fmt.Sprintf("%s;%d", animationPath, effect)

	if animation, found := am.animations.Retrieve(cachePath); found {
		return animation.(c2interface.Animation).Clone(), nil
	}

	am.Debugf(fmtLoadAnimation, animationPath, effect)

	var animation c2interface.Animation = c2asset.Animation{}
	err := am.animations.Insert(cachePath, animation, defaultCacheEntryWeight)

	return animation, err
}

// LoadComposite creates a composite object from a ObjectLookupRecord and palettePath describing it
func (am *AssetManager) LoadComposite(baseType c2enum.ObjectType, token, palettePath string) (*Composite, error) {
	am.Debugf(fmtLoadComposite, baseType, token, palettePath)

	c := &Composite{
		AssetManager: am,
		baseType:     baseType,
		basePath:     baseString(baseType),
		token:        token,
		palettePath:  palettePath,
	}

	c.SetDirection(0)

	return c, nil
}

// LoadFont loads a font the resource files
func (am *AssetManager) LoadFont(tablePath, spritePath, palettePath string) (*c2font.Font, error) {
	cachePath := fmt.Sprintf("%s;%s;%s", tablePath, spritePath, palettePath)

	if cached, found := am.fonts.Retrieve(cachePath); found {
		return cached.(*c2font.Font), nil
	}

	sheet, err := am.LoadAnimation(spritePath, palettePath)
	if err != nil {
		return nil, err
	}

	tableData, err := am.LoadFile(tablePath)
	if err != nil {
		return nil, err
	}

	am.Debugf(fmtLoadFont, tablePath, spritePath, palettePath)

	font, err := c2font.Load(tableData)
	if err != nil {
		return nil, fmt.Errorf("error while loading font table %s: %v", tablePath, err)
	}

	font.SetBackground(sheet)

	err = am.fonts.Insert(cachePath, font, defaultCacheEntryWeight)

	return font, err
}

// LoadPalette loads a palette from a given palette path
func (am *AssetManager) LoadPalette(palettePath string) (c2interface.Palette, error) {
	if cached, found := am.palettes.Retrieve(palettePath); found {
		return cached.(c2interface.Palette), nil
	}

	if types.Ext2AssetType(filepath.Ext(palettePath)) != types.AssetTypePalette {
		return nil, fmt.Errorf("not an instance of a palette: %s", palettePath)
	}

	am.Debugf(fmtLoadPalette, palettePath)

	data, err := am.LoadFile(palettePath)
	if err != nil {
		return nil, err
	}

	palette, err := c2dat.Load(data)
	if err != nil {
		return nil, err
	}

	err = am.palettes.Insert(palettePath, palette, defaultCacheEntryWeight)

	return palette, err
}

// LoadStringTable loads a string table from the given path
func (am *AssetManager) LoadStringTable(tablePath string) (c2tbl.TextDictionary, error) {
	data, err := am.LoadFile(tablePath)
	if err != nil {
		return nil, err
	}

	table, err := c2tbl.LoadTextDictionary(data)
	if err != nil {
		return table, err
	}

	am.Debugf(fmtLoadStringTable, tablePath)

	am.tables = append(am.tables, table)

	return table, err
}

// TranslateString returns the translation of the given string. The string is retrieved from
// the loaded string tables. If input value is int (e.g. from c2enum/numeric_labels.go)
// output string is translation for # + input
func (am *AssetManager) TranslateString(input interface{}) string {
	var key string

	switch s := input.(type) {
	case string:
		key = s
	case fmt.Stringer:
		key = s.String()
	case int:
		key = fmt.Sprintf("#%d", c2enum.BaseLabelNumbers(s+am.languageModifier))
	}

	for idx := range am.tables {
		if value, found := am.tables[idx][key]; found {
			return value
		}
	}

	return key
}

// LoadDataDictionary loads a txt data file
func (am *AssetManager) LoadDataDictionary(path string) (*c2txt.DataDictionary, error) {
	// we purposefully do not cache data dictionaries because we are already
	// caching the file data. The underlying csv.Reader does not implement io.Seeker,
	// so after it has been iterated through, we cannot iterate through it again.
	//
	// The easy way around this is to not cache c2txt.DataDictionary objects, and just create
	// a new instance from cached file data if/when we ever need to reload the data dict
	data, err := am.LoadFile(path)
	if err != nil {
		return nil, err
	}

	am.Debugf(fmtLoadDict, path)

	return c2txt.LoadDataDictionary(data), nil
}

// LoadRecords will load the records for the given path into the record manager.
// This is dependant on the record manager having bound a loader for the given path.
func (am *AssetManager) LoadRecords(path string) error {
	dict, err := am.LoadDataDictionary(path)
	if err != nil {
		return err
	}

	err = am.Records.Load(path, dict)
	if err != nil {
		return err
	}

	return nil
}

func (am *AssetManager) commandAssetClear([]string) error {
	am.palettes.Clear()
	am.transforms.Clear()
	am.animations.Clear()
	am.fonts.Clear()
	return nil
}
