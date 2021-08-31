# imgbed

personal image for blog

## imglink

`img2link` is the tool to convert pictures into blog links

### Usage

```text
Usage of ./img2link:
  -c    The option to do commit before convert
  -d string
        Image link domain, choose 'cdn' to use jsDelivr CDN acceleration (default "github")
  -g string
        The path of the .git folder  (default ".")
  -m string
        The message of git commit  (default "upload images")
  -s string
        Image link style, 'md' for markdown, 'url' for http (default "md")
  -t string
        The path or folder of the target image  (default ".")
```

You can directly convert the submitted image into a blog link like this:

```bash
./bin/img2link_unix -c -m "upload images" -d cdn
```

Or convert the pictures in the specified folder to blog links:

```bash
./bin/img2link_unix -t ./img/blog -s url
```
