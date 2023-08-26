<script>
	let title = "Hi, I'm Henry";
	let subtitle =
		"AI enjoyer, running enthusiast, and machine learning engineer";
</script>

	<div class="relative ml-[25vw]">
		<div class="relative justify-center">
			<div
			class="text text-overlay text-7xl lg:text-9xl"
			>
			{title}
			</div>
			<div
				class="text text-beneath text-7xl lg:text-9xl"
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

<style>
	.text {
		text-align: center;
		line-height: 1.5;
		letter-spacing: 0.1em;
		font-weight: bold;
		text-transform: uppercase;
		width: 25%;
		z-index: 2;
	}
	.text-overlay{
		position: relative;
		mix-blend-mode: color-burn;
		color: #544d4d;
	}

	.text-beneath{
        position: absolute;
        top:0;
        left:0;
		opacity: 0.2;
		color: #0f0f0f;
	}
</style>
