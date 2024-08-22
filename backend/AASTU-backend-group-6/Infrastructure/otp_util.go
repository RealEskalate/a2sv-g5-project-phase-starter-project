package infrastructure

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

var tokenStore = make(map[string]TokenData)

// TokenData holds the token information
type TokenData struct {
	Email     string
	OTP       string
	ExpiresAt time.Time
}

// Email and server configuration
const (
	smtpHost      = "smtp.gmail.com"              // SMTP server for Gmail
	smtpPort      = 465                           // Port for SMTPS (SSL/TLS)
	emailFrom     = "yordanoslegesse15@gmail.com" // Your Gmail address
	emailPassword = "bcewmdllhervddxu"            // Your app-specific password
	tokenTTL      = time.Minute * 15              // Token Time-To-Live (5 minutes)
)

// Generates a secure random OTP
func GenerateOTP() (string, error) {
	const length = 6
	otp := ""

	for i := 0; i < length; i++ {
		// Generate a random number between 0 and 9
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}

		// Append the digit to the OTP string
		otp += fmt.Sprintf("%d", n.Int64())
	}

	return otp, nil

}

// Sends the OTP email

var  SendOTPEmail = func(to, otp string) error {
	body := fmt.Sprintf(`
		Hi,

		Your OTP code is: %s

		This code will expire in 10 minutes. If you did not request this, please ignore this email.
	`, otp)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", "Eskalate G5 Blog Project", emailFrom))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(smtpHost, smtpPort, emailFrom, emailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// Handler for verifying the OTP
func verifyOTP(c *gin.Context) {
	fmt.Println("Handling OTP verification")

	var request struct {
		Email string `json:"email" binding:"required,email"`
		OTP   string `json:"otp" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the OTP
	data, exists := tokenStore[request.Email]
	if !exists || time.Now().After(data.ExpiresAt) || data.OTP != request.OTP {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	// Invalidate the OTP
	delete(tokenStore, request.Email)

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}

func main() {
	r := gin.Default()

	// Debug log
	fmt.Println("Starting server with routes /request and /verify")

	// Routes
	// r.POST("/request", requestOTP)
	r.POST("/verify", verifyOTP)

	// Start the server
	r.Run(":8080") // Starts the Gin server on port 8080
}
