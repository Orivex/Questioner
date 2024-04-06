<script>

import { goto } from '$app/navigation'
import { basePathFront} from '$lib/index.js'

</script>
<ul>
    <button on:click={() => goto(`${basePathFront}/login/`)}>Login</button>
    
</ul>