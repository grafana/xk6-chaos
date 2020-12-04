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
  p.killByName('media', victim);
}
