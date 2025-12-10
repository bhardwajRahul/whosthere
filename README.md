# whosthere

Whosthere is a TUI application that discovers devices and services on your local network, built in Go for fast, 
concurrent scanning and a clean terminal interface. I'm building this primarily for myself to deepen my understanding
of Golang and networking fundamentals. Feel free to use it, contribute to it, or suggest features! I'm open to all kinds
of suggestions and feedback.

## Configuration
A lot of behavior within whosthere can be configured to your liking. By default, whosthere will try to look for a configuration
file at `$XDG_CONFIG_HOME/whosthere/config.yaml`, or `~/.config/whosthere/config.yaml` if the [**XDG Base Directory**](https://specifications.freedesktop.org/basedir/latest/#basics)
environment variables are not set. If no configuration file is found, whosthere will create one with default values on first run.
The location of the configuration file can be overridden by providing the `--config` (`-c`) flag when starting whosthere, 
or the `WHOSTHERE_CONFIG` environment variable.

Here is an example configuration file with all available options and their default values:

```yaml
splash:
  enabled: true # show splash screen on startup
  delay: 1      # delay in seconds for the splash screen
theme:         
  # maps 1:1 to tview.Theme https://github.com/rivo/tview/blob/master/styles.go#L6
  primitive_background_color: "#000a1a"
  contrast_background_color: "#001a33"
  more_contrast_background_color: "#003366"
  border_color: "#0088ff"
  title_color: "#00ffff"
  graphics_color: "#00ffaa"
  primary_text_color: "#cceeff"
  secondary_text_color: "#6699ff"
  tertiary_text_color: "#ffaa00"
  inverse_text_color: "#000a1a"
  contrast_secondary_text_color: "#88ddff"
```

## Logging
Whosthere supports logging to a file for debugging and monitoring purposes. By default, logs are written to
`$XDG_STATE_HOME/whosthere/whosthere.log`, or `~/.local/state/whosthere/whosthere.log` if the [**XDG Base Directory**](https://specifications.freedesktop.org/basedir/latest/#basics)
environment variables are not set. The log level is set to `info` by default, but can be changed via the `WHOSTHERE_LOG`
environment variable. 

For example, to set the log level to `debug`, you can start whosthere with the following command:

```bash
WHOSTHERE_LOG=debug whosthere
```

## Platforms
This application has been tested on Linux and macOS. Windows support is not currently available, but contributions
to add Windows compatibility are welcome!

## Engine

### Scanners

...

### OUI Table
https://standards-oui.ieee.org/oui/oui.csv