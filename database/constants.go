package database

// Translation field to identify fields as subjects
// in other tables like translations.
const (
	DBTranslationFieldName        = 1
	DBTranslationFieldDescription = 2
)

// PropertyTypes
const (
	// PropertytypeText for plain text (string) properties.
	PropertytypeText uint8 = iota + 1
	// PropertytypeRichtext for rich text (string) properties.
	PropertytypeRichtext
	// PropertytypeDatetime for datetime properties.
	PropertytypeDatetime
	// PropertytypeNumeric for numeric (int) properties.
	PropertytypeNumeric
	// PropertytypeDecimal for decimal (float/real) properties.
	PropertytypeDecimal
	// PropertytypeCurrency for currency properties.
	PropertytypeCurrency
	// PropertytypeTruth for truth (boolean) properties.
	PropertytypeTruth
	// PropertytypeSingleselect for single select properties.
	PropertytypeSingleselect
	// PropertytypeMultiselect for multi select properties.
	PropertytypeMultiselect
	// PropertytypeSinglefile for single file properties.
	PropertytypeSinglefile
	// PropertytypeMultifile for multi file properties.
	PropertytypeMultifile
	// PropertytypeSinglelinkeditem for single linked item properties.
	PropertytypeSinglelinkeditem
	// PropertytypeMultiplelinkeditem for multiple linked item properties.
	PropertytypeMultiplelinkeditem
)
