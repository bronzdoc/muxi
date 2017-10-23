# Muxi - WIP

## Examples

### Simple two pane window

```yaml
windows:
  - name: sample-two-panes
    root: ~/sample
    layout: even-horizontal
    panes:
      - git status
      - rails server
```

```
.------------------.------------------.
| (0)              | (1)              |
|                  |                  |
|                  |                  |
|                  |                  |
|                  |                  |
|                  |                  |
|                  |                  |
|                  |                  |
|                  |                  |
'------------------'------------------'
```

### Simple three pane window

```yaml
windows:
  - name: sample-three-panes
    root: ~/sample
    layout: main-vertical
    panes:
      - vim
      - commands:
        - git pull
        - git status
      - rails server
```

```
.------------------.------------------.
| (0)              | (1)              |
|                  |                  |
|                  |                  |
|                  |                  |
|                  |------------------|
|                  | (2)              |
|                  |                  |
|                  |                  |
|                  |                  |
'------------------'------------------'
```

### Simple four pane window

```yaml
windows:
  - name: sample-four-panes
    root: ~/sample
    layout: tiled
    panes:
      - vim test.yml
      - go get github.com/goby-lang/goby
      - git status
      - go test
```

```
.------------------.------------------.
| (0)              | (1)              |
|                  |                  |
|                  |                  |
|                  |                  |
|------------------|------------------|
| (2)              | (3)              |
|                  |                  |
|                  |                  |
|                  |                  |
'------------------'------------------'
```
