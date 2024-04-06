<script>
        export let data
        import {_sendQuestion} from './+page.js'
</script>

<p>Your question to <b>{data.otheruser.username}:</b></p>
<input id="input" style="margin:50px;  width: 1000px;height: 50px; font-size:large;">
<button on:click={_sendQuestion(data.user.username, document.getElementById("input").value, data.otheruser.username)}>Send</button>