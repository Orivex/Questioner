<script>
    export let data
    import {_sendAnswer} from './+page.js'
</script>

<div style="display: flex;margin:200px; flex-direction:column;">
    <h1 style="color: burlywood;">{data.question.thequestion}</h1>
    <textarea id="input" style="font-size:large" rows="20" cols="100">{data.question.answer}</textarea>
    <button on:click={_sendAnswer(document.getElementById("input").value, data.question.ID)}>Send</button>
</div>