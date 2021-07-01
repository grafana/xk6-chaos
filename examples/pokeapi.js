import http from 'k6/http';
import { sleep, check } from 'k6';
import papaparse from 'https://jslib.k6.io/papaparse/5.1.1/index.js';
import { SharedArray } from 'k6/data';
import { experiments, podkiller } from 'k6/x/chaos/experiments';
import { textSummary } from 'https://jslib.k6.io/k6-summary/0.0.1/index.js';

export const options = {
    scenarios: {
        pokeapi: {
            executor: 'ramping-vus',
            exec: 'catchEmAll',
            startVUs: 1,
            stages: [
                { duration: '1m', target: 1 },
            ],
            gracefulRampDown: '1s',
        },
        chaos: {
            executor: 'per-vu-iterations',
            exec: 'killPod',
            vus: 1,
            iterations: 1,
            startTime: '5s',
        },
    },
    thresholds: {
        http_req_failed: ['rate<0.05'],
        http_req_duration: ['p(95)<100'],
    },
};

const domain = 'http://cluster.nicolevanderhoeven.com/api/v2';
const sharedData = new SharedArray("Shared Logins", function() {
    let data = papaparse.parse(open('pokemon.csv'), { header: true }).data;
    return data;
});

export function catchEmAll() {
    GetPokemon();
    ThinkTime();
}

export function GetPokemon() {
    // Get random mon from shared array
    let randomMon = sharedData[Math.floor(Math.random() * sharedData.length)];
    // console.log(JSON.stringify(randomMon) + ' selected');

    let res = http.get(domain + '/pokemon/' + randomMon.name, {tags: { name: '01_GetPokemon' }});
    check(res, {
        'is status 200': (r) => r.status === 200,
        '01-text verification': (r) => r.body.includes(randomMon.name)
    });
    sleep(Math.random() * 5);
}

export function ThinkTime() {
    sleep(Math.random() * 5);
}

export function killPod() {
    
    // Kill chosen pod.
    // podkiller.killPod('default', 'web-7d55cf8588-s6cmz');
    podkiller.killPodLike('default', 'web');
    // podkiller.killRandomPod('default');
}

export function generateChaosSummary() {
    // let victims = podkiller.getVictims();
    // let numPodsBegin = podkiller.getStartingPods('default');
    // let numPodsNow = podkiller.getNumOfPods('default');

    console.log('Podkillers object: ' + JSON.stringify(podkiller));
    let experimentType = JSON.stringify(podkiller.experiment_type);
    let numPodsBegin = JSON.stringify(podkiller.num_of_pods_before);
    let victims = JSON.stringify(podkiller.victims);
    let numPodsNow = JSON.stringify(podkiller.num_of_pods_after);

    let chaosSummary = `
    
    
    xk6-CHAOS
    -----
    EXPERIMENT: ${experimentType}.
    
    Number of pods before termination: ${numPodsBegin}
    Pod(s) terminated: ${victims}
    Number of pods at the end of the test: ${numPodsNow}

    `;

    return chaosSummary;
}

export function handleSummary(data) {
    
    return {
        'stdout': textSummary(data, { indent: ' ', enableColors: true}) + generateChaosSummary(),
        // 'raw-data.json': JSON.stringify(data)
    };
}