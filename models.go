package main

// App represents struct for the application instance.
type App struct {
	Config   *Config
	Inputs   *Inputs
	Filtered *Filtered
	Outputs  *Outputs
}

// Config represents struct for the configuration instance.
type Config struct {
	SaveFilteredPKToCSV bool      `koanf:"save_filtered_pk_to_csv"`
	ColumnsOrder        []string  `koanf:"columns_order"`
	CSVColumnSeparator  string    `koanf:"csv_column_separator"`
	API                 *API      `koanf:"api"`
	FilterColumns       []*Column `koanf:"filter_columns"`
	UpdateFields        []*Field  `koanf:"update_fields"`
}

// API represents struct for the api instance.
type API struct {
	RequestTimeout int       `koanf:"request_timeout"`
	BaseURL        string    `koanf:"base_url"`
	BaseURLSchema  string    `koanf:"base_url_schema"`
	AuthMethod     string    `koanf:"auth_method"`
	Token          string    `koanf:"token"`
	UpdateEndpoint *Endpoint `koanf:"update_endpoint"`
}

// Column represents struct for the one column instance..
type Column struct {
	ColumnName string `koanf:"column_name"`
	Condition  string `koanf:"condition"`
	Value      string `koanf:"value"`
}

// Field represents struct for the one field instance.
type Field struct {
	FieldName  string    `koanf:"field_name"`
	Values     []string  `koanf:"values"`
	Conditions []*Column `koanf:"conditions"`
}

// Endpoint represents struct for the one endpoint instance.
type Endpoint struct {
	AddPKToEndpointName bool           `koanf:"add_pk_to_endpoint_name"`
	EndpointName        string         `koanf:"endpoint_name"`
	ContentType         string         `koanf:"content_type"`
	HTTPMethod          string         `koanf:"http_method"`
	EndpointBody        map[string]any `koanf:"endpoint_body"`
}

// Inputs represents struct for the inputs instance.
type Inputs struct {
	Mapping map[string]int
	Data    [][]string
}

// Filtered represents struct for the filtered instance.
type Filtered struct {
	Data [][]string
}

// Outputs represents struct for the outputs instance.
type Outputs struct {
	Data []*Output
}

// Output represents struct for the one output instance.
type Output struct {
	ID        string
	FieldName string
	Values    []string
}
