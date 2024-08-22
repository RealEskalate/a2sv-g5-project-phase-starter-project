package emailutil

import "fmt"

func Emailtemplate(url string) string {
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
                .content a {
                    display: inline-block;
                    margin-top: 20px;
                    padding: 10px 20px;
                    color: white;
                    background-color: #4CAF50;
                    text-decoration: none;
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
                    <h1>Verify Your Account</h1>
                </div>
                <div class="content">
                    <p>Please click the button below to verify your account:</p>
                    <a href='%v'>Verify Email</a>
                    <p>Thank you for registering!</p>
                </div>
                <div class="footer">
                    <p>&copy; 2024 Your Company. All rights reserved.</p>
                </div>
            </div>
        </body>
    </html>`,
		url,
	)
}
