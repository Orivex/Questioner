export async function load(loadEvent) {
    const {fetch} = loadEvent
    const response = await fetch(`http://localhost:8080/users`);

    if(response.ok) {
        const users = await response.json();
        return { users };
    }
    else {
        return {
            status: 404,
            error: new Error("No users"),
        }
    }
}