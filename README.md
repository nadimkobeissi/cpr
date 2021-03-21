# cpr

`cpr` is a simple program that copies files while showing a progress indicator, which displays:

* Percentage copied.
* Size copied so far.
* Estimate of time remaining.

For example:

```
~/D/g/cpr ❱ ./cpr /Users/nadim/Movies/Soul/Soul.mkv soul2.mkv
16% • 2.9GB/17.4GB • 0min remaining
```

It was written because the standard `cp` does not provide this information, which is very annoying in certain cases.

`cpr` cannot do anything other than copy single files while showing the information listed above.

