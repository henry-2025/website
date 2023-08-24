<script>
	let title = "Hi, I'm Henry";
	let subtitle =
		"AI enjoyer, running enthusiast, and machine learning engineer";
</script>

<div class="h-screen">
	<div class="relative left-1/4 top-10">
		<div class="relative justify-center">
			<div
			class="text text-overlay"
			>
			{title}
			</div>
			<div
				class="text text-beneath"
			>
				{title}
			</div>
		</div>
		<div
			class="relative text-xl font-medium font-bold text-lg"
		>
			{subtitle}
		</div>
	</div>
</div>

<style>
	.text {
		text-align: center;
		line-height: 1.5;
		letter-spacing: 0.1em;
		font-weight: bold;
		text-transform: uppercase;
		width: 25%;
		font-size: 8rem;
		z-index: 2;
	}
	.text-overlay{
		position: relative;
		mix-blend-mode: color-burn;
		color: #544d4d;
	}

	.text-beneath{
		position: absolute;
		top: 0;
		left: 0;
		opacity: 0.4;
		color: #0f0f0f;
		visibility:hidden;
	}
</style>