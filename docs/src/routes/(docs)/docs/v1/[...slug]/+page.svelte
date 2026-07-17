<script lang="ts">
	import { DocPage } from "@penbot/core";
	let { data } = $props();
</script>

<DocPage component={data.component} {...data.metadata} />
