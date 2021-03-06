package template_test

import (
	"drawbridge/pkg/config/template"
	"drawbridge/pkg/utils"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestConfigTemplate_DeleteTemplate(t *testing.T) {
	t.Parallel()

	//setup
	parentPath, err := ioutil.TempDir("", "")
	defer os.RemoveAll(parentPath)

	testFilePath := path.Join(parentPath, "1.text")
	err = utils.FileWrite(testFilePath, "test content", 0666, false)
	require.NoError(t, err, "should not raise an error when writing test file.")

	fileTemplate := template.ConfigTemplate{
		FileTemplate: template.FileTemplate{
			FilePath: "{{.example}}.text",
			Template: template.Template{
				Content: "",
			},
		},
	}

	//test
	err = fileTemplate.DeleteTemplate(map[string]interface{}{
		"example":    "1",
		"config_dir": parentPath,
	})

	//assert
	require.NoError(t, err, "should not raise an error deleting filepath template")
	require.False(t, utils.FileExists(testFilePath), "test file should not be exist")
}

func TestConfigTemplate_WriteTemplate(t *testing.T) {
	t.Parallel()

	//setup
	parentPath, err := ioutil.TempDir("", "")
	defer os.RemoveAll(parentPath)

	testFilePath := path.Join(parentPath, "1.text")

	fileTemplate := template.ConfigTemplate{
		PemFilePath: "{{.example}}.pem",
		FileTemplate: template.FileTemplate{
			FilePath: "{{.example}}.text",
			Template: template.Template{
				Content: "config content",
			},
		},
	}

	//test
	actual, err := fileTemplate.WriteTemplate(map[string]interface{}{
		"example":    "1",
		"config_dir": parentPath,
		"pem_dir":    parentPath,
	}, []string{}, false)

	//assert
	require.NoError(t, err, "should not raise an error deleting filepath template")
	require.FileExists(t, actual["filepath"].(string), "test file should written ")
	require.Equal(t, testFilePath, actual["filepath"].(string), "test file path be set correctly")
}

//
//func TestFileTemplate_WriteTemplate(t *testing.T) {
//	t.Parallel()
//
//	//setup
//	parentPath, err := ioutil.TempDir("", "")
//	defer os.RemoveAll(parentPath)
//
//	testFilePathTemplate := path.Join(parentPath, "{{.example}}/{{.example2}}.text")
//	testFilePath := path.Join(parentPath, "1/2.text")
//
//	fileTemplate := template.FileTemplate{
//		FilePath: testFilePathTemplate,
//		Template: template.Template{
//			Content: "{{.content}}",
//		},
//	}
//
//	//test
//	actual, err := fileTemplate.WriteTemplate(map[string]interface{}{
//		"example": "1",
//		"example2": "2",
//		"content": "this is my content",
//	}, false)
//
//	//assert
//	require.NoError(t, err,"should not raise an error deleting filepath template")
//	require.FileExists(t, testFilePath, "should write file to correct path")
//	require.Equal(t, map[string]interface{}{"filepath": testFilePath}, actual, "should return some metadata about the template")
//}
//
//func TestFileTemplate_WriteTemplate_WhenDestinationExists(t *testing.T) {
//	t.Parallel()
//
//	//setup
//	parentPath, err := ioutil.TempDir("", "")
//	defer os.RemoveAll(parentPath)
//
//	testFilePathTemplate := path.Join(parentPath, "{{.example}}.text")
//	testFilePath := path.Join(parentPath, "1.text")
//	err = utils.FileWrite(testFilePath, "previous content", 0666, false)
//	require.NoError(t, err, "should not raise an error when writing test file.")
//
//	fileTemplate := template.FileTemplate{
//		FilePath: testFilePathTemplate,
//		Template: template.Template{
//			Content: "{{.content}}",
//		},
//	}
//
//	//test
//	_, err = fileTemplate.WriteTemplate(map[string]interface{}{
//		"example": "1",
//		"content": "this is my content",
//	}, false)
//
//	//assert
//	require.Error(t, err,"should raise an error if destination file already exists.")
//}
