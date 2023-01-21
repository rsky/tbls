package json

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k1LoW/tbls/dict"
	"github.com/k1LoW/tbls/schema"
	"github.com/tenntenn/golden"
)

func TestOutputSchema(t *testing.T) {
	s := newTestSchema(t)
	o := new(JSON)
	buf := &bytes.Buffer{}
	err := o.OutputSchema(buf, s)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	f := "json_output_schema"
	if os.Getenv("UPDATE_GOLDEN") != "" {
		golden.Update(t, testdataDir(), f, got)
		return
	}
	if diff := golden.Diff(t, testdataDir(), f, got); diff != "" {
		t.Error(diff)
	}
}

func TestEncodeAndDecode(t *testing.T) {
	s1 := newTestSchema(t)
	o := new(JSON)
	buf := &bytes.Buffer{}
	err := o.OutputSchema(buf, s1)
	if err != nil {
		t.Error(err)
	}
	s2 := &schema.Schema{}
	dec := json.NewDecoder(buf)
	if err := dec.Decode(s2); err != nil {
		t.Error(err)
	}
	if err := s2.Repair(); err != nil {
		t.Error(err)
	}

	_ = removeColumnRelations(s1)
	_ = removeColumnRelations(s2)

	opt := cmpopts.IgnoreUnexported(dict.New())

	if diff := cmp.Diff(s1, s2, opt); diff != "" {
		t.Errorf("schemas not equal\n%v", diff)
	}
}

func removeColumnRelations(s *schema.Schema) error {
	for _, t := range s.Tables {
		for _, c := range t.Columns {
			c.ParentRelations = nil
			c.ChildRelations = nil
		}
	}
	return nil
}

func testdataDir() string {
	wd, _ := os.Getwd()
	dir, _ := filepath.Abs(filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata"))
	return dir
}

func newTestSchema(t *testing.T) *schema.Schema {
	ca := &schema.Column{
		Name:     "a",
		Type:     "bigint(20)",
		Comment:  "column a",
		Nullable: false,
	}
	cb := &schema.Column{
		Name:     "b",
		Type:     "text",
		Comment:  "column b",
		Nullable: true,
	}

	ta := &schema.Table{
		Name:    "a",
		Type:    "BASE TABLE",
		Comment: "table a",
		Columns: []*schema.Column{
			ca,
			{
				Name:     "a2",
				Type:     "datetime",
				Comment:  "column a2",
				Nullable: false,
				Default: sql.NullString{
					String: "CURRENT_TIMESTAMP",
					Valid:  true,
				},
			},
		},
	}
	tb := &schema.Table{
		Name:    "b",
		Type:    "BASE TABLE",
		Comment: "table b",
		Columns: []*schema.Column{
			cb,
			{
				Name:     "b2",
				Comment:  "column b2",
				Type:     "text",
				Nullable: true,
			},
		},
	}
	r := &schema.Relation{
		Table:         ta,
		Columns:       []*schema.Column{ca},
		ParentTable:   tb,
		ParentColumns: []*schema.Column{cb},
	}
	ca.ParentRelations = []*schema.Relation{r}
	cb.ChildRelations = []*schema.Relation{r}

	s := &schema.Schema{
		Name: "testschema",
		Tables: []*schema.Table{
			ta,
			tb,
		},
		Relations: []*schema.Relation{
			r,
		},
		Functions: []*schema.Function{},
		Driver: &schema.Driver{
			Name:            "testdriver",
			DatabaseVersion: "1.0.0",
			Meta:            &schema.DriverMeta{},
		},
	}
	return s
}
