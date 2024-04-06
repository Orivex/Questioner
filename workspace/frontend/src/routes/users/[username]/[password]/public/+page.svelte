<script>
    export let data 

    import { goto } from '$app/navigation'
    import { basePathFront} from '$lib/index.js'
</script>

<p>Here are all users on Questioner. Ask them anything!</p>


{#each data.users as user}      
    <button on:click={() => goto(`${basePathFront}/users/${data.user.username}/${data.user.password}/public/${user.username}`)}>
        {user.username}</button>
{/each}
