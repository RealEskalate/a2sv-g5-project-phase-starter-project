package mail_service

import (
	"blog_api/delivery/env"
	"fmt"
)

func EmailVerificationTemplate(hostUrl string, username string, token string) string {
	link := hostUrl + "/api/" + env.ENV.ROUTE_PREFIX + "/auth/verify/email/" + username + "/" + token
	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Email Verification</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f4f4f4;
					color: #333;
					line-height: 1.6;
				}
				.container {
					max-width: 600px;
					margin: 0 auto;
					padding: 20px;
					background-color: #fff;
					border-radius: 8px;
					box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
				}
				.header {
					text-align: center;
					padding-bottom: 20px;
				}
				.header h1 {
					margin: 0;
					color: #333;
				}
				.content {
					padding: 20px;
				}
				.button {
					display: inline-block;
					padding: 10px 20px;
					margin-top: 20px;
					background-color: #007bff;
					color: #fff !important;
					font-weight: bold;
					text-decoration: none;
					border-radius: 5px;
				}
				.button:hover {
					background-color: #0056b3;
				}
				.footer {
					text-align: center;
					padding-top: 20px;
					font-size: 0.9em;
					color: #777;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Blog API</h1>
				</div>
				<div class="content">
					<p>Dear User,</p>
					<p>Thank you for registering with Blog API. Please verify your email address by clicking the button below:</p>
					<a href="%s" class="button">Verify Email</a>
					<p>If you did not create an account, no further action is required.</p>
				</div>
				<div class="footer">
					<p>&copy; 2023 Blog API. All rights reserved.</p>
				</div>
			</div>
		</body>
		</html>`, link)
}
