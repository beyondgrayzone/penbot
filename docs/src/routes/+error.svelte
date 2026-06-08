<script lang="ts">
	import { page } from "$app/state";
	import { Button } from "$lib/components/ui/button/index.js";
	import {
		Card,
		CardContent,
		CardDescription,
		CardFooter,
		CardHeader,
		CardTitle,
	} from "$lib/components/ui/card";
</script>

<div class="flex min-h-screen items-center justify-center p-4">
	<Card class="bg-background w-full max-w-md text-center">
		<CardHeader>
			<CardTitle class="text-default text-6xl font-bold">{page.status}</CardTitle>
			<CardDescription class="text-xl">{page.error?.message}</CardDescription>
		</CardHeader>
		<CardContent>
			{#if page.status >= 400 && page.status < 500}
				<p class="text-muted-foreground">
					The page you are looking for might have been removed, had its name changed, or
					is temporarily unavailable.
				</p>
			{:else}
				Something went wrong!
			{/if}
		</CardContent>
		<CardFooter class="flex justify-center">
			<Button href="/">Go to Homepage</Button>
		</CardFooter>
	</Card>
</div>
