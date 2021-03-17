# Epikins

This projects aims to help Epitech pedagogical team and students to handle automated testing in a better way.  
It's using Jenkins JSON api to fetch different information.

## Currently supported features:
* Microsoft office authentication
* Start automated tests (ATs) as a member of the pedagogical team
* Start ATs as a student
* Full role management
* Change ATs limit dynamically for a given project as a module manager
* View build log to know precisely who started a build, to who and when

# Deployement
## Azure app registration
1. Go to this link : https://docs.microsoft.com/en-us/azure/active-directory/develop/scenario-spa-app-registration#redirect-uri-msaljs-20-with-auth-code-flow
2. Follow the instructions under **Create the app registration** 
3. After that, follow the instructions under **Redirect URI: MSAL.js 2.0 with auth code flow**
4. When you need to enter a *Redirect URI* enter something that match the following pattern: `http(s)://{domain where your app will be available}/login`
5. Once you've finished, go to the **Manifest** tab, you'll find it in the panel on the left
6. To force the app to give you a 2.0 access token, change the line `"accessTokenAcceptedVersion": null` to `"accessTokenAcceptedVersion": 2`
7. Follow the instructions you'll find here: https://docs.microsoft.com/fr-fr/azure/active-directory/develop/quickstart-configure-app-expose-web-apis
8. Go to the **API permissions** tab, you'll find it in the panel on the left
9. Select **Add a permission**, go to the **My APIs** tab and add the one you created previously
10. Finally, go to the **Token configuration** tab, you'll find it in the panel on the left
11. Select **Add optionnal claims**, choose **access** type, check **email** field, select **add**, check **Turn on the Microsoft Graph email permission** field and select add.

You're app is now fully configured on Azure, congrats!

## Docker configuration
First, make sure **Docker** is installed on the machine targeted by the deployment.  
Before typing the `docker-compose up` magic command, you must fill both `.env` files you'll find in `server` and `client` directories.  

After that, you must fill in the content of `init_mongo.js` file. It'll allow you to create a first user which has access to the entire application. You must also enter a first Jenkins username and its corresponding API Key.  

After that, you can enter `docker-compose up` in your command line prompt and you're app is deployed locally!  
In order that your app can be reached from the outside world, you'll need to configure an internal redirection to localhost:8081 for the client and localhost:8082 for the server.
