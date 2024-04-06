export async function load(loadEvent) {
    const {fetch, params} = loadEvent
    const {username} = params
    const response = await fetch(`http://localhost:8080/users/${encodeURIComponent(username)}`);
    if(response.ok) {
        const user = await response.json();
        return { user };
    }
    else {
        return {
            status: 404,
            error: new Error("User not found"),
        }
    }
}