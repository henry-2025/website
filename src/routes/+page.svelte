<script>
	let title = "Hi, I'm Henry Pick";
	let subtitle = "AI enjoyer, running enthusiast, and aspiring machine learning engineer";
	let mission = ""
</script>
<div class="title-container">
	<div class="text text-under text-center text-under-blended section-title-1">{title}</div>
	<div class="text text-under text-center text-under-overlay section-title-1">{title}</div>
	<div class="subtitle font-mono font-bold text-lg">
			{subtitle}
	</div>
</div>
<style lang="scss">
	.title-container {
		position: relative;
		top: var(--position-text-top);
		left: var(--position-text-left);
		width: fit-content;
	}

	.subtitle{
		position: relative;
		top: 350px;
		font-family: 'Inconsolata';
		font-weight: 500;
		font-size: large;
	}

	:root {
		// --section-gap: calc(#{$standard_space});
		--position-text-top: 0px;
		--position-text-left: 30px;	
		--gradient-padding: 120px;
		--gradient-title-margin: 100px;
		--section-skew-Y: -12deg;
		--section-angle-sin: 0.212;
		--transform-origin-x: calc(var(--section-gap) * 0.8);
		// @include mq($container_medium_width) {
		//     --section-gap: calc(
		//         (100vw - #{$container_medium_width} + #{$standard_space} * 2) /
		//             2
		//     );
		// }
	}
	.section-title-1 {
		font-size: 64px;
		font-weight: 700;
		text-transform: uppercase;
		width: 500px;
		max-width: 80vw;
		font-size: 100px;
		width: 600px;
		font-family: monospace;
		font-weight: 200;
	}

	.text {
		line-height: 1;
		margin: var(--gradient-title-margin) 0 0 0;
		text-transform: none;
		letter-spacing: 2px;
		min-height: 200px;
		display: flex;
		align-items: flex-end;
		font-family: 'Inconsolata';
		font-weight: 700;
	}

	.text-under {
		position: absolute;
		z-index: 2;
	}
	.text-under-blended {
		color: #f32424;
		mix-blend-mode:difference;
	}

	.text-subtitle {
		position: relative;
	}

	.text-under-overlay {
		opacity: 0.5;
		color: #a05c9e;
		mix-blend-mode: color-dodge;
	}
</style>
