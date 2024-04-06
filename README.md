An app created, using Golang (with mux and gorm) as the backend API, Svelte-Kit (consists of HTML and CSS) as the frontend and MySQL as the database.


Before using the app, please create a new database in MySQL called 'questionerdb'. The user and the password should both be 'root'. (I suggest using MySQL-Workbench)

Launch backend:
Change directory (cd) to 'Questioner\workspace\backend' and execute the command in the terminal 'go run .'

Launch frontend:
Change directory (cd) to 'Questioner\workspace\frontend' and execute the command in the terminal 'npm run dev'

The packages for those commands have to be installed before using them.

Visit the website on: 'http://localhost:5173'

Caution: 
The app is insecure because no security was implemented (very loose login: Credentials are shown in the URL). But this shouldn't be a problem, as long as the servers run locally.
One day I started this app, but didn't finish it, because I wanted to stick more with SpringBoot/Java. In two days I finished it quickly for fun. So, don't expect qualitative code or a crazy-innovative app! :)
