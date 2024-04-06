    <script>
            import { goto } from '$app/navigation'
    import { basePathFront} from '$lib/index.js'
        export let data 
    </script>

<p>Hi! Here you can see questions that other users asked you.
    <button on:click={()=> location.reload()}>Reload ↻</button>
</p>

    <ul>
        {#each data.user.Questions as question} 
        <div>
            <h2>
                <span class="question">{question.thequestion}</span>
                <span style="font-size: small;">asked by {question.questioner}</span>
                <button on:click={() => goto(location.href + "/answer/" + question.ID)}>Answer</button>
            </h2>
            <p style="width:750px;">{question.answer}</p>
        </div>
        {/each}
    </ul>