# cgocheck

By default, `cgocheck=1`, which can be optimized away for pure Go code?

> These rules are checked dynamically at runtime. The checking is controlled by the cgocheck setting of the GODEBUG environment variable. The default setting is GODEBUG=cgocheck=1, which implements reasonably cheap dynamic checks. These checks may be disabled entirely using GODEBUG=cgocheck=0. Complete checking of pointer handling, at some cost in run time, is available via GODEBUG=cgocheck=2. 