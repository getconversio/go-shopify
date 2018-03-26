package goshopify

import (
	"reflect"
	"testing"
	"time"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func scriptTagTests(t *testing.T, scriptTag ScriptTag) {
	// Check that ID is assigned to the returned scriptTag
	expectedInt := 596726825
	if scriptTag.ID != expectedInt {
		t.Errorf("ScriptTag.ID returned %+v, expected %+v", scriptTag.ID, expectedInt)
	}
}

func TestScriptTagList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/script_tags.json",
		httpmock.NewStringResponder(200, `{"script_tags": [{"id":1},{"id":2}]}`))

	scriptTags, err := client.ScriptTag.List(nil)
	if err != nil {
		t.Errorf("ScriptTag.List returned error: %v", err)
	}

	expected := []ScriptTag{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(scriptTags, expected) {
		t.Errorf("ScriptTag.List returned %+v, expected %+v", scriptTags, expected)
	}
}

func TestScriptTagCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/script_tags/count.json",
		httpmock.NewStringResponder(200, `{"count": 3}`))

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/script_tags/count.json?created_at_min=2016-01-01T00%3A00%3A00Z",
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.ScriptTag.Count(nil)
	if err != nil {
		t.Errorf("ScriptTag.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("ScriptTag.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.ScriptTag.Count(CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("ScriptTag.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("ScriptTag.Count returned %d, expected %d", cnt, expected)
	}
}

func TestScriptTagGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/script_tags/1.json",
		httpmock.NewStringResponder(200, `{"script_tag": {"id":1}}`))

	scriptTag, err := client.ScriptTag.Get(1, nil)
	if err != nil {
		t.Errorf("ScriptTag.Get returned error: %v", err)
	}

	expected := &ScriptTag{ID: 1}
	if !reflect.DeepEqual(scriptTag, expected) {
		t.Errorf("ScriptTag.Get returned %+v, expected %+v", scriptTag, expected)
	}
}

func TestScriptTagCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/script_tags.json",
		httpmock.NewBytesResponder(200, loadFixture("scripttag.json")))

	scriptTag := ScriptTag{
		Src: "http://myserver.com/test.js",
	}

	returnedScriptTag, err := client.ScriptTag.Create(scriptTag)
	if err != nil {
		t.Errorf("ScriptTag.Create returned error: %v", err)
	}

	scriptTagTests(t, *returnedScriptTag)
}

func TestScriptTagUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "https://fooshop.myshopify.com/admin/script_tags/1.json",
		httpmock.NewBytesResponder(200, loadFixture("scripttag.json")))

	scriptTag := ScriptTag{
		ID:          1,
		Src: "http://myserver.com/test.js",
	}

	returnedScriptTag, err := client.ScriptTag.Update(scriptTag)
	if err != nil {
		t.Errorf("ScriptTag.Update returned error: %v", err)
	}

	scriptTagTests(t, *returnedScriptTag)
}

func TestScriptTagDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "https://fooshop.myshopify.com/admin/script_tags/1.json",
		httpmock.NewStringResponder(200, "{}"))

	err := client.ScriptTag.Delete(1)
	if err != nil {
		t.Errorf("ScriptTag.Delete returned error: %v", err)
	}
}
