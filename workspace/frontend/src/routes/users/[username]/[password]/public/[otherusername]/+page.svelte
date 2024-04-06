    <script>
            import { goto } from '$app/navigation'
            import { basePathFront} from '$lib/index.js'
        export let data 
    </script>

<h1 style="color: burlywood; margin:60px">{data.otheruser.username}
    <button on:click={() => goto(`${basePathFront}/users/${data.user.username}/${data.user.password}/public/${data.otheruser.username}/ask`)} >Ask <b>{data.otheruser.username}</b> something</button>
    <button on:click={()=> location.reload()}>Reload ↻</button>
</h1>
<ul>
    {#each data.otheruser.Questions as question} 
    <div>
        <h3>
            <span class="question">{question.thequestion}</span>
            <span style="font-size: small;">asked by {question.questioner}</span>
        </h3>
        <p>{question.answer}</p>
    </div>
    {/each}
</ul>