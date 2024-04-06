import { basePathApi } from '$lib/index.js';
import { goto } from '$app/navigation'

export async function load(loadEvent) {
    const {fetch, params} = loadEvent
    const {question_id} = params
    const response = await fetch(`${basePathApi}/questions/get/${question_id}`)
    const question = await response.json()

    return {question,}
}

export async function _sendAnswer(answer, question_id) {
    await fetch(`${basePathApi}/users/answer`, {
        method: 'POST',
        body: JSON.stringify({
            answer,
            question_id,
        })
    })

    goto(`../`);
}