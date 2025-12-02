# Overview

This is a simple Azure Function that performs URL redirection based on incoming HTTP requests.
It listens for HTTP `GET` requests and redirects the client to a specified target URL.

## Project Structure

You may wonder why the layout of this function is different from a typical Azure Function project.
This is because this function uses a custom handler and has to listen on the root path (`/`) to perform redirection correctly.
Therefore, the project structure and configuration differ from standard Azure Function templates.

### Key Differences

To listen on the root path:

- `function.json` set the `route` property of the `httpTrigger` binding to `/`. This removed the subpath (by defaul `/api/`) from the trigger URL.
- `host.json` sets the `routePrefix` to `/`. This will make the function listen to the root instead of the function name.
- The default page is disabled by setting `AzureWebJobsDisableHomepage` to `true` via app settings.

To use a custom handler:

- The function is built as a standalone executable
- `function.json` specifies the binary of the custom handler via `defaultExecutablePath`
