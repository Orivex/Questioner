<script>
  import { _registerUser, _loginUser } from './+page.js'
  let username = ''
  let password = ''
</script>

  <div style="display:flex; flex-direction:column;">
    <h1 style="color: burlywood;">Welcome to Questioner!</h1>
    <input placeholder="Username" bind:value={username} />
    <label for="css" style="color: red;"> You don't need a password when registering! It is automatically generated and then given to you</label>
    <br>
    <input placeholder="Password" bind:value={password} />
    <button type="button" on:click={_registerUser(username)}>Register</button>
    <button type="button" on:click={_loginUser(username, password)}>Login</button>
  </div>

<style>
    div {
        max-width: 400px; /* Set the maximum width of the div */
        margin: 0 auto; /* Center the div horizontally */
        padding: 200px; /* Add some padding for better appearance */
        text-align: center; /* Center the text inside the div */
    }

    input, button {
		display: block;
        margin-bottom: 10px; /* Add space between input fields and buttons */
        padding: 10px; /* Add padding to inputs and buttons */
        font-size: 16px; /* Set the font size for inputs and buttons */
    }
</style>