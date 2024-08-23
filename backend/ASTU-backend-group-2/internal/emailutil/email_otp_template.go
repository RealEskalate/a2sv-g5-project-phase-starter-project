package emailutil

import (
	"fmt"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
)

// OTPEmailTemplate generates an HTML email template with the provided OTP for user verification.
func OTPEmailTemplate(otp string,env *bootstrap.Env) string {
	return fmt.Sprintf(
		`<html>
        <head>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f4f4f4;
                    color: #333333;
                    margin: 0;
                    padding: 0;
                }
                .container {
                    width: 100%%;
                    max-width: 600px;
                    margin: 0 auto;
                    background-color: #ffffff;
                    padding: 20px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                }
                .header {
                    text-align: center;
                    padding: 10px 0;
                    background-color: #4CAF50;
                    color: white;
                }
                .content {
                    padding: 20px;
                    text-align: center;
                }
                .content p {
                    font-size: 16px;
                    line-height: 1.5;
                }
                .otp-code {
                    display: inline-block;
                    font-size: 24px;
                    font-weight: bold;
                    margin-top: 20px;
                    padding: 10px;
                    background-color: #f0f0f0;
                    border-radius: 5px;
                }
                .footer {
                    text-align: center;
                    padding: 10px 0;
                    font-size: 12px;
                    color: #999999;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <div class="header">
                    <h1>Reset Your Password</h1>
                </div>
                <div class="content">
                    <p>To reset your password, please use the following OTP:</p>
                    <div class="otp-code">%s</div>
                    <p>This OTP is valid for a %v minutes. Please do not share it with anyone.</p>
                    <p>Thank you!</p>
                </div>
                <div class="footer">
                    <p>&copy; 2024 Your Company. All rights reserved.</p>
                </div>
            </div>
        </body>
    </html>`,
		otp,
        env.PassResetCodeExpirationMin,

	)
}
