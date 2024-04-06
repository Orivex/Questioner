import { basePathApi } from '$lib/index.js';
import { goto } from '$app/navigation'

export async function _sendQuestion(questioner, thequestion, username) {
    await fetch(`${basePathApi}/ask-a-question`, {
        method: 'POST',
        body: JSON.stringify({
            questioner,
            thequestion,
            username,
        })
    })

    goto(`./`);
    
}