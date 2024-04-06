export async function load(loadEvent) {
    const {fetch, params} = loadEvent
    const {otherusername} = params
    const response = await fetch(`http://localhost:8080/users/${encodeURIComponent(otherusername)}`);
    if(response.ok) {
        const otheruser = await response.json();
        return { otheruser };
    }
    else {
        return {
            status: 404,
            error: new Error("User not found"),
        }
    }
}