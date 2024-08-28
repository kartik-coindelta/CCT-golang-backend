package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func SendLoginOTP(email string) (map[string]interface{}, error) {
	verificationCode := rand.Intn(9000) + 1000

	apiUrl := os.Getenv("SES_URL")
	htmlData := fmt.Sprintf(`
        <table bgcolor="#F2F2F2" border="0" cellpadding="0" cellspacing="0" width="100%%">
        <tbody>
          <tr>
            <td>
              <div style='margin:50px;font-size:16px;line-height:24px'>
                <table border="0" cellpadding="0" cellspacing="0" width="100%%">
                  <tbody>
                    <tr>
                      <td>
                        <table border="0" cellpadding="0" cellspacing="0" width="100%%">
                          <tbody>
                            <tr style="">
                              <td style="background-color:white;padding-top:20px;padding-bottom:30px">
                                <table border="0" cellpadding="0" cellspacing="0" width="100%%">
                                  <tbody>
                                    <tr style="display:flex;justify-content:start;align-items:center;padding-bottom:20px;">
                                      <td align="left" style="padding-top:0;padding-left:10px">
                                        <a href="https://uat.coincircletrust.com/" rel="noreferrer" target="_blank">
                                          <img src="%s" alt="CoinCirlceTrust" width="376" height="73" style="vertical-align:middle" class="CToWUd" data-bit="iit">
                                        </a>
                                      </td>
                                    </tr>
                                    <tr>
                                      <td style="font-family:Helvetica,Arial,sans-serif;font-size:16px;line-height:24px;word-break:break-word;padding-left:20px;padding-right:20px;padding-top:20px;margin-left:px;margin-right:px">
                                        Use the following OTP to login into your %s account %s.
                                      </td>
                                    </tr>
                                    <tr>
                                      <td style="font-family:Helvetica,Arial,sans-serif;font-size:16px;line-height:24px;word-break:break-word;padding-left:20px;padding-right:20px;padding-top:20px;margin-left:px;margin-right:px">
                                        <b>%d</b>
                                        <br /><br />
                                        If you don't know why you received this, you are not required to take any action. Please do not share this OTP with anyone.
                                      </td>
                                    </tr>
                                    <tr>
                                      <td style="font-family:Helvetica,Arial,sans-serif;font-size:16px;line-height:24px;word-break:break-word;padding-left:20px;padding-right:20px;padding-top:30px;margin-left:px;margin-right:px">
                                        Thanks,<br>%s Team
                                      </td>
                                    </tr>
                                  </tbody>
                                </table>
                              </td>
                            </tr>
                          </tbody>
                        </table>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </td>
          </tr>
        </tbody>
      </table>`, os.Getenv("email_banner_image"), os.Getenv("company_title"), email, verificationCode, os.Getenv("company_title"))

	body := map[string]interface{}{
		"verificationCode": verificationCode,
		"email":            email,
		"subject":          "Login OTP",
		"htmlData":         htmlData,
	}

	jsonBody, _ := json.Marshal(body)
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("Error sending email:", err)
		return map[string]interface{}{
			"message":          "Internal server error",
			"verificationCode": verificationCode,
		}, err
	}
	defer resp.Body.Close()

	// Log and handle empty response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return map[string]interface{}{
			"message":          "Internal server error",
			"verificationCode": verificationCode,
		}, err
	}
	log.Println("Response body:", string(bodyBytes))

	var result map[string]interface{}
	if len(bodyBytes) == 0 {
		result = map[string]interface{}{
			"verificationCode": verificationCode,
			"message":          "No response data",
		}
	} else if err := json.Unmarshal(bodyBytes, &result); err != nil {
		log.Println("Error decoding response:", err)
		return map[string]interface{}{
			"message":          "Internal server error",
			"verificationCode": verificationCode,
		}, err
	} else {
		if resp.StatusCode == http.StatusOK {
			result["verificationCode"] = verificationCode
			result["verificationCodeTimestamp"] = time.Now().Unix()
		} else {
			result["verificationCode"] = verificationCode
			result["message"] = "Failed to create the invite"
		}
	}

	return result, nil
}

func SendInvitation(email, url, companyName string) error {
	apiUrl := os.Getenv("SES_URL")
	htmlData := fmt.Sprintf(`
        <table bgcolor="#F2F2F2" border="0" cellpadding="0" cellspacing="0" width="100%%">
          <tbody>
            <tr>
              <td>
                <div style='margin:50px;font-size:16px;line-height:24px'>
                  <table border="0" cellpadding="0" cellspacing="0" width="100%%">
                    <tbody>
                      <tr>
                        <td>
                          <table border="0" cellpadding="0" cellspacing="0" width="100%%">
                            <tbody>
                              <tr style="">
                                <td style="background-color:white;padding-top:20px;padding-bottom:30px">
                                  <table border="0" cellpadding="0" cellspacing="0" width="100%%">
                                    <tbody>
                                      <tr style="display:flex;justify-content:start;align-items:center;padding-bottom:20px;">
                                        <td align="left" style="padding-top:0;padding-left:10px">
                                          <a href="https://uat.coincircletrust.com/" rel="noreferrer" target="_blank">
                                            <img src="%s" alt="CoinCirlceTrust" width="376" height="73" style="vertical-align:middle" class="CToWUd" data-bit="iit">
                                          </a>
                                        </td>
                                      </tr>
                                      <tr>
                                        <td style="font-family:Helvetica,Arial,sans-serif;font-size:16px;line-height:24px;word-break:break-word;padding-left:20px;padding-right:20px;padding-top:20px;margin-left:px;margin-right:px">
                                          Hi <b>%s</b>, We hope this email finds you well. We are writing to request your assistance in completing your verification process by filling in some required details through our secure online portal.
                                        </td>
                                      </tr>
                                      <tr>
                                        <td style="font-family:Helvetica,Arial,sans-serif;font-size:16px;line-height:24px;word-break:break-word;padding-left:20px;padding-right:20px;padding-top:20px;margin-left:px;margin-right:px">
                                          To complete the process, please follow the link below:<br /><br />
                                          <a href="%s" rel="noreferrer" target="_blank">Link to fill details</a>
                                        </td>
                                      </tr>
                                      <tr>
                                        <td style="font-family:Helvetica,Arial,sans-serif;font-size:16px;line-height:24px;word-break:break-word;padding-left:20px;padding-right:20px;padding-top:30px;margin-left:px;margin-right:px">
                                          Thanks,<br>%s Team
                                        </td>
                                      </tr>
                                    </tbody>
                                  </table>
                                </td>
                              </tr>
                            </tbody>
                          </table>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </td>
            </tr>
          </tbody>
        </table>`, os.Getenv("email_banner_image"), email, url, companyName)

	body := map[string]interface{}{
		"email":    email,
		"subject":  "Invitation to Register",
		"htmlData": htmlData,
	}

	log.Printf("emails: %s\n", email)
	log.Printf("companyNames: %s", companyName)

	jsonBody, _ := json.Marshal(body)
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("Error sending invite:", err)
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return err
	}

	log.Println("Response body:", string(bodyBytes))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send invite, status code: %d", resp.StatusCode)
	}

	return nil
}
