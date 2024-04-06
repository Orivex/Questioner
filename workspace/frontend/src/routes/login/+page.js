import { goto } from '$app/navigation'
import { basePathApi } from '$lib/index.js';

export async function _registerUser (username) {
    const userExistsResponse = await fetch(`${basePathApi}/auth/userExists/${encodeURIComponent(username)}`)
    const userExists = await userExistsResponse.json()
    if (userExists) {
        alert("User already exists!")
        return;
    }

    const response = await fetch(`${basePathApi}/auth/register`, {
        method: 'POST',
        body: JSON.stringify({
            username,
        })
    })
    
    const user = await response.json()

    goto(`/users/${username}/${user.password}`, { replaceState: true });
}

export async function _loginUser (username, password) {
    const response = await fetch(`${basePathApi}/auth/login`, {
        method: 'POST',
        body: JSON.stringify({
            username,
            password,
        })
    })

    if (response.ok) {
        goto(`/users/${username}/${password}`, { replaceState: true });
    } 
    else {
        alert("Invalid password or username");
    }
}