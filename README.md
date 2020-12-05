> ### ⚠️ This is a proof of concept
>
> As this is a proof of concept,  it won't be supported by the k6 team.
> It may also break in the future as xk6 evolves. USE AT YOUR OWN RISK!
> Any issues with the tool should be raised [here](https://github.com/simskij/xk6-chaos/issues).

</br>
</br>

<div align="center">

![logo](assets/logo.png)

# xk6-chaos
A k6 extension for testing for the unknown unknowns.
Built for [k6](https://github.com/loadimpact/k6) using [xk6](https://github.com/k6io/xk6).

</div>

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download `xk6`:
  ```bash
  $ go get -u github.com/k6io/xk6
  ```

2. Build the binary:
  ```bash
  $ xk6 build v0.29.0 \
      --with github.com/simskij/xk6-chaos
  ```

## Example

```javascript
import chaos from 'k6/x/chaos';
import { Pods } from 'k6/x/chaos/k8s';

export default function () {
  console.log(`Running simskij/k6-extension-chaos@${chaos.version}.`);
  const p = new Pods();
  console.log(
    `There are currently ${p.list().length} pods in the default namespace.`
  );
  killPod(p);
  console.log(
    `There are now ${p.list().length} pods in the default namespace.`
  );
}

function killPod(p) {
  const victim = p.list()[0];
  console.log(`Killing pod ${victim}`);
  p.killByName('default', victim);
}
```

Result output:

```bash
$ ./k6 run script.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: test.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

INFO[0000] Running simskij/k6-extension-chaos@v0.0.1.             source=console
INFO[0000] There are currently 33 pods in the default namespace.  source=console
INFO[0000] Killing pod chaos-webserver-54bd848884-ds2g9           source=console
INFO[0000] There are now 32 pods in the default namespace.        source=console

running (00m00.1s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.1s/10m0s  1/1 iters, 1 per VU

    data_received........: 0 B 0 B/s
    data_sent............: 0 B 0 B/s
    iteration_duration...: avg=111.72ms min=111.72ms med=111.72ms max=111.72ms p(90)=111.72ms p(95)=111.72ms
    iterations...........: 1   7.513995/s

```
