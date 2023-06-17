package main

// App ...
type App struct {
	Config   *Config
	Inputs   *Inputs
	Filtered *Filtered
	Outputs  *Outputs
}

// Config ...
type Config struct {
	SaveFilteredToCSV bool      `koanf:"save_filtered_to_csv"`
	ColumnsOrder      []string  `koanf:"columns_order"`
	CSVSeparator      string    `koanf:"csv_separator"`
	API               *API      `koanf:"api"`
	FilterColumns     []*Column `koanf:"filter_columns"`
	UpdateFields      []*Field  `koanf:"update_fields"`
}

// API ...
type API struct {
	RequestTimeout int       `koanf:"request_timeout"`
	BaseURL        string    `koanf:"base_url"`
	BaseURLSchema  string    `koanf:"base_url_schema"`
	AuthMethod     string    `koanf:"auth_method"`
	Token          string    `koanf:"token"`
	UpdateEndpoint *Endpoint `koanf:"update_endpoint"`
}

// Column ...
type Column struct {
	ColumnName string `koanf:"column_name"`
	Condition  string `koanf:"condition"`
	Value      string `koanf:"value"`
}

// Field ...
type Field struct {
	FieldName  string    `koanf:"field_name"`
	Values     []string  `koanf:"values"`
	Conditions []*Column `koanf:"conditions"`
}

// Endpoint ...
type Endpoint struct {
	AddPKToEndpointName bool           `koanf:"add_pk_to_endpoint_name"`
	EndpointName        string         `koanf:"endpoint_name"`
	ContentType         string         `koanf:"content_type"`
	HTTPMethod          string         `koanf:"http_method"`
	EndpointBody        map[string]any `koanf:"endpoint_body"`
}

// Inputs ...
type Inputs struct {
	Mapping map[string]int
	Data    [][]string
}

// Filtered ...
type Filtered struct {
	Data [][]string
}

// Outputs ...
type Outputs struct {
	Data []*Output
}

type Output struct {
	ID        string
	FieldName string
	Values    []string
}
