package mermaid

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/k1LoW/tbls/config"
	"github.com/k1LoW/tbls/testutil"
	"github.com/tenntenn/golden"
)

func TestOutputSchema(t *testing.T) {
	tests := []struct {
		hideDef  bool
		wantFile string
	}{
		{false, "mermaid_test_schema"},
		{true, "mermaid_test_schema.hidedef"},
	}
	for _, tt := range tests {
		t.Run(tt.wantFile, func(t *testing.T) {
			s := testutil.NewSchema(t)
			c, err := config.New()
			if err != nil {
				t.Error(err)
			}
			if err := c.LoadConfigFile(filepath.Join(testdataDir(), "out_test_tbls.yml")); err != nil {
				t.Error(err)
			}
			if err := c.ModifySchema(s); err != nil {
				t.Error(err)
			}
			c.ER.HideDef = tt.hideDef
			o := New(c)
			got := &bytes.Buffer{}
			if err := o.OutputSchema(got, s); err != nil {
				t.Error(err)
			}
			if os.Getenv("UPDATE_GOLDEN") != "" {
				golden.Update(t, testdataDir(), tt.wantFile, got)
				return
			}
			if diff := golden.Diff(t, testdataDir(), tt.wantFile, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestOutputSchemaTemplate(t *testing.T) {
	s := testutil.NewSchema(t)
	c, err := config.New()
	if err != nil {
		t.Error(err)
	}
	if err := c.LoadConfigFile(filepath.Join(testdataDir(), "out_templates_test_tbls.yml")); err != nil {
		t.Error(err)
	}
	// use the templates in the testdata directory
	c.Templates.Mermaid.Schema = filepath.Join(testdataDir(), c.Templates.Mermaid.Schema)
	if err := c.MergeAdditionalData(s); err != nil {
		t.Error(err)
	}
	o := New(c)
	got := &bytes.Buffer{}
	if err := o.OutputSchema(got, s); err != nil {
		t.Error(err)
	}
	f := fmt.Sprintf("mermaid_template_test_schema")
	if os.Getenv("UPDATE_GOLDEN") != "" {
		golden.Update(t, testdataDir(), f, got)
		return
	}
	if diff := golden.Diff(t, testdataDir(), f, got); diff != "" {
		t.Error(diff)
	}
}

func TestOutputTable(t *testing.T) {
	s := testutil.NewSchema(t)
	c, err := config.New()
	if err != nil {
		t.Error(err)
	}
	if err := c.LoadConfigFile(filepath.Join(testdataDir(), "out_test_tbls.yml")); err != nil {
		t.Error(err)
	}
	if err := c.MergeAdditionalData(s); err != nil {
		t.Error(err)
	}
	ta := s.Tables[0]

	o := New(c)
	got := &bytes.Buffer{}
	if err := o.OutputTable(got, ta); err != nil {
		t.Error(err)
	}
	f := fmt.Sprintf("mermaid_test_a")
	if os.Getenv("UPDATE_GOLDEN") != "" {
		golden.Update(t, testdataDir(), f, got)
		return
	}
	if diff := golden.Diff(t, testdataDir(), f, got); diff != "" {
		t.Error(diff)
	}
}

func TestOutputTableTemplate(t *testing.T) {
	s := testutil.NewSchema(t)
	c, err := config.New()
	if err != nil {
		t.Error(err)
	}
	if err := c.LoadConfigFile(filepath.Join(testdataDir(), "out_templates_test_tbls.yml")); err != nil {
		t.Error(err)
	}
	// use the templates in the testdata directory
	c.Templates.Mermaid.Table = filepath.Join(testdataDir(), c.Templates.Mermaid.Table)
	if err := c.ModifySchema(s); err != nil {
		t.Error(err)
	}

	ta := s.Tables[0]

	o := New(c)
	got := &bytes.Buffer{}
	if err := o.OutputTable(got, ta); err != nil {
		t.Error(err)
	}
	f := fmt.Sprintf("mermaid_template_test_a")
	if os.Getenv("UPDATE_GOLDEN") != "" {
		golden.Update(t, testdataDir(), f, got)
		return
	}
	if diff := golden.Diff(t, testdataDir(), f, got); diff != "" {
		t.Error(diff)
	}
}

func testdataDir() string {
	wd, _ := os.Getwd()
	dir, _ := filepath.Abs(filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata"))
	return dir
}
