<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>jhpick</title>
    <link rel="icon" href="/favicon.png" type="image/x-icon" />
    <link rel="mask-icon" href="/favicon.png" type="image/x-icon" />
  </head>

  <body>
    <header>
      <nav>
        <div class="logo-container">
          <a href="/" class="logo"> JHP </a>
        </div>
        <div class="nav-buttonset">
          <a href="/projects" class="nav-item">Projects</a>
          <a href="/contact" class="nav-item">Contact</a>
          <a href="/about" class="nav-item">About</a>
        </div>
      </nav>
    </header>
    <slot></slot>
  </body>

  <footer>
    <p>Copyright all rights reserved JHP Consulting LLC</p>
    <div class="socials">
      <div class="github">
        <a href="https://www.github.com/henry-2025">
          <svg height="24" viewBox="0 0 24 24" width="24" fill="var(--primary)">
            <path
              class="cls-1"
              d="m20.99,5.06C17.92,1.56,14.1.57,13.04.4c-.08-.01-.18-.03-.33-.05C8.86-.2,5.45,1.66,4.13,2.67-1.53,6.95-.47,19.26,3.94,21.66c2.13,1.16,6.84.81,6.84.81.14-.73.07-1.67-.57-2.16-.94.18-1.7.11-2.18.03-.53-.09-1.51-.27-2.37-1.02-.45-.4-.94-1.01-.81-1.19.14-.19.92.35,1.94.59.83.2,1.94.24,3.37-.27,0-.45.02-.9.03-1.35-.57.05-2.26.13-3.93-.94-1.4-.9-2.02-2.11-2.24-2.56-.19-.37-.93-1.97-.59-3.99.09-.52.24-1.37.86-2.26.55-.79,1.22-1.23,1.62-1.45-.04-.2-.32-1.54.13-1.78.19-.1.44.03,1.16.46,1,.6,1.4.89,1.66,1.04,1.46.86,4.01.49,5.18-.77.45-.49.52-.92.94-.97.66-.08,1.4.87,1.72,1.67.16.39.13.53.24.97.3,1.18.93,1.8,1.24,2.29,1.07,1.72,0,4.74-1.19,6.33-.93,1.26-1.09.43-2.88,1.99l.19,1.32c-.78,1.45.46,3.8.46,3.8,1.41.97,4.19-.81,5.63-2.32,2.23-2.33,2.51-5.28,2.59-6.63.29-5.09-1.22-7.36-1.99-8.24Z"
            />
          </svg>
        </a>
      </div>
      <div class="linkedin">
        <a href="https://www.linkedin.com/in/henry-pick/">
          <svg viewBox="0 0 24 24" fill="var(--primary)" width="24" height="24">
            <path
              d="m21.25,14.2c-.12-1.79.08-1.44.16-5.9.08-4.48.13-6.75-.59-7.46-1.49-1.47-3.09.34-8.78-.11-2.47-.19-6.34-.86-9.83-.03-.41.1-1.03.27-1.43.81-.86,1.18.29,2.87.78,5.31.29,1.45.09,3.08-.32,6.33-.61,4.86-1.56,6.57-.54,7.65.63.67,1.5.53,3.53.51,5.83-.05,6.7,1.04,12.66.65,3.77-.25,4.38-.95,4.66-1.35.88-1.25-.04-2.57-.3-6.41Zm-13.69-3.61c.64,4.79-.77,7.57-1.4,7.88-.04.02-.08.03-.11.04-.6.15-1.29-.43-1.62-.86-.49-.66-.3-1.33-.13-3.53.17-2.27,0-2,.15-2.83.35-1.83,1.37-4,2.08-3.91.61.08.83,1.77,1.02,3.21Zm-.38-4.82c-.54.31-1.31.23-1.8-.22-.64-.57-.73-1.65-.11-2.32.59-.64,1.6-.66,2.18-.19.81.67.63,2.2-.27,2.72Zm11.14,8.85c-.04.74-.02,3.5-1.37,4.04-.14.06-.58.23-.89.03-1-.68,1.34-4.17-.08-6.92-.14-.26-.45-.87-.94-.92-.81-.08-1.56,1.4-1.7,1.67-1.26,2.51-.14,4.62-1.41,5.52-.07.05-.51.37-.91.22-.21-.07-.62-.36-.75-2.4-.09-1.4.03-2.45.11-3.16.42-3.91.28-5.18,1.35-6.33.18-.2.57-.6.86-.5.73.24-.1,3.23.86,3.66.81.37,1.52-1.05,2.56-.92,1.59.2,2.43,3.72,2.32,6.01Z"
            />
          </svg>
        </a>
      </div>
    </div>
  </footer>
</html>

<style>
  :root {
    font-family: "Inconsolata", monospace;
    font-optical-sizing: auto;
    font-weight: 300;
    font-style: normal;
    font-variation-settings: "wdth" 100;
  }
  /*Header Style*/
  header {
    font-variation-settings: "wdth" 125;
    font-weight: 600;
  }

  .logo {
    font-size: 1.5em;
  }

  nav {
    margin: 1em 1em;
    display: flex;
    flex-direction: row;
    align-items: center;
    flex-grow: 1;
  }

  header {
    width: 100%;
    z-index: 2;
  }

  .logo-container {
    margin: 0 1em;
  }

  .nav-buttonset {
    display: flex;
    flex-direction: row;
    flex-grow: 1;
    justify-content: right;
  }

  .nav-buttonset > a {
    margin: 0 1em;
  }

  .nav-item {
    text-decoration: underline;
  }

  /*Primary Document Styles*/
  body {
    margin: 0, 10em;
    min-height: 100vh;
    display: flex;
    flex-direction: column;
  }

  /*Footer Styles*/
  footer {
    display: flex;
    flex-direction: row;
    font-variation-settings: "wdth" 125;
    justify-content: center;
    align-items: center;
    background: linear-gradient(
      to bottom,
      rgba(255, 255, 255, 0),
      rgba(255, 255, 255, 1)
    );
    z-index: 2;
  }

  .socials {
    display: flex;
    flex-direction: row;
    margin: 1em;
  }

  .socials > div {
    margin: 1em;
  }

  a {
    color: inherit;
    text-decoration: inherit;
  }
</style>
