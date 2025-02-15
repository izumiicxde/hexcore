package mail

import (
	"fmt"
	"hexcore/config"
)

func GenerateVerificationEmail(username, code string) string {
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Email Verification</title>
		<style>
			body { font-family: Arial, sans-serif; background-color: #f4f4f4; margin: 0; padding: 20px; }
			.container { max-width: 480px; background: white; padding: 25px; border-radius: 8px; 
				box-shadow: 0 4px 10px rgba(0,0,0,0.1); margin: auto; text-align: center; }
			h2 { color: #333; }
			p { color: #555; font-size: 16px; }
			.code { font-size: 24px; font-weight: bold; color: #007BFF; background: #f0f0f0; 
				padding: 10px 20px; display: inline-block; border-radius: 5px; margin-top: 10px; }
			.footer { font-size: 12px; color: #777; margin-top: 20px; }
			.button { background-color: #007BFF; color: white; padding: 12px 20px; 
				text-decoration: none; font-size: 16px; border-radius: 5px; display: inline-block; margin-top: 15px; }
			.button:hover { background-color: #0056b3; color:white; }
		</style>
	</head>
	<body>
		<div class="container">
			<h2>Email Verification</h2>
			<p>Hello <strong>%s</strong>,</p>
			<p>Use the verification code below to verify your email:</p>
			<p class="code">%s</p>
			<p>Or click the button below to verify:</p>
			<a href="%s/verify?code=%s" class="button">Verify Email</a>
			<p class="footer">If you didn't request this, please ignore this email.</p>
		</div>
	</body>
	</html>
	`, username, code, config.Envs.API_ENDPOINT, code)
}
