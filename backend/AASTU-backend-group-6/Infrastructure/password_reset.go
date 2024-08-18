package infrastructure

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

// In-memory storage for tokens (for demonstration purposes; use a database in production)
// var tokenStore = make(map[string]TokenData)

// TokenData holds the token information
// type TokenData struct {
//     Email     string
//     ExpiresAt time.Time
// }

// Email and server configuration
const (
    SmtpHost      = "smtp.gmail.com"          // Correct SMTP server for Gmail
    SmtpPort      = 465                       // Port for SMTPS (SSL/TLS)
    EmailFrom     = "yordanoslegesse15@gmail.com" // Your Gmail address
    EmailPassword = "bcewmdllhervddxu"        // Your app-specific password
    ServerHost    = "http://localhost:8080"   // Change to your domain in production
    TokenTTlL      = time.Hour                 // Token Time-To-Live
)

// Generates a secure random token
func GenerateResetToken() (string, error) {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

// Sends the password reset email
func SendResetEmail(to, token string) error {
    resetLink := fmt.Sprintf("%s/reset-password?token=%s", ServerHost, token)
    body := fmt.Sprintf(`
        Hi,

        You requested a password reset. Click the link below to reset your password:

        %s

        If you did not request this, please ignore this email.
    `, resetLink)

    m := gomail.NewMessage()
    m.SetHeader("From", EmailFrom)
    m.SetHeader("To", to)
    m.SetHeader("Subject", "Password Reset")
    m.SetBody("text/plain", body)

    d := gomail.NewDialer(SmtpHost, SmtpPort, EmailFrom, EmailPassword)

    if err := d.DialAndSend(m); err != nil {
        return err
    }

    return nil
}

// Handler for requesting a password reset
// func requestPasswordReset(c *gin.Context) {
//     var request struct {
//         Email string `json:"email" binding:"required,email"`
//     }

//     if err := c.ShouldBindJSON(&request); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//         return
//     }

//     // Generate a reset token
//     token, err := generateToken()
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate reset token"})
//         return
//     }

//     // Store the token with expiration
//     tokenStore[token] = TokenData{
//         Email:     request.Email,
//         ExpiresAt: time.Now().Add(TokenTTlL),
//     }

//     // Send the reset email
//     if err := sendResetEmail(request.Email, token); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to send reset email"})
//         return
//     }

//     c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent"})
// }

// Handler for resetting the password
// func resetPassword(c *gin.Context) {
//     var request struct {
//         Token       string `json:"token" binding:"required"`
//         NewPassword string `json:"new_password" binding:"required,min=6"`
//     }

//     if err := c.ShouldBindJSON(&request); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//         return
//     }

//     // Validate the token
//     data, exists := tokenStore[request.Token]
//     if !exists || time.Now().After(data.ExpiresAt) {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired token"})
//         return
//     }

//     // Here, update the user's password in the database (not implemented in this example)

//     // Invalidate the token
//     delete(tokenStore, request.Token)

//     c.JSON(http.StatusOK, gin.H{"message": "Password has been reset successfully"})
// }

