package crud_functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	structs "github.com/Parz1val02/cloud-cli/structs"
)

func DeleteTemplate(templateId, token string) error {
	serverPort := 4444
	requestURL := fmt.Sprintf("http://10.20.12.162:%d/templateservice/templates/%s", serverPort, templateId)
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", requestURL, nil)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	req.Header.Set("X-API-Key", token)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	var jsonresp structs.NormalResponse
	err = json.Unmarshal(body, &jsonresp)
	if err != nil {
		return fmt.Errorf("Error decoding response body: %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d, Error: %s\n", resp.StatusCode, jsonresp.Msg)
	}
	fmt.Printf("%s\n", jsonresp.Msg)
	return nil
}
