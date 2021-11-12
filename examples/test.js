import { Podkillers } from 'k6/x/chaos/experiments';
import { textSummary } from 'https://jslib.k6.io/k6-summary/0.0.1/index.js';

export default function () {
  killPod();
}

// The killPod function terminates a pod within a Kubernetes cluster according to specifications provided.
export function killPod() {
  
  // Instantiate a new Podkiller object
  const podkiller = new Podkillers();
  
  // The line below terminates a random pod in the specified namespace.
  // podkiller.killRandomPod('default');

  // The line below kills a pod within the namespace whose name exactly matches "web-7d55cf8588-7bxpv".
  // podkiller.killPod('default', 'web-7d55cf8588-7bxpv');

  // The line below terminates a pod within the namespace whose name contains "web"
  podkiller.killPodLike('default', 'web');
  
}

// The handleSummary function creates a summary of the chaos experiments after the standard k6 summary.
export function handleSummary(data) {
  return {
      'stdout': textSummary(data, { indent: ' ', enableColors: true}) + new Podkillers().generateSummary(),
  };
}