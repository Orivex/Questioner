<script>
    import { goto } from '$app/navigation'
    import { basePathFront} from '$lib/index.js'
    export let data
</script>


<div class="container">
    <button onclick="history.back()">🡸</button>
    <button on:click={() => goto(`${basePathFront}/login`)}>Logout</button>
    <button on:click={() => goto(`${basePathFront}/users/${data.user.username}/${data.user.password}`)}>My page</button>
    <button on:click={() => goto(`${basePathFront}/users/${data.user.username}/${data.user.password}/public`)}>Ask people</button>
    <span style="font-size:x-large;color:brown">
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <b>Logged in as:  {data.user.username} </b>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <b>Your password: {data.user.password}  </b>
    </span>
</div>

<slot/>