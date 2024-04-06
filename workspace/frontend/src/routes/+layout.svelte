<slot/>

<style>
    :global(.grid) {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
        gap: 10px;
    }

    :global(input) {
        background-color: brown;
        color: burlywood;
    }

    :global(textarea) {
        background-color: brown;
        color: burlywood;
    }

    :global(div) {
        font-family: Arial, Helvetica, sans-serif;
        justify-content: center;
        align-items: center;
    }

    :global(span) {
        color:burlywood;
        display: inline-block;
    }

    :global(.container) {
        font-family: Arial, Helvetica, sans-serif;
        justify-content: center;
        align-items: center;  
        background-color:burlywood;  
    }

    :global(body) {
        background-color: brown;
    }

    :global(button) {
        font-size: medium;
        border: solid 2px;
        border-radius: 75px;
        color: burlywood;
        background-color: brown;
        margin: 5px;
        height: auto;   
        width: auto;
        justify-content: center;
        align-items: center;
    }

    :global(p) {
        color: burlywood;
        padding: 20px;
        font-size: large;
    }

    :global(.question) {
        color: burlywood;
    }

    :global(.answer) {
        color: darkslategray;
    }
</style>
