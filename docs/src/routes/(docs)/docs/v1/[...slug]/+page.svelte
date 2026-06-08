<script lang="ts">
	import { DocPage } from "@bladocs/core";
	let { data } = $props();
</script>

<DocPage component={data.component} {...data.metadata} />
