# Processes
## Login
* User enters Email
* Clicks the login button
  * Post request send to backend with email in payload
* Gets redirected to a "waiting" page
* Backend sends email to the given Email-Address
  * Contains a link with the email and one time password in the URL-Parameters
  * Clicking the link will:
    * Create a new Session
    * Set the session as a cookie
    * Invalidate/Delete the one time password on the backend
    * Redirect user to homepage
* On the homepage:
  * The App will load the public VAPID key for notifications
  * Ask the User to enable Notifications
    * Accepted:
      * Sends POST request to the Backend with the subscription details and sessionID to authenticate
    * Rejected:
      * User is informed that they will not receive any automatic Articles but can still manualy load them